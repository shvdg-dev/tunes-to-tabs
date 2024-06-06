package views

import (
	. "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
)

// CreateHomePage creates a home page.
func (h *Home) CreateHomePage() Node {
	return Div(Text(h.Context.Localizer.Localize(consts.BundleHomeIntro)))
}
