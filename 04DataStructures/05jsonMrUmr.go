package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Person can be an exported variiable
type Person struct {
	FirstName  string `json:"FName"`
	LastName   string `json:"LName"`
	Age        int
	birthPlace string // not exported or optional parameter
}

func (p Person) GetFullName() (fullName string) {
	return fmt.Sprintf("FullName:\t%s, %s\n", p.LastName, p.FirstName)
}

// Marshalling: converts struct into json at already defined string/file level
func (p *Person) UnloadJson() (binData []byte) {
	bs, _ := json.Marshal(p)
	return bs
}

// UnMarshalling: accepting an input in the form of bytes which is in json and converts into struct
func (p *Person) loadJson(binData []byte) (personData *Person) {
	json.Unmarshal(binData, p)
	return p
}

// Encoding is writing to a stream.
func (p *Person) WriteToStdOut() {
	json.NewEncoder(os.Stdout).Encode(p)
}

// Reading from a stream and decoding a json and putiing it into a structure
func (p *Person) ReadFrom(strJson string) {
	strReader := strings.NewReader(strJson)
	json.NewDecoder(strReader).Decode(p)
}

func main() {

	var person1, person2, person3, person4 Person
	byteArr2D := make([][]byte, 0)

	person1 = Person{FirstName: "Randy",
		LastName:   "Paush",
		Age:        35,
		birthPlace: "USA"}
	byteArr2D = append(byteArr2D, []byte(`{"FName": "Mandy", "LName": "Raush", "Age" : 15, "birthPlace": "EUN"}`))
	byteArr2D = append(byteArr2D, []byte(`{"FName": "Raghava", "LName": "Ram", "Age" : 16, "birthPlace": "IN"}`))
	var person4Json string = `{"FName": "Ravana", "LName": "Lankeshwara", "Age" : 28, "birthPlace": "SL"}`

	_ = person1.UnloadJson()
	person2.loadJson(byteArr2D[0])
	person3.loadJson(byteArr2D[1])
	person4.ReadFrom(person4Json)

	// the observation would be the lowecase fields are not printed because the compiler ignores it.
	fmt.Printf(person1.GetFullName())
	fmt.Printf(person2.GetFullName())
	person3.WriteToStdOut()
	person4.WriteToStdOut()

}
