package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

func getTemplateItem(name string) TemplateItem {
	amount, _ := getAmountValue(name)
	date, _ := getDateValue(name)
	website, _ := getWebsiteValue(name)
	username, _ := getUsernameValue(name)
	password, _ := getPasswordValue(name)
	returnItem := TemplateItem{Name: name, Amount: amount, Date: date, Website: website, Username: username, Password: password}
	return returnItem
}

func getBudgetList() ([]string, int) {
	lastPayDate := getLastPayDay()
	nextPayDate := getNextPayDay()
	lastArray := strings.Split(lastPayDate, "-")
	nextArray := strings.Split(nextPayDate, "-")
	lastPayDay := lastArray[2]
	nextPayDay := nextArray[2]

	var budgetList []string
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db")
		log.Printf("%s", err)
		db.Close()
		return budgetList, 1
	}

	rows, err := db.Query("SELECT name FROM template where date >= " + lastPayDay + " AND date < " + nextPayDay)
	if err != nil {
		log.Println("error querying db")
		log.Printf("%s", err)
		db.Close()
		return budgetList, 2
	}

	var name string

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println("error scanning rows")
			log.Printf("%s", err)
			rows.Close()
			db.Close()
			return budgetList, 3
		}
		budgetList = append(budgetList, name)
	}
	rows.Close()
	db.Close()
	return budgetList, 0
}
