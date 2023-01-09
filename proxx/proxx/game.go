package proxx

type Game struct {
	board  IBoard
	ended  bool
	won    bool
	opened int
}

func NewGame(board IBoard) IGame {
	return &Game{board, false, false, 0}
}

func (g *Game) Click(x, y int) []ICell {
	opened := g.board.Open(x, y)
	g.opened += len(opened)
	for _, cell := range opened {
		if cell.IsBlackHole() {
			g.ended = true
			g.won = false
		}
	}
	emptyCellsCount := g.board.Width()*g.board.Height() - len(g.board.BlackHoles())
	if !g.ended && g.opened == emptyCellsCount {
		g.ended = true
		g.won = true
	}

	if g.ended {
		for _, cell := range g.board.BlackHoles() {
			g.board.Open(cell.X(), cell.Y())
			opened = append(opened, cell)
		}
	}

	return opened
}

func (g *Game) IsEnded() bool {
	return g.ended
}

func (g *Game) IsWon() bool {
	return g.won
}
