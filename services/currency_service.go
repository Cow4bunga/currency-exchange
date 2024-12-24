package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cow4bunga/currency-exchange/dbqueries"
	"github.com/Cow4bunga/currency-exchange/models"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDatabase(database *sqlx.DB) {
	db = database
}

func CreateTableIfNotExists() {
	_, err := db.Exec(dbqueries.CreateCurrenciesTable)
	if err != nil {
		fmt.Println("Error creating table:", err)
	}
}

func LoadCurrencies() {
	resp, err := http.Get("https://api.nbrb.by/exrates/rates?periodicity=0")
	if err != nil {
		fmt.Println("Error while fetching currencies:", err)
		return
	}
	defer resp.Body.Close()

	var currencies []models.Currency
	if err := json.NewDecoder(resp.Body).Decode(&currencies); err != nil {
		fmt.Println("Error decoding currencies JSON:", err)
		return
	}

	fmt.Println("Fetched currencies:", currencies)

	for _, currency := range currencies {
		fmt.Printf("Inserting currency: ID=%d, Date=%s, Abbreviation=%s, Rate=%.4f, Scale=%d, Name=%s\n",
			currency.ID, currency.Date, currency.Abbreviation, currency.OfficialRate, currency.Scale, currency.Name)
		_, err := db.Exec(dbqueries.InsertCurrency,
			currency.ID,
			currency.Date,
			currency.Abbreviation,
			currency.OfficialRate,
			currency.Scale,
			currency.Name)
		if err != nil {
			fmt.Println("Error while inserting currency data:", err)
		}
	}
}

func GetAllCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	err := db.Select(&currencies, dbqueries.SelectAllCurrencies)
	if err != nil {
		fmt.Println("Error fetching currencies:", err)
		return nil, err
	}
	return currencies, nil
}

func GetCurrencyByDate(date string) ([]models.Currency, error) {
	var currencies []models.Currency
	err := db.Select(&currencies, dbqueries.SelectCurrencyByDate, date)
	if err != nil {
		fmt.Println("Error fetching currencies by date:", err)
		return nil, err
	}
	return currencies, err
}
