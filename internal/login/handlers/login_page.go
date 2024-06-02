package handlers

import (
	"net/http"
	"tabs/internal/constants"
	inf "tabs/internal/info"
	viewData "tabs/internal/login/data"
)

// LoginPage handles the login page request.
func (l *Login) LoginPage(writer http.ResponseWriter, request *http.Request) {
	info := l.Context.Informer.NewInfo(request, inf.WithTitle(l.Context.Localizer.Localize(constants.BundleLogin)))
	data := viewData.NewLoginData()
	page := l.Views.Login.CreateLoginPage(info, data)
	l.Renderer.Render(writer, request, info, page)
}
