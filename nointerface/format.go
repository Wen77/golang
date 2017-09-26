package main

import (
	"fmt"
	"strconv"
)

type Format struct { // create a type format
	value float64
}

func (p Format) Sp() {
	i := fmt.Sprintf("%.2f", p.value)
	f, _ := strconv.ParseFloat(i, 2)
	fmt.Println(f)
}

func main() {
	p1 := Format{5.8258}
	p1.Sp()

}
