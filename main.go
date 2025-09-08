package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var Arrows = make([]Arrow, 0, 20)
var textureNames = [4]string{
	"textures/arrow.png",
}
var textures [4]rl.Texture2D

// var attack_cooldown Timer = NewTimer(0.5)

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
	draw(&g.final_texture)
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.final_texture.Texture.Width), Height: -float32(g.final_texture.Texture.Height)}
	dest := rl.Rectangle{X: 0, Y: 0, Width: float32(w), Height: float32(h)}
	rl.DrawTexturePro(g.final_texture.Texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	rl.EndDrawing()
}

func initialize() {
	for i, texturesName := range textureNames {
		textures[i] = rl.LoadTexture(texturesName)
	}
	fmt.Println(textures)

}
func update(final_texture *rl.RenderTexture2D) {
	dt := rl.GetFrameTime()
	if rl.IsKeyPressed(rl.KeySpace) {
		NewArrow(rl.NewVector2(300, 600), 150, -45, 0.3, 16) //New Arrow Test
	}

	for i := range Arrows {
		Arrows[i].update(dt)
	}


}
func draw(finalTexture *rl.RenderTexture2D) {

	DrawArrows(finalTexture)

	rl.DrawFPS(finalTexture.Texture.Width/2, 100)
}

func main() {

	game := &Game{
		baseW:  1200,
		baseH:  675,
		aspect: 16.0 / 9.0,
	}

	rl.InitWindow(game.baseW, game.baseH, "Archery Game")
	rl.SetWindowState(rl.FlagWindowResizable)

	game.final_texture = rl.LoadRenderTexture(game.baseW*2, game.baseH*2)
	defer rl.UnloadRenderTexture(game.final_texture)

	// rl.SetTargetFPS(25)
	initialize()
	for !rl.WindowShouldClose() {
		update(&game.final_texture)
		game.draw()

	}

	rl.CloseWindow()
}
