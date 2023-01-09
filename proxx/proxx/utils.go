package proxx

import (
	"strconv"
	"strings"
)

// parseBoard parses a board from a multiline-string where `*` indicates black holes.
func parseBoard(s string) IBoard {
	lines := splitBoardLines(s)
	size := len(lines)
	board := NewBoard(size)
	for y, line := range lines {
		line = strings.TrimSpace(line)
		for x, c := range strings.Split(line, " ") {
			if c == "*" {
				board.AddBlackHole(x, y)
			}
		}
	}
	return board
}

// parseCoordinates parses board string and returns a list of coordinates for found marker
func parseCoordinates(s, marker string) [][2]int {
	var coordinates [][2]int
	lines := splitBoardLines(s)
	for y, line := range lines {
		line = strings.TrimSpace(line)
		for x, c := range strings.Split(line, " ") {
			if c == marker {
				coordinates = append(coordinates, [2]int{x, y})
			}
		}
	}
	return coordinates
}

// printBoard prints a board to a multiline-string where
// 		`*` indicates black holes;
//		`?` indicates closed cells;
//		`.` indicates open cells with no adjacent black holes;
//		`1`-`8` indicates open cells with 1-8 adjacent black holes.
func printBoard(board IBoard) string {
	var s strings.Builder
	for y := 0; y < board.Height(); y++ {
		for x := 0; x < board.Width(); x++ {
			cell := board.Cell(x, y)
			if cell.IsOpen() {
				if cell.IsBlackHole() {
					s.WriteString("*")
				} else if cell.AdjacentBlackHoles() == 0 {
					s.WriteString(".")
				} else {
					s.WriteString(strconv.Itoa(cell.AdjacentBlackHoles()))
				}
			} else {
				s.WriteString("?")
			}
			s.WriteString(" ")
		}
		s.WriteString("\n")
	}
	return s.String()
}

func sanitize(s string) string {
	lines := splitBoardLines(s)
	return strings.Join(lines, "\n")
}

func splitBoardLines(s string) []string {
	lines := make([]string, 0)
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func compareBoardsStrings(a, b string) bool {
	a = sanitize(a)
	b = sanitize(b)
	return a == b
}

// revealAll reveals all cells on the board.
func revealAll(board IBoard) {
	for y := 0; y < board.Height(); y++ {
		for x := 0; x < board.Width(); x++ {
			board.Open(x, y)
		}
	}
}
