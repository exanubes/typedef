package targets

import "github.com/exanubes/typedef/internal/domain"

type OutputOptions struct {
	Filepath string
}

func Create(t string, options OutputOptions) domain.OutputTarget {
	switch t {
	case "cli":
		return NewCliTarget()
	case "file":
		return NewFileTarget(options.Filepath)
	}
	return nil
}
