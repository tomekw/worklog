package main

import (
	"fmt"
	"os"
	"time"
)

const hoursMinutesFormat = "15:04"

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: worklog start_date end_date")
		os.Exit(0)
	}

	timeFrom, err := time.Parse(hoursMinutesFormat, os.Args[1])

	if err != nil {
		fmt.Println("Invalid timeFrom time")
		os.Exit(-1)
	}

	timeTo, err := time.Parse(hoursMinutesFormat, os.Args[2])

	if err != nil {
		fmt.Println("Invalid timeTo time")
		os.Exit(-1)
	}

	if timeTo.Before(timeFrom) {
		fmt.Println("timeTo has to be after timeFrom")
		os.Exit(-1)
	}

	fmt.Println(timeTo.Sub(timeFrom))
}
