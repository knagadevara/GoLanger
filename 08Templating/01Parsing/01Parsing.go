package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tplParsed, tplFuncParsed *template.Template
var templateNames = make(map[string][]string)

type authPerson struct {
	FName string
	LName string
	Age   uint8
}

type vehicelCar struct {
	Manufacturer string
	Name         string
	Doors        int
}

type carOwner struct {
	Car   vehicelCar
	Owner authPerson
}

func makePassword(s string) string {
	return strings.ToUpper(s[:3])
}

var tempFunc = template.FuncMap{
	"mkp": makePassword,
}

func init() {
	// tpl1 := template.Must(template.New("testTemplate.tpl").Parse("This is a dynamic new template")) // Creates a new template
	// the below does not work and panics as the parsing of the file is done first and the functions are not declared yet.
	// tplTmp := template.Must(template.ParseFiles("templates/funcTemp.ftpl"))
	// tplTmp.Funcs(tempFunc)
	tplParsed = template.Must(template.ParseGlob("templates/*.tpl")) //// to parse multiple files from multiple different locations
	// parsedTemplate, err := template.ParseFiles("templates/welcomeHtml.tpl") // to parse only one file
	tplFuncParsed = template.Must(template.New("").Funcs(tempFunc).ParseFiles("templates/funcTemp.ftpl")) // Creates a Temporary template, Parces funcs into it and then Parses a template file
}

func main() {

	var carList = make([]carOwner, 0)
	somePerson := []string{"John", "Wick", "31"}
	var JohnW = authPerson{"John", "Wick", 31}
	var NickC = authPerson{"Nick", "Coleman", 50}
	var ArnieS = authPerson{"Arnold", "Sdcreenfo", 71}
	var TonySu = authPerson{"Tony", "Soprano", 41}
	var LuckyL = authPerson{"Lucky", "Luchiyano", 37}
	templateNames["Paris"] = []string{"Marquie"}
	templateNames["Germany"] = []string{"Dragana"}
	templateNames["USA"] = []string{"Hugo"}
	templateNames["Russia"] = []string{"Igor"}
	hondaS800 := vehicelCar{"Honda", "S800", 2}
	fordMustang := vehicelCar{"fordMustang", "GT500", 2}
	toyotaS1000 := vehicelCar{"Toyota", "S1000", 2}
	bmwM3 := vehicelCar{"BMW", "M3", 2}
	certioenDS1 := vehicelCar{"certioen", "DS1", 2}
	cO1 := carOwner{fordMustang, JohnW}
	cO2 := carOwner{hondaS800, NickC}
	cO3 := carOwner{certioenDS1, ArnieS}
	cO4 := carOwner{toyotaS1000, LuckyL}
	cO5 := carOwner{bmwM3, TonySu}
	carList = []carOwner{cO1, cO2, cO3, cO4, cO5}

	// for k, v := range templateNames {
	// 	//Fucking never create files like this
	// 	mkfname := fmt.Sprintf("%s", k)
	// 	hFile, err := os.Create(mkfname + ".html")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	tplParsed.ExecuteTemplate(hFile, mkfname+".tpl", v)
	// 	defer hFile.Close()
	// }

	JohnFile, err := os.Create("welcomeJohn.html")
	donRides, err := os.Create("donRides.html")
	welcomeFile, err := os.Create("welcomeHuman.html")
	jKillsFile, err := os.Create("jKills.html")
	funcTemp, err := os.Create("funcTemp.html")
	if err != nil {
		log.Fatal(err)
	}
	defer funcTemp.Close()
	defer donRides.Close()
	defer welcomeFile.Close()
	defer jKillsFile.Close()
	defer JohnFile.Close()

	// parsedTemplate.Execute(os.Stdout, nil) // to send the template over to StdOut(screen)
	// err1 := parsedTemplate.Execute(newFile, nil) // to print the template to the specified file.
	err1 := tplParsed.ExecuteTemplate(welcomeFile, "welcomeHtml.tpl", somePerson) // to print the template to the specified file.
	tplParsed.ExecuteTemplate(jKillsFile, "JKills.tpl", templateNames)            // loop over the strings as per the login in template
	tplParsed.ExecuteTemplate(JohnFile, "welcomeJohn.tpl", JohnW)                 // A simple struct
	tplParsed.ExecuteTemplate(donRides, "donRides.tpl", carList)                  // Sending a complex struct which would unpack itself in template
	tplFuncParsed.ExecuteTemplate(funcTemp, "funcTemp.ftpl", carList)             // Sending a complex struct which would unpack itself in template

	if err1 != nil {
		log.Fatal(err)
	}
}
