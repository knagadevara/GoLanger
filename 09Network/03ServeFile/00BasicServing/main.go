package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// using io.copy
// serving directly from web
func WebServe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<img src="https://upload.wikimedia.org/wikipedia/en/c/c7/Batman_Infobox.jpg" width="500" height="600"  style="vertical-align:middle">`)
}

func defaultServ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>NothingHere!</h1> <p><img src="/jl.jpg" width="500" height="600" style="vertical-align:middle"> </p>`)
	log.Default().Println("defaultServ: Done")
}

// the below function  wont work as the image needs to be unpacked and then served.
func DiskServe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<img src="/supman.png" width="500" height="600 style="vertical-align:middle">"`)
	log.Default().Println("DiskServe: Done")
}

func ServeSteel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	file, err := os.Open("supman.png")
	if err != nil {
		http.Error(w, "404 Not Found!", 404)
		log.Default().Fatalln(err)
	}
	defer file.Close()
	io.Copy(w, file)
	log.Default().Println("ServeFile: Done")
}

func ServeJL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	file, err := os.Open("jl.jpg")
	if err != nil {
		http.Error(w, "404 Not Found!", 404)
		log.Default().Fatalln(err)
	}
	defer file.Close()
	fst, err := file.Stat()
	if err != nil {
		log.Default().Fatalln(err)
	}
	http.ServeContent(w, r, fst.Name(), fst.ModTime(), file)
	log.Default().Println("JL: Done")
}

func main() {
	http.HandleFunc("/bat", WebServe)
	http.HandleFunc("/sup", DiskServe)
	http.HandleFunc("/steel", ServeSteel)
	http.HandleFunc("/", defaultServ)
	http.HandleFunc("/jl.jpg", ServeJL)
	http.ListenAndServe(":8080", nil)

}
