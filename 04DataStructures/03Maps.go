package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GetPage will search for the page in net and gives us a response
func GetPage(httpURL string) {

	httpErrorLogger := func(httpError error) {
		if httpError != nil {

			log.Fatal(httpError)
		}
	}

	httpResponse, httpError := http.Get(httpURL)
	httpErrorLogger(httpError)

	httpBody, httpError := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	httpErrorLogger(httpError)

	httpHeader := httpResponse.Header
	httpXYZ := httpResponse.Status
	fmt.Printf("%s \n", httpBody)

	fmt.Printf("%v \n", httpXYZ)

	for v, k := range httpHeader {

		fmt.Println(v, " - ", k)
	}

}

func main() {
	GetPage("https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html")

}
