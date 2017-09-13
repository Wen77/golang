package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var pay []Record
var b []uint8
var c []uint8

type Record struct {
	Name    string
	Date    time.Time
	Salaire string
}

func main() {
	pay := csvf("table.csv")
	//marshal
	j, err := json.Marshal(pay)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(j)
	fmt.Printf("%T", j)
	crf(j)

}
func csvf(filePath string) []Record {
	//open
	e, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer e.Close()
	//read
	d := csv.NewReader(e)
	rowd, err := d.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, row := range rowd {
		if i == 0 {
			continue
		}
		name := (row[0])
		date, _ := time.Parse("2006-01-02", row[1])
		salaire := (row[2])
		pay = append(pay, Record{
			Name:    name,
			Date:    date,
			Salaire: salaire,
		})
	}
	return pay
}

func crf(b []uint8) {
	//open and read
	f, err := os.OpenFile("table.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//write
	j, err := f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(j)
}
