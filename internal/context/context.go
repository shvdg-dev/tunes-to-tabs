package context

import (
	"github.com/shvdg-dev/base-pkg/database"
	"github.com/shvdg-dev/base-pkg/i18n"
	"github.com/shvdg-dev/base-pkg/logger"
	sess "github.com/shvdg-dev/base-pkg/sessions"
	ttt "github.com/shvdg-dev/tunes-to-tabs-api"
	"github.com/shvdg-dev/tunes-to-tabs/internal/info"
)

// Context represents the execution context of the application.
type Context struct {
	Localizer *i18n.Localizer
	Logger    *logger.Logger
	Informer  *info.Informer
	Sessions  *sess.Service
	API       *ttt.API
}

// NewContext initializes a new Context structure with the given dependencies.
func NewContext(database *database.Manager, localizer *i18n.Localizer) *Context {
	sessions := sess.NewService(database)
	return &Context{
		Localizer: localizer,
		Logger:    logger.NewLogger(),
		Informer:  info.NewInformer(sessions),
		Sessions:  sessions,
		API:       ttt.NewAPI(database)}
}
