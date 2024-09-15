package main

import (
	"fmt"
	"math"
)

const taxaInflacao = 2.5

func main() {
	montanteInvestido := getValorDoUsuario("Investimento: R$")
	taxaRetorno := getValorDoUsuario("%  de retorno: ")
	anos := getValorDoUsuario("Periodo em anos: ")

	valorFuturo, valorFuturoReal := calcularValores(montanteInvestido, taxaRetorno, anos)

	txtValorFuturo := fmt.Sprintf("Rendimento bruto: R$%.1f\n", valorFuturo)
	txtValorFuturoReal := fmt.Sprintf("Rendimento com inflacao: R$%.1f\n", valorFuturoReal)

	fmt.Print(txtValorFuturo, txtValorFuturoReal)
}

func getValorDoUsuario(texto string) float64 {
	var valorInserido float64
	fmt.Print(texto)
	fmt.Scan(&valorInserido)
	return valorInserido
}

func calcularValores(montanteInvestido, taxaRetorno, anos float64) (float64, float64) {
	valorFuturo := montanteInvestido * math.Pow(1+taxaRetorno/100, anos)
	valorFuturoReal := valorFuturo / math.Pow(1+taxaInflacao/100, anos)
	return valorFuturo, valorFuturoReal
}
