package contributions

import (
	"time"
)

const (
	datedisplay = "02/01/2006"
)

type Amount float64

var DiffD int

func Base(Reco []Short) ([]Amount, map[string]Amount) {
	Gross = 2711.86
	Plaftheo = 3269.00
	BaseGMPFM = 342.48
	BmutTANC = 97.23
	BmutTACAD = 97.23
	TxAbt = 0.9825
	Bccaa = Gross * TxAbt
	N = 30.00 //number of calendar days for Urssaf

	/*Données de paie*/
	a := "01/06/2016"
	b := "30/06/2016"

	//Cumuls Gross NC & CAD
	CumulGrossNC = CumulGrossNCM1 + Gross
	CumulGrossCAD = CumulGrossCADM1 + Gross

	//call sliceNC
	Out1 := make(chan Amount)
	Out2 := make(chan Amount)
	Out3 := make(chan Amount)
	Out4 := make(chan Amount)
	Out5 := make(chan Amount)
	Out6 := make(chan Amount)
	Out7 := make(chan Amount)
	Out8 := make(chan Amount)
	Out9 := make(chan Amount)
	Out10 := make(chan Amount)
	Out11 := make(chan Amount)
	Out12 := make(chan Amount)
	Out13 := make(chan Amount)
	Out14 := make(chan Amount)
	Out15 := make(chan Amount)
	Out16 := make(chan Amount)

	SliceNC(N, a, b, Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, Out10, Out11, Out12, Out13, Out14, Out15, Out16)
	Plurf, Cuf := <-Out1, <-Out2
	Gpe, Cgpe := <-Out3, <-Out4
	Nta, Cnta := <-Out5, <-Out6
	Ntb, Cntb := <-Out7, <-Out8
	Nagta, Cnagta := <-Out9, <-Out10
	Nagtb, Cnagtb := <-Out11, <-Out12
	Pnta, CptncA := <-Out13, <-Out14
	Pntb, CptncB := <-Out15, <-Out16

	//call sliceCA
	Out1C := make(chan Amount)
	Out2C := make(chan Amount)
	Out3C := make(chan Amount)
	Out4C := make(chan Amount)
	Out5C := make(chan Amount)
	Out6C := make(chan Amount)
	Out7C := make(chan Amount)
	Out8C := make(chan Amount)
	Out9C := make(chan Amount)
	Out10C := make(chan Amount)
	Out11C := make(chan Amount)
	Out12C := make(chan Amount)
	Out13C := make(chan Amount)
	Out14C := make(chan Amount)
	Out15C := make(chan Amount)
	Out16C := make(chan Amount)
	Out17C := make(chan Amount)
	Out18C := make(chan Amount)
	Out19C := make(chan Amount)
	Out20C := make(chan Amount)
	Out21C := make(chan Amount)
	Out22C := make(chan Amount)
	Out23C := make(chan Amount)
	Out24C := make(chan Amount)
	Out25C := make(chan Amount)
	Out26C := make(chan Amount)
	Out27C := make(chan Amount)
	Out28C := make(chan Amount)
	Out29C := make(chan Amount)
	Out30C := make(chan Amount)

	SliceCAD(N, a, b, Out1C, Out2C, Out3C, Out4C, Out5C, Out6C, Out7C, Out8C, Out9C, Out10C, Out11C, Out12C, Out13C, Out14C, Out15C, Out16C, Out17C, Out18C, Out19C, Out20C, Out21C, Out22C, Out23C, Out24C, Out25C, Out26C, Out27C, Out28C, Out29C, Out30C)
	Plurfc, Cufc := <-Out1C, <-Out2C
	Gpec, Cgpec := <-Out3C, <-Out4C
	Cta, Ccta := <-Out5C, <-Out6C
	Ctb, Cctb := <-Out7C, <-Out8C
	Ctc, Cctc := <-Out9C, <-Out10C
	BaseGMP, Cbgmp := <-Out11C, <-Out12C
	Cagta, Ccagta := <-Out13C, <-Out14C
	Cagtb, Ccagtb := <-Out15C, <-Out16C
	Cagtc, Ccagtc := <-Out17C, <-Out18C
	Ctap, Cctaap := <-Out19C, <-Out20C
	Ctbp, Cctbap := <-Out21C, <-Out22C
	BaseCET, CCET := <-Out23C, <-Out24C
	Pcta, CPtba := <-Out25C, <-Out26C
	Pctb, CPtbb := <-Out27C, <-Out28C
	Pctc, CPtbc := <-Out29C, <-Out30C

	//call function Count to calculate the contributions
	Ect := Count1(Reco, Plurf, Gpe, Nta, Ntb, Nagta, Nagtb, Pnta, Pntb, Plurfc, Gpec, Cta, Ctb, Ctc, BaseGMP, Cagta, Cagtb, Cagtc, Ctap, Ctbp, BaseCET, Pcta, Pctb, Pctc)
	Cumul := make([]Amount, 0)
	Cumul = append(Cumul, Cuf, Cgpe, Cnta, Cntb, Cnagta, Cnagtb, CptncA, CptncB, Cufc, Cgpec, Ccta, Cctb, Cctc, Cbgmp, Ccagta, Ccagtb, Ccagtc, Cctaap, Cctbap, CCET, CPtba, CPtbb, CPtbc)

	return Cumul, Ect
}

//create a method that is going to be use in several contributions which is based on a Plaf
func (V Amount) Prora(N Amount, a, b string) Amount {
	s, _ := time.Parse(datedisplay, a)
	t, _ := time.Parse(datedisplay, b)
	DiffD = t.Day() - s.Day() + 1
	return Amount(V) * (Amount(DiffD) / N)
}
