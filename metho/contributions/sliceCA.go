package contributions

import "time"

func (PlaftheoM Amount) CadTB() Amount {
	return PlaftheoM * 3.00
}
func (PlaftheoM Amount) CadTC() Amount {
	return PlaftheoM * 4.00
}

func SliceCAD(N Amount, a, b string, Out1C, Out2C, Out3C, Out4C, Out5C, Out6C, Out7C, Out8C, Out9C, Out10C, Out11C, Out12C, Out13C, Out14C, Out15C, Out16C, Out17C, Out18C, Out19C, Out20C, Out21C, Out22C, Out23C, Out24C, Out25C, Out26C, Out27C, Out28C, Out29C, Out30C chan Amount) {
	//***Tranches théoriques***
	//Urssaf
	PlaftheoM = Plaftheo.Prora(N, a, b)
	CumulPlaftheoM = CumulPlaftheoM1 + PlaftheoM
	//Pôle emploi
	GPM = PlaftheoM * 4.00
	GrossPEtheoM = GPM.Prora(N, a, b)
	CumulGrossPEtheo = CumulGrossPEtheoM1 + GrossPEtheoM

	//Retraite
	CADTAtheoM = PlaftheoM
	CumulCADTAtheoM = CumulCADTAtheoM1 + CADTAtheoM
	CADTBtheoM = PlaftheoM.CadTB()
	CumulCADTBtheoM = CumulCADTBtheoM1 + CADTBtheoM
	CADTCtheoM = PlaftheoM.CadTC()
	CumulCADTCtheoM = CumulCADTCtheoM1 + CADTCtheoM

	//AGFF
	AGFFCADTAtheoM = PlaftheoM
	CumulAGFFCADTAtheoM = CumulAGFFCADTAtheoM1 + AGFFCADTAtheoM
	AGFFCADTBtheoM = PlaftheoM.CadTB()
	CumulAGFFCADTBtheoM = CumulAGFFCADTBtheoM1 + AGFFCADTBtheoM
	AGFFCADTCtheoM = PlaftheoM.CadTC()
	CumulAGFFCADTCtheoM = CumulAGFFCADTCtheoM1 + AGFFCADTCtheoM

	//Prévoyance
	CADTAprvtheoM = PlaftheoM
	CumulCADTAprvtheoM = CumulCADTAprvtheoM1 + CADTAprvtheoM
	CADTBprvtheoM = PlaftheoM.CadTB()
	CumulCADTBprvtheoM = CumulCADTBprvtheoM1 + CADTBprvtheoM
	CADTCprvtheoM = PlaftheoM.CadTC()
	CumulCADTCprvtheoM = CumulCADTCprvtheoM1 + CADTCprvtheoM

	//***Tranches réelles***
	//Urssaf
	go AlgoCADUA(Gross, CumulGrossCADM1, CumulPlaftheoM1, Out1C, Out2C)

	//Pôle emploi
	go AlgoAC(Gross, CumulGrossPEM1, Out3C, Out4C)

	//Retraite
	go AlgoCADA(Gross, CumulGrossCADM1, CumulCADTAM1, Out5C, Out6C)
	go AlgoCADB(Rest1, CumulCADTBM1, Out7C, Out8C)
	go AlgoCADC(Rest3, CumulCADTCM1, Out9C, Out10C)

	//GMP
	go AlgoBPGMP(PlaftheoM, BaseGMPFM, Gross, Out11C, Out12C)

	//Agff
	go AlgoAGFFCADA(Gross, CumulGrossCADM1, CumulAGFFCADTAM1, Out13C, Out14C)
	go AlgoAGFFCADB(Rest5, CumulAGFFCADTBM1, Out15C, Out16C)
	go AlgoAGFFCADC(Rest7, CumulAGFFCADTCM1, Out17C, Out18C)

	//APEC
	go AlgoAPCA(Cta, Out19C, Out20C)
	go AlgoAPCB(Ctb, Out21C, Out22C)

	//CET
	go AlgoCET(Cta, Ctb, Ctc, Out23C, Out24C)

	//Prevoyance
	go AlgoCADprvA(Gross, CumulGrossNCM1, CumulPrvTACADM1, Out25C, Out26C)
	go AlgoCADprvB(Rest9, CumulPrvTBCADM1, Out27C, Out28C)
	go AlgoCADprvC(Rest11, CumulPrvTCCADM1, Out29C, Out30C)

	time.Sleep(5 * time.Millisecond)
}
