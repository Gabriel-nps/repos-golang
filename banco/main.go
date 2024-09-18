package main

import (
	"fmt"

	"example.com/banco/fileops"
)

const arquivoSaldo = "saldo.txt"

func main() {
	var saldoDisponivel, err = fileops.BuscaSaldoEmArquivo(arquivoSaldo, 1200)
	var opcao int
	var montanteDeposito float64
	var montanteSaque float64

	if err != nil {
		fmt.Println("ERRO")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Bem-vindo ao Banco GOLANG")

	for {
		opcoes()
		fmt.Print("Digite sua opcao: ")
		fmt.Scan(&opcao)
		fmt.Println("Voce escolheu:", opcao)

		switch opcao {
		case 1:
			fmt.Println("Seu saldo disponivel: R$", saldoDisponivel)
		case 2:
			fmt.Print("Depositar: ")
			fmt.Scan(&montanteDeposito)

			if montanteDeposito <= 0 {
				fmt.Println("Valor invalido. Precisa ser maior que 0.")
				continue
			}

			saldoDisponivel += montanteDeposito
			fmt.Println("Saldo atualizado! Novo montante:", saldoDisponivel)
			fileops.EscreverSaldoEmArquivo(saldoDisponivel, arquivoSaldo)
		case 3:
			fmt.Print("Sacar: ")

			fmt.Scan(&montanteSaque)

			if montanteSaque <= 0 {
				fmt.Println("Valor invalido. Precisa ser maior que 0.")
				continue
			}

			if montanteSaque > saldoDisponivel {
				fmt.Println("Saque invalido. Voce nao pode sacar mais que possui.")
				continue
			}

			saldoDisponivel -= montanteSaque
			fmt.Println("Saldo atualizado! Novo montante:", saldoDisponivel)
			fileops.EscreverSaldoEmArquivo(saldoDisponivel, arquivoSaldo)
		default:
			fmt.Println("Até logo!")
			fmt.Println("Obrigado por escolher o Banco GOLANG")
			return
		}
	}
}
