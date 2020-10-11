package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

var buckets = make([][]string, 2)

//GetPage will search for the page in net and gives us a response
func GetPage(httpURL string) {

	httpErrorLogger := func(httpError error) {
		if httpError != nil {

			log.Fatal(httpError)
		}
	}

	httpResponse, httpError := http.Get(httpURL)
	httpErrorLogger(httpError)

	// httpBody, httpError := ioutil.ReadAll(httpResponse.Body)
	httpHeader := httpResponse.Header
	httpXYZ := httpResponse.Status

	// reader := strings.NewReader(httpResponse.Body)
	scanner := bufio.NewScanner(httpResponse.Body)
	defer httpResponse.Body.Close()
	httpErrorLogger(httpError)
	scanner.Split(bufio.ScanWords)

	for i := 0; i <= 12; i++ {
		buckets = append(buckets, []string{})
	}

	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i, _ := range buckets {
		fmt.Println("Bucket Number: ", i, " - ", buckets[i])
	}

}

//HashBucket throws a number
func HashBucket(w string, n int) int {
	var sum int
	for _, v := range w {

		sum += int(v)
	}
	return sum % n
}

func main() {
	GetPage("https://web.ics.purdue.edu/~gchopra/class/public/pages/webdesign/05_simple.html")

}
