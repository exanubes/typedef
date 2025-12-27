package clipboard

import (
	"runtime"

	"github.com/exanubes/typedef/internal/domain"
)

func New() domain.Clipboard {

	switch runtime.GOOS {
	case "darwin":
		return NewMacosClipboard()
	case "linux":
		return NewLinuxClipboard()
	}

	return nil
}
