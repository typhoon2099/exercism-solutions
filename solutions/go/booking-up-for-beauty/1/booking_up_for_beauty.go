package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedTime, _ := time.Parse("1/02/2006 15:04:05", date)

    return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	timeDate, _ := time.Parse("January 2, 2006 15:04:05", date)

    return timeDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	timeDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

    fmt.Println(timeDate)

    return timeDate.Hour() >= 12 && timeDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	timeDate, _ := time.Parse("1/2/2006 15:04:05", date)

    return fmt.Sprintf("You have an appointment on %s.", timeDate.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().Year()

    return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
