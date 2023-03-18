package controllers

import (
	"fmt"
	"loginPage/db"
	"loginPage/helpers"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var login = "index.html"
var loginPath = "templates/index.html"

var Home = "homepage.html"
var Homepath = "templates/homepage.html"

var cookieId string

// loginfun
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN")
	if helpers.CheckSession(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//template creation
	temp, err := template.New(login).ParseFiles(loginPath)

	if helpers.CheckError(err) {
		return
	}

	temp.Execute(w, db.LoginMessage)

}

// submitfun
func Submit(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SUBMIT")

	if helpers.CheckSession(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	formEmail := r.PostFormValue("email")
	formPass := r.PostFormValue("password")

	user, ok := db.Users[formEmail]

	// user not found
	if !ok {
		db.LoginMessage.Message = "User not found"
		db.LoginMessage.Color = "text-danger"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("check")
	//user found then chech pass with our user pass
	if formPass != user.Pass {

		db.LoginMessage.Color = "text-danger"
		db.LoginMessage.Message = "incorrect password"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//sessiontToken := "tester"
	sessiontToken := uuid.NewString()
	sessiontTime := time.Now().Add(2 * time.Minute)
	db.SessionToken = sessiontToken

	db.Sessions[sessiontToken] = db.Session{
		Name:   "Sreenath",
		Expire: sessiontTime,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   sessiontToken,
		Expires: sessiontTime,
	})

	//clear all login error messages

	//user ok password ok
	fmt.Println("sub last")
	HomePage(w, r)

}

func HomePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HOMEPAGE")

	if !helpers.CheckSession(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	temp, err := template.New(Home).ParseFiles(Homepath)

	if helpers.CheckError(err) {
		return
	}

	ses, _ := db.Sessions[db.SessionToken]

	temp.Execute(w, ses)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("LOGOUT")

	if !helpers.CheckSession(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	delete(db.Sessions, db.SessionToken)
	db.SessionToken = ""
	Login(w, r)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ERRORPAGE")

	if helpers.CheckSession(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
