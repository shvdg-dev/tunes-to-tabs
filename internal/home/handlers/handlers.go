package handlers

import (
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
	rend "github.com/shvdg-dev/tunes-to-tabs/internal/renderer"
	"github.com/shvdg-dev/tunes-to-tabs/internal/views"
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
