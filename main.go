package main

import (
	"net/http"
	"os"

	"github.com/darkcorner17375/go-gin-mysql-restfulAPI/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB = config.InitGorm()
	log          = logrus.New()
)

func init() {
	_ = godotenv.Load(".env")
	switch os.Getenv("GIN_MODE") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	config.InitConfig()

}

func main() {

	defer config.CloseDatabaseConnection(db)

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

	if os.Getenv("PORT") != "" || os.Getenv("HOST") != "" {
		port := os.Getenv("PORT")
		router.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}

}
