package db_connection

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// host     = "localhost"
	host     = "postgres" //service name for discovery
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "dac"
)

type Record struct {
	Index      int    `json:"index,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Cn         string `json:"cn,omitempty"`
	Certs      string `json:"cert,omitempty"`
	Flash_time string `json:"flash_time,omitempty"`
}

func Get_cert(cn string) ([]byte, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlStatement := `SELECT * FROM "dac-store" WHERE cn=$1;`
	var record Record
	row := db.QueryRow(sqlStatement, cn)
	err = row.Scan(&record.Index, &record.Mac, &record.Cn, &record.Certs, &record.Flash_time)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return nil, errors.New("no rows were returned")
	case nil:
		// fmt.Println(record.certs)
		// return record, nil
		// Return the record as JSON
		jsonData, jsonErr := json.Marshal(record)
		if jsonErr != nil {
			return nil, jsonErr
		}
		fmt.Println(string(jsonData))
		return jsonData, nil
	default:
		panic(err)
	}
}
