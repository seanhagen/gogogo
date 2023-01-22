package gogo

type Result struct {
	blackPieces int
	whitePieces int
}

// Pieces ...
func (r Result) Pieces() (int, int) {
	return r.blackPieces, r.whitePieces
}
