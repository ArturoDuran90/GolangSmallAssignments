package main

import (
	"fmt"
	"time"
)

type FixedHoliday struct {
	name  string
	month time.Month
	day   int
	year  int
}

type MovingHoliday struct {
	name      string
	month     time.Month
	dayOfWeek time.Weekday
	instance  int
	year      int
}

func newFixedHoliday(name string, month time.Month, day int, year int) FixedHoliday {
	return FixedHoliday{name, month, day, year}
}

func newMovingHoliday(name string, month time.Month, dayOfWeek string, instance int, year int) MovingHoliday {
	day := time.Monday
	switch dayOfWeek {
	case "sunday":
		day = time.Sunday
	case "monday":
		day = time.Monday
	case "tuesday":
		day = time.Tuesday
	case "wednesday":
		day = time.Wednesday
	case "thursday":
		day = time.Thursday
	case "friday":
		day = time.Friday
	case "saturday":
		day = time.Saturday
	}
	return MovingHoliday{name, month, day, instance, year}
}

func (fh FixedHoliday) printHoliday() {
	date := time.Date(fh.year, fh.month, fh.day, 0, 0, 0, 0, time.UTC)
	weekday := date.Weekday().String()
	fmt.Printf("%s is on %s, %s %02d in %d\n", fh.name, weekday, fh.month.String(), fh.day, fh.year)
}

func (mh MovingHoliday) printHoliday() {
	t := time.Date(mh.year, mh.month, 1, 0, 0, 0, 0, time.UTC)
	for {
		if t.Weekday() == mh.dayOfWeek {
			mh.instance--
			if mh.instance == 0 {
				break
			}
		}
		t = t.AddDate(0, 0, 1)
	}
	fmt.Printf("%s is on %s, %s %02d in %d\n", mh.name, mh.dayOfWeek.String(), mh.month.String(), t.Day(), mh.year)
}

func main() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanln(&year)

	newYear := newFixedHoliday("- New Year's Day", time.January, 1, year)
	independenceDay := newFixedHoliday("- Independence Day", time.July, 4, year)
	halloween := newFixedHoliday("- Halloween", time.October, 31, year)
	mothersDay := newMovingHoliday("- Mother's Day", time.May, "sunday", 2, year)
	fathersDay := newMovingHoliday("- Father's Day", time.June, "sunday", 3, year)
	thanksgiving := newMovingHoliday("- Thanksgiving", time.November, "thursday", 4, year)

	newYear.printHoliday()
	independenceDay.printHoliday()
	halloween.printHoliday()
	mothersDay.printHoliday()
	fathersDay.printHoliday()
	thanksgiving.printHoliday()
}
