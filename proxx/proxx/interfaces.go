package proxx

type ICell interface {
	// Id returns the cell's unique id.
	Id() string
	// X returns the cell's x coordinate.
	X() int
	// Y returns the cell's y coordinate.
	Y() int
	// IsBlackHole returns true if the cell is a black hole.
	IsBlackHole() bool
	// SetBlackHole sets the cell to be a black hole.
	SetBlackHole()
	// IsOpen returns true if the cell is open.
	IsOpen() bool
	// SetOpen sets the cell to be open.
	SetOpen()
	// AdjacentBlackHoles returns the number of adjacent black holes.
	AdjacentBlackHoles() int
	// AddAdjacentBlackHole increases the number of adjacent black holes by one.
	AddAdjacentBlackHole()
}

type IBoard interface {
	// Width returns the width of the board.
	Width() int
	// Height returns the height of the board.
	Height() int
	// Cell returns the cell at the given position.
	Cell(x, y int) ICell
	// Open reveals the cell(s) at the given position. Returns list of cells that were opened.
	Open(x, y int) []ICell
	// Fill fills the board with black holes.
	Fill(count int)
	// AddBlackHole adds a black hole to the board.
	AddBlackHole(x, y int)
	// BlackHoles returns the list of cells that are black holes.
	BlackHoles() []ICell
}

type IGame interface {
	// Click reveals the cell(s) at the given position. Returns list of cells that were opened.
	// Also updates game status.
	Click(x, y int) []ICell
	// IsEnded returns true if the game is ended.
	IsEnded() bool
	// IsWon returns true if the game is won.
	IsWon() bool
}
