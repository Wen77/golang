# DATE : Parse, Format  and calculate the duration between two dates from date to date

```
package main

import (
	"fmt"
	"time"
)

const (
	datedisplay = "02/01/2006"
)

func main() {
	a := "01/09/2017"
	b := "30/05/2010"

	/*Parse the value a and b according to the layout "datedisplay"
	: mean recognize those values as dates*/

	t, _ := time.Parse(datedisplay, a)
	s, _ := time.Parse(datedisplay, b)
	fmt.Println(t, s)

	//Format t and s
	fmt.Printf("%s %s\n", t.Format("02/01/2006"), s.Format("02/01/2006"))

	//Find the difference between two days : indicating the number of day , month and year
	diffD := s.Day() - t.Day() + 1
	diffM := t.Month() - s.Month()
	diffY := t.Year() - s.Year()

	//Notice diffM is type Month which underlying type is an int
	fmt.Println("seniority of", diffD, "days", int(diffM), "months and", diffY, "years")
}
