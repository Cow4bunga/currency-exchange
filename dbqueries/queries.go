package dbqueries

const (
	CreateCurrenciesTable = `
    CREATE TABLE IF NOT EXISTS currencies (
        id INT NOT NULL,
        date DATE NOT NULL,
        abbreviation VARCHAR(10) NOT NULL,
        official_rate DECIMAL(10, 4) NOT NULL,
        scale INT NOT NULL,
        name VARCHAR(255) NOT NULL,
        PRIMARY KEY (id, date) 
    );`

	InsertCurrency = `
    INSERT INTO currencies (id, date, abbreviation, official_rate, scale, name) 
    VALUES (?, ?, ?, ?, ?, ?);`

	SelectAllCurrencies = `
    SELECT * FROM currencies;`

	SelectCurrencyByDate = `
    SELECT * FROM currencies WHERE date = ?;`
)
