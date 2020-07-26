package model

import (
	"Attendance/pkg/db"
)

func InsertTimeData(timeID,yearDate,RealTime string ,state int )error{
	stmt, err := db.Conn.Prepare("INSERT INTO timeManagement VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(timeID,state,yearDate,RealTime)
	return err
}