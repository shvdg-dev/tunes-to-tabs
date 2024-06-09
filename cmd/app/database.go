package main

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs/internal/constants"
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
)

// createDatabase initializes the database connection by retrieving the database URL from the environment.
func createDatabase() *logic.DatabaseManager {
	URL := logic.GetEnvValueAsString(constants.KeyDatabaseURL)
	return logic.NewDatabaseManager(constants.ValueDatabaseDriver, URL)
}

// prepareDatabase prepares the database by creating tables and inserting data.
func prepareDatabase(context *ctx.Context) {
	context.API.Users.CreateUsersTable()
	//TODO: context.Sessions.CreateSessionsTable()
	insertAdmin(context)
}

// insertAdmin inserts an administrator user into the database.
func insertAdmin(context *ctx.Context) {
	email := logic.GetEnvValueAsString(constants.KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(constants.KeyAdminInitialPassword)
	context.API.Users.InsertUser(email, password)
}
