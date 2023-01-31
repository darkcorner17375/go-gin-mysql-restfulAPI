package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

type configServer struct {
	Port string
	Host string
}

var ServerInfo configServer

func initServerConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		ServerInfo.Port = ""
		ServerInfo.Host = ""
	case gin.DebugMode:
		ServerInfo.Port = os.Getenv("PORT")
		ServerInfo.Host = os.Getenv("HOST")
	}
}
