package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// forma mais verbosa de fazer e quando não precisamos de todos os campos
	contaDaKarlla := ContaCorrente{titular: "Karlla", numeroAgencia: 589, numeroConta: 123456, saldo: 325.5}
	fmt.Println(contaDaKarlla)

	// usada quando precisamos preencher todos os campos
	contaDaLadjane := ContaCorrente{"Ladjane", 231, 654321, 829.7}
	fmt.Println(contaDaLadjane)

	// usada quando não precisamos peencher todos os campos
	contaDoYoda := ContaCorrente{titular: "yoda", saldo: 22.9}
	fmt.Println(contaDoYoda)

	// utiliza de ponteiros para referenciar a variável utilizada
	var contaDaGabi *ContaCorrente
	contaDaGabi = new(ContaCorrente)
	contaDaGabi.titular = "Gabi"
	contaDaGabi.saldo = 600
	fmt.Println(contaDaGabi)
}
