package model

import (
	"Attendance/pkg/db"
)

func InsertAttendanceTime(timeID,studentNumber,state string,timeInformation [5]int )error{
	stmt, err := db.Conn.Prepare("INSERT INTO timeManagement VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(timeID,studentNumber,state,timeInformation[0],timeInformation[1],timeInformation[2],timeInformation[3],timeInformation[4])
	return err
}