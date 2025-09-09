package main

import (
	// "fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var game *Game = &Game{
	baseW:  1200,
	baseH:  675,
	aspect: 16.0 / 9.0,
}

func main() {

	game.initialize()

	defer rl.UnloadRenderTexture(game.final_texture)

	// rl.SetTargetFPS(25)
	for !rl.WindowShouldClose() {
		game.update()
		game.stabilizeRatioAndDraw()

	}

	rl.CloseWindow()

}
