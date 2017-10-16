package contributions

var arrcotacadee string
var arrcotacader string
var agfftacadee string
var agfftacader string
var prevtacadee string
var prevtacader string

//Cot ARRCO TA CAD
func (t Behav) ArrTACAD(Cta Amount) Amount {
	return Cta * t.Taux
}

//Cot AGFF TA CAD
func (t Behav) AgffTACAD(Cagta Amount) Amount {
	return Cagta * t.Taux
}

//Cot AGIRC TB CAD
func (t Behav) AgiTBCAD(Ctb Amount) Amount {
	return Ctb * t.Taux
}

//GMP
func (t Behav) GMP(BaseGMP Amount) Amount {
	return BaseGMP * t.Taux
}

//Cot AGFF TB CAD
func (t Behav) AgffTBCAD(Cagtb Amount) Amount {
	return Cagtb * t.Taux
}

//Cot AGIRC TC CAD
func (t Behav) AgiTCCAD(Ctc Amount) Amount {
	return Ctc * t.Taux
}

//CET
func (t Behav) CET(BaseCET Amount) Amount {
	return BaseCET * t.Taux
}

//Cot AGFF TC CAD
func (t Behav) AgffTCCAD(Cagtc Amount) Amount {
	return Cagtc * t.Taux
}

//Cot APEC TA
func (t Behav) ApecTA(Ctaap Amount) Amount {
	return Ctap * t.Taux
}

//Cot APEC TB
func (t Behav) ApecTB(Ctbap Amount) Amount {
	return Ctbp * t.Taux
}

//Cot prev TA CAD
func (t Behav) PrevTACAD(Pcta Amount) Amount {
	return Pcta * t.Taux
}

//Cot prev TB CAD
func (t Behav) PrevTBCAD(Pctb Amount) Amount {
	return Pctb * t.Taux
}

//Cot prev TC CAD
func (t Behav) PrevTCCAD(Pctc Amount) Amount {
	return Pctc * t.Taux
}

//Cot mut TA CAD
func (t Behav) MutTACAD(BmutTACAD Amount) Amount {
	return BmutTACAD * t.Taux
}
