package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
)

var (
	dbProtocol         = "tcp"
	dbAddress          = "127.0.0.1"
	dbUser             = "root"
	dbPassword         = "password"
	dbName             = "testdb"
	dbTable            = "testtb"
	dumpDir            = "dumps"
	dumpFilenameFormat = fmt.Sprintf("%s-060102T150405", dbName)
)

func main() {

	// Create  a new config
	config := mysql.NewConfig()

	// Set the connection parameters
	config.Net = dbProtocol
	config.Addr = dbAddress
	config.User = dbUser
	config.Passwd = dbPassword
	config.DBName = dbName

	// Create a new connection
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Select all results from the testdb database and testb table
	query := fmt.Sprintf("SELECT * FROM %s", dbTable)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate through the results
	records := []string{}
	var record string
	for rows.Next() {
		err := rows.Scan(&record)
		if err != nil {
			fmt.Println(err)
			return
		}
		records = append(records, record)
	}
	fmt.Println(records)

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}

	// Dump database to file
	dumpError := dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", dumpError)
		return
	}
}
