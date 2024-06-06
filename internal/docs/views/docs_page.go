package views

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
)

// CreateDocsPage creates a documentation page.
func (d *Docs) CreateDocsPage() Node {
	return Div(Text(d.Context.Localizer.Localize(consts.BundleDocsIntro)))
}
