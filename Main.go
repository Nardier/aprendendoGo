// setxkbmap -model abnt2 -layout br
package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

type livraria struct {
	livro, autor, genero string
}

func main() {

	livraria := []string{"Nardier", "Barbosa", "de", "Lira", "Sampaio"}

	for {
		opcaoMenu := menu()

		switch opcaoMenu {
		case 1:
			limpaTela()
			primeiraOpcao(livraria)
		case 2:
			limpaTela()
			segundaOpcao(livraria)
		case 3:
			limpaTela()
			livraria = terceiraOpcao(livraria)
		case 4:
			limpaTela()
			quartaOpcao(livraria)
		case 5:
			limpaTela()
			quintaOpcao(livraria)
		case 9:
			limpaTela()
		case 0:
			return
		default:
		}
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
	fmt.Println("9- Limpar tela")
	fmt.Println("0- Sair")
	fmt.Println("Escolha uma opção")
	fmt.Scan(&menuOpcao)

	return menuOpcao
}

func primeiraOpcao(livraria []string) {
	var posicaoLivro int
	fmt.Println("Selecione a posição do livro de entre 1 e", len(livraria))
	fmt.Scan(&posicaoLivro)
	posicaoLivro = posicaoLivro - 1

	if posicaoLivro > 5 || posicaoLivro < 0 {
		fmt.Println("Opção inválida")
	} else {
		fmt.Println("Nº:", posicaoLivro+1, "Livro ->", livraria[posicaoLivro])
	}
	return
}

func segundaOpcao(livraria []string) {
	var i = 1
	for j := 0; j < len(livraria); j++ {
		fmt.Println("Nº:", i, "Livro ->", livraria[j])
		i++
	}
}

func quartaOpcao(vet []string) {
	var posicaoLivro int
	var novoLivro string
	fmt.Println("Selecione a posição do livro de entre 1 e 5 para sobrescrever")
	fmt.Scan(&posicaoLivro)
	posicaoLivro = posicaoLivro - 1

	if posicaoLivro > len(vet) {
		fmt.Println("Opção inválida")
	} else {
		fmt.Println("Digite o novo nome do livro")
		fmt.Scan(&novoLivro)
		vet[posicaoLivro] = novoLivro
		fmt.Println("Alterado com sucesso")
	}
	return
}

func terceiraOpcao(livraria []string) []string {
	var novoLivro string
	fmt.Println("Digite o nome do livro que deseja cadastrar")
	fmt.Scan(&novoLivro)
	livraria = append(livraria, novoLivro)
	fmt.Println("Alterado com sucesso")
	return livraria
}

func quintaOpcao(livraria []string) {
	var posLivro int
	segundaOpcao(livraria)
	fmt.Println("Digite o numero do livro que deseja excluir")
	fmt.Scan(&posLivro)
	posLivro = posLivro - 1
	copy(livraria[posLivro:], livraria[posLivro+1:])
	livraria[len(livraria)-1] = ""
	livraria = livraria[:len(livraria)-1]
}
func limpaTela() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func debug() {
	fmt.Println("Chegou")
	var any int
	fmt.Scan(&any)
}
