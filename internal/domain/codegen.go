package domain

import "fmt"

type Format int

const (
	GOLANG = iota
	TYPESCRIPT
	ZOD
	JSDOC
)

func (f Format) String() string {
	switch f {
	case GOLANG:
		return "golang"

	case TYPESCRIPT:
		return "typescript"

	case ZOD:
		return "typescript-zod"

	case JSDOC:
		return "jsdoc"

	}

	return fmt.Sprintf("%d", f)
}
