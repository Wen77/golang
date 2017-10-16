package contributions

import (
	"fmt"
	//"sort"
	"strconv"

	"golang.org/x/text/number"
)

type Format struct { // create a type format
	value Amount
}

type Contrib struct {
	Code     string
	Libel    string
	Base     Amount
	TauxS    Amount
	MontantS Amount
	TauxP    Amount
	MontantP Amount
}

type ContribA struct {
	Code     string
	Libel    string
	Base     Amount
	TauxS    number.Formatter
	MontantS Amount
	TauxP    number.Formatter
	MontantP Amount
}

func (n Format) Sp() Amount {
	i := fmt.Sprintf("%.2f", n.value)
	f, _ := strconv.ParseFloat(i, 2)
	return Amount(f)
}
func (n Format) Pct() number.Formatter {
	d := number.Percent(n, number.MaxFractionDigits(2))
	return d
}
func Printer(Reco []Short) {

	//launch and calculate the contributions with base0.go,cotis1.go,cotis2.go and count1.go
	Cumul, Ect := Base(Reco)

	//we obtain intermediate base1.go that are necessary to certain contributions
	Bcsgcsa, Bufs := Base1(Ect)
	//fmt.Println(Bcsgcsa, Bufs)

	//with base1.go calculate the last contributions using cotis3.go with count2.go
	//fmt.Println(Reco)
	Ect, Ee, Er := Count2(Reco, Ect, Bcsgcsa, Bufs) //retenues salariales et cotisations employeurs
	out7 := make(chan []Contrib)
	out8 := make(chan []ContribA)

	//We display the contributions on the paysheet with rubrik.go
	go Rubrik(Reco, Ect, Bcsgcsa, Bufs, out7, out8)
	Amount, Arr := <-out7, <-out8
	fmt.Println(Amount)
	fmt.Println(Arr)
	net := Gross - Ee
	fmt.Println("The net salary is :", net)
	fmt.Println(Ee, Er)
	fmt.Println(Cumul)
}
