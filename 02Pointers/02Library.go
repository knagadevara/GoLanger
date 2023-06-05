package main

import (
	"time"
)

type Book struct {
	id                  uint16
	Title               string
	Author              string
	Pages               uint16
	Language            string
	BookAvailableToLend bool
	timeLent            *timeStamps
}

type mContact struct {
	workPhone uint16
	altPhone  uint16
	eMail     string
}

type Member struct {
	id             uint16
	Name           string
	currentLimit   uint16
	validityTill   *time.Time
	ContactDetails *mContact
}

type timeStamps struct {
	checkIn  time.Time
	checkOut time.Time
}

type Library struct {
	LentBooks  map[*Member]*Book
	libContact *mContact
}

func getFakeTime(year int, timeObj *time.Time) (TimeArray []time.Time) {
	timeArray := make([]time.Time, 12)
	for i := 1; i < 13; i++ {
		dateOld := time.Date(year, time.January, i, 0, 0, 0, 0, timeObj.Location())
		timeArray = append(timeArray, dateOld)
		// dateOld + dateOld.Month()
	}
	return timeArray
}

func main() {
	var TestTime time.Time
	fakeTimeArray := getFakeTime(2020, &TestTime)
	for _, tdime := range fakeTimeArray {
		println(tdime.Format("2020-01-01"))
	}

	bookArray := make([]Book, 10)
	memberArray := make([]Member, 5)
	library := make(map[Library]timeStamps)

}
