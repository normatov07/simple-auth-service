package main

import (
	"github.com/joho/godotenv"
	"github.com/normatov07/auth-service/app/api/rest"
	"github.com/normatov07/auth-service/common/utils"
	"github.com/normatov07/auth-service/db/postgres"
)

func init() {
	err := godotenv.Load()
	utils.LoadLogs()
	if err != nil {
		panic(err)
	}

	utils.SetMode()
}

func main() {
	postgres.InitConn()
	defer postgres.Close()

	app := rest.GetServer()

	app.RunHTTP()

	defer utils.LogFile.Close()
}
