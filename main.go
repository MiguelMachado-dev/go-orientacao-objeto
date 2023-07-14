package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println(*contaDaCris == *contaDaCris2)
}
