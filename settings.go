package main

import (
	"github.com/picardrulez/lcars"
	"io"
	"net/http"
	"strconv"
)

type Settings struct {
	PeriodLength int
	PeriodFormat string
	StartDate    string
}

func settingsHandler(w http.ResponseWriter, r *http.Request) {
	settings, _ := getSettings()
	content := `
	<h1>Settings</h1>
	<br/>
	<table<tr><td>Period Length</td><td>Period FOrmat</td><td>Start Date</td><td></td></tr>
	<form method="post" action="/updatesettings">
	<tr><td><input type="text" size="5" id="periodlength" name="periodlength" value="` + strconv.Itoa(settings.PeriodLength) + `"></td><td>
	<input type="text" size="5" id="periodformat" name="periodformat" value="` + settings.PeriodFormat + `"></td><td>
	<input type="text" size="5" id="startdate" name="startdate" value="` + settings.StartDate + `"></td><td>
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

	settings := Settings{PeriodLength: periodlength, PeriodFormat: periodformat, StartDate: startdate}
	if userName != "" {
		res := updateSettings(settings)
		if res != 0 {
			io.WriteString(w, "an error occured updating settings")
		}
		http.Redirect(w, r, "/settings", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
