package login

import (
	"github.com/go-chi/chi/v5"
	consts "github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
	hand "github.com/shvdg-dev/tunes-to-tabs/internal/handlers"
)

// Login is used for routing and handling regarding logging in or out.
type Login struct {
	Context  *ctx.Context
	Handlers *hand.Handlers
}

// NewLogin creates a new Login instance.
func NewLogin(context *ctx.Context, handlers *hand.Handlers) *Login {
	return &Login{Context: context, Handlers: handlers}
}

// SetupRouter sets up the router for the Login struct.
func (l *Login) SetupRouter(router chi.Router) {
	router.Get(consts.PathLogin, l.Handlers.Login.LoginPage)
	router.Post(consts.PathLogin, l.Handlers.Login.LoggingIn)
	router.Get(consts.PathLogout, l.Handlers.Login.LoggingOut)
}
