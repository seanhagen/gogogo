package gogo

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	MinBoardSize int = 4
	MaxBoardSize int = 26

	emptySpace rune = 'X'
	blackPiece rune = 'B'
	whitePiece rune = 'W'

	blackPlayer string = "black"
	whitePlayer string = "white"

	charset string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ3456789"
	codeLen int    = 4
)

type Board struct {
	size          int
	code          string
	board         []rune
	nextPiece     rune
	currentPlayer string
	rand          *rand.Rand
}

// NewBoard ...
func NewBoard(boardSize int) (Board, error) {
	if boardSize < MinBoardSize {
		return Board{}, fmt.Errorf("%q smaller minimum board size %q", boardSize, MinBoardSize)
	}

	if boardSize > MaxBoardSize {
		return Board{}, fmt.Errorf("%q larger than maximum board size %q", boardSize, MaxBoardSize)
	}

	return Board{
		size:          boardSize,
		board:         buildBoard(boardSize),
		rand:          rand.New(rand.NewSource(time.Now().UnixNano())),
		currentPlayer: blackPlayer,
		nextPiece:     blackPiece,
	}, nil
}

// String ...
func (b Board) String() string {
	if b.board == nil {
		return "Invalid Board"
	}

	sb := bytes.NewBuffer(nil)
	vert := "%s "
	if b.size > 9 {
		vert = "%2s "
	}

	for i := b.size; i >= 1; i-- {
		sb.WriteString(fmt.Sprintf(vert, fmt.Sprintf("%v", i)))
		for j := 1; j <= b.size; j++ {
			idx := b.coordsToIdx(i, j)
			fmt.Printf("(%v,%v) -> %v,%q", i, j, idx, string(b.board[idx]))
			if j == 1 {
				sb.WriteRune(b.board[idx])
			} else {
				sb.WriteString(fmt.Sprintf("%2s", string(b.board[idx])))
			}
			fmt.Printf("\t")
		}
		fmt.Printf("\n")
		sb.WriteString("\n")
	}

	if b.size > 9 {
		sb.WriteString(fmt.Sprintf("%4s", "A"))
	} else {
		sb.WriteString(fmt.Sprintf("%3s", "A"))
	}

	for i := 1; i < b.size; i++ {
		sb.WriteString(fmt.Sprintf("%2s", string(rune(65+i))))
	}

	return sb.String()
}

// Code ...
func (b *Board) Code() string {
	if len(b.code) > 0 {
		return b.code
	}
	c := make([]byte, codeLen)
	for i := range c {
		c[i] = charset[b.rand.Intn(len(charset))]
	}
	b.code = string(c)
	return b.code
}

// CurrentPlayer ...
func (b *Board) CurrentPlayer() string {
	return b.currentPlayer
}

// Place ...
func (b *Board) Place(input string) (Result, error) {
	if len(input) < 2 {
		return Result{}, fmt.Errorf("invalid input %q", input)
	}

	idx, err := b.canPlaceAt(input)
	if err != nil {
		return Result{}, err
	}
	b.board[idx] = b.nextPiece

	numBlackPieces, numWhitePieces := b.advanceToNextTurn()
	return Result{blackPieces: numBlackPieces, whitePieces: numWhitePieces}, nil
}

// validCoordinates ...
func (b Board) validCoordinates(i, j int) bool {
	//fmt.Printf("validCoordinates(%v, %v) =>  len(b.board): %v\n", i, j, idx, len(b.board))
	return i > 0 && j > 0 && i <= b.size && j <= b.size
	// if i < 0 || j < 0 || i > b.size ||
	// idx := b.coordsToIdx(i, j)
	// return idx > 0 && idx < len(b.board)
}

// pieceAt ...
func (b Board) pieceAt(i, j int) coord {
	if !b.validCoordinates(i, j) {
		return coord{}
	}
	idx := b.coordsToIdx(i, j)
	fmt.Printf("board: \n%v\n", b.String())

	fmt.Printf("pieceAt(%v,%v) => %v (idx: %v, size: %v)\n", i, j, string(b.board[idx]), idx, b.size)
	return coord{i, j, b.board[idx]}
}

// getNeighbours ...
func (b Board) getNeighbours(i, j int) []coord {
	look := []coord{
		{x: i - 1, y: j},
		{x: i + 1, y: j},
		{x: i, y: j - 1},
		{x: i, y: j + 1},
	}
	var zero coord
	current := b.pieceAt(i, j)
	fmt.Printf("checking for neighbours of (%v,%v)\n", i, j)
	neighbours := []coord{}
	for _, l := range look {
		check := b.pieceAt(l.x, l.y)
		//fmt.Printf("check == zero? %v\n", check == zero)
		if !check.SamePos(zero) && check.val == current.val {
			fmt.Printf("\tneighbour (%v,%v) -> %v\n", l.x, l.y, check)
			neighbours = append(neighbours, check)
		}
	}

	return neighbours
}

// advanceToNextTurn ...
func (b *Board) advanceToNextTurn() (int, int) {
	if b.nextPiece == blackPiece {
		b.nextPiece = whitePiece
		b.currentPlayer = whitePlayer
	} else {
		b.nextPiece = blackPiece
		b.currentPlayer = blackPlayer
	}

	numBlackPieces, numWhitePieces := 0, 0
	for _, p := range b.board {
		if p == blackPiece {
			numBlackPieces++
		} else if p == whitePiece {
			numWhitePieces++
		}
	}

	return numBlackPieces, numWhitePieces
}

// coordsToIdx ...
func (b *Board) coordsToIdx(i, j int) int {

	i -= 1
	j -= 1

	// swap to make things work
	i, j = j, i
	return ((i * b.size) + j)
}

// inputToIdx ...
func (b Board) inputToIdx(input string) (int, error) {
	bits := strings.Split(input, "")
	vert := bits[0]
	horz := strings.Join(bits[1:], "")

	y := strings.Index(charset, vert)
	x, err := strconv.Atoi(horz)
	if err != nil {
		return -1, fmt.Errorf("invalid horizontal position %q", horz)
	}
	y += 1

	if x <= 0 || x > b.size {
		return -1, fmt.Errorf("invalid horizontal position %q for board size %v", input, b.size)
	}
	if y <= 0 || y > b.size {
		return -1, fmt.Errorf("invalid vertical position %q for board size %v", input, b.size)
	}

	return b.coordsToIdx(x, y), nil
}

// canPlaceAt ...
func (b Board) canPlaceAt(input string) (int, error) {
	idx, err := b.inputToIdx(input)
	if err != nil {
		return -1, fmt.Errorf("can't place at %q, %w", input, err)
	}

	// simple check -- is there a piece there?
	if p := b.board[idx]; p != emptySpace {
		return 0, fmt.Errorf("position %q already occupied", input)
	}

	return idx, nil
}

// getStrings ...
func (b Board) getString(x, y int, pieceType rune) []int {
	fmt.Printf("getString(%v, %v, %v)\n", x, y, string(pieceType))
	piece := b.pieceAt(x, y)

	if piece.val != pieceType {
		fmt.Printf("looking for %q, piece at (%v,%v) is %q\n", string(pieceType), x, y, string(piece.val))
		return nil
	}

	visited := map[coord]bool{}
	queue := []coord{piece}

	var current coord

	for len(queue) > 0 && len(queue) < len(b.board) {
		current, queue = pop(queue)
		fmt.Printf("current (%v,%v -> %v): %v \n", current.x, current.y, string(current.val), current)
		fmt.Printf("queue: %v\n", queue)

		visited[current] = true
		neighbours := b.getNeighbours(current.x, current.y)
		fmt.Printf("neighbours: \n%v\n", neighbours)

		fmt.Printf("\n----------\n")

		for _, n := range neighbours {
			if _, ok := visited[n]; !ok {
				queue = append(queue, n)
			}
		}

	}

	fmt.Printf("queue: \n")
	spew.Dump(queue)

	fmt.Printf("\ncurrent: %v\n", current)
	fmt.Printf("visited: \n")
	spew.Dump(visited)

	return nil
}

// buildBoard ...
func buildBoard(size int) []rune {
	len := size * size
	arr := make([]rune, len)
	for i := 0; i < len; i++ {
		arr[i] = 'X'
	}
	return arr
}

type coord struct {
	x, y int
	val  rune
}

// coordFromIndex ...
func coordFromIndex(idx, size int, val rune) coord {
	if idx < 0 || idx > size {
		return coord{}
	}
	x := idx / size
	y := idx % size

	return coord{x: x, y: y, val: val}
}

// String ...
func (c coord) String() string {
	x := string(rune(65 + c.x))
	return fmt.Sprintf("{%v%d: %v}", x, c.y, string(c.val))
}

// AsPosition ...
func (c coord) AsPosition() string {
	x := string(rune(65 + c.x))
	return fmt.Sprintf("%s%v", x, c.y)
}

// Index ...
func (c coord) Index(size int) int {
	i := c.x - 1
	j := c.y - 1

	// swap to make things work
	i, j = j, i

	return ((i * size) + j)
}

// SamePos ...
func (c coord) SamePos(other coord) bool {
	return c.x == other.x && c.y == other.y
}

func pop[T any](queue []T) (T, []T) {
	if len(queue) == 0 {
		var noop T
		return noop, queue[0:0]
	} else if len(queue) == 1 {
		return queue[0], queue[0:0]
	}
	return queue[0], queue[1:]
}
