package handler

import (
	"Attendance/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)
func HandleResistAttend()gin.HandlerFunc{
		return func(c *gin.Context){
			//fmt.Println("aaaa")
			const layout = "2006-01-02 15:04"
			time := time.Now()
			yearDate:=time.Format(layout)
			minute:=strconv.Itoa(time.Minute())
			hour:= strconv.Itoa(time.Hour())
			realTime :=hour +":" + minute

			if err:=model.InsertAttendanceTime(yearDate,realTime);err!=nil{
				log.Println(err)
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
				return
			}
			//fmt.Println(reflect.TypeOf(yearDate))
			fmt.Println(yearDate)
			fmt.Println(realTime)
			c.JSON(http.StatusOK, "")
			return
	}
}