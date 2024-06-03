package views

import (
	ctx "tabs/internal/context"
	. "tabs/internal/docs/views"
	. "tabs/internal/document/navbar"
	. "tabs/internal/error/views"
	. "tabs/internal/home/views"
	. "tabs/internal/jobs/views"
	. "tabs/internal/login/views"
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
