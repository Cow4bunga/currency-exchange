package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Cow4bunga/currency-exchange/services"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")

		var currencies interface{}

		var err error
		if date == "" {
			currencies, err = services.GetAllCurrencies()
		} else {
			currencies, err = services.GetCurrencyByDate(date)
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(currencies)
	}).Methods("GET")
}
