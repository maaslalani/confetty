package physics

import (
	"math"
)

type Position Vector
type Velocity Vector
type Acceleration Vector

var Gravity = Acceleration{
	X: 0,
	Y: 9.81,
}

type Motion struct {
	pos Position
	vel Velocity
	acc Acceleration
}

type Physics struct {
	current Motion
	initial Motion
	fps     float64
}

type Vector struct {
	X float64
	Y float64
}

func New(pos, vel, acc Vector, fps float64) *Physics {
	motion := Motion{
		pos: Position(pos),
		vel: Velocity(vel),
		acc: Acceleration(acc),
	}
	return &Physics{
		initial: motion,
		current: motion,
		fps:     fps,
	}
}

func (p *Physics) Reset() {
	p.current = p.initial
}

func (p *Physics) Update() {
	p.current.pos.Y, p.current.vel.Y = p.current.pos.Y+p.current.vel.Y/p.fps, p.current.vel.Y+p.current.acc.Y/p.fps
	p.current.pos.X, p.current.vel.X = p.current.pos.X+p.current.vel.X/p.fps, p.current.vel.X+p.current.acc.X/p.fps
}

func (p Physics) PosX() int {
	return int(math.Round(p.current.pos.X))
}

func (p Physics) PosY() int {
	return int(math.Round(p.current.pos.Y))
}
