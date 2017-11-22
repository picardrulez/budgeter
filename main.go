package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/picardrulez/lcars"
	"log"
	"net/http"
	"os"
	"sort"
)

var VERSION = "v0.5.5"
var LOGFILE string = "/var/log/budget"
var DEFAULTAUTH = "c@me0c@nd1"
var mymenu = lcars.Menu{Items: []string{"/viewTemplate|View Template", "/editTemplate|Edit Template", "/viewBudget|View Budget", "/editBudget|Edit Budget", "/createUser|Create User", "/changePassword|Change Password"}}
var lcarssettings = lcars.Settings{Title: "Budgeter", TopColor: "dodger-blue-alt", BottomColor: "hopbush", MenuColor: "tan", Menu: true}

//create cookie generator
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

//create router
var router = mux.NewRouter()

func main() {
	//log stuff
	var _, err = os.Stat(LOGFILE)
	if os.IsNotExist(err) {
		var file, err = os.Create(LOGFILE)
		checkError(err)
		defer file.Close()
	}
	f, err := os.OpenFile(LOGFILE, os.O_WRONLY|os.O_APPEND, 0644)
	checkError(err)
	defer f.Close()
	log.SetOutput(f)
	log.Println("-----------------------------------")
	startup()
	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/main", mainHandler)
	router.HandleFunc("/viewTemplate", viewTemplateHandler)
	router.HandleFunc("/editTemplate", editTemplateHandler)
	router.HandleFunc("/edittemplateprocessor", edittemplateprocessor).Methods("POST")
	router.HandleFunc("/deletetemplateprocessor", deletetemplateprocessor).Methods("POST")
	router.HandleFunc("/addtotemplate", addtotemplateHandler).Methods("POST")
	router.HandleFunc("/viewBudget", viewBudgetHandler)
	router.HandleFunc("/editBudget", editBudgetHandler)
	router.HandleFunc("/createUser", createUserHandler)
	router.HandleFunc("/usercreation", usercreationHandler)

	http.Handle("/", router)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	log.Println("budgeter " + VERSION + " started")
	http.ListenAndServe(":8000", nil)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	content := `
	<h1>Budget Stuff</h1>
	<br/>
	<br/>
	blah blah blah
	`
	t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
	t.Execute(w, createPage)
}

func alphabetizeList(memberList []string) []string {
	sort.Strings(memberList)
	return memberList
}
