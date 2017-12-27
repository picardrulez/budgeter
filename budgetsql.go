package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
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
	settings, _ := getSettings()
	lastPayDate := getLastPayDay(settings)
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

func getCurrentBudget() ([]string, int) {
	var budgetList []string
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db")
		log.Printf("%s", err)
		db.Close()
		return budgetList, 1
	}

	rows, err := db.Query("SELECT name, ispaid FROM budget")
	if err != nil {
		log.Println("error querying db")
		log.Printf("%s", err)
		db.Close()
		return budgetList, 2
	}

	var name string
	var ispaid string

	for rows.Next() {
		err = rows.Scan(&name, &ispaid)
		if err != nil {
			log.Println("error scannign rows")
			log.Printf("%s", err)
			rows.Close()
			db.Close()
			return budgetList, 3
		}
		budgetList = append(budgetList, name+":"+ispaid)
	}
	rows.Close()
	db.Close()
	return budgetList, 0
}

func removePaidItems() int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("delete from budget where ispaid='true'")
	if err != nil {
		log.Println("error preparing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := stmt.Exec()
	if err != nil {
		log.Println("error executing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows deleted from budget")
	db.Close()
	return 0
}

func budgetInsert(newBudget []string) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	for _, k := range newBudget {
		stmt, err := db.Prepare("INSERT INTO budget(name, ispaid) values(?,?)")
		if err != nil {
			log.Println("error preparing insert statement")
			log.Printf("%s", err)
			db.Close()
			return 2
		}
		_, err = stmt.Exec(k, "false")
		if err != nil {
			log.Println("error executing insert statement")
			log.Printf("%s", err)
			db.Close()
			return 3
		}
	}
	db.Close()
	return 0
}

func isBudgetEmpty() bool {
	//database, _ := sql.Open("sqlite3", "./budget.db")
	//budgetrows, err := database.Query("SELECT * from template")
	//if err != nil {
	//	log.Println("error running select against budget")
	//	log.Printf("%s", err)
	//	return false
	//}
	//budgetcount := rowCounter(budgetrows)
	budgetcount := newRowCounter("budget")
	if budgetcount > 0 {
		return false
	} else {
		return true
	}
}

func payItem(item string, pay int) {
	db, err := sql.Open("sqlite3", "./budget.db")
	var mybool string
	if pay == 0 {
		mybool = "false"
	} else if pay == 1 {
		mybool = "true"
	}

	if err != nil {
		log.Println("error opening db for pay update")
		log.Printf("%s", err)
		db.Close()
		return
	}
	stmt, err := db.Prepare("update budget set ispaid = ? where name = ?")
	if err != nil {
		log.Println("error preparing pay update statement")
		log.Printf("%s", err)
		db.Close()
		return
	}
	res, err := stmt.Exec(mybool, item)
	if err != nil {
		log.Println("error updating paid item")
		log.Printf("%s", err)
		db.Close()
		return
	}
	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return

}
