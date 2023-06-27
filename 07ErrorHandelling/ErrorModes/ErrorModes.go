package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type distanceToDST struct {
	up int
	dn int
}

type DirError struct {
	dst *distanceToDST
	err error
}

func (e *DirError) Error() string {
	return fmt.Sprintf("\tDirError!: %v UP:%d - DN:%d\n", e.err, e.dst.up, e.dst.dn)
}

func (d *distanceToDST) DistanceTravelled() (int, error) {
	if d.dn < d.up {
		return -1, &DirError{d, fmt.Errorf("Overshoot or Stuck!")}
	}
	return d.up - d.dn, nil
}

func createLogger() {
	nf, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Error Cannot Create a log file: %v", err)
	}
	log.SetOutput(nf)
}

func getHalf(num int64) (int64, error) {
	if num == 0 {

		return -1, errors.New("0 cannot be divided!")
	}
	return num / 2, nil
}

func OpenSomeFile(fileName string) {
	someFile, err := os.Open(fileName)
	if err != nil {
		// log.Printf("Cannot Open file: %v\n", err)
		log.Fatalf("Cannot Open file: %v\n", err)
		// log.Panicf("Cannot Open file: %v\n", err)
	}
	defer someFile.Close()
}

func init() {
	createLogger()
}

func main() {

	// half, err := getHalf(0)
	// if err != nil {
	// 	log.Fatalf("ERROR!: %v\n", err)
	// }
	// fmt.Printf("Half: %d\n", half)
	var dst = distanceToDST{130, 120}
	getDistance, err := dst.DistanceTravelled()
	if err != nil {
		log.Printf("%v\n", err)
	}
	fmt.Printf("Distance togo: %d\n", getDistance)

}
