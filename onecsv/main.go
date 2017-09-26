package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Record struct {
	Name    string
	Date    time.Time
	Salaire float64
	Bonus   float64
}

func main() {

	records := prs("table.csv")

	for _, record := range records {

		sbrut := make(map[string]float64)
		sbrut[record.Name] = 100 + record.Salaire + record.Bonus
		for key, val := range sbrut {

			fmt.Println(key, "-", val)
		}
	}
}

func prs(filePath string) []Record {
	//open csv
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	//read csv
	rdr := csv.NewReader(src)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	records := make([]Record, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}

		var name string
		name = row[0]
		//parse csv
		date, _ := time.Parse("2006-01-02", row[1])
		salaire, _ := strconv.ParseFloat(row[2], 64)
		bonus, _ := strconv.ParseFloat(row[3], 64)

		records = append(records, Record{
			Name:    name,
			Date:    date,
			Salaire: salaire,
			Bonus:   bonus,
		})
	}
	return records
}
