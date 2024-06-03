package views

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	consts "tabs/internal/constants"
)

// CreateJobsPage creates a home page.
func (j *Jobs) CreateJobsPage() Node {
	return Div(Text(j.Context.Localizer.Localize(consts.BundleJobsIntro)))
}
