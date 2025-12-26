package domain

type InputReader interface {
	Read() (string, error)
}
