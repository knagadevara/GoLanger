package main

import "fmt"

type CarPart struct {
	name    string
	atPoint uint8
}

const assemblyPoints uint8 = 3

var assemblyLine = make([]CarPart, assemblyPoints)

func printPointData1(part []CarPart) {
	for i := 0; i < len(part); i++ {
		var data CarPart = part[i]
		fmt.Printf("Part %v\t%v\n", data.atPoint, data.name)
	}
}

func printPointData2(part []CarPart) {
	for _, elem := range part {
		fmt.Printf("Part %v\t%v\n", elem.atPoint, elem.name)
	}
}

func main() {
	assemblyLine = []CarPart{{"Chafts", 1}, {"Camms", 2}, {"Pistons", 3}}
	printPointData1(assemblyLine)
	assemblyLine = append(assemblyLine, CarPart{"Cylinders", 4}, CarPart{"CylinderHead", 5}, CarPart{"OilSeal", 6})
	printPointData2(assemblyLine[4:6])
}
