package jobs

import (
	"github.com/go-chi/chi/v5"
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	hand "github.com/shvdg-dev/tunes-to-tabs/internal/handlers"
)

// Jobs is used for routing and handling regarding the jobs page.
type Jobs struct {
	Handlers *hand.Handlers
}

// NewJobs creates a new instance of Jobs.
func NewJobs(handlers *hand.Handlers) *Jobs {
	return &Jobs{Handlers: handlers}
}

// SetupRouter sets up the routes for the Jobs struct.
func (h *Jobs) SetupRouter(router chi.Router) {
	router.Get("/", h.Handlers.Jobs.JobsPage)
	router.Get(consts.PathJobs, h.Handlers.Jobs.JobsPage)
	router.Post(consts.PathPlaylistToTabs, h.Handlers.Jobs.PlaylistToTabs)
}
