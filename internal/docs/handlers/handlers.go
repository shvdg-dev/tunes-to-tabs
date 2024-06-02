package handlers

import (
	ctx "tabs/internal/context"
	rend "tabs/internal/renderer"
	"tabs/internal/views"
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
