package contributions

import (
	"fmt"
)

var tx map[string]float64
var PlafretTANC float64
var BmutTANC float64

func Calc(Co []Contrib) {

	for _, cotis := range Co {

		tx := make(map[string]float64)

		tx[cotis.Libel] = cotis.Taux

		for Key, Val := range tx {
			fmt.Println(Key, "-", Val) //libelle and associated rate
			fmt.Printf("%T\n", Key)
			PlafretTANC = 2800.00
			BmutTANC = 100.00
			s := Behav{Val}
			Printer(Key, s)
		}
	}
}
