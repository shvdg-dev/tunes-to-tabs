package context

import (
	"tabs/internal/info"
	"tabs/pkg/base/database"
	"tabs/pkg/base/i18n"
	loggr "tabs/pkg/base/logger"
	sess "tabs/pkg/base/sessions"
	ttt "tabs/pkg/ttt-api"
)

// Context represents the execution context of the application.
type Context struct {
	Localizer *i18n.Localizer
	Logger    *loggr.Logger
	Informer  *info.Informer
	Sessions  *sess.Service
	API       *ttt.API
}

// NewContext initializes a new Context structure with the given dependencies.
func NewContext(database *database.Manager, localizer *i18n.Localizer) *Context {
	sessions := sess.NewService(database)
	return &Context{
		Localizer: localizer,
		Logger:    loggr.NewLogger(),
		Informer:  info.NewInformer(sessions),
		Sessions:  sessions,
		API:       ttt.NewAPI(database)}
}
