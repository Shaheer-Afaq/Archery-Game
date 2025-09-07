package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Arrows []Arrow

type Game struct {
	baseW         int32
	baseH         int32
	aspect        float32
	final_texture rl.RenderTexture2D
}

func (g *Game) draw() {
	w := rl.GetScreenWidth()
	h := rl.GetScreenHeight()
	currentAspect := float32(w) / float32(h)

	if currentAspect > g.aspect {
		w = int(float32(h) * g.aspect)
		rl.SetWindowSize(w, h)
	} else if currentAspect < g.aspect {
		h = int(float32(w) / g.aspect)
		rl.SetWindowSize(w, h)
	}

	rl.BeginTextureMode(g.final_texture)
	rl.ClearBackground(rl.RayWhite)
	draw()
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.final_texture.Texture.Width), Height: -float32(g.final_texture.Texture.Height)}
	dest := rl.Rectangle{X: 0, Y: 0, Width: float32(w), Height: float32(h)}
	rl.DrawTexturePro(g.final_texture.Texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
	fmt.Println(w, h)

	rl.EndDrawing()
}

func initialize() {
	

}
func update() {

}
func draw() {
	
}

func main() {
	game := &Game{
		baseW:  1920,
		baseH:  1080,
		aspect: 16.0/9.0,
	}

	rl.InitWindow(game.baseW, game.baseH, "Archery Game")
	rl.SetWindowState(rl.FlagWindowResizable)

	game.final_texture = rl.LoadRenderTexture(game.baseW, game.baseH)
	defer rl.UnloadRenderTexture(game.final_texture)

	initialize()
	for !rl.WindowShouldClose() {
		update()
		game.draw()
	}

	rl.CloseWindow()
}
