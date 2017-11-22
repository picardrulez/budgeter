package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

func insertUser(username string, password string) int {
	log.Println("inserting username: " + username + "password: " + password)
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("INSERT INTO users(username, password) values(?,?)")
	if err != nil {
		log.Println("error preparing insert statement")
		log.Printf("%s", err)
		db.Close()
		return 2
	}
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Println("error executing insert statement")
		log.Printf("%s", err)
		db.Close()
		return 3
	}
	db.Close()
	return 0
}

func getUserList() ([]string, int) {
	var userlist []string
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db")
		log.Printf("%s", err)
		db.Close()
		return userlist, 1
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("error querying db")
		log.Printf("%s", err)
		db.Close()
		return userlist, 2
	}
	var username string
	var password string

	for rows.Next() {
		err = rows.Scan(&username, &password)
		if err != nil {
			log.Println("error scannign rows")
			log.Printf("%s", err)
			rows.Close()
			db.Close()
			return userlist, 3
		}
		userlist = append(userlist, username)
	}
	rows.Close()
	db.Close()
	return userlist, 0
}

func getPassword(username string) (string, int) {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return "", 1
	}
	stmt, err := db.Prepare("SELECT password FROM users where username = ?")
	if err != nil {
		log.Println("error preparing password select statement")
		log.Printf("%s", err)
		db.Close()
		return "", 2
	}
	defer stmt.Close()
	var password string
	err = stmt.QueryRow(username).Scan(&password)
	if err != nil {
		log.Println("error scanning row for password")
		log.Printf("%s", err)
		db.Close()
		return "", 3
	}
	db.Close()
	return password, 0
}

func updatePassword(username string, password string) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	userupdate, err := db.Prepare("update users set password=? where username=?")
	if err != nil {
		log.Println("error preparing")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := userupdate.Exec(password, username)
	if err != nil {
		log.Println("error updating")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return 0
}

func deleteUser(username string) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("delete from users where username=?")
	if err != nil {
		log.Println("error preparing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := stmt.Exec(username)
	if err != nil {
		log.Println("error executing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return 0
}

func userExists(username string) bool {
	log.Println("checking if user " + username + " exists")
	log.Println("opening db")
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
	}
	log.Println("running query")
	stmt, err := db.Prepare("SELECT username FROM users where username = ?")
	if err != nil {
		log.Println("error preparing select statement")
		log.Printf("%s", err)
		db.Close()
		return false
	}
	defer stmt.Close()
	var retusername string
	err = stmt.QueryRow(username).Scan(&retusername)
	if err != nil {
		log.Println("error scanning")
		log.Printf("%s", err)
		db.Close()
		return false
	}
	log.Println("retusername is: " + retusername)
	db.Close()
	return true
}

func templateList() []TemplateItem {
	var templatereturn []TemplateItem
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
	}
	rows, err := db.Query("SELECT * FROM template")
	if err != nil {
		log.Println("error querying db")
		log.Printf("%s", err)
		rows.Close()
		db.Close()
		return templatereturn
	}
	var name string
	var amount int
	var date int

	for rows.Next() {
		err = rows.Scan(&name, &amount, &date)
		if err != nil {
			log.Println("eror scanning rows")
			log.Printf("%s", err)
			rows.Close()
			db.Close()
			return templatereturn
		}
		thisitem := TemplateItem{Name: name, Amount: amount, Date: date}
		templatereturn = append(templatereturn, thisitem)
	}
	rows.Close()
	db.Close()
	return templatereturn
}

func addToTemplate(name string, amount int, date int) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("INSERT INTO template(name, amount, date) values(?,?,?)")
	if err != nil {
		log.Println("error preparing insert statemnt")
		log.Printf("%s", err)
		db.Close()
		return 2
	}
	_, err = stmt.Exec(name, amount, date)
	if err != nil {
		log.Println("error executing insert statemtnt")
		log.Printf("%s", err)
		db.Close()
		return 3
	}
	db.Close()
	return 0
}

func updateTemplate(name string, amount int, date int) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for insert")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("update template set amount=?, date=? where name=?")
	if err != nil {
		log.Println("error preparing")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := stmt.Exec(amount, date, name)
	if err != nil {
		log.Println("error updating")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return 0
}
func deleteTemplate(name string) int {
	db, err := sql.Open("sqlite3", "./budget.db")
	if err != nil {
		log.Println("error opening db for removal")
		log.Printf("%s", err)
		db.Close()
		return 1
	}
	stmt, err := db.Prepare("delete from template where name=?")
	if err != nil {
		log.Println("error preparing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 2
	}

	res, err := stmt.Exec(name)
	if err != nil {
		log.Println("error executing delete statement")
		log.Printf("%s", err)
		db.Close()
		return 3
	}

	affect, _ := res.RowsAffected()
	log.Println(strconv.FormatInt(affect, 10) + " rows affected")
	db.Close()
	return 0
}

func rowCounter(rows *sql.Rows) (count int) {
	for rows.Next() {
		_ = rows.Scan(&count)
	}
	return count
}
