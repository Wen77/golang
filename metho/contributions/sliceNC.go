package contributions

import "time"

func (PlaftheoM Amount) NcTB() Amount {
	return PlaftheoM * 2.00
}
func SliceNC(N Amount, a, b string, Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, Out10, Out11, Out12, Out13, Out14, Out15, Out16 chan Amount) {
	//***Tranches théoriques***
	//Urssaf
	PlaftheoM = Plaftheo.Prora(N, a, b)
	CumulPlaftheoM = CumulPlaftheoM1 + PlaftheoM

	//Pôle emploi
	GPM = PlaftheoM * 4.00
	GrossPEtheoM = GPM.Prora(N, a, b)
	CumulGrossPEtheo = CumulGrossPEtheoM1 + GrossPEtheoM

	//Retraite NC
	NCTAtheoM = PlaftheoM
	CumulNCTAtheoM = CumulNCTAtheoM1 + NCTAtheoM
	NCTBtheoM = PlaftheoM.NcTB()
	CumulNCTBtheoM = CumulNCTBtheoM1 + NCTBtheoM

	//AGFF
	AGFFNCTAtheoM = PlaftheoM
	CumulAGFFNCTAtheoM = AGFFNCTAtheoM
	AGFFNCTBtheoM = AGFFNCTAtheoM.NcTB()
	CumulAGFFNCTBtheoM = CumulAGFFNCTBtheoM1 + AGFFNCTBtheoM

	//Prévoyance
	NCTAprvtheoM = PlaftheoM
	CumulNCTAprvtheoM = CumulNCTAprvtheoM1 + NCTAprvtheoM
	NCTBprvtheoM = NCTAprvtheoM.NcTB()
	CumulNCTBprvtheoM = CumulNCTBprvtheoM1 + NCTBprvtheoM

	//***Tranches réelles***
	//Urssaf Plafonnée
	go AlgoNCUA(Gross, CumulGrossNCM1, CumulPlaftheoM1, Out1, Out2)

	//Pôle emploi
	go AlgoAC(Gross, CumulGrossPEM1, Out3, Out4)

	//Retraite
	go AlgoNCA(Gross, CumulGrossNCM1, CumulNCTAM1, Out5, Out6)
	go AlgoNCB(Rest0, CumulNCTBM1, Out7, Out8)

	//AGFF
	go AlgoAGFFNCA(Gross, CumulGrossNCM1, CumulNCTAM1, Out9, Out10)
	go AlgoAGFFNCB(Rest2, CumulNCTBM1, Out11, Out12)

	//Prevoyance
	go AlgoNCprvA(Gross, CumulGrossNCM1, CumulPrvTANCM1, Out13, Out14)
	go AlgoNCprvB(Rest4, CumulPrvTBNCM1, Out15, Out16)
	time.Sleep(5 * time.Millisecond)
}
