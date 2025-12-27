package targets

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
)

type ClipboardTarget struct {
	clipboard domain.Clipboard
}

func NewClipboardTarget(clipboard domain.Clipboard) *ClipboardTarget {
	return &ClipboardTarget{clipboard: clipboard}
}

func (target *ClipboardTarget) Send(code string) error {
	if err := target.clipboard.Write(code); err != nil {
		return err
	}
	fmt.Println("\nOutput saved to clipboard")
	return nil
}
