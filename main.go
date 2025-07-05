package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/confetty/confetti"
	"github.com/maaslalani/confetty/fireworks"
)

func main() {
	var err error
	var model tea.Model
	modelType := flag.String("m", "confetti", "model type: confetti|fireworks")
	autoShooting := flag.Bool("a", false, "enable auto shooting")
	autoShootingInterval := flag.Uint("i", 1, "auto shooting interval 1~10 sec")
	flag.Parse()
	if *autoShootingInterval < 1 {
		*autoShootingInterval = 1
	}
	if *autoShootingInterval > 10 {
		*autoShootingInterval = 10
	}
	if *modelType == "fireworks" {
		model = fireworks.InitialModel(*autoShooting, *autoShootingInterval)
	} else {
		model = confetti.InitialModel(*autoShooting, *autoShootingInterval)
	}
	p := tea.NewProgram(model, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
