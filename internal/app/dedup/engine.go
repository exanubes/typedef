package dedup

import "github.com/exanubes/typedef/internal/domain"

type DedupEngine struct {
	type_pool map[string]*domain.NamedType
}

func New() *DedupEngine {
	return &DedupEngine{
		type_pool: make(map[string]*domain.NamedType),
	}
}

func (engine *DedupEngine) Get(input *domain.ObjectType) *domain.NamedType {
	hash := engine.hash(input.Canonical())
	if result, ok := engine.type_pool[hash]; ok {
		return result
	}

	return nil
}

func (engine *DedupEngine) Add(input *domain.NamedType) {
	hash := engine.hash(input.Identity.Canonical())

	engine.type_pool[hash] = input
}

func (engine *DedupEngine) hash(value string) string {
	// TODO: hashing algorithm
	return value
}
