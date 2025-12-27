package clipboard

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type LinuxClipboard struct{}

func NewLinuxClipboard() *LinuxClipboard {
	return &LinuxClipboard{}
}

func (clipboard LinuxClipboard) Read() (string, error) {
	cmd := read_cmd()
	if cmd == nil {
		return "", fmt.Errorf("Clipboard support disabled.")
	}
	var result bytes.Buffer

	cmd.Stdout = &result

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return result.String(), nil
}

func (clipboard LinuxClipboard) Write(input string) error {
	cmd := write_cmd()

	if cmd == nil {
		return fmt.Errorf("Clipboard support disabled.")
	}

	cmd.Stdin = strings.NewReader(input)

	return cmd.Run()
}

func read_cmd() *exec.Cmd {
	if has("wl-paste") {
		return exec.Command("wl-paste")
	}

	if has("xclip") {
		return exec.Command("xclip", "-selection", "clipboard", "-o")
	}

	return nil
}

func write_cmd() *exec.Cmd {
	if has("wl-copy") {
		return exec.Command("wl-copy")
	}

	if has("xclip") {
		return exec.Command("xclip", "-selection", "clipboard")
	}

	return nil
}

func has(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
