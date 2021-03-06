package main

import (
	"github.com/picardrulez/lcars"
	"io"
	"net/http"
	"strconv"
)

type TemplateItem struct {
	Name     string
	Amount   int
	Date     int
	Website  string
	Username string
	Password string
}

func viewTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateArray := templateList()
	content := `
	<h1>Template</h1>
	<br/>
	<table><tr><td>Name</td><td>Amount</td><td>Date</td><td>Website</td><td>Username</td><td>Password</td><td>          </td></tr>
	`
	for _, k := range templateArray {
		content = content + "<tr><td>" + k.Name + "</td><td>" + strconv.Itoa(k.Amount) + "</td><td>" + strconv.Itoa(k.Date) + "</td><td>" + k.Website + "</td><td>" + k.Username + "</td><td>" + k.Password + "</td><td></td></tr>"
	}
	content = content + `
	<form method="post" action="/addtotemplate">
		<tr><td><label for="name">Name</label>
		<input type="text" size="15" id="name" name="name"</td><td>
		<label for="amount">Amount</label>
		<input type="text" size="5" id="amount" name="amount"></td><td>
		<label for="date">Date</label>
		<input type="text" size="5" id="date" name="date"></td><td>
		<label for="website">Website</label>
		<input type="text" size="20" id="website" name="website"></td><td>
		<label for="username">Username</label>
		<input type="text" size="20" id="username" name="username"></td><td>
		<label for="password">Password</label>
		<input type="text" size="20" id="password" name="password"></td><td>
		<button type="submit">Add</button></td></tr>
	</form>
	</table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func addtotemplateHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	name := r.FormValue("name")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	date, _ := strconv.Atoi(r.FormValue("date"))
	website := r.FormValue("website")
	username := r.FormValue("username")
	password := r.FormValue("password")
	templateitem := TemplateItem{Name: name, Amount: amount, Date: date, Website: website, Username: username, Password: password}
	if userName != "" {
		res := addToTemplate(templateitem)
		if res != 0 {
			io.WriteString(w, "an erro roccured adding item")
		}
		http.Redirect(w, r, "/viewTemplate", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func editTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateArray := templateList()
	content := `
	<h1>Edit Template</h1>
	<br/>
	<table><tr><td>Name</td><td>Amount</td><td>Date</td><td>Website</td><td>Username</td><td>Password</td></tr>
	`
	for _, k := range templateArray {
		content = content + "<tr><td>" + k.Name + `</td><td>
	<FORM METHOD="post" action="/edittemplateprocessor">
	<input type="hidden" name="name" value="` + k.Name + `">
		<input type="text" size="5" id="amount" name="amount" value="` + strconv.Itoa(k.Amount) + `"></td><td>
		<input type="text" size="5" id="date" name="date" value="` + strconv.Itoa(k.Date) + `"></td><td>
		<input type="text" size="15" id="website" name="website" value="` + k.Website + `"></td><td>
		<input type="text" size="15" id="username" name="username" value="` + k.Username + `"></td><td>
		<input type="text" size="15" id="password" name="password" value="` + k.Password + `"></td><td>
		<button type="submit">Submit</button>
		</form></td><td>
		<FORM METHOD="post" action="/deletetemplateprocessor">
		<input type="hidden" name="name" value="` + k.Name + `">
		<button type="submit">Delete</button>
		</form></td></tr>
		`
	}
	content = content + `
	</table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func edittemplateprocessor(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	name := r.FormValue("name")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	date, _ := strconv.Atoi(r.FormValue("date"))
	website := r.FormValue("website")
	username := r.FormValue("username")
	password := r.FormValue("password")
	templateitem := TemplateItem{Name: name, Amount: amount, Date: date, Website: website, Username: username, Password: password}
	if userName != "" {
		res := updateTemplate(templateitem)
		if res != 0 {
			io.WriteString(w, "an error occured updating item "+name)
		}
		http.Redirect(w, r, "/editTemplate", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func deletetemplateprocessor(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	name := r.FormValue("name")
	if userName != "" {
		res := deleteTemplate(name)
		if res != 0 {
			io.WriteString(w, "an error occured deleting item "+name)
		}
		http.Redirect(w, r, "/editTemplate", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
