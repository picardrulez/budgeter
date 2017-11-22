package main

import (
	"github.com/picardrulez/lcars"
	"net/http"
)

func viewBudgetHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Budget</h1>
	<br/>
	<br/>
	budget stuff goes here
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func editBudgetHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Edit Budget</h1>
	<br/>
	<br/>
	budget edit stuff goes here
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}
