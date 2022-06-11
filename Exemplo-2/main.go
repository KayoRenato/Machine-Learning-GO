package main

import "fmt"

func main() {

	//variavel do tipo int
	var numero int
	numero = 100
	println(numero)

	//variavel do tipo float
	var preco float32 = 25.
	println(preco)

	//variavel tipo string
	palavra := "Kayo Renato"
	println(palavra)

	//criando lista em Go
	lista_numeros_float := []float32{8.17, 12.23, 32.33}

	fmt.Println(lista_numeros_float)
	fmt.Println(len(lista_numeros_float))

	//atribuir valor a lista
	lista_numeros_float = append(lista_numeros_float, 12.11)

	fmt.Println(lista_numeros_float)
	fmt.Println(len(lista_numeros_float))

	//substituicao de valor da lista
	lista_numeros_float[0] = 22.22

	fmt.Println(lista_numeros_float)
	fmt.Println(len(lista_numeros_float))

}
