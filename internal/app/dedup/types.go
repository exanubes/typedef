package dedup

import "github.com/exanubes/typedef/internal/domain"

type Engine interface {
	Get(*domain.ObjectType) *domain.NamedType
	Add(*domain.NamedType)
}
