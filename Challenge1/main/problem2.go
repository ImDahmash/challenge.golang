package main

import (
	"log"
	"math/rand"
	"time"
)

func problem2() {

	log.Printf("problem2: started at time %v --------------------------------------------", time.Since(start))

	for inx := 0; inx < 10; inx++ {

		go printRandom2(inx)

	}

	// Here, We make the main routine: fun(problem2) sleeps for 10 Seconds
	// All subroutines are working in concurrency way.
	// Each random number needs 1 second to be printed
	// We have 10 subroutines, so we will print 10 random numbers for each routine (10*10).
	time.Sleep(10 * time.Second)

	log.Printf("problem2: finised at time %v --------------------------------------------", time.Since(start))
}

func printRandom2(slot int) {

	//Here we loop over each routine and print random numbers through inx loop from 0 to 10
	//We used time.sleep for 1 second to wait 1 second to print each random number per routine(slot)
	for inx := 0; inx < 10; inx++ {

		log.Printf("problem2: slot=%03d count=%05d rand=%f at time %v", slot, inx, rand.Float32(), time.Since(start))
		time.Sleep(time.Second)
	}
}
