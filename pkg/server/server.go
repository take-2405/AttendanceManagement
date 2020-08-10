
package server

import (
	"Attendance/pkg/middleware"
	"Attendance/pkg/server/handler"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init(){
	Server = gin.Default()
	Server.GET("/Auth", handler.HandleAuthCreate())
	Server.GET("/Attendance", middleware.Authenticate(handler.HandleResistAttend()))
	//Server.GET("/BreakIN", handler.HandleResistBreakIN())
	//Server.GET("/BreakOUT", handler.HandleResistBreakOUT())
	//Server.GET("/Retired", middleware.Authenticate(handler.HandleResistRetired()))
}

