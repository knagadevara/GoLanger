package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "plain/text ; charset=utf-8")
	http.SetCookie(w, &http.Cookie{Name: "HomePage", Value: "RandomValue"})
	cookie, err := r.Cookie("counter")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{Name: "counter", Value: "0"}
	}
	counter, _ := strconv.Atoi(cookie.Value)
	counter++
	cookie.Value = strconv.Itoa(counter)
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "<html><head><h1>%v</h1></head></html>", counter)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "plain/text ; charset=utf-8")
	allCookies := r.Cookies()
	var content string
	for _, v := range allCookies {
		if v.Name != "" {
			htmlPg := fmt.Sprintf("<p>%v : %v</p>", v.Name, v.Value)
			content += htmlPg
		}
	}
	htmlPg := fmt.Sprintf("<html><head>%v</head></html>", content)
	fmt.Fprintln(w, htmlPg)
}

func expireCookie(w http.ResponseWriter, r *http.Request) {
	allCookies := r.Cookies()
	for _, v := range allCookies {
		v.MaxAge = -1
		http.SetCookie(w, v)
	}
	http.Redirect(w, r, "/getCookie", http.StatusSeeOther)
}

func main() {
	servMux := http.NewServeMux()
	servMux.HandleFunc("/setCookie", homePage)
	servMux.HandleFunc("/getCookie", getCookie)
	servMux.HandleFunc("/delCookie", expireCookie)
	http.ListenAndServe(":8081", servMux)
}
