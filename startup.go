package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func startup() int {
	database, _ := sql.Open("sqlite3", "./budget.db")
	usersTableCreate, err := database.Prepare("CREATE TABLE IF NOT EXISTS users (username TEXT, password TEXT)")
	if err != nil {
		log.Println("error preparing create users table statement")
		log.Printf("%s", err)
		return 1
	}
	usersTableCreate.Exec()

	templateTableCreate, err := database.Prepare("CREATE TABLE IF NOT EXISTS template (name TEXT, amount INT, date INT, website TEXT, username TEXT, password TEXT)")
	if err != nil {
		log.Println("error preparing create template table statement")
		log.Printf("%s", err)
		return 1
	}
	templateTableCreate.Exec()

	budgetTableCreate, err := database.Prepare("CREATE TABLE IF NOT EXISTS budget (name TEXT, ispaid BOOL default 0)")
	if err != nil {
		log.Println("error preparing create budget table statement")
		log.Printf("%s", err)
		return 1
	}
	budgetTableCreate.Exec()

	settingsTableCreate, err := database.Prepare("CREATE TABLE IF NOT EXISTS settings (periodlength INT, periodformat TEXT, startdate TEXT, currentpayday TEXT)")
	if err != nil {
		log.Println("error preparing create settings tale statement")
		log.Printf("%s", err)
		return 1
	}
	settingsTableCreate.Exec()

	settingsrows, err := database.Query("SELECT * from settings")
	if err != nil {
		log.Println("error running select against settings")
		log.Printf("%s", err)
		return 1
	}
	settingscount := rowCounter(settingsrows)
	if settingscount < 1 {
		settingsInsert, err := database.Prepare("INSERT INTO settings VALUES(1, 'Days', '01-01-2017', '01-01-2017')")
		if err != nil {
			log.Println("error preparing db insert for settings")
			log.Printf("%s", err)
			return 1
		}
		settingsInsert.Exec()
	}

	if userExists("admin") != true {
		hashedpass := encryptPassword(DEFAULTAUTH)
		insertUser("admin", hashedpass)
	} else {
	}

	adminpass, _ := getPassword("admin")
	passres := checkPass(DEFAULTAUTH, adminpass)
	if passres != true {
		log.Println("admin pass does not match, resetting pass")
		hashedpass := encryptPassword(DEFAULTAUTH)
		log.Println("updating admin password")
		passres := updatePassword("admin", hashedpass)
		if passres != 0 {
			log.Println("an error occured creating admin password")
			return 2
		}
	} else {
	}
	userlist, _ := getUserList()
	for _, k := range userlist {
		log.Println(k)
	}
	return 0
}

func budgetCheck() {
	if isOutOfPayPeriod() {
		createreturn := createNewPayPeriod()
		if createreturn > 0 {
			log.Println("error creating new pay period")
		}
	}
}
