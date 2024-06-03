package views

import "tabs/internal/context"

// Jobs is used for views regarding the home page.
type Jobs struct {
	Context *context.Context
}

// NewJobs initializes a new instance of Jobs.
func NewJobs(context *context.Context) *Jobs {
	return &Jobs{Context: context}
}
