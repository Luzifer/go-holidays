package main

import (
	"fmt"
	"strings"
	"time"
)

type (
	iCalendar struct {
		Events []iCalendarEvent
	}

	iCalendarEvent struct {
		Summary string
		Date    time.Time
		UID     string
	}
)

func (i iCalendar) String() string {
	lines := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		fmt.Sprintf("PRODID:-//Luzifer//holiday-api %s//EN", version),
	}

	for _, evt := range i.Events {
		lines = append(lines, evt.String())
	}

	lines = append(lines, "END:VCALENDAR")

	return strings.Join(lines, "\r\n")
}

func (i iCalendarEvent) String() string {
	return strings.Join([]string{
		"BEGIN:VEVENT",
		strings.Join([]string{"DTSTAMP", time.Now().UTC().Format("20060102T150405Z")}, ":"),
		strings.Join([]string{"SUMMARY", i.Summary}, ":"),
		strings.Join([]string{"DTSTART;VALUE=DATE", i.Date.Format("20060102")}, ":"),
		strings.Join([]string{"UID", i.UID}, ":"),
		"DURATION:P1D",
		"END:VEVENT",
	}, "\r\n")
}
