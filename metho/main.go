package main

import (
	//"fmt"
	"log"
	"time"

	"gotraining.com/try/metho/contributions"
)

var Bcsgcsa, Bufs float64
var Nc []string
var Gross float64

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %dms", name, elapsed.Nanoseconds()/1000)
}

func main() {
	start := time.Now()
	//open and read csv file with rate.go
	r := make(chan []contributions.Short)
	go contributions.Ctrib("rate.csv", r)
	Reco := <-r
	//fmt.Println(Reco)
	contributions.Printer(Reco)
	elapsed := time.Since(start)
	log.Printf("metho took %dmillisnd", elapsed.Nanoseconds()/1000000)
}
