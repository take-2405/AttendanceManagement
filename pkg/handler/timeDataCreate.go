package handler

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateTimeData()[3]string{
	var rValue  [3]string
	//const layout = "2006-01-02 15:04"
	const layout = "2006-01-02"
	time := time.Now()
	rValue[1] =time.Format(layout)
	minute:=strconv.Itoa(time.Minute())
	hour:= strconv.Itoa(time.Hour())
	rValue[2]=hour +":" + minute
	//attendTime:=yearDate+":"+realTime

	// UUIDでユーザIDを生成する
	timeID, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "timeID is error")
		//ここどうしよう
		//return
	}
	rValue[0]=timeID.String()
	return rValue
}
