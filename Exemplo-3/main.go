package main

import (
	"fmt"
)

func main() {

	// string vazia com 3 posições
	s := make([]string, 3)
	fmt.Println("Lista Vazia: ", s)

	//Atribui valores a cada elemento da lista
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Nova Lista: ", s)

	// Adicionando valores na lista
	s = append(s, "Kayo")
	s = append(s, "Erika")

	fmt.Println("Lista com Itens add: ", s)

	// Criacao de nova lista
	c := make([]string, len(s))

	// Copia de uma lista existente
	copy(c, s)
	fmt.Println("Lista Copiada: ", c)
	fmt.Println("Tamanho da Lista Copiada: ", len(c))

	// Slice da lista

	fmt.Println("Lista Fatiada até posicao 1", s[:2])          //a b
	fmt.Println("Lista Fatiada da posicao 3 em diante", s[3:]) //Kayo Erika
	fmt.Println("Lista Fatiada da posicao 2 ate 3", s[2:4])    // c Kayo

	// Criando uma lista com valores inicializados
	t := []string{"g", "h", "i"}
	fmt.Println("Lista Valores Inicializados: ", t)

	//Criação de Matriz
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Matriz :", twoD)
}
