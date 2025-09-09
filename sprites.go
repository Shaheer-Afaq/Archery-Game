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

func (a *Arrow) update() {
	if a.frictionTimer.Update() {
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

func (a *Arrow) GetHitbox(texture *rl.Texture2D) rl.Vector2 {
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

	game.Arrows = append(game.Arrows, Arrow{
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
func DrawArrows() {
	for i := 0; i < len(game.Arrows); {
		if game.Arrows[i].pos.X > float32(game.final_texture.Texture.Width) || game.Arrows[i].pos.Y > float32(game.final_texture.Texture.Height) {
			game.Arrows[i] = game.Arrows[len(game.Arrows)-1]
			game.Arrows = game.Arrows[:len(game.Arrows)-1]

		} else {
			game.Arrows[i].draw(&game.textures[0])
			i++
		}
	}
}
func UpdateArrows() {
	for i := range game.Arrows {
		game.Arrows[i].update()
	}
}

type ProgressBar struct {
	width        float32
	height       float32
	pos          rl.Vector2
	progress     float32 //a value between 0 and 1
	color        rl.Color
	bgcolor      rl.Color
	visible      bool
	border_width float32
	vertical     bool
}

func (pb *ProgressBar) update() {
	pb.progress = max(0.0, min(pb.progress, 1.0))
}

func (pb *ProgressBar) draw() {
	if pb.visible {
		rl.DrawRectangle(int32(pb.pos.X), int32(pb.pos.Y), int32(pb.width), int32(pb.height), pb.bgcolor)

		if pb.vertical {
			bar_height := (pb.height - pb.border_width*2) * pb.progress
			rl.DrawRectangleRec(
				rl.NewRectangle(
					pb.pos.X+pb.border_width,
					pb.pos.Y+pb.height-pb.border_width-bar_height,
					pb.width-pb.border_width*2,
					bar_height,
				),
				pb.color,
			)
		} else {
			rl.DrawRectangle(
				int32(pb.pos.X+pb.border_width),
				int32(pb.pos.Y+pb.border_width),
				int32(float32(pb.width-pb.border_width*2)*pb.progress),
				int32(pb.height-pb.border_width*2), pb.color)

		}
	}
}

func NewProgressBar(width float32, height float32, pos rl.Vector2, progress float32, color rl.Color, bgcolor rl.Color, border_width float32, vertical bool) ProgressBar {
	pb := ProgressBar{
		width: width, height: height, pos: pos, progress: progress, color: color, bgcolor: bgcolor,
		visible: true, border_width: border_width, vertical: vertical,
	}
	game.ProgressBars = append(game.ProgressBars, pb)
	return pb
}
