package fireworks

import (
	"math"
	"math/rand"
	"time"

	"github.com/maaslalani/confetty/array"
	"github.com/maaslalani/confetty/simulation"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	framesPerSecond = 60.0
	numParticles    = 50
)

var (
	colors     = []string{"#a864fd", "#29cdff", "#78ff44", "#ff718d", "#fdff6a"}
	characters = []string{"+", "*", "•"}
	head_char  = "▄"
	tail_char  = "│"
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/framesPerSecond, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

type model struct {
	system *simulation.System
}

func SpawnShoot(width, height int) simulation.Particle {
	color := lipgloss.Color(array.Sample(colors))
	v := float64(rand.Intn(15) + 15.0)

	x := rand.Float64() * float64(width)

	p := simulation.Particle{
		Physics: simulation.NewProjectile(
			simulation.FPS(framesPerSecond),
			simulation.Point{X: x, Y: float64(height)},
			simulation.Vector{X: 0, Y: -v},
			simulation.Vector(simulation.TerminalGravity),
		),
		Char:          lipgloss.NewStyle().Foreground(color).Render(head_char),
		TailChar:      lipgloss.NewStyle().Foreground(color).Render(tail_char),
		Shooting:      true,
		ExplosionCall: SpawnExplosion,
	}
	return p
}

func SpawnExplosion(x, y float64, width, height int) []simulation.Particle {
	color := lipgloss.Color(array.Sample(colors))
	v := float64(rand.Intn(10) + 20.0)

	particles := []simulation.Particle{}

	for i := 0; i < numParticles; i++ {
		p := simulation.Particle{
			Physics: simulation.NewProjectile(
				simulation.FPS(framesPerSecond),
				simulation.Point{X: x, Y: y},
				simulation.Vector{X: math.Cos(float64(i)) * v, Y: math.Sin(float64(i)) * v / 2},
				simulation.Vector(simulation.TerminalGravity),
			),
			Char:     lipgloss.NewStyle().Foreground(color).Render(array.Sample(characters)),
			Shooting: false,
		}
		particles = append(particles, p)
	}
	return particles
}

func InitialModel() model {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	return model{system: &simulation.System{
		Particles: []simulation.Particle{SpawnShoot(width, height)},
		Frame: simulation.Frame{
			Width:  width,
			Height: height,
		},
	}}
}

// Init initializes the confetti after a small delay
func (m model) Init() tea.Cmd {
	return animate()
}

// Update updates the model every frame, it handles the animation loop and
// updates the particle physics every frame
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		m.system.Particles = append(m.system.Particles, SpawnShoot(m.system.Frame.Width, m.system.Frame.Height))

		return m, nil
	case frameMsg:
		m.system.Update()
		return m, animate()
	case tea.WindowSizeMsg:
		m.system.Frame.Width = msg.Width
		m.system.Frame.Height = msg.Height
		return m, nil
	default:
		return m, nil
	}
}

// View displays all the particles on the screen
func (m model) View() string {
	return m.system.Render()
}
