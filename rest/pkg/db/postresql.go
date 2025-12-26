package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
)

// DBConnect establishes a connection to the database and returns a sql.DB instance
func DBConnect(db *sql.DB, userName string, password string, dbHostName string, dbPort string, dbName string) (*sql.DB, error) {
	// TODO: Implement database connection logic here
	var err error
	if db == nil {
		dsn := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			userName,
			password,
			dbHostName,
			dbPort,
			dbName,
		)

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %v", err)
		}
	}

	return db, nil
} // func DBConnect

// TableExists checks if a table with the given name exists in the database
func TableExists(db *sql.DB, tableName string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM information_schema.tables
		WHERE table_name = $1
	`
	var count int
	err := db.QueryRow(query, tableName).Scan(&count) // Query the db for the tableName

	// If error is returned, then return false with the error
	if err != nil {
		return false, err
	}
	// Return true if a result on tableName was found with count > 0
	return count > 0, nil
} // func TableExists

func TableInit(db *sql.DB, tableName string) error {
	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
    id          SERIAL PRIMARY KEY,
	soda  		VARCHAR(50)  NOT NULL UNIQUE,
	amount   	BIGINT  NOT NULL CHECK (amount >= 0)
	);`, tableName)

	// Execute the table creation query
	result, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table %s: %v", tableName, err)
	}
	fmt.Printf("tableInit succeeded with the following result: %s", result)
	return nil
} // func TableInit

func TableAddSoda(db *sql.DB, tableName string, soda string, amount uint) (int, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (soda, amount)
		VALUES ($1, $2)
		RETURNING id, soda, amount;
	`, tableName)
	var id int
	err := db.QueryRow(query, soda, amount).Scan(&id, &soda, &amount)
	if err != nil {
		return 0, fmt.Errorf("failed to add soda to table %s: %v", tableName, err)
	}
	return id, nil
} // func TableAddamount

func TableGetSoda(db *sql.DB, tableName string, id int) (*sql.Rows, error) {
	query := fmt.Sprintf(`
		SELECT id, soda, amount
		FROM %s
		WHERE id = $1;
	`, tableName)

	response, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get amount from table %s: %v", tableName, err)
	}
	return response, nil
} // func TableGetamount

func TableUpdateSodaAmount(db *sql.DB, tableName string, id int, soda string, amount uint) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET amount = $1
		WHERE id = $2;
	`, tableName)
	_, err := db.Exec(query, amount, id)
	return err
} // func TableUpdateamount

func TableDeleteSoda(db *sql.DB, tableName string, id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1;`, tableName)
	_, err := db.Exec(query, id)
	return err
} // func TableDeleteSoda

func DBClose(db *sql.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
} // func DBClose
