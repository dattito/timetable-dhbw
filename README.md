# Timetable DHBW

by Dattito | David Siregar

## Problem

The timetable of the DHBW (Mannheim) can be looked up as a Google Calender / ICS subscription.
However, the calener events include the names for the rooms of the lectures, but they do not include the
address of the DHBW, which I thought would be nice to have so my smartphone automatically picks it up
and e.g. reminds me when to leave to be there on time.

## Solution

Instead of subscribing to the Google Calender of the DHBW, I created a script that also acts like an ics-server (http-server).
But every time a client requests the ics-file, the script fetches the latest version of the timetable from the original
Google Calender (ORIGINAL_ICS_URL) and adds the address of the DHBW to the location of the events (only the events that are not online).

## Hosting of this projects source code

This project is hosted on [git.datti.to](https://git.datti.to/dattito/timetable-dhbw)
and has a mirror to [github.com](https://github.com/dattito/timetable-dhbw).

## Usage with Docker

```bash
docker run -d -p 3000:3000 -e ORIGINAL_ICS_URL=https://calendar.google.com/calendar/ical/.../public/basic.ics git.datti.to/dattito/timetable-dhbw:1.1.1
```

## Versioning

timetable-dhbw follows [Semantic Versioning](https://semver.org/). Docker Containers are tagged with the version number.
The `latest` tag always points to the latest commit on the `main` branch, which is not necessarily the latest release.

## Further development

If there is a interest for it, I could extend the script and deploy it publicly for all DHBW students.
For that, contact me on github or via email (contact@relay.datti.to).
