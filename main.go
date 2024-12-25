package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Cow4bunga/currency-exchange/config"
	"github.com/Cow4bunga/currency-exchange/routes"
	"github.com/Cow4bunga/currency-exchange/services"
	"github.com/Cow4bunga/currency-exchange/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var portNumber int = 8080

func loadConfig() (config.Config, error) {
	var cfg config.Config
	data, err := os.ReadFile("config.json")
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}
	defer db.Close()

	services.InitDatabase(db)
	services.CreateTableIfNotExists()
	utils.StartScheduler()

	router := mux.NewRouter()
	routes.SetupRoutes(router)

	log.Printf("Server running on port %d\n", portNumber)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), router); err != nil {
		log.Fatal("Server did not start: ", err)
	}
}
