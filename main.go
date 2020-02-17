package main

import (
	"github.com/ZaneWithSpoon/fathomBack/api"
	"github.com/ZaneWithSpoon/fathomBack/config"
	"github.com/ZaneWithSpoon/fathomBack/db"
	"github.com/ZaneWithSpoon/fathomBack/types"
)

func main() {
	var dbService *db.DbService

	if config.IsDev() {
		dbService = db.GetDBService(
			"localhost",
			5432,
			"postgres",
			"postgres",
		)
	} else {
		dbService = db.GetDBService(
			"{production_server_url}.rds.amazonaws.com",
			5432,
			"postgres",
			"postgres",
		)
	}

	//start DB
	db.Start(dbService)

	//run DB Migrations
	types.MigrateUsers()

	//Start API
	api.StartAPI(dbService)
}