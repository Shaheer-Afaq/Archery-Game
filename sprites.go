package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Arrow struct {
	pos      rl.Vector2
	power    float32
	velocity *rl.Vector2
	angle    float32
	gravity  float32
	friction float32
}

func (a *Arrow) update(dt float32) {

}
func (a *Arrow) draw(dt float32, texture *rl.Texture2D) {
	rl.DrawTextureEx(*texture, a.pos, a.angle, 1, rl.White)
}

func NewArrow(pos rl.Vector2, power float32, angle float32) Arrow {
	var velocity rl.Vector2
	velocity.X = power * float32(math.Cos(float64(angle)))
	velocity.Y = power * float32(math.Sin(float64(angle)))

	Arrows.append(Arrow{
		pos:      pos,
		power:    power,
		velocity: &velocity,
		angle:    angle,
		gravity:  -9.81,
		friction: 0.9,
	})
}
