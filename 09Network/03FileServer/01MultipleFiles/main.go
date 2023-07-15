package main

import (
	"io"
	"net/http"
)

func serveSkull(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="GraySkull.jpg" width="500" height="600"  style="vertical-align:middle">`)
}

func serveJL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("DataSource", "Web")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="justiceLeague.jpg" width="500" height="600"  style="vertical-align:middle">`)
}

// serving a complete directory
func main() {
	rootFileDir := http.Dir("../CommonObjects/images")
	srvRtDrHndlr := http.FileServer(rootFileDir)
	smux := http.NewServeMux()
	// File Server serving the contents of the directory,
	// because of index.html the files will not be displayed.
	smux.Handle("/", srvRtDrHndlr)
	// these are the alternate routes to go to images
	smux.HandleFunc("/gs", serveSkull) // without instanciating file server this will return an empty box of image
	smux.HandleFunc("/jl", serveJL)    // without instanciating file server this will return an empty box of image

	http.ListenAndServe(":8081", smux)
}
