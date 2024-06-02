package error

import (
	"github.com/go-chi/chi/v5"
	hand "tabs/internal/handlers"
)

// Error is used for routing and handling regarding errors.
type Error struct {
	Handlers *hand.Handlers
}

// NewError returns a new instance of Error.
func NewError(handlers *hand.Handlers) *Error {
	return &Error{Handlers: handlers}
}

// SetupRouter sets up the error router.
func (e *Error) SetupRouter(router chi.Router) {
	router.NotFound(e.Handlers.Error.NotFound())
}
