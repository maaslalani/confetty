//go:build !windows
// +build !windows

package fireworks

import (
	"golang.org/x/term"
)

func InitialModel() model {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	return InitialModelWithSize(width, height)
}
