package proxx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_FillDistribution(t *testing.T) {
	distribution := make(map[string]int)
	for i := 0; i < 1000; i++ {
		board := NewBoard(3)
		board.Fill(5)
		assert.Len(t, board.BlackHoles(), 5, "invalid number of black holes")
		for _, cell := range board.BlackHoles() {
			distribution[cell.Id()]++
		}
	}
	assert.Len(t, distribution, 9, "invalid distribution")
	for _, count := range distribution {
		assert.InDelta(t, 550, count, 50, "invalid distribution")
	}
}

func TestBoard_AdjacentBlackHoles(t *testing.T) {
	board := parseBoard(`
		* . . . . . *
		. * . . . * .
		. . * * * . .
		. . * . * . .
		. . * * * . .
		. * . . . * .
		* . . . . . *
	`)
	revealAll(board)
	expected := `
		* 2 1 . 1 2 * 
		2 * 3 3 3 * 2 
		1 3 * * * 3 1 
		. 3 * 8 * 3 . 
		1 3 * * * 3 1 
		2 * 3 3 3 * 2 
		* 2 1 . 1 2 *
	`
	result := printBoard(board)
	if !compareBoardsStrings(result, expected) {
		t.Errorf("invalid board:\n%s\nexpected:\n%s\n", result, expected)
	}
}

func TestBoard_OpenEmpty(t *testing.T) {
	board := parseBoard(`
		. . . * . . . .
		. . . . * . * *
		. . . . . . . .
		* . . . . . . .
		. * . . . . . .
		. . * * * * . .
		. . . . . . . .
		* . . . . . . .
	`)
	opened := board.Open(0, 0)
	openedIds := make([]string, 0)
	for _, cell := range opened {
		openedIds = append(openedIds, cell.Id())
	}
	expectedToBeOpened := parseCoordinates(`
		@ @ @ * . . . .
		@ @ @ @ * . * *
		@ @ @ @ @ @ @ @
		* @ @ @ @ @ @ @
		. * @ @ @ @ @ @
		. . * * * * @ @
		. @ @ @ @ @ @ @
		* @ @ @ @ @ @ @
	`, "@")
	expectedIds := make([]string, 0)
	for _, coord := range expectedToBeOpened {
		id := cellId(coord[0], coord[1])
		expectedIds = append(expectedIds, id)
	}
	assert.ElementsMatch(t, expectedIds, openedIds, "invalid opened cells")

	expected := `
        . . 1 ? ? ? ? ? 
        . . 1 2 ? ? ? ? 
        1 1 . 1 1 2 2 2 
        ? 2 1 . . . . . 
        ? ? 3 3 3 2 1 . 
        ? ? ? ? ? ? 1 . 
        ? 2 2 3 3 2 1 . 
        ? 1 . . . . . . 
	`
	result := printBoard(board)
	if !compareBoardsStrings(result, expected) {
		t.Errorf("invalid board:\n%s\nexpected:\n%s\n", result, expected)
	}
}
