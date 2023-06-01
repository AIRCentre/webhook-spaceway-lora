package mysqldriver

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type mysqlDriver struct {
	db *sql.DB
}

func New(username, password, hostname, port, database string) (*mysqlDriver, error) {
	// Create a new TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Set the TLS config in the MySQL driver
	mysql.RegisterTLSConfig("custom", tlsConfig)

	// Setup connection URI
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=custom", username, password, hostname, port, database)

	// Connect to the MySQL server with TLS
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	return &mysqlDriver{db}, err
}

func (d mysqlDriver) Query(query string, args ...interface{}) ([]map[string][]byte, error) {
	rows, err := d.db.Query(query, args...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	mapSlice, err := rowsToMaps(rows)
	if err != nil {
		return nil, err
	}

	return mapSlice, nil
}

func (d mysqlDriver) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func rowsToMaps(rows *sql.Rows) ([]map[string][]byte, error) {
	// Get the column names from the rows
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	// Make a slice to hold the result
	result := make([]map[string][]byte, 0)

	// Make a slice to hold the column values
	values := make([][]byte, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Loop through the rows
	for rows.Next() {
		// Scan the row into the values slice
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		// Make a map to hold the column name/value pairs
		rowMap := make(map[string][]byte)

		// Loop through the values and add them to the map
		for i, value := range values {
			rowMap[columns[i]] = value
		}

		// Add the map to the result slice
		result = append(result, rowMap)
	}

	return result, nil
}
