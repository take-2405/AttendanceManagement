package other

import (
	"Attendance/pkg/server/model"
	"Attendance/pkg/timedata"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func HandleResistBreakOUT()gin.HandlerFunc{
	return func(c *gin.Context){
		// Contextから認証済みのユーザIDを取得
		studentNumber := c.GetString("studentNumber")
		if len(studentNumber) == 0 {
			log.Println(errors.New("studentNumber is empty"))
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var timeInformation [5]int
		timeInformation =timedata.CreateTimeData()

		//attendTime:=yearDate+":"+realTime
		timeID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "timeID is error")
			return
		}
		if err:= model.InsertAttendanceTime(timeID.String(),studentNumber,"BreakIN",timeInformation);err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.JSON(http.StatusOK, "")
		return
	}
}
