package utils

import (
	"time"

	"github.com/Cow4bunga/currency-exchange/services"
)

func StartScheduler() {
	ticker := time.NewTicker(24 * time.Hour)

	go func() {
		for {
			select {
			case <-ticker.C:
				services.LoadCurrencies()
			}
		}
	}()

	// Load immediately at application startup
	services.LoadCurrencies()
}
