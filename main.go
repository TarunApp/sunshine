package main

import (
	"fmt"
	"time"
)

func timer(t time.Duration) {
	fmt.Println("Timer starting")
	timer := time.NewTimer(t * time.Second)
	<-timer.C


	fmt.Println("Time ended after:", t, "minutes")
}

func main() {

	timer(5)

}
