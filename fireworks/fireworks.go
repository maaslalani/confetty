// TODO: Lots of code duplication between fireworks and confetti extract to a
// `particle system` package
package fireworks

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/maaslalani/confetty/array"
	"github.com/maaslalani/confetty/physics"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	framesPerSecond = 60.0
	numParticles    = 75
)

var (
	colors     = []string{"#fdff6a", "#ff718d"}
	characters = []string{"•"}
	// characters = []string{"▄", "▀"}
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/framesPerSecond, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func wait(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return nil
	}
}

// Fireworks model
type model struct {
	particles []*Particle
	viewport  viewport.Model
}

type Particle struct {
	char    string
	physics *physics.Physics
	radius  float64
}

func spawn() []*Particle {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	color := lipgloss.Color(array.Sample(colors))
	v := float64(rand.Intn(10) + 20.0)
	r := float64(rand.Intn(5) + 10)

	particles := []*Particle{}

	x := rand.Float64() * float64(width)
	y := rand.Float64() * float64(height)

	for i := 0; i < numParticles; i++ {
		p := &Particle{
			physics: physics.New(
				physics.Point{X: x, Y: y},
				physics.Vector{X: math.Cos(float64(i)) * v, Y: math.Sin(float64(i)) * v / 2},
				physics.Vector{Y: 2},
				framesPerSecond,
			),
			char: lipgloss.NewStyle().
				Foreground(color).
				Render(array.Sample(characters)),
			radius: r,
		}

		particles = append(particles, p)
	}
	return particles
}

func InitialModel() model {
	return model{particles: spawn()}
}

// Init initializes the confetti after a small delay
func (m model) Init() tea.Cmd {
	return tea.Sequentially(wait(time.Second/2), animate())
}

// Update updates the model every frame, it handles the animation loop and
// updates the particle physics every frame
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

		// frame animation
	case frameMsg:
		particlesVisible := numParticles
		for _, p := range m.particles {
			p.physics.Update()

			y := p.physics.PosY()
			x := p.physics.PosX()

			// Particle is out of view
			if y < 0 || y >= m.viewport.Height-1 || x < 0 || x >= m.viewport.Width-1 {
				particlesVisible -= 1
				continue
			}

			// Particle has reached its distance from the radius.
			// In the fireworks simulation, firework particles fade after reaching a certain point
			// just like in real life, so we don't render them if they've passed a certain distance
			if p.physics.Displacement() > p.radius {
				particlesVisible -= 1
				continue
			}
		}

		if particlesVisible <= 0 {
			m.particles = spawn()
		}

		return m, animate()

	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		return m, nil

	default:
		return m, nil
	}
}

// View displays all the particles on the screen
func (m model) View() string {
	height := m.viewport.Height
	width := m.viewport.Width
	if height <= 0 || width <= 0 {
		return ""
	}

	var out strings.Builder

	grid := make([][]string, m.viewport.Height)
	for i := range grid {
		grid[i] = make([]string, m.viewport.Width)
	}

	for _, p := range m.particles {
		y := p.physics.PosY()
		x := p.physics.PosX()

		if y < 0 || x < 0 || y >= height-1 || x >= width-1 {
			continue
		}

		grid[y][x] = p.char
	}

	// Print out grid
	for i := range grid {
		for _, col := range grid[i] {
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
