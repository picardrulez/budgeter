package main

import ()

func getTemplateItem(name string) TemplateItem {
	amount, _ := getAmountValue(name)
	date, _ := getDateValue(name)
	website, _ := getWebsiteValue(name)
	username, _ := getUsernameValue(name)
	password, _ := getPasswordValue(name)
	returnItem := TemplateItem{Name: name, Amount: amount, Date: date, Website: website, Username: username, Password: password}
	return returnItem
}
