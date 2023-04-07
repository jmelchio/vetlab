package main

import (
	"flag"
	"fmt"

	"github.com/jmelchio/vetlab/vetlabcmd"
)

func main() {
	var dbHost, dbUser, dbPassword, dbName string
	var dbPort int
	var autoMigrate, useSSL bool

	flag.StringVar(&dbHost, "dbHost", "localhost", "Database hostname")
	flag.IntVar(&dbPort, "dbPort", 5432, "Database port")
	flag.StringVar(&dbUser, "dbUser", "postgres", "Database username")
	flag.StringVar(&dbPassword, "dbPassword", "password", "Database password")
	flag.StringVar(&dbName, "dbName", "vetlab", "Database name")
	flag.BoolVar(&autoMigrate, "autoMigrate", true, "Initiate or auto-migrate the schema")
	flag.BoolVar(&useSSL, "useSSL", false, "Use SSL connection to DB")

	flag.Parse()

	fmt.Printf("Starting the program with dbHost=%s, dbport=%d, dbUser=%s, dbPassword=%s, dbName=%s, autoMigrate=%v, useSSL=%v",
		dbHost, dbPort, dbUser, dbPassword, dbName, autoMigrate, useSSL)

	vetlabcmd.Run(dbHost, dbPort, dbUser, dbPassword, dbName, autoMigrate, useSSL)
}
