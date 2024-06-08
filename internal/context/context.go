package context

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	ttt "github.com/shvdg-dev/tunes-to-tabs-api"
	"github.com/shvdg-dev/tunes-to-tabs/internal/info"
)

// Context represents the execution context of the application.
type Context struct {
	Localizer *logic.Localizer
	Logger    *logic.Logger
	Informer  *info.Informer
	API       *ttt.API
}

// NewContext initializes a new Context structure with the given dependencies.
func NewContext(database *logic.DatabaseManager, localizer *logic.Localizer) *Context {
	return &Context{
		Localizer: localizer,
		Logger:    logic.NewLogger(),
		Informer:  info.NewInformer(logic.NewSessionManager(database)),
		API:       ttt.NewAPI(database)}
}
