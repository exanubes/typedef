package utils

type IdProvider struct {
	current int
}

func (provider *IdProvider) Next() int {
	provider.current += 1
	return provider.current
}
func (provider *IdProvider) Reset() {
	provider.current = 0
}
