package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/picardrulez/lcars"
	"log"
	"net/http"
	"strconv"
)

type Settings struct {
	PeriodLength  int
	PeriodFormat  string
	StartDate     string
	CurrentPayDay string
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	settings, _ := getSettings()
	content := `
	<h1>Settings</h1>
	<br/>
	<table<tr><td>Period Length</td><td>Period Format</td><td>Start Date</td><td></td></tr>
	<form method="post" action="/updatesettings">
	<tr><td><input type="text" size="5" id="periodlength" name="periodlength" value="` + strconv.Itoa(settings.PeriodLength) + `"></td><td>
	<input type="text" size="5" id="periodformat" name="periodformat" value="` + settings.PeriodFormat + `"></td><td>
	<input type="date" size="5" id="startdate" name="startdate" value="` + settings.StartDate + `"></td><td>
	<button type="submit">Update</button></td></tr>
	</form>
	</table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func updatesettingsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("in updatesettingsHandler")
	userName := getUserName(r)
	periodlength, _ := strconv.Atoi(r.FormValue("periodlength"))
	periodformat := r.FormValue("periodformat")
	startdate := r.FormValue("startdate")
	settings := Settings{PeriodLength: periodlength, PeriodFormat: periodformat, StartDate: startdate}

	log.Println("checking if user is logged in")
	if userName != "" {
		log.Println("user is logged in, updating settings")
		updateSettings(settings)
		log.Println("running temp settings select")
		db, err := sql.Open("sqlite3", "./budget.db")
		if err != nil {
			log.Println("error opening db for insert")
			log.Printf("%s", err)
			db.Close()
			return
		}
		rows, err := db.Query("SELECT * from settings")
		if err != nil {
			log.Println("error running query")
			log.Printf("%s", err)
			db.Close()
			return
		}
		var periodLength int
		var periodFormat string
		var startDate string
		var currentPayDay string

		for rows.Next() {
			err = rows.Scan(&periodLength, &periodFormat, &startDate, &currentPayDay)
			if err != nil {
				log.Println("error scanning rows")
				log.Printf("%s", err)
				return
			}
			log.Println(strconv.Itoa(periodLength))
			log.Println(periodFormat)
			log.Println(startDate)
		}
		http.Redirect(w, r, "/", 302)
	} else {
		log.Println("user not logged in ")
		http.Redirect(w, r, "/settings", 302)
	}
}
