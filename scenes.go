package main

// "math"

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScene() {

}
func UpdateScene() {

}

type Scene struct {
	name        string
	update_func func()
	draw_func   func()
}

func (s *Scene) Update() {
	s.update_func()
}

func (s *Scene) Draw() {
	s.draw_func()
}

func NewScene(name string, update_func func(), draw_func func()) {

	game.Scenes[name] = &Scene{
		name:        name,
		update_func: update_func,
		draw_func:   draw_func,
	}
}

// var UpdateStartScene = func() {

// }

// var DrawStartScene = func() {
// 	rl.DrawCircle(400, 400, 40, rl.Red)
// }

func UpdateStartScene() {

}

func DrawStartScene() {
	rl.DrawCircle(400, 400, 40, rl.Red)
}
func UpdateMainScene() {

}

func DrawMainScene() {
	rl.DrawCircle(400, 400, 40, rl.Blue)
}
func UpdateEndScene() {

}

func DrawEndScene() {
	rl.DrawCircle(400, 400, 40, rl.Green)
}
