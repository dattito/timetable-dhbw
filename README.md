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

## Usage

I published the docker image on github (gcr.io/dattito/timetable-dhbw). You can pull it and add the environment variable `ORIGINAL_ICS_URL` to the docker container (with the url of the original ics-file).
Then you can access the new ics-file on port 3000 of the container (on the root path) and subscribe to it in your calendar app, after publishing it to the internet.

## Further development

If there is a interest for it, I could extend the script and deploy it publicly for all DHBW students.
For that, contact me on github or via email (contact@relay.datti.to).
