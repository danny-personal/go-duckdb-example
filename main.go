package main

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM read_csv_auto('hoge.csv')")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Print the results
	for rows.Next() {
		var column0 string
		var column1 int
		var column2 string

		err = rows.Scan(&column0, &column1, &column2)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("column0: %s\n", column0)
	}
}
