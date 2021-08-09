package simulation

import (
	"fmt"
	"strings"

	"github.com/maaslalani/confetty/physics"
)

type System struct {
	Frame     Frame
	Particles []Particle
}

type Particle struct {
	Char    string
	Physics *physics.Physics
	Hidden  bool
}

type Frame struct {
	Width  int
	Height int
}

func (s *System) Update() {
	for _, p := range s.Particles {
		if p.Hidden {
			continue
		}

		if !s.Visible(p) {
			p.Hidden = true
			continue
		}

		p.Physics.Update()
	}
}

func (s *System) Visible(p Particle) bool {
	y := p.Physics.PosY()
	x := p.Physics.PosX()
	return y >= 0 && y < s.Frame.Height-1 && x >= 0 && x < s.Frame.Width-1
}

func (s *System) Render() string {
	var out strings.Builder
	plane := make([][]string, s.Frame.Height)
	for i := range plane {
		plane[i] = make([]string, s.Frame.Width)
	}
	for _, p := range s.Particles {
		if s.Visible(p) {
			plane[p.Physics.PosY()][p.Physics.PosX()] = p.Char
		}
	}
	for i := range plane {
		for _, col := range plane[i] {
			if col == "" {
				fmt.Fprint(&out, " ")
			} else {
				fmt.Fprint(&out, col)
			}
		}
		fmt.Fprint(&out, "\n")
	}
	return out.String()
}
