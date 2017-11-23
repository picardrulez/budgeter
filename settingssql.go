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

func updateSettings(settings Settings) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("update settings set periodlength = ?, periodformat = ?, startdate = ?")
	if err != nil {
		log.Println("error preparing ")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := stmt.Exec(settings.PeriodLength, settings.PeriodFormat, settings.StartDate)
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