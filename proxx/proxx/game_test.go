package proxx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_SimpleWin(t *testing.T) {
	board := parseBoard(`
		. . . . . . .
		. . . . . . .
		. . * * * . .
		. . * * * . .
		. . * * * . .
		. . . . . . .
		. . . . . . .
	`)
	game := NewGame(board)
	game.Click(0, 0)
	expected := `
        . . . . . . . 
        . 1 2 3 2 1 . 
        . 2 * * * 2 . 
        . 3 * * * 3 . 
        . 2 * * * 2 . 
        . 1 2 3 2 1 . 
        . . . . . . . 
	`
	result := printBoard(board)
	if !compareBoardsStrings(result, expected) {
		t.Errorf("invalid board:\n%s\nexpected:\n%s\n", result, expected)
	}
	assert.True(t, game.IsEnded(), "game should be ended")
	assert.True(t, game.IsWon(), "game should be won")
}

func TestGame_SimpleLose(t *testing.T) {
	board := parseBoard(`
		. . . . . . .
		. . . . . . .
		. . * * * . .
		. . * * * . .
		. . * * * . .
		. . . . . . .
		. . . . . . .
	`)
	game := NewGame(board)
	game.Click(3, 3)
	expected := `
        ? ? ? ? ? ? ? 
        ? ? ? ? ? ? ? 
        ? ? * * * ? ? 
        ? ? * * * ? ? 
        ? ? * * * ? ? 
        ? ? ? ? ? ? ? 
        ? ? ? ? ? ? ? 
	`
	result := printBoard(board)
	if !compareBoardsStrings(result, expected) {
		t.Errorf("invalid board:\n%s\nexpected:\n%s\n", result, expected)
	}
	assert.True(t, game.IsEnded(), "game should be ended")
	assert.False(t, game.IsWon(), "game should be lost")
}

func TestGame_Playing(t *testing.T) {
	board := parseBoard(`
		* * . . . * *
		* * . . . * *
		. . * * * . .
		. . * . * . .
		. . * * * . .
		* * . . . * *
		* * . . . * *
	`)
	game := NewGame(board)
	game.Click(3, 0)
	assert.False(t, game.IsEnded(), "game should not be ended")
	game.Click(0, 3)
	assert.False(t, game.IsEnded(), "game should not be ended")
	game.Click(6, 3)
	assert.False(t, game.IsEnded(), "game should not be ended")
	game.Click(3, 6)
	assert.False(t, game.IsEnded(), "game should not be ended")
	game.Click(3, 3)
	expected := `
        * * 2 . 2 * * 
        * * 4 3 4 * * 
        2 4 * * * 4 2 
        . 3 * 8 * 3 . 
        2 4 * * * 4 2 
        * * 4 3 4 * * 
        * * 2 . 2 * * 
	`
	result := printBoard(board)
	if !compareBoardsStrings(result, expected) {
		t.Errorf("invalid board:\n%s\nexpected:\n%s\n", result, expected)
	}
	assert.True(t, game.IsEnded(), "game should be ended")
	assert.True(t, game.IsWon(), "game should be won")
}
