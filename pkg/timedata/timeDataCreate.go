package timedata

import "time"

func CreateTimeData()[5]int{
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
	return timeInformation
}
