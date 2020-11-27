package main

import (
	"encoding/json"
	"fmt"
)

// Person can be an exported variiable
type Person struct {
	FirstName  string
	LastName   string
	Age        int
	birthPlace string // not exported or optional parameter
}

func main() {

	person1 := Person{FirstName: "Randy",
		LastName:   "Paush",
		Age:        35,
		birthPlace: "USA"}
	// Marshalling the fields into json
	bs, _ := json.Marshal(person1)
	fmt.Printf("Type: %T \nValue: %v \n", bs, bs)
	fmt.Println(string(bs))
	// the observation would be the lowecase fields are not printed because the compiler ignores it.

	var person2 Person
	bsUnm := []byte(`{"FirstName": "Mandy", "LastName": "Raush", "Age" : 15, "birthPlace": "EUN"}`)
	json.Unmarshal(bsUnm, &person2)
	fmt.Printf("FirstName Value: %v \n", person2.FirstName)
	fmt.Printf("Type: %T \nValue: %v \n", person2, person2)
}
