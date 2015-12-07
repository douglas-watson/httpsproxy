# HTTPS Proxy for google docs

* Douglas Watson
* November 2015

I built a Shiny application that fetches data from a Google Spreadsheet. This
works great on my laptop, but fails on shinyapps.io, where https requests are blocked.
This small google app engine project works as an https to http proxy.

Running this requires the Google App engine Golang SDK. Download it and place it
in in ~/apps/go_appengine/. Activate the development environment with:

  source env_setup.sh

Then run the development environment with:

  goapp serve

To publish to App Engine, create a new app on
https://console.developers.google.com, and follow the instructions there.
