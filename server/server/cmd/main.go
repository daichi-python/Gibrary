package main

import (
	"os"
	"server/cmd/api"
	"server/cmd/database"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)

	err := godotenv.Load("../.env")
	if err != nil {
		logrus.WithField("error", err).Error("Failed to read env file.")
		return
	}

	dbController := database.Setup()
	apiController := api.NewAPIController(dbController)

	router := api.Setup(*apiController)

	router.Run(":4000")
}
