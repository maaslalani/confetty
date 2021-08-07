package confetty

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/maaslalani/confetty/pkg/physics"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const (
	fps = 60.0
)

var colors = []lipgloss.Color{
	lipgloss.Color("#a864fd"),
	lipgloss.Color("#29cdff"),
	lipgloss.Color("#78ff44"),
	lipgloss.Color("#ff718d"),
	lipgloss.Color("#fdff6a"),
}

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func waitASec(ms int) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Millisecond * time.Duration(ms))
		return nil
	}
}

type model struct {
	particles []*Particle
	viewport  viewport.Model
}

type Particle struct {
	char    string
	physics *physics.Physics
	color   lipgloss.Color
}

var characters = []string{"▄", "▀", "█"} // "▓", "▒", "░"}

func InitialModel() model {
	particles := []*Particle{}

	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 25; i++ {
		x := float64(width / 2)
		y := float64(height / 2)

		p := &Particle{
			char: lipgloss.NewStyle().Foreground(colors[rand.Intn(len(colors))]).Render(characters[rand.Intn(len(characters))]),
			physics: physics.New(
				physics.Vector{X: x, Y: y},
				physics.Vector{X: (rand.Float64() - 0.5) * 100, Y: (rand.Float64() - 0.5) * 100},
				physics.Vector(physics.Gravity),
				fps,
			),
		}
		particles = append(particles, p)
	}

	return model{particles: particles}
}

func (m model) Init() tea.Cmd {
	return tea.Sequentially(waitASec(500), animate())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
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
		if y < 0 || x < 0 || x >= m.viewport.Width-1 || y >= m.viewport.Height-1 {
			continue
		}
		grid[y][x] = p.char
	}
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
