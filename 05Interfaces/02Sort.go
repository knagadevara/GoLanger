package main

import (
	"fmt"
	"sort"
	"strings"
)

type WORD []string
type NUMBER []int

type Sorter interface {
	sortOn()
	reverseOn()
}

func (w *WORD) Len() int           { return len(*w) }
func (w *WORD) Less(i, j int) bool { return i < j }
func (w *WORD) Swap(i, j int)      { i, j = j, i }

func (n *NUMBER) Len() int           { return len(*n) }
func (n *NUMBER) Less(i, j int) bool { return i < j }
func (n *NUMBER) Swap(i, j int)      { i, j = j, i }

func convertToLower(str []string) {
	for i, s := range str {
		str[i] = strings.ToLower(s)
	}
}

func (w *WORD) sortOn() {
	sort.StringSlice(*w).Sort()
	fmt.Println(*w)
}

func (w *WORD) reverseOn() {
	sort.Sort(sort.Reverse(sort.StringSlice(*w)))
	fmt.Println(*w)
}

func (n *NUMBER) sortOn() {
	sort.IntSlice(*n).Sort()
	fmt.Println(*n)
}

func (n *NUMBER) reverseOn() {
	sort.Sort(sort.Reverse(sort.IntSlice(*n)))
	fmt.Println(*n)
}

func smartSort(s Sorter) {
	s.sortOn()
}

func smartReverse(s Sorter) {
	s.reverseOn()
}

func main() {

	unsortedString := "This name is only the final element of the path the base name not the entire path"
	var listOfWords WORD = WORD(strings.Split(unsortedString, " "))
	convertToLower(listOfWords)
	listOfWords.sortOn()
	listOfWords.reverseOn()
	unSortedIntegers := NUMBER([]int{31, 0, 54, 12, 78, 99})
	unSortedIntegers.sortOn()
}
