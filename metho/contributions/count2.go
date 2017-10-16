package contributions

//"fmt"

var SommeS, SommeP Amount

func Count2(Reco []Short, Ect map[string]Amount, Bcsgcsa, Bufs Amount) (map[string]Amount, Amount, Amount) {

	for _, cotis := range Reco {
		t := Behav{cotis.Taux}
		if cotis.Libel == "forfsocer" {
			Ect["forfsocer"] = t.Forfsoc(Bufs)
		} else if cotis.Libel == "csgdaaee" {
			Ect["csgdaaee"] = t.Ccdaa(Bccaa)
		} else if cotis.Libel == "csgdsaee" {
			Ect["csgdsaee"] = t.Ccdsa(Bcsgcsa)
		} else if cotis.Libel == "csgrdaaee" {
			Ect["csgrdaaee"] = t.Cci(Bccaa)
		} else if cotis.Libel == "csgrdsaee" {
			Ect["csgrdsaee"] = t.Ccisa(Bcsgcsa)
		}
	}
	SommeS := SomS(Ect)
	SommeP := SomP(Ect)
	//fmt.Println(SommeS, SommeP)
	return Ect, SommeS, SommeP
}
func SomS(Ect map[string]Amount) Amount {
	for k, v := range Ect {
		if k == "urfmalee" || k == "urfvtee" || k == "urfvtaee" {
			SommeS += v
		} else if k == "acee" || k == "arrcotancee" || k == "agfftancee" {
			SommeS += v
		} else if k == "prevtancee" || k == "muttancee" || k == "agfftancee" {
			SommeS += v
		} else if k == "csgdaaee" || k == "csgdsaee" || k == "csgrdaaee" || k == "csgrdsaee" {
			SommeS += v
		}
	}
	return SommeS
}
func SomP(Ect map[string]Amount) Amount {
	for k, v := range Ect {
		if k == "urfmaler" || k == "urfvter" || k == "urfvtaer" {
			SommeP += v
		} else if k == "urfafer" || k == "urfater" || k == "urftrspter" {
			SommeP += v
		} else if k == "urffnaler" || k == "csaer" || k == "cser" {
			SommeP += v
		} else if k == "acer" || k == "agser" || k == "arrcotancer" {
			SommeP += v
		} else if k == "agfftancer" || k == "prevtancer" || k == "muttancer" || k == "forfsocer" {
			SommeP += v
		}
	}
	return SommeP
}
