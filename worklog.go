package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const hoursMinutesFormat = "15:04"

func init() {
	log.SetFlags(0)
}

func main() {
	if len(os.Args) != 3 {
		log.Println("Usage: worklog start_date end_date")
		return
	}

	timeFrom, err := time.Parse(hoursMinutesFormat, os.Args[1])

	if err != nil {
		log.Fatalln("Invalid timeFrom time")
	}

	timeTo, err := time.Parse(hoursMinutesFormat, os.Args[2])

	if err != nil {
		log.Fatalln("Invalid timeTo time")
	}

	if timeTo.Before(timeFrom) {
		log.Fatalln("timeTo has to be after timeFrom")
	}

	fmt.Println(timeTo.Sub(timeFrom))
}
