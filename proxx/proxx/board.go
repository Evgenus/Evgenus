package proxx

import "math/rand"

type Board struct {
	width  int
	height int
	cells  [][]ICell
}

func NewBoard(size int) IBoard {
	cells := make([][]ICell, size)
	for y := 0; y < size; y++ {
		cells[y] = make([]ICell, size)
		for x := 0; x < size; x++ {
			cells[y][x] = NewCell(x, y, false)
		}
	}
	return &Board{size, size, cells}
}

func (b *Board) Width() int {
	return b.width
}

func (b *Board) Height() int {
	return b.height
}

func (b *Board) Cell(x, y int) ICell {
	return b.cells[y][x]
}

func (b *Board) Open(x, y int) []ICell {
	cell := b.Cell(x, y)
	queue := []ICell{cell}
	opened := make([]ICell, 0)
	for len(queue) > 0 {
		nextQueue := make([]ICell, 0)
		for _, current := range queue {
			if current.IsOpen() {
				continue
			}
			current.SetOpen()
			opened = append(opened, current)
			if current.AdjacentBlackHoles() == 0 {
				for _, adj := range b.adjacentCells(current) {
					nextQueue = append(nextQueue, adj)
				}
			}
		}
		queue = nextQueue
	}
	return opened
}

func (b *Board) Fill(count int) {
	empty := b.emptyCells()
	rand.Shuffle(len(empty), func(i, j int) { empty[i], empty[j] = empty[j], empty[i] })
	chosen := empty[:count]
	for _, cell := range chosen {
		b.AddBlackHole(cell.X(), cell.Y())
	}
}

func (b *Board) AddBlackHole(x, y int) {
	cell := b.Cell(x, y)
	if cell.IsBlackHole() {
		return
	}
	cell.SetBlackHole()
	for _, adj := range b.adjacentCells(cell) {
		adj.AddAdjacentBlackHole()
	}
}

func (b *Board) BlackHoles() []ICell {
	res := make([]ICell, 0)
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			cell := b.cells[y][x]
			if cell.IsBlackHole() {
				res = append(res, cell)
			}
		}
	}
	return res
}

func (b *Board) emptyCells() []ICell {
	res := make([]ICell, 0)
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			cell := b.cells[y][x]
			if !cell.IsBlackHole() {
				res = append(res, cell)
			}
		}
	}
	return res
}

var neighbors = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (b *Board) adjacentCells(cell ICell) []ICell {
	res := make([]ICell, 0)
	for _, n := range neighbors {
		x := cell.X() + n[0]
		y := cell.Y() + n[1]
		if x < 0 || x >= b.width {
			continue
		}
		if y < 0 || y >= b.height {
			continue
		}
		res = append(res, b.cells[y][x])
	}
	return res
}
