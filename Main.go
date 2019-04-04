package main

import "fmt"

func main() {
	var livraria [5]string
	livraria[0] = "Nardier"
	livraria[1] = "Barbosa"
	livraria[2] = "de"
	livraria[3] = "Lira"
	livraria[4] = "Sampaio"

	//var ponteir [...]*string = &livraria[5]
	opcaoMenu := menu()

	switch opcaoMenu {
	case 1:
		primeiraOpcao(livraria)
	case 2:
		segundaOpcao(livraria)
	case 4:
		quartaOpcao(livraria)
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

func primeiraOpcao(vet [5]string) {
	var posicaoLivro int
	fmt.Println("Selecione a posição do livro de entre 1 e 5")
	fmt.Scan(&posicaoLivro)
	posicaoLivro = posicaoLivro - 1

	if posicaoLivro > 5 || posicaoLivro < 0 {
		fmt.Println("Opção inválida")
	} else {
		fmt.Println("Nº:", posicaoLivro+1, "Livro ->", vet[posicaoLivro])
	}

}

func segundaOpcao(vet [5]string) {
	for j := 0; j < 5; j++ {
		fmt.Println("Nº:", j, "Livro ->", vet[j])
	}
}

func quartaOpcao(vet [5]string) {
	var posicaoLivro int
	var novoLivro string
	fmt.Println("Selecione a posição do livro de entre 1 e 5 para sobrescrever")
	fmt.Scan(&posicaoLivro)
	posicaoLivro = posicaoLivro - 1

	if posicaoLivro > 5 || posicaoLivro < 0 {
		fmt.Println("Opção inválida")
	} else {
		fmt.Println("Digite o novo nome do livro")
		fmt.Scan(&novoLivro)
		vet[posicaoLivro] = novoLivro
		fmt.Println("Alterado com sucesso")
	}
	segundaOpcao(vet)
}
