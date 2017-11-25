package main

import (
	"github.com/picardrulez/lcars"
	"net/http"
	"strconv"
)

type Settings struct {
	PeriodLength int
	PeriodFormat string
	StartDate    string
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Settings</h1>
	<br/>
	<table<tr><td>Period Length</td><td>Period Format</td><td>Start Date</td><td></td></tr>
	<form method="post" action="/updatesettings">
	<tr><td><input type="text" size="5" id="periodlength" name="periodlength" value="` + strconv.Itoa(PERIODLENGTH) + `"></td><td>
	<input type="text" size="5" id="periodformat" name="periodformat" value="` + PERIODFORMAT + `"></td><td>
	<input type="text" size="5" id="startdate" name="startdate" value="` + strconv.Itoa(STARTDATE[0]) + "/" + strconv.Itoa(STARTDATE[1]) + "/" + strconv.Itoa(STARTDATE[2]) + `"></td><td>
	<button type="submit">Update</button></td></tr>
	</form>
	</table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func updatesettingsHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	periodlength, _ := strconv.Atoi(r.FormValue("periodlength"))
	periodformat := r.FormValue("periodformat")
	startdate := r.FormValue("startdate")

	if userName != "" {
		http.Redirect(w, r, "/settings", 302)
	} else {
		PERIODLENGTH = periodlength
		PERIODFORMAT = periodformat
		STARTDATE[0] = startdate[0]
		STARTDATE[1] = startdate[1]
		STARTDATE[2] = startdate[2]
		http.Redirect(w, r, "/", 302)
	}
}
