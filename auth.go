package main

import (
	"fmt"
	"github.com/picardrulez/lcars"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
)

const indexPage = `
<form method="post" action="/login">
	<input type="text" id="name" name="name" autofocus="autofocus" placeholder="username">
	<input type="password" id="password" name="password" placeholder="password">
	<button type="submit">Login</button>
</form>
`

func rootHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		http.Redirect(w, r, "/main", 302)
	} else {
		content := `
		<html><head></head><body bgcolor='black'>
		<center>
		<br/><br/><br/>
		<br/>
		<img src='/resources/ufp2.gif'>
		<br/>
		<br/>
		`
		content = content + indexPage + `
		<font color="white">Level Alpha One Security Clearence required for access.</font>
		`
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	}
}

func getUserName(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		realPass, _ := getPassword(name)
		hashcheck := checkPass(pass, realPass)
		if hashcheck {
			setSession(name, w)
			redirectTarget = "/main"
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func checkPass(password string, hash string) bool {
	log.Println("bytifying password and hash")
	log.Println("password is " + password)
	log.Println("hash is " + hash)
	bytePass := []byte(password)
	byteHash := []byte(hash)

	log.Println("comparing passes with bcrypt")
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		log.Println("an error occured")
		log.Println(err)
		return false
	} else {
		return true
	}
}

func encryptPassword(password string) string {
	log.Println("bytifying password")
	log.Println("password is " + password)
	bytePassword := []byte(password)

	log.Println("turning bytepassword into hash")
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("an error occured hashing password")
	}
	log.Println("returning hashedpassword: " + string(hashedPassword))
	return string(hashedPassword)
}

func passwordHandler(w http.ResponseWriter, r *http.Request) {
	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	userName := getUserName(r)
	if userName != "" && oldPass != "" {
		realPass, _ := getPassword(userName)
		hashcheck := checkPass(oldPass, realPass)
		if hashcheck {
			hashpass := encryptPassword(newPass)
			postres := insertUser(userName, hashpass)
			if postres != 0 {
				io.WriteString(w, "an error occured posting new password")
				return
			}
		} else {
			io.WriteString(w, "old password is incorrect")
			return
		}
	}
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
