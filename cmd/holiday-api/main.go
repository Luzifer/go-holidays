package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Luzifer/go-holidays/holidays"
	"github.com/Luzifer/rconfig/v2"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	cfg = struct {
		Listen         string `flag:"listen" default:":3000" description:"IP/Port to listen on"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	//go:embed index.html
	indexHTML []byte

	version = "dev"
)

func init() {
	if err := rconfig.Parse(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("holiday-api %s\n", version)
		os.Exit(0)
	}
}

func main() {
	r := mux.NewRouter()
	parts := []string{"", "{country-code:[a-z-]+}", "{year:[0-9]{4}}", "{month:[0-9]{2}}", "{day:[0-9]{2}}"}
	for i := 2; i <= len(parts); i++ {
		p := strings.Join(parts[:i], "/")
		r.HandleFunc(p, handleHolidays)
		r.HandleFunc(strings.Join([]string{p, "{format:[a-z]+}"}, "."), handleHolidays)
	}
	r.HandleFunc("/", handleReadme)

	srv := &http.Server{
		Addr:         cfg.Listen,
		Handler:      r,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}
	log.Println(srv.ListenAndServe())
}

func handleHolidays(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		countryCode = vars["country-code"]
		format      = vars["format"]
		year        = time.Now().Year()
	)

	if format == "" {
		format = "json"
	}

	if y, ok := vars["year"]; ok {
		var err error
		if year, err = strconv.Atoi(y); err != nil {
			http.Error(res, "You need to specify the year as a 4 character number", http.StatusBadRequest)
			return
		}
	}

	check := strings.TrimRight(strings.Join([]string{strconv.Itoa(year), vars["month"], vars["day"]}, "-"), "-")

	days, err := holidays.GetHolidays(countryCode, year)
	if err != nil {
		http.Error(res, "An error ocurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	outputSet := []holidays.Holiday{}
	for _, h := range days {
		if strings.HasPrefix(h.Date, check) {
			outputSet = append(outputSet, h)
		}
	}

	switch format {
	case "ics":
		cal := iCalendar{}
		for _, h := range outputSet {
			cal.Events = append(cal.Events, iCalendarEvent{
				Summary: h.Name,
				Date:    h.ParsedDate,
				UID:     fmt.Sprintf("%s_%s@hoiday-api.fyi", countryCode, h.ParsedDate.Format("20060102")),
			})
		}

		res.Header().Set("Content-Type", "text/calendar")
		fmt.Fprintln(res, cal.String())

	case "json":
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(outputSet)

	default:
		http.Error(res, fmt.Sprintf("Unknown format: %s", format), http.StatusBadRequest)
		return
	}
}

func handleReadme(res http.ResponseWriter, r *http.Request) {
	res.Write(indexHTML)
}
