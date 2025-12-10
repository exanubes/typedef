package targets

import "github.com/exanubes/typedef/internal/domain"

type FactoryOptions struct {
	Filepath string
}

type TargetFactory struct{}

func (factory TargetFactory) Create(t string, options FactoryOptions) domain.OutputTarget {
	switch t {
	case "cli":
		return NewCliTarget()
	case "file":
		return NewFileTarget(options.Filepath)
	case "clipboard":
		return NewClipboardTarget()
	}
	return nil
}

type Factory interface {
	Create(string, FactoryOptions) domain.OutputTarget
}
