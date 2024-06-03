package handlers

import (
	ctx "tabs/internal/context"
	rend "tabs/internal/renderer"
	"tabs/internal/views"
)

// Jobs is used for handlers regarding the Jobs page.
type Jobs struct {
	Context  *ctx.Context
	Views    *views.Views
	Renderer *rend.Renderer
}

// NewJobs creates a new instance of the Jobs.
func NewJobs(context *ctx.Context, views *views.Views, renderer *rend.Renderer) *Jobs {
	return &Jobs{Context: context, Views: views, Renderer: renderer}
}
