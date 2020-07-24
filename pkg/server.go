package pkg

import (
	"AttendanceMagegement/pkg/handler"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init(){
	Server = gin.Default()
	Server.GET("/resist", handler.HandleResistCreate())
}