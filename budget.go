package main

import (
	"github.com/picardrulez/lcars"
	"net/http"
	"strconv"
)

func viewBudgetHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Budget</h1>
	<br/>
	<br/>
	budget stuff goes here
	`
	amount, _ := getAmountValue("Netflix")
	date, _ := getDateValue("Netflix")
	website, _ := getWebsiteValue("Netflix")
	username, _ := getUsernameValue("Netflix")
	password, _ := getPasswordValue("Netflix")
	content = content + `
	<br/>
	|NETFLIX|
	<br/>
	Amount: ` + strconv.Itoa(amount) + `
	<br/>
	Date:  ` + strconv.Itoa(date) + `
	<br/>
	Website: ` + website + `
	<br/>
	Username: ` + username + `
	<br/>
	Password: ` + password + `
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
