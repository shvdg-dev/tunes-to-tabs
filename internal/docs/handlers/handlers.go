package handlers

import (
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
	rend "github.com/shvdg-dev/tunes-to-tabs/internal/renderer"
	"github.com/shvdg-dev/tunes-to-tabs/internal/views"
)

// Docs is used for handlers regarding logging in or out.
type Docs struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewDocs creates a new instance of the Docs.
func NewDocs(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Docs {
	return &Docs{Context: context, Views: views, Renderer: renderer}
}
