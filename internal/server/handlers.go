package server

import (
	"encoding/json"
	"log"
	"net/http"
)

const logFormatString = "handleCalculatePost: %v\n"

func handleCalculatePost(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqAirports [][]string
		err := json.NewDecoder(r.Body).Decode(&reqAirports)
		if err != nil {
			log.Printf(logFormatString, err)

			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)

			return
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				log.Printf(logFormatString, err)
			}
		}()

		startEnd, err := s.FindStartEndAirports(reqAirports)
		if err != nil {
			log.Printf(logFormatString, err)

			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)

			return
		}

		if err := json.NewEncoder(w).Encode(startEnd); err != nil {
			log.Printf(logFormatString, err)

			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)

			return
		}
	}
}
