package contributions

func AlgoNCUA(Gross, CumulGrossNCM1, CumulPlaftheoM1 Amount, Out1, Out2 chan Amount) {
	if Gross <= Amount(PlaftheoM) || CumulPlaftheoM1+Plafurf <= CumulGrossNCM1+Gross {
		Plafurf = Gross
		CumulPlaftheoM = CumulPlaftheoM1 + Plafurf
	} else {
		Plafurf = Amount(PlaftheoM)
		CumulPlaftheoM = CumulPlaftheoM1 + Plafurf
	}
	Out1 <- Plafurf
	Out2 <- CumulPlaftheoM
}

func AlgoAC(Gross, CumulGrossPEM1 Amount, Out3, Out4 chan Amount) {
	if Gross <= GrossPEtheoM {
		GrossPE = Gross
		CumulGrossPE = CumulGrossPEM1 + GrossPE
	} else {
		GrossPE = GrossPEtheoM
		CumulGrossPE = CumulGrossPEM1 + GrossPE
	}
	Out3 <- GrossPE
	Out4 <- CumulGrossPE
}

func AlgoNCA(Gross, CumulGrossNCM1, CumulNCTAM1 Amount, Out5, Out6 chan Amount) Amount {
	if Gross <= NCTAtheoM || CumulNCTAM1+NCTAtheoM <= CumulGrossNCM1+Gross {
		NCTA = Gross
		Rest0 = Gross - NCTA
		CumulNCTA = CumulNCTAM1 + NCTA
	} else {
		NCTA = NCTAtheoM
		Rest0 = Gross - NCTA
		CumulNCTA = CumulNCTAM1 + NCTA
	}
	Out5 <- NCTA
	Out6 <- CumulNCTA
	return Rest0
}

func AlgoNCB(Rest0, CumulNCTBM1 Amount, Out7, Out8 chan Amount) {
	if Rest0 <= NCTBtheoM || CumulNCTBM1+NCTBtheoM <= CumulGrossNCM1+Gross {
		NCTB = Rest0
		CumulNCTB = CumulNCTBM1 + NCTB
	} else {
		NCTB = NCTBtheoM
		CumulNCTB = CumulNCTBM1 + NCTB
	}
	Out7 <- NCTB
	Out8 <- CumulNCTB
}
func AlgoAGFFNCA(Gross, CumulGrossNCM1, CumulAGFFNCTAM1 Amount, Out9, Out10 chan Amount) Amount {
	if Gross <= AGFFNCTAtheoM || CumulAGFFNCTAM1+AGFFNCTAtheoM <= CumulGrossNCM1+Gross {
		AGFFNCTA = Gross
		Rest2 = Gross - AGFFNCTA
		CumulAGFFNCTA = CumulAGFFNCTAM1 + AGFFNCTA
	} else {
		AGFFNCTA = AGFFNCTAtheoM
		Rest2 = Gross - AGFFNCTA
		CumulAGFFNCTA = CumulAGFFNCTAM1 + AGFFNCTA
	}
	Out9 <- AGFFNCTA
	Out10 <- CumulAGFFNCTA
	return Rest2
}

func AlgoAGFFNCB(Rest2, CumulNCTBM1 Amount, Out11, Out12 chan Amount) {
	if Rest2 <= AGFFNCTBtheoM || CumulAGFFNCTBM1+AGFFNCTBtheoM <= CumulGrossNCM1+Gross {
		AGFFNCTB = Rest2
		CumulAGFFNCTB = CumulAGFFNCTBM1 + AGFFNCTB
	} else {
		AGFFNCTB = AGFFNCTBtheoM
		CumulAGFFNCTB = CumulAGFFNCTBM1 + AGFFNCTB
	}
	Out11 <- AGFFNCTB
	Out12 <- CumulAGFFNCTB
}

func AlgoNCprvA(Gross, CumulGrossNCM1, CumulPrvTANCM1 Amount, Out13, Out14 chan Amount) Amount {
	if Gross <= NCTAprvtheoM || CumulPrvTANCM1+NCTAprvtheoM <= CumulGrossNCM1+Gross {
		PrvTANC = Gross
		Rest4 = Gross - PrvTANC
		CumulPrvTANC = CumulPrvTANCM1 + PrvTANC
	} else {
		PrvTANC = NCTAprvtheoM
		Rest4 = Gross - PrvTANC
		CumulPrvTANC = CumulPrvTANCM1 + PrvTANC
	}
	Out13 <- PrvTANC
	Out14 <- CumulPrvTANC
	return Rest4
}

func AlgoNCprvB(Rest4, CumulPrvTBNCM1 Amount, Out15, Out16 chan Amount) {
	if Rest4 <= NCTBprvtheoM || CumulPrvTBNCM1+NCTBtheoM <= CumulGrossNCM1+Gross {
		PrvTBNC = Rest4
		CumulPrvTBNC = CumulPrvTBNCM1 + PrvTBNC
	} else {
		PrvTBNC = NCTBprvtheoM
		CumulPrvTBNC = CumulPrvTBNCM1 + PrvTBNC
	}
	Out15 <- PrvTBNC
	Out16 <- CumulPrvTBNC
}
