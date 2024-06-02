package middleware

import (
	"net/http"
	"strings"
	consts "tabs/internal/constants"
	ctx "tabs/internal/context"
	inf "tabs/internal/info"
	rend "tabs/internal/renderer"
	vi "tabs/internal/views"
)

// Middleware is used for handling middleware.
type Middleware struct {
	Context  *ctx.Context
	Views    *vi.Views
	Renderer *rend.Renderer
}

// NewMiddleware creates a new instance of the Middleware.
func NewMiddleware(context *ctx.Context, views *vi.Views, renderer *rend.Renderer) *Middleware {
	return &Middleware{Context: context, Views: views, Renderer: renderer}
}

// Authentication enforces authentication for protected resources.
func (m *Middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		isAuthenticated := m.Context.Sessions.IsAuthenticated(request)
		if !IsResourceAccessible(request.URL.Path) && !isAuthenticated {
			writer.WriteHeader(http.StatusUnauthorized)
			title := m.Context.Localizer.Localize(consts.BundleNotAuthenticatedTitle)
			info := m.Context.Informer.NewInfo(request, inf.WithTitle(title))
			notAuthPage := m.Views.Error.CreateNotAuthenticatedPage()
			m.Renderer.Render(writer, request, info, notAuthPage)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

// IsResourceAccessible checks if the given path is accessible.
func IsResourceAccessible(path string) bool {
	return strings.HasPrefix(path, consts.PathPublic) || strings.HasPrefix(path, consts.PathLogin)
}
