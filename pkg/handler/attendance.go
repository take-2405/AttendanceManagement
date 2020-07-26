package handler

import (
	"Attendance/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
func HandleResistAttend()gin.HandlerFunc{
		return func(c *gin.Context){
			data:=CreateTimeData()
			if err:=model.InsertTimeData(data[0],data[1],data[2],1);err!=nil{
				log.Println(err)
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
				return
			}
			c.JSON(http.StatusOK, "")

	}
}