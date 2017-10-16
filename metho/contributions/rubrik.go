package contributions

//import "fmt"

var Bsa1, Bsa2, Bsa3, Bsa4, Bsa5, Bsa6, Bsa7, Bsa8, Bsa9, Bsa10, Bsa11, Bsa12, Bsa13, Bsa14, Bsa15, Bsa16, Bsa17, Bsa18, Bsa19, Bsa20 ContribA
var Bs1, Bs2, Bs3, Bs4, Bs5, Bs6, Bs7, Bs8, Bs9, Bs10, Bs11, Bs12, Bs13, Bs14, Bs15, Bs16, Bs17, Bs18, Bs19, Bs20 Contrib

func Rubrik(Reco []Short, Ecte map[string]Amount, Bcsgcsa, Bufs Amount, out5 chan []Contrib, out6 chan []ContribA) {
	Bs := make([]Contrib, 0)
	Bsa := make([]ContribA, 0)
	Nc := []string{
		"PAY/COT/N/100",
		"PAY/COT/N/110",
		"PAY/COT/N/120",
		"PAY/COT/N/130",
		"PAY/COT/N/140",
		"PAY/COT/N/150",
		"PAY/COT/N/160",
		"PAY/COT/N/170",
		"PAY/COT/N/180",
		"PAY/COT/N/200",
		"PAY/COT/N/210",
		"PAY/COT/N/300",
		"PAY/COT/N/400",
		"PAY/COT/N/700",
		"PAY/COT/N/800",
		"PAY/COT/N/900",
		"PAY/COT/N/910",
		"PAY/COT/N/920",
		"PAY/COT/N/930",
		"PAY/COT/N/940",
	}
	//create a slice type Contrib
	Cotisa := make(map[string]Amount)
	for _, v := range Reco {
		Cotisa[v.Libel] = v.Taux
	}
	for _, val := range Nc {
		if val == "PAY/COT/N/100" {
			Bsa1 := ContribA{
				Code:     val,
				Libel:    "Urssaf maladie",
				Base:     Gross,
				TauxS:    Format{Cotisa["urfmalee"]}.Pct(),
				MontantS: Format{Ecte["urfmalee"]}.Sp(),
				TauxP:    Format{Cotisa["urfmaler"]}.Pct(),
				MontantP: Format{Ecte["urfmaler"]}.Sp(),
			}
			Bs1 := Contrib{
				Code:     val,
				Libel:    "Urssaf maladie",
				Base:     Gross,
				TauxS:    Cotisa["urfmalee"],
				MontantS: Ecte["urfmalee"],
				TauxP:    Cotisa["urfmaler"],
				MontantP: Ecte["urfmaler"],
			}
			Bs = append(Bs, Bs1)
			Bsa = append(Bsa, Bsa1)
		} else if val == "PAY/COT/N/110" {
			Bsa2 := ContribA{
				Code:     val,
				Libel:    "Urssaf vieillesse Totalité",
				Base:     Gross,
				TauxS:    Format{Cotisa["urfvtee"]}.Pct(),
				MontantS: Format{Ecte["urfvtee"]}.Sp(),
				TauxP:    Format{Cotisa["urfvter"]}.Pct(),
				MontantP: Format{Ecte["urfvter"]}.Sp(),
			}
			Bs2 := Contrib{
				Code:     val,
				Libel:    "Urssaf vieillesse Totalité",
				Base:     Gross,
				TauxS:    Cotisa["urfvtee"],
				MontantS: Ecte["urfvtee"],
				TauxP:    Cotisa["urfvter"],
				MontantP: Ecte["urfvter"],
			}
			Bs = append(Bs, Bs2)
			Bsa = append(Bsa, Bsa2)
		} else if val == "PAY/COT/N/120" {
			Bsa3 := ContribA{
				Code:     val,
				Libel:    "Urssaf vieillesse TA",
				Base:     Plafurf,
				TauxS:    Format{Cotisa["urfvtaee"]}.Pct(),
				MontantS: Format{Ecte["urfvtaee"]}.Sp(),
				TauxP:    Format{Cotisa["urfvtaer"]}.Pct(),
				MontantP: Format{Ecte["urfvtaer"]}.Sp(),
			}
			Bs3 := Contrib{
				Code:     val,
				Libel:    "Urssaf vieillesse TA",
				Base:     Plafurf,
				TauxS:    Cotisa["urfvtaee"],
				MontantS: Ecte["urfvtaee"],
				TauxP:    Cotisa["urfvtaer"],
				MontantP: Ecte["urfvtaer"],
			}
			Bs = append(Bs, Bs3)
			Bsa = append(Bsa, Bsa3)
		} else if val == "PAY/COT/N/130" {
			Bsa4 := ContribA{
				Code:     val,
				Libel:    "Urssaf Allocation familliale",
				Base:     Gross,
				TauxP:    Format{Cotisa["urfafer"]}.Pct(),
				MontantP: Format{Ecte["urfafer"]}.Sp(),
			}
			Bs4 := Contrib{
				Code:     val,
				Libel:    "Urssaf Allocation familliale",
				Base:     Gross,
				TauxP:    Cotisa["urfafer"],
				MontantP: Ecte["urfafer"],
			}
			Bs = append(Bs, Bs4)
			Bsa = append(Bsa, Bsa4)
		} else if val == "PAY/COT/N/140" {
			Bsa5 := ContribA{
				Code:     val,
				Libel:    "Urssaf Accident de travail",
				Base:     Gross,
				TauxP:    Format{Cotisa["urfater"]}.Pct(),
				MontantP: Format{Ecte["urfafer"]}.Sp(),
			}
			Bs5 := Contrib{
				Code:     val,
				Libel:    "Urssaf Accident de travail",
				Base:     Gross,
				TauxP:    Cotisa["urfater"],
				MontantP: Ecte["urfafer"],
			}
			Bs = append(Bs, Bs5)
			Bsa = append(Bsa, Bsa5)
		} else if val == "PAY/COT/N/150" {
			Bsa6 := ContribA{
				Code:     val,
				Libel:    "Urssaf Versement transport",
				Base:     Gross,
				TauxP:    Format{Cotisa["urftrspter"]}.Pct(),
				MontantP: Format{Ecte["urftrspter"]}.Sp(),
			}
			Bs6 := Contrib{
				Code:     val,
				Libel:    "Urssaf Versement transport",
				Base:     Gross,
				TauxP:    Cotisa["urftrspter"],
				MontantP: Ecte["urftrspter"],
			}
			Bs = append(Bs, Bs6)
			Bsa = append(Bsa, Bsa6)
		} else if val == "PAY/COT/N/160" {
			Bsa7 := ContribA{
				Code:     val,
				Libel:    "Urssaf FNAL",
				Base:     Gross,
				TauxP:    Format{Cotisa["urffnaler"]}.Pct(),
				MontantP: Format{Ecte["urffnaler"]}.Sp(),
			}
			Bs7 := Contrib{
				Code:     val,
				Libel:    "Urssaf FNAL",
				Base:     Gross,
				TauxP:    Cotisa["urffnaler"],
				MontantP: Ecte["urffnaler"],
			}
			Bs = append(Bs, Bs7)
			Bsa = append(Bsa, Bsa7)
		} else if val == "PAY/COT/N/170" {
			Bsa8 := ContribA{
				Code:     val,
				Libel:    "Contribution solidarité autonomie",
				Base:     Gross,
				TauxP:    Format{Cotisa["csaer"]}.Pct(),
				MontantP: Format{Ecte["csaer"]}.Sp(),
			}
			Bs8 := Contrib{
				Code:     val,
				Libel:    "Contribution solidarité autonomie",
				Base:     Gross,
				TauxP:    Cotisa["csaer"],
				MontantP: Ecte["csaer"],
			}
			Bs = append(Bs, Bs8)
			Bsa = append(Bsa, Bsa8)
		} else if val == "PAY/COT/N/180" {
			Bsa9 := ContribA{
				Code:     val,
				Libel:    "Cotisation syndicale",
				Base:     Gross,
				TauxP:    Format{Cotisa["cser"]}.Pct(),
				MontantP: Format{Ecte["cser"]}.Sp(),
			}
			Bs9 := Contrib{
				Code:     val,
				Libel:    "Cotisation syndicale",
				Base:     Gross,
				TauxP:    Cotisa["cser"],
				MontantP: Ecte["cser"],
			}
			Bs = append(Bs, Bs9)
			Bsa = append(Bsa, Bsa9)
		} else if val == "PAY/COT/N/200" {
			Bsa10 := ContribA{
				Code:     val,
				Libel:    "Assurance chômage",
				Base:     GrossPE,
				TauxS:    Format{Cotisa["acee"]}.Pct(),
				MontantS: Format{Ecte["acee"]}.Sp(),
				TauxP:    Format{Cotisa["acer"]}.Pct(),
				MontantP: Format{Ecte["acer"]}.Sp(),
			}
			Bs10 := Contrib{
				Code:     val,
				Libel:    "Assurance chômage",
				Base:     GrossPE,
				TauxS:    Cotisa["acee"],
				MontantS: Ecte["acee"],
				TauxP:    Cotisa["acer"],
				MontantP: Ecte["acer"],
			}
			Bs = append(Bs, Bs10)
			Bsa = append(Bsa, Bsa10)
		} else if val == "PAY/COT/N/210" {
			Bsa11 := ContribA{
				Code:     val,
				Libel:    "Calcul AGS",
				Base:     GrossPE,
				TauxP:    Format{Cotisa["agser"]}.Pct(),
				MontantP: Format{Ecte["agser"]}.Sp(),
			}
			Bs11 := Contrib{
				Code:     val,
				Libel:    "Calcul AGS",
				Base:     GrossPE,
				TauxP:    Cotisa["agser"],
				MontantP: Ecte["agser"],
			}
			Bs = append(Bs, Bs11)
			Bsa = append(Bsa, Bsa11)
		} else if val == "PAY/COT/N/300" {
			Bsa12 := ContribA{
				Code:     val,
				Libel:    "ARRCO TA NC",
				Base:     PlafretTANC,
				TauxS:    Format{Cotisa["arrcotancee"]}.Pct(),
				MontantS: Format{Ecte["arrcotancee"]}.Sp(),
				TauxP:    Format{Cotisa["arrcotancer"]}.Pct(),
				MontantP: Format{Ecte["arrcotancer"]}.Sp(),
			}
			Bs12 := Contrib{
				Code:     val,
				Libel:    "ARRCO TA NC",
				Base:     PlafretTANC,
				TauxS:    Cotisa["arrcotancee"],
				MontantS: Ecte["arrcotancee"],
				TauxP:    Cotisa["arrcotancer"],
				MontantP: Ecte["arrcotancer"],
			}
			Bs = append(Bs, Bs12)
			Bsa = append(Bsa, Bsa12)
		} else if val == "PAY/COT/N/400" {
			Bsa13 := ContribA{
				Code:     val,
				Libel:    "AGFF TA NC",
				Base:     PlafretTANC,
				TauxS:    Format{Cotisa["agfftancee"]}.Pct(),
				MontantS: Format{Ecte["agfftancee"]}.Sp(),
				TauxP:    Format{Cotisa["agfftancer"]}.Pct(),
				MontantP: Format{Ecte["agfftancer"]}.Sp(),
			}
			Bs13 := Contrib{
				Code:     val,
				Libel:    "AGFF TA NC",
				Base:     PlafretTANC,
				TauxS:    Cotisa["agfftancee"],
				MontantS: Ecte["agfftancee"],
				TauxP:    Cotisa["agfftancer"],
				MontantP: Ecte["agfftancer"],
			}
			Bs = append(Bs, Bs13)
			Bsa = append(Bsa, Bsa13)
		} else if val == "PAY/COT/N/700" {
			Bsa14 := ContribA{
				Code:     val,
				Libel:    "Prev TA NC",
				Base:     Pnta,
				TauxS:    Format{Cotisa["prevtancee"]}.Pct(),
				MontantS: Format{Ecte["prevtancee"]}.Sp(),
				TauxP:    Format{Cotisa["prevtancer"]}.Pct(),
				MontantP: Format{Ecte["prevtancer"]}.Sp(),
			}
			Bs14 := Contrib{
				Code:     val,
				Libel:    "Prev TA NC",
				Base:     Pnta,
				TauxS:    Cotisa["prevtancee"],
				MontantS: Ecte["prevtancee"],
				TauxP:    Cotisa["prevtancer"],
				MontantP: Ecte["prevtancer"],
			}
			Bs = append(Bs, Bs14)
			Bsa = append(Bsa, Bsa14)
		} else if val == "PAY/COT/N/800" {
			Bsa15 := ContribA{
				Code:     val,
				Libel:    "Mut TA NC",
				Base:     BmutTANC + BmutTACAD,
				TauxS:    Format{Cotisa["muttancee"]}.Pct(),
				MontantS: Format{Ecte["muttancee"]}.Sp(),
				TauxP:    Format{Cotisa["muttancer"]}.Pct(),
				MontantP: Format{Ecte["muttancer"]}.Sp(),
			}
			Bs15 := Contrib{
				Code:     val,
				Libel:    "Mut TA NC",
				Base:     BmutTANC + BmutTACAD,
				TauxS:    Cotisa["muttancee"],
				MontantS: Ecte["muttancee"],
				TauxP:    Cotisa["muttancer"],
				MontantP: Ecte["muttancer"],
			}
			Bs = append(Bs, Bs15)
			Bsa = append(Bsa, Bsa15)
		} else if val == "PAY/COT/N/900" {
			Bsa16 := ContribA{
				Code:     val,
				Libel:    "Urssaf forfait social",
				Base:     Bufs,
				TauxP:    Format{Cotisa["forfsocer"]}.Pct(),
				MontantP: Format{Ecte["forfsocer"]}.Sp(),
			}
			Bs16 := Contrib{
				Code:     val,
				Libel:    "Urssaf forfait social",
				Base:     Bufs,
				TauxP:    Cotisa["forfsocer"],
				MontantP: Ecte["forfsocer"],
			}
			Bs = append(Bs, Bs16)
			Bsa = append(Bsa, Bsa16)
		} else if val == "PAY/COT/N/910" {
			Bsa17 := ContribA{
				Code:     val,
				Libel:    "Csg déductible avec abt",
				Base:     Bccaa,
				TauxS:    Format{Cotisa["csgdaaee"]}.Pct(),
				MontantS: Format{Ecte["csgdaaee"]}.Sp(),
			}
			Bs17 := Contrib{
				Code:     val,
				Libel:    "Csg déductible avec abt",
				Base:     Bccaa,
				TauxS:    Cotisa["csgdaaee"],
				MontantS: Ecte["csgdaaee"],
			}
			Bs = append(Bs, Bs17)
			Bsa = append(Bsa, Bsa17)
		} else if val == "PAY/COT/N/920" {
			Bsa18 := ContribA{
				Code:     val,
				Libel:    "Csg déductible sans abt",
				Base:     Bcsgcsa,
				TauxS:    Format{Cotisa["csgdsaee"]}.Pct(),
				MontantS: Format{Ecte["csgdsaee"]}.Sp(),
			}
			Bs18 := Contrib{
				Code:     val,
				Libel:    "Csg déductible sans abt",
				Base:     Bcsgcsa,
				TauxS:    Cotisa["csgdsaee"],
				MontantS: Ecte["csgdsaee"],
			}
			Bs = append(Bs, Bs18)
			Bsa = append(Bsa, Bsa18)
		} else if val == "PAY/COT/N/930" {
			Bsa19 := ContribA{
				Code:     val,
				Libel:    "Csg CRDS imposable avec abt",
				Base:     Bccaa,
				TauxS:    Format{Cotisa["csgrdaaee"]}.Pct(),
				MontantS: Format{Ecte["csgrdaaee"]}.Sp(),
			}
			Bs19 := Contrib{
				Code:     val,
				Libel:    "Csg CRDS imposable avec abt",
				Base:     Bccaa,
				TauxS:    Cotisa["csgrdaaee"],
				MontantS: Ecte["csgrdaaee"],
			}
			Bs = append(Bs, Bs19)
			Bsa = append(Bsa, Bsa19)
		} else if val == "PAY/COT/N/940" {
			Bsa20 := ContribA{
				Code:     val,
				Libel:    "Csg CRDS imposable sans abt",
				Base:     Bcsgcsa,
				TauxS:    Format{Cotisa["csgrdsaee"]}.Pct(),
				MontantS: Format{Ecte["csgrdsaee"]}.Sp(),
			}
			Bs20 := Contrib{
				Code:     val,
				Libel:    "Csg CRDS imposable sans abt",
				Base:     Bcsgcsa,
				TauxS:    Cotisa["csgrdsaee"],
				MontantS: Ecte["csgrdsaee"],
			}
			Bs = append(Bs, Bs20)
			Bsa = append(Bsa, Bsa20)
		}

	}

	out5 <- Bs
	out6 <- Bsa
}
