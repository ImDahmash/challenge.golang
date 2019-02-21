package main

import (
	"log"
	"math/rand"
	"time"
)

func problem1() {

	log.Printf("problem1: started at time %v --------------------------------------------", time.Since(start))

	for inx := 0; inx < 10; inx++ {

		go printRandom1(inx)

	}

	// Here, We make the main routine: fun(problem1) sleeps for 10 Seconds
	// All subroutines are working in concurrency way.
	// By this way, we can quit all go-routines after a total of exactly 100 random numbers have been printed.
	// Each random number needs 1 second to be printed
	// We have 10 subroutines, so we will print 10 random numbers for each routine (10*10) = 100 random numbers.
	time.Sleep(10 * time.Second)
	log.Printf("problem1: 100 random numbers have been printed at time %v --------------------------------------------", time.Since(start))
	log.Printf("problem1: finised at time %v --------------------------------------------", time.Since(start))
}

func printRandom1(slot int) {

	//
	// Do not change 25 into 10!
	//

	//Here we loop over each routine and print random numbers through inx loop from 0 to 24
	//We used time.sleep for 1 second to wait 1 second to print each random number per routine(slot)
	//log.Printf("problem1: counter=%d of 25 random numbers at time %v --------------------------------------------",slot+1, time.Since(start))
	for inx := 0; inx < 25; inx++ {

		log.Printf("problem1: slot=%03d count=%05d rand=%f at time %v ", slot, inx, rand.Float32(), time.Since(start))
		time.Sleep(time.Second)

	}
}
