package contributions

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Short struct {
	Libel string
	Taux  Amount
}

func Ctrib(filePath string, r chan []Short) {
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
	//create a slice type Short
	records := make([]Short, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}
		var libel string
		libel = row[0]
		//parse csv
		taux, _ := strconv.ParseFloat(row[1], 64)
		records = append(records, Short{
			Libel: libel,
			Taux:  Amount(taux),
		})
	}
	r <- records
	close(r)
}
