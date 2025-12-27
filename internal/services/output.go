package services

import (
	"fmt"

	"github.com/exanubes/typedef/internal/domain"
	"github.com/exanubes/typedef/internal/infrastructure/targets"
)

type OutputService struct {
	outputs   targets.Factory
	clipboard domain.Clipboard
}

func NewOutputService(factory targets.Factory, clipboard domain.Clipboard) *OutputService {
	return &OutputService{outputs: factory, clipboard: clipboard}
}

func (service *OutputService) Send(payload string, options domain.OutputOptions) error {
	target := service.outputs.Create(options.Target, targets.FactoryOptions{Filepath: options.Path, Clipboard: service.clipboard})

	if target == nil {
		return fmt.Errorf("Target '%s' is not supported", options.Target)
	}

	if err := target.Send(payload); err != nil {
		return fmt.Errorf("Failed to send payload to %s target:\n %w", options.Target, err)
	}

	return nil
}
