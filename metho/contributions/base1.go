package contributions

//import "fmt"

var Bcsgcsa, Bufs Amount

// Base CSG CRDS sans abattement

func Base1(Ect map[string]Amount) (Amount, Amount) {
	for lib, cotis := range Ect {
		if lib == "muttancer" {
			Bcsgcsa += cotis
			Bufs += cotis
		} else if lib == "prevtancer" {
			Bcsgcsa += cotis
			Bufs += cotis
		}
	}
	return Bcsgcsa, Bufs
}
