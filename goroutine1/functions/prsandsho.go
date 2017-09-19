// Package functions contains functions calling in main.go
package functions

import (
	"encoding/csv"
	//"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Name    string
	Date    time.Time
	Salaire float64
}
type Bonus struct {
	Name  string
	Prime float64
}

//Prs returns records
func Prs(filePath string, r chan []Record) {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	//fmt.Printf("%T", rows)
	if err != nil {
		log.Fatalln(err)
	}
	rec := make([]Record, 0, len(rows))
	for r, row := range rows {
		if r == 0 {
			continue
		}
		var name string
		name = row[0]
		date, _ := time.Parse("2006-01-02", row[1])
		salaire, _ := strconv.ParseFloat(row[2], 64)

		rec = append(rec, Record{
			Name:    name,
			Date:    date,
			Salaire: salaire,
		})
	}
	r <- rec
}

//Sho returns bonus
func Sho(filePath string, b chan []Bonus) {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()
	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	bon := make([]Bonus, 0, len(rows))
	for b, row := range rows {
		if b == 0 {
			continue
		}
		var name string
		name = row[0]
		primes, _ := strconv.ParseFloat(row[1], 64)
		bon = append(bon, Bonus{
			Name:  name,
			Prime: primes,
		})
	}
	b <- bon
}
