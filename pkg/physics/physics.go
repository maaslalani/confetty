package physics

import "math"

type Position Vector
type Velocity Vector
type Acceleration Vector

var Gravity = Acceleration{
	X: 0,
	Y: 9.81,
}

type Physics struct {
	pos Position
	vel Velocity
	acc Acceleration
	fps float64
}

type Vector struct {
	X float64
	Y float64
}

func New(pos, vel, acc Vector, fps float64) *Physics {
	return &Physics{
		pos: Position(pos),
		vel: Velocity(vel),
		acc: Acceleration(acc),
		fps: fps,
	}
}

func (p *Physics) Update() {
	p.pos.Y, p.vel.Y = p.pos.Y+p.vel.Y/p.fps, p.vel.Y+p.acc.Y/p.fps
	p.pos.X, p.vel.X = p.pos.X+p.vel.X/p.fps, p.vel.X+p.acc.X/p.fps
}

func (p Physics) PosX() int {
	return int(math.Round(p.pos.X))
}

func (p Physics) PosY() int {
	return int(math.Round(p.pos.Y))
}
