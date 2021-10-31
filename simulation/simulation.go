package simulation

import (
	"fmt"
	"strings"
	"time"
)

type System struct {
	Frame     Frame
	Particles []Particle
}

type Particle struct {
	Char          string
	TailChar      string
	Physics       *Projectile
	Hidden        bool
	Shooting      bool
	ExplosionCall func(x, y float64, width, height int) []Particle
}

type Frame struct {
	Width  int
	Height int
}

func FPS(n int) float64 {
	return (time.Second / time.Duration(n)).Seconds()
}

func RemoveParticleFromArray(s []Particle, i int) []Particle {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *System) Update() {
	for i := len(s.Particles) - 1; i >= 0; i-- {
		p := s.Particles[i]
		pos := p.Physics.Position

		// if the shooting particle is slow enough then hide it and call the explosion function
		if !p.Hidden && p.Shooting && p.Physics.Velocity.Y > -3 {
			p.Hidden = true
			if p.ExplosionCall != nil {
				s.Particles = append(s.Particles, p.ExplosionCall(p.Physics.Position.X, p.Physics.Position.Y, s.Frame.Width, s.Frame.Height)...)
			}
		}

		// remove particles that are hidden or out of the side/bottom of the frame
		if p.Hidden || pos.X > float64(s.Frame.Width) || pos.X < 0 || pos.Y > float64(s.Frame.Height) {
			s.Particles = RemoveParticleFromArray(s.Particles, i)
		} else {
			s.Particles[i].Physics.Update()
		}
	}
}

func (s *System) Visible(p Particle) bool {
	y := int(p.Physics.Position.Y)
	x := int(p.Physics.Position.X)
	return !p.Hidden && y >= 0 && y < s.Frame.Height-1 && x >= 0 && x < s.Frame.Width-1
}

func (s *System) Render() string {
	var out strings.Builder
	plane := make([][]string, s.Frame.Height)
	for i := range plane {
		plane[i] = make([]string, s.Frame.Width)
	}
	for _, p := range s.Particles {
		if s.Visible(p) {
			plane[int(p.Physics.Position.Y)][int(p.Physics.Position.X)] = p.Char
			if p.Shooting {
				l := -int(p.Physics.Velocity.Y)
				for i := 1; i < l; i++ {
					y := int(p.Physics.Position.Y) + i
					if y > 0 && y < s.Frame.Height-1 {
						plane[y][int(p.Physics.Position.X)] = p.TailChar
					}
				}
			}
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
