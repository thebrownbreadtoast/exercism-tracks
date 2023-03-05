package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	timeObj, _ := time.Parse("1/2/2006 15:04:05", date)

	return timeObj
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointmentObj, _ := time.Parse("January 2, 2006 15:04:05", date)

	currentTimeObj := time.Now()

	return currentTimeObj.After(appointmentObj)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointmentObj, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return (appointmentObj.Hour() >= 12) && (appointmentObj.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTimeObj := Schedule(date)

	appointmentTimeStr := appointmentTimeObj.Format("Monday, January 2, 2006, at 15:04.")

	appointmentStr := fmt.Sprintf("You have an appointment on %s", appointmentTimeStr)

	return appointmentStr
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversaryDateObj, _ := time.Parse("2006-01-02 15:04:00 +0000 UTC", "2023-09-15 00:00:00 +0000 UTC")

	return anniversaryDateObj
}
