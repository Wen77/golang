package main

import (
	"fmt"
	"time"
)
type Month int
var daysBefore = [...]int32{
  	0,
  	31,
  	31 + 28,
  	31 + 28 + 31,
  	31 + 28 + 31 + 30,
  	31 + 28 + 31 + 30 + 31,
  	31 + 28 + 31 + 30 + 31 + 30,
  	31 + 28 + 31 + 30 + 31 + 30 + 31,
  	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
  	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
  	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
  	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
  	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
  }
	const (
        January Month = 1 + iota
        February
        March
        April
        May
        June
        July
        August
        September
        October
        November
        December
)
const (
			Year = 2017
)
const (
	datedisplay = "02/01/2006"
)

func main() {
	a := "01/02/2017"

	//Parse the value a and b according to the layout "datedisplay"
	t, _ := time.Parse(datedisplay, a)
	fmt.Printf("%s %s\n", t.Format("02/01/2006"))

	c:=Month(t.Month())
	d:=t.Year()

 	fmt.Println(daysIn(c,d))

}
func daysIn(a Month,b int) int {
  	return int(daysBefore[a] - daysBefore[a-1])
}
