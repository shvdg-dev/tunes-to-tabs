package main

import (
	ctx "tabs/internal/context"
	"tabs/pkg/base/database"
	"tabs/pkg/base/environment"
)

// createDatabase initializes the database connection by retrieving the database URL from the environment.
func createDatabase() *database.Manager {
	URL := environment.GetValueAsString(databaseUrlKey)
	return database.NewManager(databaseDriver, URL)
}

// prepareDatabase prepares the database by creating tables and inserting data.
func prepareDatabase(context *ctx.Context) {
	context.API.Users.CreateUsersTable()
	context.Sessions.CreateSessionsTable()
	insertAdmin(context)
}

// insertAdmin inserts an administrator user into the database.
func insertAdmin(context *ctx.Context) {
	email := environment.GetValueAsString(adminInitialEmailKey)
	password := environment.GetValueAsString(adminInitialPasswordKey)
	context.API.Users.InsertUser(email, password)
}
