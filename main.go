package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/useurmind/tut-microservice-basics/pkg/database"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()

	db, err := database.GetConnection()
	if err != nil {
		log.Error().Err(err).Msg("Could not connect to database")
		os.Exit(1)
	}
	err = database.MigrateDatabase(db.DB)
	if err != nil {
		log.Error().Err(err).Msg("Could not migrate database")
		os.Exit(1)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run()
}

