package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	baseW         int32
	baseH         int32
	aspect        float32
	final_texture rl.RenderTexture2D
	Arrows        []Arrow
	textureNames  [4]string
	textures      [4]rl.Texture2D
	ProgressBars  []ProgressBar
}

func (g *Game) initialize() {
	rl.InitWindow(g.baseW, g.baseH, "Archery Game")
	rl.SetWindowState(rl.FlagWindowResizable)
	g.final_texture = rl.LoadRenderTexture(g.baseW*2, g.baseH*2)
	g.textureNames = [4]string{
		"textures/arrow.png",
	}
	for i, texturesName := range g.textureNames {
		g.textures[i] = rl.LoadTexture(texturesName)
	}
	fmt.Println("Loaded Textures")

	g.Arrows = make([]Arrow, 0, 5)
	g.ProgressBars = make([]ProgressBar, 0, 2)
	g.ProgressBars = append(g.ProgressBars, NewProgressBar(20, 400, rl.NewVector2(100, 100), 1, rl.Red, rl.Black, 5, true))
	g.ProgressBars = append(g.ProgressBars, NewProgressBar(400, 20, rl.NewVector2(100, 100), 1, rl.Red, rl.Black, 5, false))

	// var Timebar ProgressBar = NewProgressBar(20, 400, rl.NewVector2(100, 100), 1, rl.Red, rl.Black, 5, true)

}
func (g *Game) update(final_texture *rl.RenderTexture2D) {
	dt := rl.GetFrameTime()
	if rl.IsKeyPressed(rl.KeySpace) && len(g.Arrows) < 5 {
		NewArrow(rl.NewVector2(300, 600), 150, -45, 0.3, 16) //New Arrow Test
	}

	for i := range g.Arrows {
		g.Arrows[i].update(dt)
	}
	g.ProgressBars[0].progress += rl.GetMouseDelta().X * 0.005
	g.ProgressBars[0].pos.Y += rl.GetMouseDelta().Y * 0.6
	g.ProgressBars[0].update()

}
func (g *Game) StabilizeWindowRatio() {
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
	g.draw(&g.final_texture)
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.final_texture.Texture.Width), Height: -float32(g.final_texture.Texture.Height)}
	dest := rl.Rectangle{X: 0, Y: 0, Width: float32(w), Height: float32(h)}
	rl.DrawTexturePro(g.final_texture.Texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	rl.EndDrawing()
}
func (g *Game) draw(finalTexture *rl.RenderTexture2D) {

	DrawArrows(finalTexture)
	game.ProgressBars[1].draw()

	rl.DrawFPS(finalTexture.Texture.Width/2, 100)
}

var game *Game

func main() {

	game = &Game{
		baseW:  1200,
		baseH:  675,
		aspect: 16.0 / 9.0,
	}

	game.initialize()

	defer rl.UnloadRenderTexture(game.final_texture)

	// rl.SetTargetFPS(25)
	for !rl.WindowShouldClose() {
		game.update(&game.final_texture)
		game.StabilizeWindowRatio()

	}

	rl.CloseWindow()
}
