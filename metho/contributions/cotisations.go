package contributions

import (
	"fmt"
)

type Contrib struct {
	Libel string
	Taux  float64
}
type Behav struct {
	Taux float64
}

var arrcotancee string
var arrcotancer string

//Cot ARRCO TA NC

func (t Behav) ArrTANC(PlafretTANC float64) float64 {
	return PlafretTANC * t.Taux
}

//Cot mut TA NC
func (t Behav) MutTANC(BmutTANC float64) float64 {
	return BmutTANC * t.Taux
}

//create a type Coti which defines an interface which has four methods
type Coti interface {
	ArrTANC(PlafretTANC float64) float64
	MutTANC(BmutTANC float64) float64
}

/*As we range on tx, we could use the Key of variable tx in calc.go as a boolean
element to print the appropriate amount*/

func Printer(Key string, z Coti) {
	if Key == "arrcotancee" || Key == "arrcotancer" {
		fmt.Println(z.ArrTANC(PlafretTANC)) //amount
	} else {
		fmt.Println(z.MutTANC(BmutTANC)) //amount
	}
}
