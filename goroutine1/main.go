package main

import (
	"fmt"
	"strconv"

	"github.com/Wen77/golang/goroutine1/functions"
)

func main() {
	r := make(chan []functions.Record)
	b := make(chan []functions.Bonus)
	go functions.Prs("table.csv", r)
	go functions.Sho("primes.csv", b)
	records, bonus := <-r, <-b
	//fmt.Printf("%T\n", records, bonus)
	for _, record := range records {
		sbrut := make(map[string]float64)
		for _, prime := range bonus {

			if record.Name == prime.Name {

				sbrut[record.Name] = record.Salaire + prime.Prime

				for keys, f := range sbrut {

					round := make(map[string]float64)
					i := fmt.Sprintf("%.2f", f)
					f, _ := strconv.ParseFloat(i, 2)
					round[keys] = f
					fmt.Println(round)
				}
			}
		}
	}
}
