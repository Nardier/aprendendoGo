package main

import "fmt"

func main() {
	var livraria [5]int
	opcaoMenu := menu()

	switch opcaoMenu {
	case 1:
		primeiraOpcao(livraria)
	default:
	}
}

func menu() int {

	var menuOpcao int

	fmt.Println("######## Bem vindo à livraria ######## ")
	fmt.Println("1- Buscar livro")
	fmt.Println("2- Listar livros")
	fmt.Println("3- Cadastrar livros")
	fmt.Println("4- Editar livros")
	fmt.Println("5- Excluir livros")
	fmt.Println("0- Sair")
	fmt.Println("Escolha uma opção")
	fmt.Scan(&menuOpcao)

	return menuOpcao
}

func primeiraOpcao(vet [5]int) {
	fmt.Println("Ok")
}
