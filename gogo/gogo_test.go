package gogo

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGameBoard(t *testing.T) {
	tests := []struct {
		size   int
		output string
		ok     bool
	}{
		// smallest board size is 4x4, so any boards smaller than that
		// return an error
		{size: 0},
		{size: 1},
		{size: 2},
		{size: 3},
		// a board size of 4 is valid, so we should get a board and be
		// able to get the string representation
		{
			size: 4,
			output: `4 X X X X
3 X X X X
2 X X X X
1 X X X X
  A B C D`,
			ok: true,
		},

		{
			size: 5,
			output: `5 X X X X X
4 X X X X X
3 X X X X X
2 X X X X X
1 X X X X X
  A B C D E`,
			ok: true,
		},

		{
			size: 10,
			output: `10 X X X X X X X X X X
 9 X X X X X X X X X X
 8 X X X X X X X X X X
 7 X X X X X X X X X X
 6 X X X X X X X X X X
 5 X X X X X X X X X X
 4 X X X X X X X X X X
 3 X X X X X X X X X X
 2 X X X X X X X X X X
 1 X X X X X X X X X X
   A B C D E F G H I J`,
			ok: true,
		},

		// to make life easier on myself, 26 is the largest board allowed
		{size: 27},

		// negative size should obviously be an error
		{size: -1},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v size %v ok %v", i, tt.size, tt.ok), func(t *testing.T) {
			var board Board
			var err error

			board, err = NewBoard(tt.size)
			if !tt.ok {
				assert.Error(t, err)
				assert.Equal(t, "Invalid Board", board.String())
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.output, board.String())
		})
	}
}

func TestBoardsGenerateBoardCode(t *testing.T) {
	tests := []struct {
		seed   int64
		expect string
	}{
		{2, "KMYC"},
		{3, "5RGY"},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v seed %v expect %v", i, tt.seed, tt.expect), func(t *testing.T) {
			board, err := NewBoard(9)
			require.NoError(t, err)

			board.rand = rand.New(rand.NewSource(tt.seed))
			code := board.Code()

			assert.Len(t, code, 4)
			assert.Equal(t, tt.expect, code)
			assert.Equal(t, tt.expect, board.Code())
		})
	}
}

func TestPlacePieceOnBoard(t *testing.T) {
	tests := []struct {
		size   int
		place  string
		valid  bool
		expect string
	}{
		{size: 4, place: ""},
		{size: 4, place: "A"},
		{size: 4, place: "1"},
		{size: 4, place: "A0"},
		{size: 4, place: "A-1"},
		{size: 4, place: "A5"},
		{size: 4, place: "E1"},
		{size: 4, place: "E5"},
		{
			size:  4,
			place: "A1",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X X X
1 B X X X
  A B C D`,
		}, {
			size:  4,
			place: "A2",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 B X X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "A3",
			valid: true,
			expect: `4 X X X X
3 B X X X
2 X X X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "A4",
			valid: true,
			expect: `4 B X X X
3 X X X X
2 X X X X
1 X X X X
  A B C D`,
		},

		{
			size:  4,
			place: "B1",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X X X
1 X B X X
  A B C D`,
		}, {
			size:  4,
			place: "B2",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X B X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "B3",
			valid: true,
			expect: `4 X X X X
3 X B X X
2 X X X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "B4",
			valid: true,
			expect: `4 X B X X
3 X X X X
2 X X X X
1 X X X X
  A B C D`,
		},

		{
			size:  4,
			place: "C1",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X X X
1 X X B X
  A B C D`,
		}, {
			size:  4,
			place: "C2",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X B X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "C3",
			valid: true,
			expect: `4 X X X X
3 X X B X
2 X X X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "C4",
			valid: true,
			expect: `4 X X B X
3 X X X X
2 X X X X
1 X X X X
  A B C D`,
		},

		{
			size:  4,
			place: "D1",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X X X
1 X X X B
  A B C D`,
		}, {
			size:  4,
			place: "D2",
			valid: true,
			expect: `4 X X X X
3 X X X X
2 X X X B
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "D3",
			valid: true,
			expect: `4 X X X X
3 X X X B
2 X X X X
1 X X X X
  A B C D`,
		}, {
			size:  4,
			place: "D4",
			valid: true,
			expect: `4 X X X B
3 X X X X
2 X X X X
1 X X X X
  A B C D`,
		},
		{
			size:  5,
			place: "E1",
			valid: true,
			expect: `5 X X X X X
4 X X X X X
3 X X X X X
2 X X X X X
1 X X X X B
  A B C D E`,
		},
		{
			size:  19,
			place: "A19",
			valid: true,
			expect: `19 B X X X X X X X X X X X X X X X X X X
18 X X X X X X X X X X X X X X X X X X X
17 X X X X X X X X X X X X X X X X X X X
16 X X X X X X X X X X X X X X X X X X X
15 X X X X X X X X X X X X X X X X X X X
14 X X X X X X X X X X X X X X X X X X X
13 X X X X X X X X X X X X X X X X X X X
12 X X X X X X X X X X X X X X X X X X X
11 X X X X X X X X X X X X X X X X X X X
10 X X X X X X X X X X X X X X X X X X X
 9 X X X X X X X X X X X X X X X X X X X
 8 X X X X X X X X X X X X X X X X X X X
 7 X X X X X X X X X X X X X X X X X X X
 6 X X X X X X X X X X X X X X X X X X X
 5 X X X X X X X X X X X X X X X X X X X
 4 X X X X X X X X X X X X X X X X X X X
 3 X X X X X X X X X X X X X X X X X X X
 2 X X X X X X X X X X X X X X X X X X X
 1 X X X X X X X X X X X X X X X X X X X
   A B C D E F G H I J K L M N O P Q R S`,
		},
		{
			size:  19,
			place: "S1",
			valid: true,
			expect: `19 X X X X X X X X X X X X X X X X X X X
18 X X X X X X X X X X X X X X X X X X X
17 X X X X X X X X X X X X X X X X X X X
16 X X X X X X X X X X X X X X X X X X X
15 X X X X X X X X X X X X X X X X X X X
14 X X X X X X X X X X X X X X X X X X X
13 X X X X X X X X X X X X X X X X X X X
12 X X X X X X X X X X X X X X X X X X X
11 X X X X X X X X X X X X X X X X X X X
10 X X X X X X X X X X X X X X X X X X X
 9 X X X X X X X X X X X X X X X X X X X
 8 X X X X X X X X X X X X X X X X X X X
 7 X X X X X X X X X X X X X X X X X X X
 6 X X X X X X X X X X X X X X X X X X X
 5 X X X X X X X X X X X X X X X X X X X
 4 X X X X X X X X X X X X X X X X X X X
 3 X X X X X X X X X X X X X X X X X X X
 2 X X X X X X X X X X X X X X X X X X X
 1 X X X X X X X X X X X X X X X X X X B
   A B C D E F G H I J K L M N O P Q R S`,
		},
		{
			size:  19,
			place: "S19",
			valid: true,
			expect: `19 X X X X X X X X X X X X X X X X X X B
18 X X X X X X X X X X X X X X X X X X X
17 X X X X X X X X X X X X X X X X X X X
16 X X X X X X X X X X X X X X X X X X X
15 X X X X X X X X X X X X X X X X X X X
14 X X X X X X X X X X X X X X X X X X X
13 X X X X X X X X X X X X X X X X X X X
12 X X X X X X X X X X X X X X X X X X X
11 X X X X X X X X X X X X X X X X X X X
10 X X X X X X X X X X X X X X X X X X X
 9 X X X X X X X X X X X X X X X X X X X
 8 X X X X X X X X X X X X X X X X X X X
 7 X X X X X X X X X X X X X X X X X X X
 6 X X X X X X X X X X X X X X X X X X X
 5 X X X X X X X X X X X X X X X X X X X
 4 X X X X X X X X X X X X X X X X X X X
 3 X X X X X X X X X X X X X X X X X X X
 2 X X X X X X X X X X X X X X X X X X X
 1 X X X X X X X X X X X X X X X X X X X
   A B C D E F G H I J K L M N O P Q R S`,
		},
	}

	for i, x := range tests {
		tt := x
		t.Run(
			fmt.Sprintf("test %v board size %v place %v valid %v", i, tt.size, tt.place, tt.valid),
			func(t *testing.T) {
				board, err := NewBoard(tt.size)
				require.NoError(t, err)

				var res Result
				res, err = board.Place(tt.place)

				if !tt.valid {
					assert.Error(t, err, "expected error for input %q", tt.place)
					return
				}

				require.NoError(t, err, "expected no error for input %q", tt.place)

				expectBlackPieces := 1
				expectWhitePieces := 0

				gotBlack, gotWhite := res.Pieces()
				assert.Equal(t, expectBlackPieces, gotBlack)
				assert.Equal(t, expectWhitePieces, gotWhite)

				got := board.String()
				assert.Equal(t, tt.expect, got, "expected board:\n%v\n\ngot board:\n%v\n", tt.expect, got)
			})
	}
}

func TestHandleAlternatingPlayerInputs(t *testing.T) {
	tests := []struct {
		size              int
		inputs            []string
		valid             bool
		expectBoard       string
		expectPlayer      string
		expectBlackPieces int
		expectWhitePieces int
	}{
		// simple test to see if the player gets updated correctly
		{
			size:   4,
			inputs: []string{"A1"},
			valid:  true,
			expectBoard: `4 X X X X
3 X X X X
2 X X X X
1 B X X X
  A B C D`,
			expectPlayer:      whitePlayer,
			expectBlackPieces: 1,
			expectWhitePieces: 0,
		},
		// now check that white piece gets placed properly
		{
			size:   4,
			inputs: []string{"A1", "A2"},
			valid:  true,
			expectBoard: `4 X X X X
3 X X X X
2 W X X X
1 B X X X
  A B C D`,
			expectPlayer:      blackPlayer,
			expectBlackPieces: 1,
			expectWhitePieces: 1,
		},
		// check that you can't place on top of an existing piece
		{
			size:   4,
			inputs: []string{"A1", "A1"},
		},
		{
			size:   4,
			inputs: []string{"B1", "B2", "A2", "A3", "C2", "C3", "B3"},
			valid:  true,
			expectBoard: `4 X X X X
3 W B W X
2 B X B X
1 X B X X
  A B C D`,
			expectPlayer:      whitePlayer,
			expectBlackPieces: 4,
			expectWhitePieces: 2,
		},
	}

	for i, x := range tests {
		tt := x
		t.Run(
			fmt.Sprintf("test %v inputs %v", i, strings.Join(tt.inputs, "_")),
			func(t *testing.T) {
				board, err := NewBoard(tt.size)
				require.NoError(t, err)

				var merr error

				var gotWhitePieces int
				var gotBlackPieces int

				for _, in := range tt.inputs {
					res, err := board.Place(in)
					if err != nil {
						merr = multierror.Append(merr, err)
					}
					gotBlackPieces, gotWhitePieces = res.Pieces()
				}

				if !tt.valid {
					assert.Error(t, merr)
					return
				}

				require.NoError(t, merr)

				got := board.String()
				assert.Equal(t, tt.expectBoard, got, "expected board:\n%v\n\ngot board:\n%v\n", tt.expectBoard, got)

				assert.Equal(t, tt.expectPlayer, board.CurrentPlayer())
				assert.Equal(t, tt.expectBlackPieces, gotBlackPieces, "wrong number of black pieces")
				assert.Equal(t, tt.expectWhitePieces, gotWhitePieces, "wrong number of white pieces")
			})
	}
}

func TestFindPieceStringOnBoard(t *testing.T) {
	tests := []struct {
		size          int
		boardState    func(b *Board)
		check         coord
		expectStrings []int
	}{
		// no pieces, shouldn't return any strings
		{
			size:       4,
			boardState: func(_ *Board) {},
			check:      coord{1, 1, blackPiece},
		},
		// two pieces, not next to each other, should return two strings
		{
			size: 4,
			boardState: func(b *Board) {
				b.board[1] = blackPiece
				b.board[9] = blackPiece

				// return out
			},
			check:         coord{2, 1, blackPiece},
			expectStrings: []int{1},
		},
		// two pieces, that ARE next to each other, should return a single string
		{
			size: 4,
			boardState: func(b *Board) {
				b.board[5] = blackPiece
				b.board[9] = blackPiece
			},
			check:         coord{2, 1, blackPiece},
			expectStrings: []int{5, 9},
		},
	}

	for i, x := range tests {
		tt := x
		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
			board, err := NewBoard(tt.size)
			require.NoError(t, err)

			tt.boardState(&board)

			// _, err = board.Place(tt.check.AsPosition())
			// require.NoError(t, err)

			gotStrings := board.getString(tt.check.x, tt.check.y, tt.check.val)
			assert.Equal(t, tt.expectStrings, gotStrings)
		})
	}
}

/*

   A4 B4 C4 D4    1,4  2,4  3,4  4,4    3  7  11 15

   A3 B3 C3 D3    1,3  2,3  3,3  4,3    2  6  10 14
               =>                    =>
   A2 B2 C2 D2    1,2  2,2  3,2  4,2    1  5  9  13

   A1 B1 C1 D1    1,1  2,1  3,1  4,1    0  4  8  12


   X B X X
   X X B X
   X X B X
   X B B X

   [
     "B4",
     "C3,C2,C1,B1"
   ]

   toCheck = [0,1,2,..,14,15]
   checked = map[int]bool{}

   lookDirs = [
          [-1,0],
   [0,-1],        [0,1],
          [1,0]
   ]

   for i in toCheck {
     checked[i]= true


     for l in lookDirs {

     }
   }



*/
