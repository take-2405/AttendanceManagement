
package pkg

import (
	"github.com/gin-gonic/gin"
	"Attendance/pkg/handler"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init(){
	Server = gin.Default()
	Server.GET("/Attendance", handler.HandleResistAttend())
}

