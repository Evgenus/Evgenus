package main

import (
	"proxx/proxx"
	"syscall/js"
)

var game proxx.IGame

func main() {
	done := make(chan struct{})
	js.Global().Set("createGame", js.FuncOf(createGame))
	js.Global().Set("gameStatus", js.FuncOf(gameStatus))
	js.Global().Set("clickCell", js.FuncOf(clickCell))
	<-done
}

func createGame(this js.Value, args []js.Value) any {
	size := args[0].Int()
	blackHoles := args[1].Int()

	board := proxx.NewBoard(size)
	board.Fill(blackHoles)
	game = proxx.NewGame(board)

	return nil
}

func gameStatus(this js.Value, args []js.Value) any {
	status := map[string]interface{}{
		"isWon":   game.IsWon(),
		"isEnded": game.IsEnded(),
	}
	return js.ValueOf(status)
}

func clickCell(this js.Value, args []js.Value) any {
	x := args[0].Int()
	y := args[1].Int()
	cells := game.Click(x, y)
	cellsData := make([]any, 0)
	for _, cell := range cells {
		cellData := map[string]any{
			"x":                  cell.X(),
			"y":                  cell.Y(),
			"isBlackHole":        cell.IsBlackHole(),
			"isOpen":             cell.IsOpen(),
			"adjacentBlackHoles": cell.AdjacentBlackHoles(),
		}
		cellsData = append(cellsData, cellData)
	}
	return js.ValueOf(cellsData)
}
