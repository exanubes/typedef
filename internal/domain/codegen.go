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

func ParseFormat(input string) (Format, error) {
	switch input {
	case "go", "golang":
		return GOLANG, nil
	case "typescript", "ts":
		return TYPESCRIPT, nil
	case "zod", "ts-zod", "ts_zod", "typescript_zod", "typescript-zod":
		return ZOD, nil
	case "jsdoc":
		return JSDOC, nil
	default:
		return -1, fmt.Errorf("%s is not a valid format", input)
	}
}

const DEFAULT_OUTPUT_PATH = "./typedef.txt"
