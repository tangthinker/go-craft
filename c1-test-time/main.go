package main

import (
	"log"
	"time"
)

func main() {
	defer TrackTime(time.Now())

	for i := 0; i < 1000000000; i++ {
		k := 1
		k++
	}

}

func TrackTime(pre time.Time) time.Duration {
	elapsed := time.Since(pre)
	log.Println("Time Elapsed:", elapsed)
	return elapsed
}
