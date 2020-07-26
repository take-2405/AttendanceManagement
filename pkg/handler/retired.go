package handler

import (
	"Attendance/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HandleResistRetired()gin.HandlerFunc{
	return func(c *gin.Context){
		//data 1:timeID 2:year and month and date 3:hour and minutes
		data:=CreateTimeData()
		if err:=model.InsertTimeData(data[0],data[1],data[2],4);err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.JSON(http.StatusOK, "")
		// レスポンスに必要な情報を詰めて返却
		//c.JSON(http.StatusOK, view.UserCharacterGetResponse{Characters: viewUserCharacters})
	}
}