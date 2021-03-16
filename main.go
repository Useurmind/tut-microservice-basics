package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/useurmind/tut-microservice-basics/pkg/database"
	"github.com/useurmind/tut-microservice-basics/pkg/web"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	err := database.MigrateDatabase()
	handleError(err, "migrate_database")
	

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	web.HandlePeople("/people", r)

	r.Run()
}

func handleError(err error, action string) {
	if err != nil {
		log.Error().Err(err).Msgf("Error during %s", action)
		os.Exit(1)
	}
}