package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/picardrulez/lcars"
	"log"
	"net/http"
)

func dropTablesHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Drop Databases</h1>
	<br/>
	<table><tr><td>Table</td><td></td></tr>
	<tr><td>Settings</td><td><FORM METHOD="post" action="/dropTablesProcessor"><input type="hidden" name="database" value="settings"><button type="submit">Drop</button></form></td></tr>
	<tr><td>Template</td><td><FORM METHOD="post" action="/dropTablesProcessor"><input type="hidden" name="database" value="template"><button type="submit">Drop</button></form></td></tr>
	<tr><td>Budget</td><td><FORM METHOD="post" action="/dropTablesProcessor"><input type="hidden" name="database" value="budget"><button type="submit">Drop</button></form></td></tr></table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func dropTablesProcessor(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	database := r.FormValue("database")
	if userName != "" {
		tabledrop := "DROP TABLE " + database
		database, _ := sql.Open("sqlite3", "./budget.db")
		if _, err := database.Exec(tabledrop); err != nil {
			log.Print(err)
		}
		database.Close()
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
