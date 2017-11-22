package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func getSettings() (Settings, bool) {
	var settings Settings
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db")
		log.Printf("%s", err)
		return settings, true
	}
	stmt, err := db.Prepare("SELECT periodlength, periodformat, startdate FROM settings")
	if err != nil {
		log.Println("error preparing select settings")
		log.Printf("%s", err)
		db.Close()
		return settings, true
	}

	defer stmt.Close()
	err = stmt.QueryRow().Scan(&settings.PeriodLength, &settings.PeriodFormat, &settings.StartDate)
	if err != nil {
		log.Println("error scanning row for settings")
		log.Printf("%s", err)
		db.Close()
		return settings, true
	}
	db.Close()
	return settings, false
}
