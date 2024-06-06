package handlers

import (
	"github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	inf "github.com/shvdg-dev/tunes-to-tabs/internal/info"
	"net/http"
)

// DocsPage handles the documentation page request.
func (d *Docs) DocsPage(writer http.ResponseWriter, request *http.Request) {
	title := d.Context.Localizer.Localize(constants.BundleDocs)
	info := d.Context.Informer.NewInfo(request, inf.WithTitle(title))
	page := d.Views.Docs.CreateDocsPage()
	d.Renderer.Render(writer, request, info, page)
}
