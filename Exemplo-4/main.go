package main

import "fmt"

// Criação de uma Struct
type pessoa struct {
	nome  string
	idade int
}

// Função para nova pessoa
func novaPessoa(name string) *pessoa {
	p := pessoa{nome: name, idade: 34}

	return &p

}

func main() {
	fmt.Println(novaPessoa("Kayo"))

	fmt.Println(pessoa{"Bob", 12})

	fmt.Println(pessoa{nome: "Erika"})

	fmt.Println(pessoa{idade: 100})

	fmt.Println(&pessoa{"Ana", 12})

	fmt.Println(novaPessoa("Zico"))

	s := pessoa{"Maria", 56}

	fmt.Println("Pessoa: ", s)

	sp := &s

	fmt.Println("Ponteiro para idade de Pessoa: ", sp.idade)

	sp.idade = 11

	fmt.Println("Idade Atribuida de Pessoa: ", sp.idade)

}
