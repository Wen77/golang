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

type Record struct {
	Name    string
	Date    time.Time
	Salaire string
}

func main() {
	pay := csvf("table.csv")
	fmt.Printf("%T\n", pay)
	fmt.Println(marsh(b))
}
func csvf(filePath string) []Record {
	//open
	c, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()
	//read
	d := csv.NewReader(c)
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
func marsh(s []uint8) []uint8 {
	//marshal
	b, err := json.Marshal(pay)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	return b
}
