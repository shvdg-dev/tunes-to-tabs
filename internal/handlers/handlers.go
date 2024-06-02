package handlers

import (
	ctx "tabs/internal/context"
	. "tabs/internal/docs/handlers"
	. "tabs/internal/error/handlers"
	. "tabs/internal/home/handlers"
	. "tabs/internal/login/handlers"
	"tabs/internal/renderer"
	"tabs/internal/views"
)

// Handlers represents a collection of handlers.
type Handlers struct {
	Home  *Home
	Docs  *Docs
	Login *Login
	Error *Error
}

// NewHandlers creates a new instance of the Handlers structure.
func NewHandlers(context *ctx.Context, views *views.Views, renderer *renderer.Renderer) *Handlers {
	return &Handlers{
		Home:  NewHome(context, views, renderer),
		Login: NewLogin(context, views, renderer),
		Docs:  NewDocs(context, views, renderer),
		Error: NewError(context, views, renderer)}
}
