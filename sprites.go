package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Arrow struct {
	pos           rl.Vector2
	power         float32
	velocity      rl.Vector2
	angle         float32
	scale         float32
	gravity       float32
	friction      float32
	frictionTimer Timer
}

func (a *Arrow) update(dt float32) {
	if a.frictionTimer.Update(dt) {
		a.velocity.X *= a.friction
	}
	a.velocity.Y += a.gravity * dt * 10

	if math.Abs(float64(a.velocity.X)) < 1 {
		a.velocity.X = 0
	}
	if a.velocity.Y > 350 {
		a.velocity.Y = 350
	}
	a.pos.X += a.velocity.X * dt * 10
	a.pos.Y += a.velocity.Y * dt * 10
	a.angle = float32(math.Atan2(float64(a.velocity.Y), float64(a.velocity.X))) * rl.Rad2deg

}
func (a *Arrow) draw(texture *rl.Texture2D) {
	rl.DrawTexturePro(
		*texture,
		rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height)),
		rl.NewRectangle(a.pos.X, a.pos.Y, float32(texture.Width)*a.scale, float32(texture.Height)*a.scale),
		rl.NewVector2(float32(texture.Width)*a.scale/2, float32(texture.Height)*a.scale/2),
		a.angle,
		rl.White,
	)
}

func (a *Arrow) get_hitbox(texture *rl.Texture2D) rl.Vector2 {
	cosA := float32(math.Cos(float64(a.angle) * rl.Deg2rad))
	sinA := float32(math.Sin(float64(a.angle) * rl.Deg2rad))
	offset := rl.NewVector2(float32(texture.Width)*a.scale/2, 0.0)

	return rl.NewVector2(
		a.pos.X+offset.X*cosA-offset.Y*sinA,
		a.pos.Y+offset.X*sinA+offset.Y*cosA,
	)
}

func NewArrow(pos rl.Vector2, power float32, angle float32, scale float32, gravity float32) {
	var velocity rl.Vector2
	velocity.X = power * float32(math.Cos(rl.Deg2rad*float64(angle)))
	velocity.Y = power * float32(math.Sin(rl.Deg2rad*float64(angle)))

	Arrows = append(Arrows, Arrow{
		pos:           pos,
		power:         power,
		velocity:      velocity,
		angle:         angle,
		scale:         scale,
		gravity:       gravity,
		friction:      0.98,
		frictionTimer: NewTimer(1.0 / 30.0),
	})

}

func DrawArrows(finalTexture *rl.RenderTexture2D) {
	for i := 0; i < len(Arrows); {
		if Arrows[i].pos.X > float32(finalTexture.Texture.Width) || Arrows[i].pos.Y > float32(finalTexture.Texture.Height) {
			Arrows[i] = Arrows[len(Arrows)-1]
			Arrows = Arrows[:len(Arrows)-1]

		} else {
			Arrows[i].draw(&textures[0])
			i++
		}
	}
}
