package contributions

var forfsocer string
var csgdaaee string
var csgdsaee string
var csgrdaaee string
var csgrdsaee string

//Calcul Forfait social courant
func (f Behav) Forfsoc(Bufs Amount) Amount {
	return Bufs * f.Taux
}

//Calcul CSG CRDS déductible avec abattement
func (f Behav) Ccdaa(Bccaa Amount) Amount {
	return Bccaa * f.Taux
}

//Calcul CSG CRDS déductible sans abattement
func (f Behav) Ccdsa(Bcsgcsa Amount) Amount {
	return Bcsgcsa * f.Taux
}

//Calcul CSG CRDS imposable avec abattement
func (f Behav) Cci(Bccaa Amount) Amount {
	return Bccaa * f.Taux
}

//Calcul CSG CRDS imposable sans abattement
func (f Behav) Ccisa(Bcsgcsa Amount) Amount {
	return Bcsgcsa * f.Taux
}

//create a type Coti which defines an interface which has four methods
type Coti interface {
	Urfmal(Gross Amount) Amount
	Urfvt(Gross Amount) Amount
	UrfvTA(Plurf Amount) Amount
	UrfAF(Gross Amount) Amount
	UrfAT(Gross Amount) Amount
	Urftrspt(Gross Amount) Amount
	UrfFNAL(Gross Amount) Amount
	Csa(Gross Amount) Amount
	Cs(Gross Amount) Amount
	Ac(Gpe Amount) Amount
	Ags(Gpe Amount) Amount
	ArrTANC(Nta Amount) Amount
	ArrTBNC(Ntb Amount) Amount
	AgffTANC(Nagta Amount) Amount
	AgffTBNC(Nagtb Amount) Amount
	PrevTANC(Pnta Amount) Amount
	PrevTBNC(Pntb Amount) Amount
	ArrTACAD(Cta Amount) Amount
	AgffTACAD(Cagta Amount) Amount
	AgiTBCAD(Ctb Amount) Amount
	GMP(BaseGMP Amount) Amount
	AgffTBCAD(Cagtb Amount) Amount
	AgiTCCAD(Ctc Amount) Amount
	CET(BaseCET Amount) Amount
	AgffTCCAD(Cagtc Amount) Amount
	ApecTA(Ctaap Amount) Amount
	ApecTB(Ctbap Amount) Amount
	PrevTACAD(Pcta Amount) Amount
	PrevTBCAD(Pctb Amount) Amount
	PrevTCCAD(Pctc Amount) Amount
	MutTANC(BmutTANC Amount) Amount
	MutTACAD(BmutTACAD Amount) Amount
	Forfsoc(Bufs Amount) Amount
	Ccdaa(Bccaa Amount) Amount
	Ccdsa(Bcsgcsa Amount) Amount
	Cci(Bccaa Amount) Amount
	Ccisa(Bcsgcsa Amount) Amount
}
