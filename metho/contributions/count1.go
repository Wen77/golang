package contributions

import (
	"time"
)

//import "fmt"

//import "fmt"

/*import (
	"fmt"
	//"strconv"
)*/

func Count1(Reco []Short, Plurf, Gpe, Nta, Ntb, Nagta, Nagtb, Pnta, Pntb, Plurfc, Gpec, Cta, Ctb, Ctc, BaseGMP, Cagta, Cagtb, Cagtc, Ctap, Ctbp, BaseCET, Pcta, Pctb, Pctc Amount) map[string]Amount {
	Recor := make(map[string]Amount)
	for _, cotis := range Reco {
		t := Behav{cotis.Taux}
		if cotis.Libel == "urfmalee" {
			Recor["urfmalee"] = t.Urfmal(Gross)
		} else if cotis.Libel == "urfmaler" {
			Recor["urfmaler"] = t.Urfmal(Gross)
		} else if cotis.Libel == "urfvtee" {
			Recor["urfvtee"] = t.Urfvt(Gross)
		} else if cotis.Libel == "urfvter" {
			Recor["urfvter"] = t.Urfvt(Gross)
		} else if cotis.Libel == "urfvtaee" {
			Recor["urfvtaee"] = t.UrfvTA(Plurf)
		} else if cotis.Libel == "urfvtaer" {
			Recor["urfvtaer"] = t.UrfvTA(Plurf)
		} else if cotis.Libel == "urfafer" {
			Recor["urfafer"] = t.UrfAF(Gross)
		} else if cotis.Libel == "urfater" {
			Recor["urfater"] = t.UrfAT(Gross)
		} else if cotis.Libel == "urftrspter" {
			Recor["urftrspter"] = t.Urftrspt(Gross)
		} else if cotis.Libel == "urffnaler" {
			Recor["urffnaler"] = t.UrfFNAL(Gross)
		} else if cotis.Libel == "csaer" {
			Recor["csaer"] = t.Csa(Gross)
		} else if cotis.Libel == "cser" {
			Recor["cser"] = t.Cs(Gross)
		} else if cotis.Libel == "acee" {
			Recor["acee"] = t.Ac(Gpe)
		} else if cotis.Libel == "acer" {
			Recor["acer"] = t.Ac(Gpe)
		} else if cotis.Libel == "agser" {
			Recor["agser"] = t.Ac(Gpe)
		} else if cotis.Libel == "arrcotancee" {
			Recor["arrcotancee"] = t.ArrTANC(Nta)
		} else if cotis.Libel == "arrcotancer" {
			Recor["arrcotancer"] = t.ArrTANC(Nta)
		} else if cotis.Libel == "arrcotbncee" {
			Recor["arrcotbncee"] = t.ArrTBNC(Ntb)
		} else if cotis.Libel == "arrcotbncer" {
			Recor["arrcotbncer"] = t.ArrTBNC(Ntb)
		} else if cotis.Libel == "agfftancee" {
			Recor["agfftancee"] = t.AgffTANC(Nagta)
		} else if cotis.Libel == "agfftancer" {
			Recor["agfftancer"] = t.AgffTANC(Nagta)
		} else if cotis.Libel == "agfftbncee" {
			Recor["agfftbncee"] = t.AgffTBNC(Nagtb)
		} else if cotis.Libel == "agfftbncer" {
			Recor["agfftbncer"] = t.AgffTBNC(Nagtb)
		} else if cotis.Libel == "prevtancee" {
			Recor["prevtancee"] = t.PrevTANC(Pnta)
		} else if cotis.Libel == "prevtancer" {
			Recor["prevtancer"] = t.PrevTANC(Pnta)
		} else if cotis.Libel == "prevtbncee" {
			Recor["prevtbncee"] = t.PrevTBNC(Pntb)
		} else if cotis.Libel == "prevtbncer" {
			Recor["prevtbncer"] = t.PrevTBNC(Pntb)
		} else if cotis.Libel == "muttancee" {
			Recor["muttancee"] = t.MutTANC(BmutTANC)
		} else if cotis.Libel == "muttancer" {
			Recor["muttancer"] = t.MutTANC(BmutTANC)
		} else if cotis.Libel == "muttacadee" {
			Recor["muttacadee"] = t.MutTACAD(BmutTACAD)
		} else if cotis.Libel == "muttacader" {
			Recor["muttacader"] = t.MutTACAD(BmutTACAD)
		}
	}
	//fmt.Println(Recor)
	time.Sleep(15 * time.Millisecond)
	return Recor
}
