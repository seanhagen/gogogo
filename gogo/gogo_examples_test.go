package gogo

// func ExampleCreateGame() {
// 	// so the example always generates the same code
// 	rand.Seed(2)

// 	// on a server somewhere
// 	board, err := NewBoard(5)
// 	if err != nil {
// 		// handle the error
// 	}

// 	// The board code is used when looking up games to add a player
// 	code := board.Code()
// 	fmt.Printf("Code: %q\n", code)

// 	// String() outputs the board in a pretty basic string
// 	// representation. When no pieces have been placed, empty points are
// 	// represented by 'X'.
// 	fmt.Printf("Board:\n%v\n", board.String())

// 	// As the board has just been created ( ie, it's the start of a game
// 	// ), the first player should be Black.
// 	currentPlayer := board.CurrentPlayer()
// 	fmt.Printf("Current player: %q", currentPlayer.String())
// 	currentPlayerPieces := currentPlayer.Pieces()
// 	fmt.Printf(", pieces on the board: %v\n", len(currentPlayerPieces))

// 	// Points on the board are referenced on the horizontal axis by
// 	// letters, and on the vertical axis by numbers. 'A0' is the bottom
// 	// left position on the board, and E5 is the top right on a 5x5 board.
// 	result, err := board.Place("A0")
// 	// Err will be non-nil if the placement was invalid, such as
// 	// attempting to place on top of another piece, or attempting to
// 	// place a piece off the board.
// 	if err != nil {
// 		// handle the error
// 	}

// 	// The main return value returns a struct with information about the
// 	// state of the game; whether or not it's over, how many of each
// 	// player's pieces are on the board, and how many pieces have been
// 	// captured.
// 	fmt.Printf("Game Over: %q\n", result.GameOver())
// 	b, w := result.Pieces()
// 	fmt.Printf("Pieces on board; Black: %v, White: %v\n", b, w)
// 	cb, cw := result.Captured()
// 	fmt.Printf("# of captured Black pieces: %v\n", cb)
// 	fmt.Printf("# of captured White pieces: %v\n", cw)

// 	// Now that black has played a piece, it's White's turn.
// 	fmt.Printf("Current player: %q\n", board.CurrentPlayer())

// 	// And we should see Black's piece on the board.
// 	fmt.Printf("Board:\n%v\n", board.String())

// 	// Output:
// 	// Code: 'A3KJ'
// 	// Board:
// 	// 5 XXXXX
// 	// 4 XXXXX
// 	// 3 XXXXX
// 	// 2 XXXXX
// 	// 1 XXXXX
// 	//   ABCDE
// 	// Current player: 'black', pieces on the board: 0
// 	// Game Over: 'false'
// 	// Pieces on board; Black: 1, White: 0
// 	// # of captured Black pieces: 0
// 	// # of captured White pieces: 0
// 	// Current player: 'white'
// 	// Board:
// 	// 5 XXXXX
// 	// 4 XXXXX
// 	// 3 XXXXX
// 	// 2 XXXXX
// 	// 1 BXXXX
// 	//   ABCDE
// }
