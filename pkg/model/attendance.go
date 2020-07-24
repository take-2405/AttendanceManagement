package model

import "Attendance/pkg/db"

func InsertAttendanceTime(yearDate,RealTime string)error{
	stmt, err := db.Conn.Prepare("INSERT INTO attendanceTime VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(yearDate,RealTime)
	return err
}