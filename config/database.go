package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

type configDatabase struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

var Database configDatabase

func initDatabaseConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		Database.Host = ""
		Database.Port = ""
		Database.Name = ""
		Database.Username = ""
		Database.Password = ""
	case gin.DebugMode:
		Database.Host = os.Getenv("DB_HOST")
		Database.Port = os.Getenv("DB_PORT")
		Database.Name = os.Getenv("DB_NAME")
		Database.Username = os.Getenv("DB_User")
		Database.Password = os.Getenv("DB_PASSWORD")
	}
}
