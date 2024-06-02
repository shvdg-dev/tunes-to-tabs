package handlers

import (
	ctx "tabs/internal/context"
	rend "tabs/internal/renderer"
	"tabs/internal/views"
)

// Home is used for handlers regarding the home page.
type Home struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewHome creates a new instance of the Home.
func NewHome(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Home {
	return &Home{Context: context, Views: views, Renderer: renderer}
}
