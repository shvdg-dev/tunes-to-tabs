package home

import (
	"github.com/go-chi/chi/v5"
	consts "tabs/internal/constants"
	hand "tabs/internal/handlers"
)

// Home is used for routing and handling regarding the home page.
type Home struct {
	Handlers *hand.Handlers
}

// NewHome creates a new instance of the Home.
func NewHome(handlers *hand.Handlers) *Home {
	return &Home{Handlers: handlers}
}

// SetupRouter sets up the routes for the Home struct.
func (h *Home) SetupRouter(router chi.Router) {
	router.Get("/", h.Handlers.Home.HomePage)
	router.Get(consts.PathHome, h.Handlers.Home.HomePage)
}
