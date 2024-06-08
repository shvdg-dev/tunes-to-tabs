package main

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
)

// createDatabase initializes the database connection by retrieving the database URL from the environment.
func createDatabase() *logic.DatabaseManager {
	URL := logic.GetValueAsString(databaseUrlKey)
	return logic.NewDatabaseManager(databaseDriver, URL)
}

// prepareDatabase prepares the database by creating tables and inserting data.
func prepareDatabase(context *ctx.Context) {
	context.API.Users.CreateUsersTable()
	//TODO: context.Sessions.CreateSessionsTable()
	insertAdmin(context)
}

// insertAdmin inserts an administrator user into the database.
func insertAdmin(context *ctx.Context) {
	email := logic.GetValueAsString(adminInitialEmailKey)
	password := logic.GetValueAsString(adminInitialPasswordKey)
	context.API.Users.InsertUser(email, password)
}
