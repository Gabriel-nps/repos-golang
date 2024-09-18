package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func BuscaSaldoEmArquivo(arquivo string, valorPadrao float64) (float64, error) {
	data, err := os.ReadFile(arquivo)

	if err != nil {
		return valorPadrao, errors.New("arquivo nao encontrado")
	}

	txtSaldo := string(data)
	saldo, err := strconv.ParseFloat(txtSaldo, 64)

	if err != nil {
		return 1000, errors.New("falha em converter o saldo armazenado")
	}

	return saldo, nil
}

func EscreverSaldoEmArquivo(saldo float64, arquivo string) {
	txtSaldo := fmt.Sprint(saldo)
	os.WriteFile(arquivo, []byte(txtSaldo), 0644)
}
