package main

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go index.html

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	holidays "github.com/Luzifer/go-holidays"
	"github.com/Luzifer/rconfig"
	"github.com/gorilla/mux"
)

var (
	cfg = struct {
		Listen         string `flag:"listen" default:":3000" description:"IP/Port to listen on"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

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
	r.HandleFunc("/{country-code:[a-z-]+}/{year:[0-9]{4}}", handleHolidays)
	r.HandleFunc("/{country-code:[a-z-]+}", handleHolidays)
	r.HandleFunc("/", handleReadme)
	http.ListenAndServe(cfg.Listen, r)
}

func handleHolidays(res http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var (
		countryCode = vars["country-code"]
		year        = time.Now().Year()
	)

	if y, ok := vars["year"]; ok {
		var err error
		if year, err = strconv.Atoi(y); err != nil {
			http.Error(res, "You need to specify the year as a 4 character number", http.StatusBadRequest)
			return
		}
	}

	holidays, err := holidays.GetHolidays(countryCode, year)
	if err != nil {
		http.Error(res, "An error ocurred: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(holidays)
}

func handleReadme(res http.ResponseWriter, r *http.Request) {
	readme, _ := Asset("index.html")
	res.Write(readme)
}
