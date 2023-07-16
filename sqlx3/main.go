package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// open and connect at the same time, panicing on error
	var db *sqlx.DB = sqlx.MustConnect("sqlite3", ":memory:")

	// force a connection and test that it worked
	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	schema := `CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer);`

	// execute a query on the server
	result, err := db.Exec(schema)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// or, you can use MustExec, which panics on error
	cityState := `INSERT INTO place (country, telcode) VALUES (?, ?)`
	countryCity := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	db.MustExec(cityState, "Hong Kong", 852)
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johannesburg", 27)

	// // fetch all places from the db
	// rows, err := db.Query("SELECT country, city, telcode FROM place")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // iterate over each row
	// for rows.Next() {
	// 	var country string
	// 	// note that city can be NULL, so we use the NullString type
	// 	var city sql.NullString
	// 	var telcode int
	// 	err = rows.Scan(&country, &city, &telcode)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(country, city, telcode)
	// }

	type Place struct {
		Country       string
		City          sql.NullString
		TelephoneCode int `db:"telcode"`
	}

	// rows, err := db.Queryx("SELECT * FROM place")
	// for rows.Next() {
	// 	var p Place
	// 	err = rows.StructScan(&p)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Printf("%#v\n", p)
	// }
	// // check the error from rows
	// err = rows.Err()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	p := Place{}
	pp := []Place{}

	// this will pull the first place directly into p
	err = db.Get(&p, "SELECT * FROM place LIMIT 1")
	fmt.Printf("%#v\n", p)
	// this will pull places with telcode > 50 into the slice pp
	err = db.Select(&pp, "SELECT * FROM place WHERE telcode > ?", 50)
	fmt.Printf("%#v\n", pp)
	// they work with regular types as well
	var id int
	err = db.Get(&id, "SELECT count(*) FROM place")
	fmt.Printf("%#v\n", id)
	// fetch at most 10 place names
	var names []string
	err = db.Select(&names, "SELECT name FROM place LIMIT 10")
	fmt.Printf("%#v\n", names)
}
