package clipboard

import (
	"bytes"
	"os/exec"
	"strings"
)

type DarwinClipboard struct{}

func NewMacosClipboard() *DarwinClipboard {
	return &DarwinClipboard{}
}

func (clipboard DarwinClipboard) Read() (string, error) {
	cmd := exec.Command("pbpaste")
	var result bytes.Buffer

	cmd.Stdout = &result

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return result.String(), nil
}

func (clipboard DarwinClipboard) Write(input string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(input)

	return cmd.Run()
}
