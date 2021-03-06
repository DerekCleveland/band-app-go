package rest

import (
	"band-app-go/pkg/insert"
	"band-app-go/pkg/listing"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func Handler(i insert.Service, l listing.Service) http.Handler {
	router := mux.NewRouter()

	// Insert handlers
	router.HandleFunc("/insertband", handleInsertBand(i))

	// Listing handlers
	router.HandleFunc("/listband", handleListBand(l))

	// Health check handlers
	router.HandleFunc("/healthz", handleHealthCheck)

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

func handleListBand(l listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Band name from URL params
		bandName := r.URL.Query().Get("Band")

		tmpBand := listing.Band{
			BandName: bandName,
		}

		band, err := l.SelectBand(tmpBand)
		if err != nil {
			log.Error().Msgf("Error listing band: %+v", err)
			http.Error(w, fmt.Sprintf("[ERROR] band: %+v not found", bandName), http.StatusInternalServerError)
			return
		}

		log.Info().Msgf("Listing band: %+v", band)

		// TODO response is including "band" in all the fields. This is because of the struct but even with the json flag its not removing
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", band)))
	}
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

		log.Info().Msgf("Inserting band: %+v", tempBand)
		i.InsertBand(tempBand)
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
	log.Info().Msg("responded - OK")
}
