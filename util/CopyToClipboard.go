package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CopyToClipboard is for Linux
func CopyToClipboard(text string) error {
	command := exec.Command("xclip")
	command.Args = []string{
		"-selection",
		"clipboard",
	}
	command.Stdin = bytes.NewReader([]byte(text))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting xclip command: %w", err)
	}

	err := command.Wait()
	if err != nil {
		return fmt.Errorf("error running xclip %w", err)
	}

	return nil
}
