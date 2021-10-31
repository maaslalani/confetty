package simulation

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/harmonica"
)

type System struct {
	Frame     Frame
	Particles []Particle
}

type Particle struct {
	Char    string
	Physics *harmonica.Projectile
	Hidden  bool
}

type Frame struct {
	Width  int
	Height int
}

func RemoveParticleFromArray(s []Particle, i int) []Particle {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *System) Update() {
	for _, p := range s.Particles {
		p.Physics.Update()
	}
}

func (s *System) Visible(p Particle) bool {
	y := int(p.Physics.Position().Y)
	x := int(p.Physics.Position().X)
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
			plane[int(p.Physics.Position().Y)][int(p.Physics.Position().X)] = p.Char
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
