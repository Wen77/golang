package contributions

//Cot ARRCO TA NC
func (t Behav) ArrTANC(Nta Amount) Amount {
	return Nta * t.Taux
}

//Cot ARRCO TB NC
func (t Behav) ArrTBNC(Ntb Amount) Amount {
	return Ntb * t.Taux
}

//Cot AGFF TA NC
func (t Behav) AgffTANC(Nagta Amount) Amount {
	return Nagta * t.Taux
}

//Cot AGFF TB NC
func (t Behav) AgffTBNC(Nagtb Amount) Amount {
	return Nagtb * t.Taux
}

//Cot prev TA NC
func (t Behav) PrevTANC(Pnta Amount) Amount {
	return Pnta * t.Taux
}

//Cot prev TB NC
func (t Behav) PrevTBNC(Pntb Amount) Amount {
	return Pntb * t.Taux
}

//Cot mut TA NC
func (t Behav) MutTANC(BmutTANC Amount) Amount {
	return BmutTANC * t.Taux
}
