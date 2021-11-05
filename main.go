package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/confetty/confetti"
	"github.com/maaslalani/confetty/fireworks"
)

func main() {
	var err error
	var model tea.Model
	if len(os.Args) > 1 && os.Args[1] == "fireworks" {
		model = fireworks.InitialModel()
	} else {
		model = confetti.InitialModel()
	}
	p := tea.NewProgram(model, tea.WithAltScreen())
	err = p.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
