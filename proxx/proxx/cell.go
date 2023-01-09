package proxx

import "fmt"

type Cell struct {
	x, y               int
	blackHole          bool
	open               bool
	adjacentBlackHoles int
}

func NewCell(x, y int, blackHole bool) ICell {
	return &Cell{x, y, blackHole, false, 0}
}

func (c *Cell) Id() string {
	return cellId(c.x, c.y)
}

func cellId(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

func (c *Cell) X() int {
	return c.x
}

func (c *Cell) Y() int {
	return c.y
}

func (c *Cell) IsBlackHole() bool {
	return c.blackHole
}

func (c *Cell) SetBlackHole() {
	c.blackHole = true
}

func (c *Cell) IsOpen() bool {
	return c.open
}

func (c *Cell) SetOpen() {
	c.open = true
}

func (c *Cell) AdjacentBlackHoles() int {
	return c.adjacentBlackHoles
}

func (c *Cell) AddAdjacentBlackHole() {
	c.adjacentBlackHoles++
}
