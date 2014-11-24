package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func argsToTimes(args []string) (time.Time, time.Time) {
	const hoursMinutesFormat = "15:04"

	if len(args) != 3 {
		log.Fatalln("Usage: worklog start_time end_time")
	}

	startTime, err := time.Parse(hoursMinutesFormat, args[1])
	if err != nil {
		log.Fatalln("Invalid startTime time")
	}

	endTime, err := time.Parse(hoursMinutesFormat, args[2])
	if err != nil {
		log.Fatalln("Invalid endTime time")
	}

	if endTime.Before(startTime) {
		log.Fatalln("endTime has to be after startTime")
	}

	return startTime, endTime
}

func duration(startTime, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}

func init() {
	log.SetFlags(0)
}

func main() {
	startTime, endTime := argsToTimes(os.Args)

	fmt.Println(duration(startTime, endTime))
}
