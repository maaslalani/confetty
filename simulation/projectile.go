package simulation

type Projectile struct {
	Position     Point
	Velocity     Vector
	Acceleration Vector
	DeltaTime    float64
}

type Point struct {
	X, Y, Z float64
}

type Vector struct {
	X, Y, Z float64
}

var Gravity = Vector{0, -9.81, 0}
var TerminalGravity = Vector{0, 9.81, 0}

func NewProjectile(del float64, pos Point, vel, acc Vector) *Projectile {
	return &Projectile{pos, vel, acc, del}
}

func (p *Projectile) Update() {
	p.Position.X += (p.Velocity.X * p.DeltaTime)
	p.Position.Y += (p.Velocity.Y * p.DeltaTime)
	p.Position.Z += (p.Velocity.Z * p.DeltaTime)

	p.Velocity.X += (p.Acceleration.X * p.DeltaTime)
	p.Velocity.Y += (p.Acceleration.Y * p.DeltaTime)
	p.Velocity.Z += (p.Acceleration.Z * p.DeltaTime)
}
