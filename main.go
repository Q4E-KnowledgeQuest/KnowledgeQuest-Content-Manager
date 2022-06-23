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
	"github.com/pterm/pterm"
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

	banner()

	server.Start(*portPtr, *logPtr)

}

func banner() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("KnowledgeQuest: Content Manager"))
	pterm.Info.Println("(c)2022 by Akhil Datla")
}
