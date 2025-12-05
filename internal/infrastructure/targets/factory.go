package targets

import "github.com/exanubes/typedef/internal/domain"

func Create(t string) domain.OutputTarget {
	switch t {
	case "json":
		return NewCliTarget()
	}
	return nil
}
