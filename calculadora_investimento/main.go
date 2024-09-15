package main

import (
	"fmt"
	"math"
)

func main() {
	var montanteInvestido = 1000
	var taxaRetorno = 5.5
	var anos = 10

	var valorFuturo = float64(montanteInvestido) * math.Pow(1+taxaRetorno/100, float64(anos))
	fmt.Println("Rendimento: R$", valorFuturo)
}
