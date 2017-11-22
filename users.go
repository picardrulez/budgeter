package main

import (
	"github.com/picardrulez/lcars"
	"io"
	"net/http"
	"strings"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		content := `
		<h1>Create User</h1>
		<form method="post" action="/usercreation">
			<label for="username">User</label>
			<input type="text" id="username" name="username">
			<label for="pass">Password</label>
			<input type="password" id="pass" name="pass">
			<br/>
			<button type="submit">Create User</button>
		</form>
		`
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func usercreationHandler(w http.ResponseWriter, r *http.Request) {
	newuser := strings.ToLower(r.FormValue("username"))
	password := r.FormValue("pass")
	hashedpass := encryptPassword(password)
	userres := insertUser(newuser, hashedpass)
	if userres != 0 {
		io.WriteString(w, "an error occured creating user password")
		return
	}
	http.Redirect(w, r, "/createUser", 302)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		userList, _ := getUserList()
		sortList := alphabetizeList(userList)
		content := `
		<h1>User List</h1>
		<br/>
		<table> <tr><td>USER</td></tr>
		`
		for _, k := range sortList {
			content = content + "<tr><td>" + k + "</td></tr>"
		}
		content = content + "</table>"
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		content := `
		<h1>Change Password</h1>
		<form method="post" action="/passchange">
			<labelfor="oldPass">Old Password</label>
			<input type="password" id="oldPass" name="oldPass">
			<label for="newPass">New Password</label>
			<input type="password" id="newPass" name="newPass">
			<button type="submit">Update</button>
		</form>
		`
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func adminChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		userList, _ := getUserList()
		sortList := alphabetizeList(userList)
		content := `
		<h1>Admin Change Password</h1>
		<br/>
		<form method="post" action="/passEntry">
			<select name="user">
		`
		for _, k := range sortList {
			content = content + `<option value="` + k + `">` + k + `</option`
		}
		content = content + `
		</select>
		<input type="password" name="newpass" id="newpass">
		<button type="submit">Submit</button>
		</form>
		`
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func passEntryHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	user := r.FormValue("user")
	password := r.FormValue("newpass")
	hashedpass := encryptPassword(password)
	if userName != "" {
		postres := updatePassword(user, hashedpass)
		if postres != 0 {
			io.WriteString(w, "an error occured updating password")
			return
		}
		http.Redirect(w, r, "/adminChangePassword", 302)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		userList, _ := getUserList()
		sortList := alphabetizeList(userList)
		content := `
		<h1>Delete User</h1>
		<br/>
		<form method="post" action="/userdel">
			<select name="user">
		`
		for _, k := range sortList {
			content = content + `<option value="` + k + `">` + k + `</option>`
		}
		content = content + `
		</select>
		<button type="submit">Delete User</button>
		</form>
		`
		t, createPage := lcars.MakePage(content, mymenu, lcarssettings)
		t.Execute(w, createPage)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func userdelHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	passdelres := deleteUser(user)
	if passdelres != 0 {
		io.WriteString(w, "an error occured deleting user")
		return
	}
	http.Redirect(w, r, "/deleteUser", 302)
}
