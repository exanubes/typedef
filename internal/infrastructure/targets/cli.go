package targets

import "fmt"

type CliTarget struct {
}

func NewCliTarget() *CliTarget {
	return &CliTarget{}
}

func (target *CliTarget) Send(code string) error {
	fmt.Println(code)
	return nil
}
