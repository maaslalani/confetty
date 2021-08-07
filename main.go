package main

import (
	"fmt"
	"os"

	"github.com/maaslalani/confetty/confetty"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(confetty.InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
