//go:build windows
// +build windows

package fireworks

import (
	"syscall"

	"golang.org/x/term"
)

func InitialModel() model {
	h, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}
	width, height, err := term.GetSize(int(h))
	if err != nil {
		panic(err)
	}
	return InitialModelWithSize(width, height)
}
