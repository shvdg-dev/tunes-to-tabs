package handlers

import (
	ctx "tabs/internal/context"
	rend "tabs/internal/renderer"
	"tabs/internal/views"
)

// Login is used for handlers regarding logging in or out.
type Login struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewLogin creates a new instance of the Login.
func NewLogin(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Login {
	return &Login{Context: context, Views: views, Renderer: renderer}
}
