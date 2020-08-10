package handler

import (
	"Attendance/pkg/server/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func HandleResistAttend()gin.HandlerFunc{
		return func(c *gin.Context){

			// Contextから認証済みのユーザIDを取得
			studentNumber := c.GetString("studentNumber")
			if len(studentNumber) == 0 {
				log.Println(errors.New("studentNumber is empty"))
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
				return
			}

			//const layout = "2006-01-02 15:04"
			//const layout = "2006-01-02"
			var timeInformation [5]int
			time := time.Now()
			//yearDate:=time.Format(layout)
			timeInformation[0]=time.Year()
			timeInformation[1]=int(time.Month())
			timeInformation[2]=time.Day()
			timeInformation[3]=time.Hour()
			timeInformation[4]= time.Minute()

			//UUIDでユーザIDを生成する
			timeID, err := uuid.NewRandom()
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, "timeID is error")
				return
			}
			if err:= model.InsertAttendanceTime(timeID.String(),studentNumber,"active",timeInformation);err!=nil{
				log.Println(err)
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
				return
			}
			fmt.Println("ok")
			c.JSON(http.StatusOK, "")
			return
	}
}