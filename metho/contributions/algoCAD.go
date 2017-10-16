package contributions

func AlgoCADUA(Gross, CumulGrossCADM1, CumulPlaftheoM1 Amount, Out1, Out2 chan Amount) {
	if Gross <= PlaftheoM || CumulPlaftheoM1+Plafurf <= CumulGrossCADM1+Gross {
		Plafurf = Gross
		CumulPlaftheoM = CumulPlaftheoM1 + Plafurf
	} else {
		Plafurf = PlaftheoM
		CumulPlaftheoM = CumulPlaftheoM1 + Plafurf
	}
	Out1 <- Plafurf
	Out2 <- CumulPlaftheoM
}

func AlgoCADA(Gross, CumulGrossCADM1, CumulCADTAM1 Amount, Out5, Out6 chan Amount) Amount {
	if Gross <= CADTAtheoM || CumulCADTAM1+Gross <= CumulGrossCADM1+Gross {
		CADTA = Gross
		Rest1 = Gross - CADTA
		CumulCADTA = CumulCADTAM1 + CADTA
	} else {
		CADTA = CADTAtheoM
		Rest1 = Gross - CADTA
		CumulCADTA = CumulCADTAM1 + CADTA
	}
	Out5 <- CADTA
	Out6 <- CumulCADTA
	return Rest1
}

func AlgoCADB(Rest1, CumulCADTBM1 Amount, Out7, Out8 chan Amount) Amount {
	if Rest1 <= CADTBtheoM || CumulCADTBM1+CADTBtheoM <= CumulGrossCADM1+Gross {
		CADTB = Rest1
		CumulCADTB = CumulCADTBM1 + CADTB
	} else {
		CADTB = CADTBtheoM
		Rest3 = Rest1 - CADTB
		CumulCADTB = CumulCADTBM1 + CADTB
	}
	Out7 <- CADTB
	Out8 <- CumulCADTB
	return Rest3
}

func AlgoCADC(Rest3, CumulCADTCM1 Amount, Out9, Out10 chan Amount) {
	if Rest3 <= CADTCtheoM || CumulCADTCM1+CADTCtheoM <= CumulGrossCADM1+Gross {
		CADTC = Rest3
		CumulCADTC = CumulCADTCM1 + CADTC
	} else {
		CADTC = CADTCtheoM
		CumulCADTC = CumulCADTCM1 + CADTC
	}
	Out9 <- CADTC
	Out10 <- CumulCADTC
}

func AlgoBPGMP(PlaftheoM, BaseGMPFM, Gross Amount, Out11, Out12 chan Amount) {
	if PlaftheoM+BaseGMPFM > Gross {
		if PlaftheoM+BaseGMPFM-Gross <= CADTB {
			BaseG = 0
			CumulBaseG = CumulBaseGM1 + BaseG
		} else {
			BaseG = PlaftheoM + BaseGMPFM - Gross - CADTB
			CumulBaseG = CumulBaseGM1 + BaseG
		}
		BaseG = 0
		CumulBaseG = CumulBaseGM1 + BaseG
	}
	Out11 <- BaseG
	Out12 <- CumulBaseG
}

func AlgoAGFFCADA(Gross, CumulGrossCADM1, CumulAGFFCADTAM1 Amount, Out13, Out14 chan Amount) Amount {
	if Gross <= AGFFCADTAtheoM || CumulAGFFCADTAM1+Gross <= CumulGrossCADM1+Gross {
		AGFFCADTA = Gross
		Rest5 = Gross - AGFFCADTA
		CumulAGFFCADTA = CumulAGFFCADTAM1 + AGFFCADTA
	} else {
		AGFFCADTA = AGFFCADTAtheoM
		Rest5 = Gross - AGFFCADTA
		CumulAGFFCADTA = CumulAGFFCADTAM1 + AGFFCADTA
	}
	Out13 <- AGFFCADTA
	Out14 <- CumulAGFFCADTA
	return Rest5
}

func AlgoAGFFCADB(Rest5, CumulAGFFCADTBM1 Amount, Out15, Out16 chan Amount) Amount {
	if Rest5 <= AGFFCADTBtheoM || CumulAGFFCADTBM1+CADTBtheoM <= CumulGrossCADM1+Gross {
		AGFFCADTB = Rest5
		CumulAGFFCADTB = CumulAGFFCADTBM1 + AGFFCADTB
	} else {
		AGFFCADTB = AGFFCADTBtheoM
		Rest7 = Rest5 - AGFFCADTB
		CumulAGFFCADTB = CumulAGFFCADTBM1 + AGFFCADTB
	}
	Out15 <- AGFFCADTB
	Out16 <- CumulAGFFCADTB
	return Rest7
}

func AlgoAGFFCADC(Rest7, CumulAGFFCADTCM1 Amount, Out17, Out18 chan Amount) {
	if Rest7 <= AGFFCADTCtheoM || CumulAGFFCADTCM1+CADTCtheoM <= CumulGrossCADM1+Gross {
		AGFFCADTC = Rest7
		CumulAGFFCADTC = CumulAGFFCADTCM1 + AGFFCADTC
	} else {
		AGFFCADTC = AGFFCADTCtheoM
		CumulAGFFCADTC = CumulAGFFCADTCM1 + AGFFCADTC
	}
	Out17 <- AGFFCADTC
	Out18 <- CumulAGFFCADTC
}

//APEC
func AlgoAPCA(Cta Amount, Out19, Out20 chan Amount) {
	Ctaap = Cta
	CumulCtaap = CumulCtaapM1 + Ctaap
	Out19 <- Ctaap
	Out20 <- CumulCtaap
}

func AlgoAPCB(Ctb Amount, Out21, Out22 chan Amount) {
	Ctbap = Ctb
	CumulCtbap = CumulCtbapM1 + Ctbap
	Out21 <- Ctbap
	Out22 <- CumulCtbap
}

//CET
func AlgoCET(Cta, Ctb, Ctc Amount, Out23, Out24 chan Amount) {
	BaseC = Cta + Ctb + Ctc
	CumulCet = CumulCetM1 + BaseC
	Out23 <- BaseC
	Out24 <- CumulCet
}

func AlgoCADprvA(Gross, CumulGrossNCM1, CumulPrvTACADM1 Amount, Out25, Out26 chan Amount) Amount {
	if Gross <= CADTAprvtheoM || CumulPrvTACADM1+CADTAprvtheoM <= CumulGrossCADM1+Gross {
		PrvTACAD = Gross
		Rest9 = Gross - PrvTACAD
		CumulPrvTACAD = CumulPrvTACADM1 + PrvTACAD
	} else {
		PrvTACAD = CADTAprvtheoM
		Rest9 = Gross - PrvTACAD
		CumulPrvTACAD = CumulPrvTACADM1 + PrvTACAD
	}
	Out25 <- PrvTACAD
	Out26 <- CumulPrvTACAD
	return Rest9
}

func AlgoCADprvB(Rest9, CumulPrvTBCADM1 Amount, Out27, Out28 chan Amount) Amount {
	if Rest9 <= CADTBprvtheoM || CumulPrvTBCADM1+PrvTBCAD <= CumulGrossCADM1+Gross {
		PrvTBCAD = Rest9
		CumulPrvTBCAD = CumulPrvTBCADM1 + PrvTBCAD
	} else {
		PrvTBCAD = CADTBprvtheoM
		Rest11 = Rest9 - PrvTBCAD
		CumulPrvTBCAD = CumulPrvTBCADM1 + PrvTBCAD
	}
	Out27 <- PrvTBCAD
	Out28 <- CumulPrvTBCAD
	return Rest11
}

func AlgoCADprvC(Rest11, CumulPrvTCCADM1 Amount, Out29, Out30 chan Amount) {
	if Rest11 <= CADTCprvtheoM || CumulPrvTCCADM1+PrvTCCAD <= CumulGrossCADM1+Gross {
		PrvTCCAD = Rest11
		CumulPrvTCCAD = CumulPrvTCCADM1 + PrvTCCAD
	} else {
		PrvTCCAD = CADTCprvtheoM
		CumulPrvTCCAD = CumulPrvTCCADM1 + PrvTCCAD
	}
	Out29 <- PrvTCCAD
	Out30 <- CumulPrvTCCAD
}
