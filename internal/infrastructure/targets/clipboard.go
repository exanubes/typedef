package targets

import (
	"fmt"

	"golang.design/x/clipboard"
)

type ClipboardTarget struct {
}

func NewClipboardTarget() *ClipboardTarget {
	return &ClipboardTarget{}
}

func (target *ClipboardTarget) Send(code string) error {
	clipboard.Write(clipboard.FmtText, []byte(code))
	fmt.Println("\nOutput saved to clipboard")
	return nil
}
