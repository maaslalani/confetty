package confetty

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/maaslalani/confetty/pkg/array"
	"github.com/maaslalani/confetty/pkg/physics"

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
	colors     = []string{"#a864fd", "#29cdff", "#78ff44", "#ff718d", "#fdff6a"}
	characters = []string{"█", "▓", "▒", "░", "▄", "▀"}
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

// Confetti model
type model struct {
	particles []*Particle
	viewport  viewport.Model
}

type Particle struct {
	char    string
	physics *physics.Physics
}

func InitialModel() model {
	particles := []*Particle{}

	width, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	for i := 0; i < numParticles; i++ {
		x := float64(width / 2)
		y := float64(-1)

		p := &Particle{
			physics: physics.New(
				physics.Vector{X: x + (float64(width/4) * (rand.Float64() - 0.5)), Y: y},
				physics.Vector{X: (rand.Float64() - 0.5) * 100, Y: rand.Float64() * 50},
				physics.Vector(physics.Gravity),
				framesPerSecond,
			),
			char: lipgloss.NewStyle().
				Foreground(lipgloss.Color(array.Sample(colors))).
				Render(array.Sample(characters)),
		}

		particles = append(particles, p)
	}

	return model{particles: particles}
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
		for _, p := range m.particles {
			p.physics.Update()
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
	if m.viewport.Height <= 0 || m.viewport.Width <= 0 {
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

		if y < 0 {
			continue
		}

		// Particle is out of view
		if y >= m.viewport.Height-1 || x < 0 || x >= m.viewport.Width-1 {
			// loop motion
			p.physics.Reset()
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
