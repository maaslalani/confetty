package cmd

import (
	"fmt"
	"os"

	"github.com/maaslalani/confetty/confetti"
	"github.com/maaslalani/confetty/fireworks"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "confetty",
	Short: "Confetti in your TTY",
	Long:  `Confetty gives your confetti and fireworks in your terminal`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(confetti.InitialModel(), tea.WithAltScreen())
		return p.Start()
	},
}

var confettiCmd = &cobra.Command{
	Use:     "confetti",
	Aliases: []string{"confetty"},
	Short:   "Confetti in your TTY",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(confetti.InitialModel(), tea.WithAltScreen())
		return p.Start()
	},
}

var fireworksCmd = &cobra.Command{
	Use:   "fireworks",
	Short: "Fireworks in your TTY",
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(fireworks.InitialModel(), tea.WithAltScreen())
		return p.Start()
	},
}

func init() {
	rootCmd.AddCommand(confettiCmd)
	rootCmd.AddCommand(fireworksCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
