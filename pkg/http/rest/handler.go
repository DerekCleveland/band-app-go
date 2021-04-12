package rest

import (
	"band-app-go/pkg/insert"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func Handler(i insert.Service) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/insertband", handleInsertBand(i))

	// Debugging and profiling
	//router.HandleFunc("/debug/pprof/", pprof.Index)
	//router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//router.HandleFunc("/debug/pprof/trace", pprof.Trace)
	//router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//router.HandleFunc("debug/pprof/index", pprof.Index)
	//router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	//router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//router.Handle("/debug/pprof/block", pprof.Handler("block"))

	return router
}

func handleInsertBand(i insert.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Band name from URL params
		bandName := r.URL.Query().Get("Band")
		bandRating, err := strconv.ParseFloat(r.URL.Query().Get("Rating"), 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "[FATAL] error reading and converting rating form value")
			return
		}
		bandGenre := r.URL.Query().Get("Genre")

		tempBand := insert.Band{
			BandName:   bandName,
			BandRating: bandRating,
			BandGenre:  bandGenre,
		}

		err = i.CheckIfBandExists(tempBand)
		if err == insert.ErrDuplicate {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "[ERROR] band already exists")
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "[FATAL] error reading and converting rating form value")
			return
		}

		log.Info("Inserting band:", tempBand)
		i.InsertBand(tempBand)
	}
}
