module github.com/Luzifer/go-holidays/cmd/holiday-api

go 1.17

replace github.com/Luzifer/go-holidays/holidays => ../../holidays

require (
	github.com/Luzifer/go-holidays/holidays v0.0.0-20210219180756-c48fac09a2e4
	github.com/Luzifer/rconfig/v2 v2.4.0
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	gopkg.in/validator.v2 v2.0.0-20210331031555-b37d688a7fb0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
