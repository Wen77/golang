package contributions

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func Ctrib(filePath string) []Contrib {
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
	//create a slice type Contrib
	Co := make([]Contrib, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		var libel string
		libel = row[0]
		taux, _ := strconv.ParseFloat(row[1], 64)

		Co = append(Co, Contrib{
			Libel: libel,
			Taux:  taux,
		})
	}
	return Co
}
