package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
	"uuid" // this was added by github.com/google/uuid
)

type AppUser struct {
	FirstName    string
	LastName     string
	EmailAddress string
	UserName     string
}

type CheckIfLoggedIn struct {
	user       AppUser
	isLoggedIn bool
}

var parsedTemplate *template.Template
var userDB = map[string]AppUser{}    // username , AppUser
var dbSessions = map[string]string{} // session ID, username

func init() {
	parsedTemplate = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", index)
	serveMux.HandleFunc("/bar", bar)
	serveMux.HandleFunc("/signup", signUp)
	serveMux.HandleFunc("/parseDetails", userSignup)
	serveMux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8081", serveMux)
}

func CheckUserExists(r *http.Request) (string, bool) {
	cookieSession, err := r.Cookie("session") // session should only exist if the user is created
	if err != nil {
		log.Default().Println("FUNC:CheckUserExists\tUser Session Does Not Exist")
		return "", false // sending user dose not exist and an empty user name
	} else {
		log.Default().Println("FUNC:CheckUserExists\tUser Session Does Exist")
		return cookieSession.Value, true
	}
}

func IsNameUnique(r *http.Request) (string, bool) {
	cookieId, ok := CheckUserExists(r) // session should only exist if the user is created
	// if user exist's then check if username is unique
	user, ok := dbSessions[cookieId]
	if !ok {
		log.Default().Println("FUNC:IsNameUnique\tUser Name is Unique! Creating User")
		return user, true
	} else {
		log.Default().Println("FUNC:IsNameUnique\tUser Name is Not Unique! Please choose a different username ->")
		return "", false
	}
}

func IsLoggedIn(r *http.Request) (string, bool) {
	cookieId, ok := CheckUserExists(r) // session should only exist if the user is created
	// if user exist's then check if username is unique
	username, ok := dbSessions[cookieId]
	if ok {
		log.Default().Println("FUNC:IsLoggedIn\tUser is Logged in, Session is Active.")
		return username, ok
	} else {
		log.Default().Println("FUNC:IsLoggedIn\tUser is not Logged in, Session is Inactive.")
		return "", false
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	username, ok := IsLoggedIn(r)
	if ok {
		executeForm(w, "index.html", userDB[username])
		return
	} else {
		executeForm(w, "index.html", nil)
		return
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	username, ok := IsLoggedIn(r)
	if ok {
		parsedTemplate.ExecuteTemplate(w, "bar.html", userDB[username])
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	_, ok := IsLoggedIn(r)
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		// parsedTemplate.ExecuteTemplate(w, "bar.html", userDB[username])
	} else {
		executeForm(w, "signup.html", nil)
	}
}

func executeForm(w http.ResponseWriter, templateName string, data any) {
	err := parsedTemplate.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Printf("Executed Template: %s", templateName)
	}
}

func parseFormGetUser(r *http.Request) *AppUser {
	var formData url.Values
	err := r.ParseForm()
	if err != nil {
		log.Default().Fatalln(err)
	} else {
		log.Default().Println("Executed ParsedForm, Populated values into PostForm and Form.")
	}
	switch r.Method {
	case http.MethodGet:
		formData = r.Form
	case http.MethodPost:
		formData = r.PostForm
	default:
		formData.Add("FirstName", r.FormValue("FirstName"))
		formData.Add("LastName", r.FormValue("LastName"))
		formData.Add("EmailAddress", r.FormValue("EmailAddress"))
	}
	var endUser AppUser = AppUser{
		FirstName:    formData.Get("FirstName"),
		LastName:     formData.Get("LastName"),
		EmailAddress: formData.Get("EmailAddress"),
		UserName:     formData.Get("FirstName") + formData.Get("LastName"),
	}
	return &endUser
}

func userSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	userDetails := parseFormGetUser(r)
	uname, ok := IsNameUnique(r)
	if ok {
		_, err := r.Cookie("session")
		if err != nil {
			fmt.Fprintln(w, "Creating User!")
			sessionToken := uuid.NewString()
			expiresAt := time.Now().Add(120 * time.Second)
			sessionCoocie := &http.Cookie{
				Name: "session", Value: sessionToken,
				HttpOnly: true, Expires: expiresAt}
			http.SetCookie(w, sessionCoocie)
			log.Default().Printf("Creating Session: %s", sessionCoocie)
			dbSessions[sessionToken] = userDetails.UserName
			log.Default().Println(dbSessions[sessionToken])
		}
		userDB[uname] = *userDetails
		executeForm(w, "loggedIn.html", userDetails)
		log.Default().Println(userDB[uname])
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
