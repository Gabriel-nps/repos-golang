package main

import (
	"fmt"
	"math"
)

func main() {
	var montanteInvestido float64
	var anos float64
	var taxaRetorno float64
	const taxaInflacao = 2.5

	fmt.Print("Investimento: R$")
	fmt.Scan(&montanteInvestido)

	fmt.Print("Periodo em anos: ")
	fmt.Scan(&anos)

	fmt.Print("%  de retorno: ")
	fmt.Scan(&taxaRetorno)

	valorFuturo := montanteInvestido * math.Pow(1+taxaRetorno/100, anos)
	valorFuturoReal := valorFuturo / math.Pow(1+taxaInflacao/100, anos)

	txtValorFuturo := fmt.Sprintf("Rendimento bruto: R$%.1f\n", valorFuturo)
	txtValorFuturoReal := fmt.Sprintf("Rendimento com inflacao: R$%.1f\n", valorFuturoReal)

	fmt.Print(txtValorFuturo, txtValorFuturoReal)
}
