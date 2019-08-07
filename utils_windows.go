// +build windows

package ishell

import (
	"github.com/Mrs4s/readline"
)

func clearScreen(s *Shell) error {
	return readline.ClearScreen(s.writer)
}
