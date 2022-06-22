/*
 * File: main.go
 * File Created: Wednesday, 22nd June 2022 12:14:20 am
 * Author: Akhil Datla
 * © 2022 Akhil Datla
 */

package main

import (
	"flag"
	"main/components/dbmanager"
	"main/server"
)

func main() {
	dbnamePtr := flag.String("dbname", "contentrepository.db", "name of the database")
	portPtr := flag.Int("port", 8080, "port to listen on")
	logPtr := flag.Bool("log", false, "enable logging")

	flag.Parse()

	err := dbmanager.Open(*dbnamePtr)
	if err != nil {
		panic(err)
	}

	defer dbmanager.Close()

	server.Start(*portPtr, *logPtr)

}