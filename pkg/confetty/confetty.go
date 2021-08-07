package confetty

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

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
	char  string
	acc   Acceleration
	vel   Velocity
	pos   Position
	color lipgloss.Color
}

type Position struct {
	x float64
	y float64
}

type Velocity struct {
	x float64
	y float64
}

type Acceleration struct {
	x float64
	y float64
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
			pos:  Position{x: x, y: y},
			vel:  Velocity{x: (rand.Float64() - 0.5) * 100, y: (rand.Float64() - 0.5) * 100},
			acc:  Acceleration{x: 0, y: 9.8},
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
			p.pos.y, p.vel.y = p.pos.y+p.vel.y/fps, p.vel.y+p.acc.y/fps
			p.pos.x, p.vel.x = p.pos.x+p.vel.x/fps, p.vel.x+p.acc.x/fps
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
		y := int(math.Round(p.pos.y))
		x := int(math.Round(p.pos.x))
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
