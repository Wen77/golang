package main

import (
	"github.com/Wen77/golang/metho/contributions"
)

func main() {
	Co := contributions.Ctrib("rate.csv")
	contributions.Calc(Co)
}
