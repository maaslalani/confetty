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
	if len(os.Args) > 1 && os.Args[1] == "fireworks" {
		p := tea.NewProgram(fireworks.InitialModel(), tea.WithAltScreen())
		err = p.Start()
	} else {
		p := tea.NewProgram(confetti.InitialModel(), tea.WithAltScreen())
		err = p.Start()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
