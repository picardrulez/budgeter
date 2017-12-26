package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

func getSettings() (Settings, bool) {
	var settings Settings
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db")
		log.Printf("%s", err)
		return settings, true
	}
	stmt, err := db.Prepare("SELECT * FROM settings")
	if err != nil {
		log.Println("error preparing select settings")
		log.Printf("%s", err)
		db.Close()
		return settings, true
	}

	var periodlength int
	var periodformat string
	var startdate string
	var currentpayday string

	err = stmt.QueryRow().Scan(&periodlength, &periodformat, &startdate, &currentpayday)
	if err != nil {
		log.Println("error scanning row for settings")
		log.Printf("%s", err)
		db.Close()
		defer stmt.Close()
		return settings, true
	}
	db.Close()
	defer stmt.Close()
	settings = Settings{PeriodLength: periodlength, PeriodFormat: periodformat, StartDate: startdate, CurrentPayDay: currentpayday}
	return settings, false
}

func updateSettings(settings Settings) int {
	log.Println("updating currentpayday")
	settings.CurrentPayDay = getLastPayDay(settings)
	log.Println("CurrentPayDay is " + settings.CurrentPayDay)
	log.Println("opening budget.db")
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	log.Println("preparing statement")
	stmt, err := db.Prepare("update settings set periodlength = ?, periodformat = ?, startdate = ?, currentpayday = ?")
	if err != nil {
		log.Println("error preparing ")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	log.Println("executing statement")
	res, err := stmt.Exec(settings.PeriodLength, settings.PeriodFormat, settings.StartDate, settings.CurrentPayDay)
	if err != nil {
		log.Println("error updating")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return 0
}
