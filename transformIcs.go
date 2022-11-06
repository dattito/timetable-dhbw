package main

import (
	"fmt"
	"net/http"
	"strings"

	ics "github.com/arran4/golang-ical"
)

func transformCalender(ics *ics.Calendar) {
	events := ics.Events()

	for _, event := range events {
		if !strings.Contains(event.GetProperty("LOCATION").Value, "online") {
			//event.SetProperty("SUMMARY", fmt.Sprintf("Raum: %s \n %s", event.GetProperty("LOCATION").Value, event.GetProperty("SUMMARY").Value))
			roomName := ""
			if event.GetProperty("LOCATION").Value != "" {
				roomName = fmt.Sprintf("%s, ", event.GetProperty("LOCATION").Value)
			}
			event.SetLocation(fmt.Sprintf("%sCoblitzallee 1-9, 68163 Mannheim, Deutschland", roomName))
		}
	}
}

func getOriginalIcsFile(originalIcsUrl string) (*ics.Calendar, error) {

	resp, err := http.Get(originalIcsUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	cal, err := ics.ParseCalendar(resp.Body)

	if err != nil {
		return nil, err
	}

	return cal, nil
}

func getNewIcsFile(originalIcsUrl string) (*ics.Calendar, error) {
	cal, err := getOriginalIcsFile(originalIcsUrl)
	if err != nil {
		return nil, err
	}

	transformCalender(cal)

	return cal, nil
}
