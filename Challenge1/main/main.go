package main

import (
	"log"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func main() {

	problem1()

	log.Printf("\n\n")

	problem2()

}
