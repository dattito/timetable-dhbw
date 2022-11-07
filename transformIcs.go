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
		if event.GetProperty("DESCRIPTION") == nil {
			event.SetDescription("")
		}

		// Add address to location
		if strings.Contains(event.GetProperty("LOCATION").Value, "online") || strings.Contains(event.GetProperty("SUMMARY").Value, "online") {
			location := event.GetProperty("LOCATION").Value
			if location == "" {
				location = "online"
			}
			event.SetDescription(fmt.Sprintf("(%s)\n%s", location, event.GetProperty("DESCRIPTION").Value))
			event.SetLocation("")
		} else {
			roomName := ""
			if event.GetProperty("LOCATION").Value != "" {
				roomName = fmt.Sprintf("%s, ", event.GetProperty("LOCATION").Value)
			}
			event.SetLocation(fmt.Sprintf("%sCoblitzallee 1-9, 68163 Mannheim, Deutschland", roomName))
		}

		// Sometimes, "online" is wrote into the summary
		if strings.Contains(event.GetProperty("SUMMARY").Value, "online") {
			if !strings.Contains(event.GetProperty("DESCRIPTION").Value, "online") {
				event.SetDescription(fmt.Sprintf("(%s)\n%s", "online", event.GetProperty("DESCRIPTION").Value))
			}
			event.SetSummary(strings.Replace(event.GetProperty("SUMMARY").Value, "online", "", -1))
		}

		// Move name of lecturer to summary
		summary := event.GetProperty("SUMMARY").Value
		lecturer, summaryWithoutLecturer := splitLecturerFromString(summary)
		if lecturer != "" {
			event.SetSummary(summaryWithoutLecturer)
			event.SetDescription(fmt.Sprintf("Dozent: %s\n%s", lecturer, event.GetProperty("DESCRIPTION").Value))
		}
	}
}

func splitLecturerFromString(text string) (string, string) {
	split := strings.Split(text, " ")

	lecturerTitleIndex := -1

	for i, t := range split {
		if strings.Contains(t, "Fr") ||
			strings.Contains(t, "Hr") ||
			strings.Contains(t, "Dr") ||
			strings.Contains(t, "Prof") {
			lecturerTitleIndex = i
			break
		}
	}

	if lecturerTitleIndex == -1 {
		return "", text
	}

	return strings.Join(split[lecturerTitleIndex:lecturerTitleIndex+2], " "), strings.Join(split[:lecturerTitleIndex], " ") + " " + strings.Join(split[lecturerTitleIndex+2:], " ")
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
