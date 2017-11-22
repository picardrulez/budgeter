package main

import (
	"github.com/picardrulez/lcars"
	"io"
	"net/http"
	"strconv"
)

type TemplateItem struct {
	Name   string
	Amount int
	Date   int
}

func viewTemplateHandler(w http.ResponseWriter, r *http.Request) {
	templateArray := templateList()
	content := `
	<h1>Template</h1>
	<br/>
	<table><tr><td>Name</td><td>Amount</td><td>Date</td></tr>
	`
	for _, k := range templateArray {
		content = content + "<tr><td>" + k.Name + "</td><td>" + strconv.Itoa(k.Amount) + "</td><td>" + strconv.Itoa(k.Date) + "</td></tr>"
	}
	content = content + "</table>"
	content = content + `
	<form method="post" action="/addtotemplate">
		<labelfor="name">Name</label>
		<input type="text" id="name" name="name"
		<label for="amount">Amount</label>
		<input type="text" id="amount" name="amount">
		<label for="date">Date</label>
		<input type="text" id="date" name="date">
		<button type="submit">Add</button>
	</form>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func addtotemplateHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	name := r.FormValue("name")
	amount, _ := strconv.Atoi(r.FormValue("amount"))
	date, _ := strconv.Atoi(r.FormValue("date"))
	if userName != "" {
		res := addToTemplate(name, amount, date)
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
	<table><tr><td>Name</td><td>Amount</td><td>Date<td></tr>
	`
	for _, k := range templateArray {
		content = content + "</tr><td>" + k.Name + `</td><td>
	<FORM METHOD="post" action="/edittemplateprocessor"
	<input type="hidden" name="name" value="` + k.Name + `">
		<input type="text" size="5" id="amount" name="amount" value="` + strconv.Itoa(k.Amount) + `"></td><td>
		<input type="text" size="5" id="date" name="date" value="` + strconv.Itoa(k.Date) + `"></td><td>
		<button type="submit">Submit</button>
		</form></td></tr>
		`
	}
	content = content + `
	</table>
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

//func edittemplateprocessor (w http.ResponseWriter, r *http.Request) {
//	userName := getUserName(r)
//	name := r.FormValue("name")
//	amount, _ := strconv.Atoi(r.FormValue("amount"))
//	date, _ := strconv.Atoi(r.FormValue("date"))
//	if userName != "' {
//		res := add
//	}'"
//}
