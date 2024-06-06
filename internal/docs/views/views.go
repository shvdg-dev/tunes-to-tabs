package views

import "github.com/shvdg-dev/tunes-to-tabs/internal/context"

// Docs is used for views regarding documentation.
type Docs struct {
	Context *context.Context
}

// NewDocs creates a new Docs instance with the given context.
func NewDocs(context *context.Context) *Docs {
	return &Docs{Context: context}
}
