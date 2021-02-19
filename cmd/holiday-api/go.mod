module github.com/Luzifer/go-holidays/cmd/holiday-api

go 1.15

replace github.com/Luzifer/go-holidays/holidays => ../../holidays

require (
	github.com/Luzifer/go-holidays/holidays v0.0.0-00010101000000-000000000000
	github.com/Luzifer/rconfig/v2 v2.2.1
	github.com/gorilla/mux v1.8.0
)
