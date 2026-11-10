package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	layout := "1/2/2006 15:04:05"
	timeObj, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return timeObj
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// "July 25, 2019 13:45:00"
	layout := "January 2, 2006 15:04:05"
	timeObj, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	return timeObj.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// "Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 2, 2006 15:04:05"
	timeObj, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	hour := timeObj.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// "7/25/2019 13:45:00"
	layoutIn := "1/2/2006 15:04:05"
	timeObj, err := time.Parse(layoutIn, date)
	if err != nil {
		panic(err)
	}

	// You have an appointment on Thursday, July 25, 2019, at 13:45
	layoutOut := "You have an appointment on Monday, January 2, 2006, at 15:04."
	return timeObj.Format(layoutOut)
}

// AnniversaryDate returns a Time with this year's anniversary.
// time.Date(2112, time.February, 29, 11, 59, 59, 0, time.UTC)
func AnniversaryDate() time.Time {
	thisYear := time.Now().Year()
	anniversary := time.Date(thisYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary

}
