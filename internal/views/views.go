package views

import (
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/docs/views"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/document/navbar"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/error/views"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/home/views"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/jobs/views"
	. "github.com/shvdg-dev/tunes-to-tabs/internal/login/views"
)

// Views represents a collection of views.
type Views struct {
	Navbar *Navbar
	Home   *Home
	Docs   *Docs
	Login  *Login
	Jobs   *Jobs
	Error  *Error
}

// NewViews creates a new instance of the Views structure.
func NewViews(context *ctx.Context) *Views {
	return &Views{
		Navbar: NewNavBar(context),
		Home:   NewHome(context),
		Docs:   NewDocs(context),
		Login:  NewLogin(context),
		Jobs:   NewJobs(context),
		Error:  NewError(context),
	}
}
