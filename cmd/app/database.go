package main

import (
	"github.com/shvdg-dev/base-pkg/database"
	"github.com/shvdg-dev/base-pkg/utils"
	ctx "github.com/shvdg-dev/tunes-to-tabs/internal/context"
)

// createDatabase initializes the database connection by retrieving the database URL from the environment.
func createDatabase() *database.Manager {
	URL := utils.GetValueAsString(databaseUrlKey)
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
	email := utils.GetValueAsString(adminInitialEmailKey)
	password := utils.GetValueAsString(adminInitialPasswordKey)
	context.API.Users.InsertUser(email, password)
}
