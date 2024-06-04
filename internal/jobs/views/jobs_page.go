package views

import (
	icons "github.com/eduardolat/gomponents-lucide"
	. "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	consts "tabs/internal/constants"
)

// CreateJobsPage creates a home page.
func (j *Jobs) CreateJobsPage() Node {
	return Div(
		Header(Text(j.Context.Localizer.Localize(consts.BundleJobsIntro))),
		Div(Class("pt-4 flex flex-col space-y-3"),
			j.CreatePlaylistToTabsForm()))
}

// CreatePlaylistToTabsForm creates a form for processing a playlist to tabs.
func (j *Jobs) CreatePlaylistToTabsForm() Node {
	return FormEl(hx.Post(""),
		Div(Class("flex flex-col space-y-2"),
			j.CreatePlaylistLinkField(),
			j.CreateSubmitButton()))
}

// CreatePlaylistLinkField creates a field for entering a playlist link.
func (j *Jobs) CreatePlaylistLinkField() Node {
	return Div(
		Label(Text(j.Context.Localizer.Localize(consts.BundlePlaylistFieldTitle))),
		Label(components.Classes{"input": true, "input-bordered": true, "flex": true, "items-center": true, "gap-2": true, "max-w-md": true},
			icons.ListMusic(components.Classes{"stroke-current": true}),
			Input(Class("grow"),
				Attr("type", "text"),
				Attr("name", "playlist-link"),
				Placeholder(j.Context.Localizer.Localize(consts.BundlePlaylistFieldPlaceholder)))))
}

// CreateSubmitButton creates the button for submitting a form.
func (j *Jobs) CreateSubmitButton() Node {
	return Button(Class("btn btn-primary btn-md w-20"), Type("submit"),
		Text(j.Context.Localizer.Localize(consts.BundleSubmit)))
}
