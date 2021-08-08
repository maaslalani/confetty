package physics_test

import (
	"testing"
	"time"

	. "github.com/maaslalani/confetty/physics"
)

const fps = 60

func simulate(p *Physics, d time.Duration) {
	frames := int(d.Seconds() * fps)
	for i := 0; i < frames; i++ {
		p.Update()
	}
}

func TestNew(t *testing.T) {
	x := 8
	y := 20

	physics := New(Point{float64(x), float64(y)}, Vector{1, 1}, Vector(Gravity), 60)
	if x != physics.PosX() {
		t.Fatal("x coordinate unexpected")
	}

	if y != physics.PosY() {
		t.Fatal("y coordinate unexpected")
	}
}

func TestUpdate(t *testing.T) {
	physics := New(Point{0, 0}, Vector{5, 5}, Vector(Gravity), float64(fps))

	// coordinates is the location at which the object should be
	// after i+1 seconds of simulation.
	coordinates := []Point{
		{5, 10},
		{10, 29},
		{15, 59},
		{20, 98},
		{25, 147},
		{30, 206},
		{35, 275},
	}

	for _, c := range coordinates {
		simulate(physics, time.Second)

		x := physics.PosX()
		y := physics.PosY()

		if x != int(c.X) {
			t.Logf("Want: %d, Got: %d", x, int(c.X))
			t.Fatal("x coordinate unexpected")
		}
		if y != int(c.Y) {
			t.Logf("Want: %d, Got: %d", y, int(c.Y))
			t.Fatal("y coordinate unexpected")
		}
	}
}

func TestDisplacement(t *testing.T) {
	tt := []struct {
		x   int
		y   int
		vel Vector
		d   float64
	}{
		{x: 5, y: 5, vel: Vector{5, 10}, d: 11.180339887498933},
		{x: 0, y: 0, vel: Vector{1, 1}, d: 1.414213562373097},
	}

	for _, tc := range tt {
		physics := New(Point{float64(tc.x), float64(tc.y)}, tc.vel, Vector{}, float64(fps))

		simulate(physics, time.Second)

		if physics.Displacement() != tc.d {
			t.Log(physics.Displacement())
			t.Fatal("expected displacement to be 15")
		}
	}
}

func TestReset(t *testing.T) {
	x := 5
	y := 10
	vel := Vector{20, 20}

	physics := New(Point{float64(x), float64(y)}, vel, Vector(Gravity), float64(fps))

	simulate(physics, time.Second)

	// store the current position after 1 second of simulation we will simulate 1
	// second again and ensure that the object reaches the same position, we
	// don't care what the position is but if the object reaches the same
	// position again then we ensure that the velocity and acceleration was reset
	cx := physics.PosX()
	cy := physics.PosY()

	physics.Reset()

	if physics.PosX() != x {
		t.Fatal("expected x to be reset")
	}

	if physics.PosY() != y {
		t.Fatal("expected y to be reset")
	}

	// here we are simply checking that the object reaches the same position
	// after being reset in the same amount of time. This ensures velocity and
	// acceleration were reset without checking them explicitly
	simulate(physics, time.Second)

	if physics.PosX() != cx {
		t.Fatal("expected simulation to be repeatable")
	}

	if physics.PosY() != cy {
		t.Fatal("expected simulation to be repeatable")
	}
}
