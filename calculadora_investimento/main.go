package main

import (
	"fmt"
	"math"
)

func main() {
	var montanteInvestido float64 = 1000
	var anos float64 = 10
	taxaRetorno := 5.5

	var valorFuturo = montanteInvestido * math.Pow(1+taxaRetorno/100, anos)
	fmt.Println("Rendimento: R$", valorFuturo)
}
