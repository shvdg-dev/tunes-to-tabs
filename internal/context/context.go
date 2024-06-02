package context

import (
	"tabs/internal/info"
	sess "tabs/internal/sessions"
	"tabs/internal/users"
	"tabs/pkg/database"
	"tabs/pkg/i18n"
	loggr "tabs/pkg/logger"
)

// Context represents the execution context of the application.
type Context struct {
	Localizer *i18n.Localizer
	Logger    *loggr.Logger
	Informer  *info.Informer
	Sessions  *sess.Service
	Users     *users.Service
}

// NewContext initializes a new Context structure with the given dependencies.
func NewContext(database *database.Connection, localizer *i18n.Localizer) *Context {
	sessions := sess.NewService(database)
	return &Context{
		Localizer: localizer,
		Logger:    loggr.NewLogger(),
		Informer:  info.NewInformer(sessions),
		Sessions:  sessions,
		Users:     users.NewService(database)}
}
