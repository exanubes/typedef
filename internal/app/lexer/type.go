package lexer

import "github.com/exanubes/typedef/internal/domain"

type Lexer interface {
	NextToken() domain.Token
}

type Factory interface {
	Create(t, input string) Lexer
}
