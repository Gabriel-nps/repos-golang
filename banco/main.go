package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const arquivoSaldo = "saldo.txt"

func buscaSaldoEmArquivo() (float64, error) {
	data, err := os.ReadFile(arquivoSaldo)

	if err != nil {
		return 1000, errors.New("arquivo de saldo nao encontrado")
	}

	txtSaldo := string(data)
	saldo, err := strconv.ParseFloat(txtSaldo, 64)

	if err != nil {
		return 1000, errors.New("falha em converter o saldo armazenado")
	}

	return saldo, nil
}

func escreverSaldoEmArquivo(saldo float64) {
	txtSaldo := fmt.Sprint(saldo)
	os.WriteFile(arquivoSaldo, []byte(txtSaldo), 0644)
}

func main() {
	var saldoDisponivel, err = buscaSaldoEmArquivo()
	var opcao int
	var montanteDeposito float64
	var montanteSaque float64

	if err != nil {
		fmt.Println("ERRO")
		fmt.Println(err)
		fmt.Println("----------")
	}

	for {
		fmt.Println("Bem-vindo ao Banco GOLANG")
		fmt.Println("O que voce gostaria de fazer?")

		fmt.Println("1. Conferir saldo")
		fmt.Println("2. Depositar")
		fmt.Println("3. Sacar")
		fmt.Println("4. Sair")
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
			escreverSaldoEmArquivo(saldoDisponivel)
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
			escreverSaldoEmArquivo(saldoDisponivel)
		default:
			fmt.Println("Até logo!")
			fmt.Println("Obrigado por escolher o Banco GOLANG")
			return
		}
	}
}
