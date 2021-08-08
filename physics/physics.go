package physics

import (
	"math"
)

// Position is the location of an object on a 2-dimensional plane
type Position Point

// Velocity is the velocity vector of an object's motion
type Velocity Vector

// Acceleration is the acceleration vector of an object's motion
type Acceleration Vector

// Gravity is the acceleration of gravity
// Downward (+) by g m/s²
var Gravity = Acceleration{
	X: 0,
	Y: 9.81,
}

// Motion represents an objects motion
// it keeps track of the position, velocity, and acceleration
type Motion struct {
	pos Position
	vel Velocity
	acc Acceleration
}

// Physics tracks the current motion and initial motion of an object along with
// fps to account for the Update in frames rather than per second
type Physics struct {
	current Motion
	initial Motion
	fps     float64
}

// Vector represents a magnitude and a direction in the form of a Point
// from the origin (0, 0)
type Vector struct {
	X float64
	Y float64
}

// Point is a coordinate on a 2-dimensional plane
type Point struct {
	X float64
	Y float64
}

// Distance calculates the euclidean distance between two points
func (a Point) Distance(b Point) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

// New initialize a physics simulation with simple motion
func New(pos Point, vel, acc Vector, fps float64) *Physics {
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

// Reset resets the current motion back to the initial
func (p *Physics) Reset() {
	p.current = p.initial
}

// Update increases the position of the motion by the velocity
// and increases the velocity by the acceleration
func (p *Physics) Update() {
	p.current.pos.X += p.current.vel.X / p.fps
	p.current.pos.Y += p.current.vel.Y / p.fps

	p.current.vel.X += p.current.acc.X / p.fps
	p.current.vel.Y += p.current.acc.Y / p.fps
}

// Displacement calculates the displacement between the current position and
// its initial position
func (p Physics) Displacement() float64 {
	return Point(p.initial.pos).Distance(Point(p.current.pos))
}

// PosX returns the integer value of the current x coordinate for motion
// not to be confused with Posix :D
func (p Physics) PosX() int {
	return int(math.Round(p.current.pos.X))
}

// PosY returns the integer value of the current y coordinate for motion
func (p Physics) PosY() int {
	return int(math.Round(p.current.pos.Y))
}
