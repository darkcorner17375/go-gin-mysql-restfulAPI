package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	//
	log.SetLevel(log.TraceLevel)
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Error loading .env file")
	}

}

func main() {

	//新增gin引擎
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong123321",
		})
	})

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
	}))

	log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")

	if os.Getenv("PORT") != "" || os.Getenv("HOST") != "" {
		port := os.Getenv("PORT")
		router.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

}
