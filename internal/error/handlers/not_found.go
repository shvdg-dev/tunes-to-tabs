package handlers

import (
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	doc "github.com/shvdg-dev/tunes-to-tabs/internal/info"
	"net/http"
)

// NotFound handles HTTP requests for a 404 Not Found error, by rendering the appropriate error page.
func (e *Error) NotFound() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		title := e.Context.Localizer.Localize(consts.BundlePageNotFoundTitle)
		info := e.Context.Informer.NewInfo(request, doc.WithTitle(title))
		page := e.Views.Error.CreatePageNotFound()
		e.Renderer.Render(writer, request, info, page)
	}
}
