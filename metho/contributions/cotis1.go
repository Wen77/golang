package contributions

type Behav struct {
	Taux Amount
}

//Calcul Urssaf maladie
func (t Behav) Urfmal(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul Urssaf vieillesse Totalité
func (t Behav) Urfvt(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul Urssaf vieillesse TA
func (t Behav) UrfvTA(Plurf Amount) Amount {
	return Plurf * t.Taux
}

//Calcul allocation familliale
func (t Behav) UrfAF(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul Accident de travail
func (t Behav) UrfAT(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul versement transport
func (t Behav) Urftrspt(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul FNAL
func (t Behav) UrfFNAL(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul contribution solidarité autonomie
func (t Behav) Csa(Gross Amount) Amount {
	return Gross * t.Taux
}

//Calcul cotisation syndicale
func (t Behav) Cs(Gross Amount) Amount {
	return Gross * t.Taux
}

//Assurance chômage
func (t Behav) Ac(Gpe Amount) Amount {
	return Gpe * t.Taux
}

//Calcul AGS
func (t Behav) Ags(Gpe Amount) Amount {
	return Gpe * t.Taux
}
