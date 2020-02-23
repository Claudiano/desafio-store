package services

import (
	"testing"
)

func TestFindBallanceById(t *testing.T) {

	var service = ServiceAccount{}

	res, err := service.FindBallanceById(15)
	if err != nil {
		t.Fatal(err)
	}

	if res != 0 {
		t.Errorf("Erro ao buscar saldo de clientes: esperava-se %v venho %v",
			0, res)
	}

}
