package docs

import (
	"github.com/go-chi/chi/v5"
	consts "tabs/internal/constants"
	hand "tabs/internal/handlers"
)

// Docs is used for routing and handling regarding documentation.
type Docs struct {
	Handlers *hand.Handlers
}

// NewDocs creates a new instance of the Docs struct.
func NewDocs(handlers *hand.Handlers) *Docs {
	return &Docs{Handlers: handlers}
}

// SetupRouter sets up the routes for the Docs struct.
func (d *Docs) SetupRouter(router chi.Router) {
	router.Get(consts.PathDocs, d.Handlers.Docs.DocsPage)
}
