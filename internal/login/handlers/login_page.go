package handlers

import (
	"github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	inf "github.com/shvdg-dev/tunes-to-tabs/internal/info"
	viewData "github.com/shvdg-dev/tunes-to-tabs/internal/login/data"
	"net/http"
)

// LoginPage handles the login page request.
func (l *Login) LoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(constants.BundleLogin)))
	data := viewData.NewLoginData()
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}
