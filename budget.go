package main

import (
	"github.com/picardrulez/lcars"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func viewBudgetHandler(w http.ResponseWriter, r *http.Request) {
	settings, _ := getSettings()
	budgetCheck()
	nextPayDay := getNextPayDay()
	lastpayday := getLastPayDay(settings)
	budgetList, _ := getCurrentBudget()
	amountTotal := 0
	content := `
	<h1>Budget</h1>
	<br/>
	Pay Period:  ` + lastpayday + ` - ` + nextPayDay + `
	<br/>
	<FORM METHOD='post' action='/forceBudgetCreation'><button type='submit'>Force Budget Refresh</button></form><br/>
	<table><tr><td>Name</td><td>Amount</td><td>Date</td><td>Website</td><td>Username</td><td>Password</td><td>Paid</td><td></td><td></td></tr>
	`
	for _, k := range budgetList {
		itemArray := strings.Split(k, ":")
		itemName := itemArray[0]
		itemPaid := itemArray[1]
		log.Println("itemPaid is " + itemPaid)
		currentItem := getTemplateItem(itemName)
		amountTotal = amountTotal + currentItem.Amount
		content = content + "<tr><td>" + currentItem.Name + "</td><td>" + strconv.Itoa(currentItem.Amount) + "</td><td>" + strconv.Itoa(currentItem.Date) + "</td><td>" + currentItem.Website + "</td><td>" + currentItem.Username + "</td><td>" + currentItem.Password + "</td><td>" + itemPaid + "</td><td><FORM METHOD='post' action='/payItemProcessor'>"
		if itemPaid == "false" {
			islate := isLate(currentItem.Date)
			content = content + "<input type='hidden' name='item' value='" + currentItem.Name + "'><input type='hidden' name='pay' value='1'><button type='submit'>Paid</button></form></td><td>"
			if islate {
				content = content + "<text='red'>LATE</text></td></tr>"
			} else {
				content = content + "</td></tr>"
			}

		} else {
			content = content + "<input type='hidden' name='item' value='" + currentItem.Name + "'><input type='hidden' name='pay' value='0'><button type='submit'>Unpay</button></form></td><td></td></tr>"
		}
	}
	content = content + "<tr><td>TOTAL:</td><td>" + strconv.Itoa(amountTotal) + "</td><td></td><td></td><td></td><td></td><td></td><td></td></tr>"
	content = content + "</table>"

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

func isPayDay(datecheck string, settings Settings) bool {
	payday := settings.StartDate
	length := settings.PeriodLength
	dateArray := strings.Split(payday, "-")
	year, _ := strconv.Atoi(dateArray[0])
	monthint, _ := strconv.Atoi(dateArray[1])
	month := time.Month(monthint)
	day, _ := strconv.Atoi(dateArray[2])
	start := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	checkArray := strings.Split(datecheck, "-")
	checkYear, _ := strconv.Atoi(checkArray[0])
	checkMonthInt, _ := strconv.Atoi(checkArray[1])
	checkMonth := time.Month(checkMonthInt)
	checkDay, _ := strconv.Atoi(checkArray[2])
	check := time.Date(checkYear, checkMonth, checkDay, 0, 0, 0, 0, time.UTC)
	diff := check.Sub(start)
	days := diff.Hours() / 24
	returnbool := int(days)%length == 0
	return returnbool
}

func isOutOfPayPeriod() bool {
	settings, _ := getSettings()
	nowdate := time.Now()
	nowdateformatted := nowdate.Format("2006-01-02")
	nowArray := strings.Split(nowdateformatted, "-")
	nowYear, _ := strconv.Atoi(nowArray[0])
	nowMonthInt, _ := strconv.Atoi(nowArray[1])
	nowMonth := time.Month(nowMonthInt)
	nowDay, _ := strconv.Atoi(nowArray[2])
	nowcheck := time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.UTC)

	currentPayDay := settings.CurrentPayDay
	cpdArray := strings.Split(currentPayDay, "-")
	cpdYear, _ := strconv.Atoi(cpdArray[0])
	cpdMonthInt, _ := strconv.Atoi(cpdArray[1])
	cpdMonth := time.Month(cpdMonthInt)
	cpdDay, _ := strconv.Atoi(cpdArray[2])
	cpdcheck := time.Date(cpdYear, cpdMonth, cpdDay, 0, 0, 0, 0, time.UTC)

	diff := cpdcheck.Sub(nowcheck)
	days := diff.Hours() / 24

	if days > 13 {
		return true
	} else {
		return false
	}
}

func getLastPayDay(settings Settings) string {
	nowdate := time.Now()
	hazpayday := false
	var lastpayday string
	for hazpayday == false {
		nowdateformated := nowdate.Format("2006-01-02")
		paydaycheck := isPayDay(nowdateformated, settings)
		if paydaycheck {
			lastpayday = nowdateformated
			hazpayday = true
		} else {
			nowdate = nowdate.AddDate(0, 0, -1)
		}
	}
	return lastpayday
}

func getNextPayDay() string {
	settings, _ := getSettings()
	lastPayDay := getLastPayDay(settings)
	lastDate, err := time.Parse("2006-01-02", lastPayDay)
	if err != nil {
		log.Println(err)
		return ""
	}
	nextdate := lastDate.AddDate(0, 0, 14)
	nextFormatted := nextdate.Format("2006-01-02")
	return nextFormatted
}

func createNewPayPeriod() int {
	newBudget, _ := getBudgetList()
	removeItemsReturn := removePaidItems()
	if removeItemsReturn > 0 {
		log.Println("error removing paid items from budget")
		return 1
	}
	insertReturn := budgetInsert(newBudget)
	if insertReturn > 0 {
		log.Println("error inserting new budget")
		return 2
	}
	return 0
}

func payItemProcessor(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	item := r.FormValue("item")
	pay, _ := strconv.Atoi(r.FormValue("pay"))
	if userName != "" {
		payItem(item, pay)
		http.Redirect(w, r, "/viewBudget", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func budgetCreationProcessor(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		response := createNewPayPeriod()
		if response > 0 {
			log.Println("error creating pay period")
		}
		http.Redirect(w, r, "/viewBudget", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
