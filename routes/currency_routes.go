package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Cow4bunga/currency-exchange/services"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")

		log.Printf("Received request for /currencies with date: %s\n", date)

		var currencies interface{}
		var err error

		if date == "" {
			currencies, err = services.GetAllCurrencies()
		} else {
			currencies, err = services.GetCurrencyByDate(date)
		}

		if err != nil {
			log.Println("Error fetching currencies:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Successfully retrieved currencies: %+v\n", currencies)

		if err := json.NewEncoder(w).Encode(currencies); err != nil {
			log.Println("Error encoding currencies to JSON:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}).Methods("GET")
}
