package main

import (
	"github.com/picardrulez/lcars"
	"net/http"
	"strconv"
)

func viewBudgetHandler(w http.ResponseWriter, r *http.Request) {
	thisItem := getTemplateItem("Netflix")
	content := `
	<h1>Budget</h1>
	<br/>
	<br/>
	budget stuff goes here
	`
	content = content + `
	<br/>
	` + thisItem.Name + `
	<br/>
	Amount: ` + strconv.Itoa(thisItem.Amount) + `
	<br/>
	Date:  ` + strconv.Itoa(thisItem.Date) + `
	<br/>
	Website: ` + thisItem.Website + `
	<br/>
	Username: ` + thisItem.Username + `
	<br/>
	Password: ` + thisItem.Password + `
	<br/>
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

func isPayDay(int) bool {
	return true
}
