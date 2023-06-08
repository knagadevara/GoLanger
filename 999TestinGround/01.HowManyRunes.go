package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var countString []string
var catInChinese string

func toUpper(s *string) int {
	runeCount := utf8.RuneCountInString(*s)
	*s = strings.ToUpper(*s)
	return runeCount
}

func main() {
	countString = os.Args
	fmt.Printf("%v%v\n", countString[1], strings.Repeat("!", toUpper(&countString[1])))
	catInChinese = "猫"
	fmt.Printf("%v\nLenth of string is:%v\nCharacters in String: %v\n",
		catInChinese,
		len(catInChinese),
		utf8.RuneCountInString(catInChinese))
}
