package main

import (
	// "fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var dt float32
var starttime time.Time = time.Now()
var time_elapsed float32

type Game struct {
	baseW         int32
	baseH         int32
	width         int32
	height        int32
	aspect        float32
	final_texture rl.RenderTexture2D
	Arrows        []Arrow
	textureNames  [4]string
	textures      [4]rl.Texture2D
	ProgressBars  []ProgressBar
	scene         string
	Scenes        map[string]*Scene
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

	g.Arrows = make([]Arrow, 0, 5)
	g.ProgressBars = make([]ProgressBar, 0, 2)
	g.Scenes = make(map[string]*Scene, 3)

	NewScene("start", UpdateStartScene, DrawStartScene)
	NewScene("main", UpdateMainScene, DrawMainScene)
	NewScene("end", UpdateEndScene, DrawEndScene)

	g.scene = "start"
}
func (g *Game) stabilizeRatioAndDraw() {
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
	g.draw()
	rl.EndTextureMode()

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	src := rl.Rectangle{X: 0, Y: 0, Width: float32(g.final_texture.Texture.Width), Height: -float32(g.final_texture.Texture.Height)}
	dest := rl.Rectangle{X: 0, Y: 0, Width: float32(w), Height: float32(h)}
	rl.DrawTexturePro(g.final_texture.Texture, src, dest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	rl.EndDrawing()
}
func (g *Game) update() {
	time_elapsed = float32(time.Since(starttime).Milliseconds())
	dt = rl.GetFrameTime()
	g.width = g.final_texture.Texture.Width
	g.height = g.final_texture.Texture.Height

	g.Scenes[g.scene].Update() //Update current scene
	UpdateArrows()

	if rl.IsKeyPressed(rl.KeyA) {
		g.scene = "start"
	}else if rl.IsKeyPressed(rl.KeyS) {
		g.scene = "main"
	}else if rl.IsKeyPressed(rl.KeyD) {
		g.scene = "end"
	}

	// if rl.IsKeyPressed(rl.KeySpace) && len(g.Arrows) < 5 {
	// 	NewArrow(rl.NewVector2(300, 600), 150, -45, 0.3, 16) //New Arrow Test
	// }

	// NewArrow(rl.NewVector2(300, 600), 150, -45, 0.3, 16) //New Arrow Test
	// for _, scene := range g.Scenes {
	// 	if scene.name == g.scene {
	// 		scene.Update()
	// 	}
	// }

}
func (g *Game) draw() {
	g.Scenes[g.scene].Draw() //Draw current scene
	// DrawStartScene()
	DrawArrows()
	rl.DrawFPS(g.final_texture.Texture.Width/2, 100)
}