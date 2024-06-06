package handlers

import (
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/docs/handlers"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/error/handlers"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/home/handlers"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/jobs/handlers"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/login/handlers"
	"github.com/shvdg-dev/tunes-to-tabs/internal/renderer"
	"github.com/shvdg-dev/tunes-to-tabs/internal/views"
)

// Handlers represents a collection of handlers.
type Handlers struct {
	Home  *Home
	Docs  *Docs
	Login *Login
	Jobs  *Jobs
	Error *Error
}

// NewHandlers creates a new instance of the Handlers structure.
func NewHandlers(context *ctx.Context, views *views.Views, renderer *renderer.Renderer) *Handlers {
	return &Handlers{
		Home:  NewHome(context, views, renderer),
		Docs:  NewDocs(context, views, renderer),
		Login: NewLogin(context, views, renderer),
		Jobs:  NewJobs(context, views, renderer),
		Error: NewError(context, views, renderer)}
}
