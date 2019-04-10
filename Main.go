// setxkbmap -model abnt2 -layout br
package main

import (
	"fmt"
	"os"
	"os/exec"
	//_ "github.com/mattn/go-sqlite3"
)

type livroTipo struct {
	livro, autor, genero                                                  string
	anoPublicacao, quantidadeEstoque, vendidos, alugados, paginas, edicao int
}

func main() {

	livrariaSlice := []livroTipo{}

	for {
		opcaoMenu := menu()

		switch opcaoMenu {
		case 1:
			limpaTela()
			primeiraOpcao(livrariaSlice)
		case 2:
			limpaTela()
			segundaOpcao(livrariaSlice)
		case 3:
			limpaTela()
			livrariaSlice = append(livrariaSlice, terceiraOpcao())
			limpaTela()
		case 4:
			limpaTela()
			quartaOpcao(livrariaSlice)
		case 5:
			limpaTela()
			quintaOpcao(livrariaSlice)
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

func primeiraOpcao(livrariaSlice []livroTipo) {
	var posicaoLivro int
	fmt.Println("Selecione a posição do livro de entre 1 e", len(livrariaSlice))
	fmt.Scan(&posicaoLivro)
	if posicaoLivro > len(livrariaSlice) || posicaoLivro == 0 {
		fmt.Println("Opção inválida")
	} else {
		posicaoLivro = posicaoLivro - 1
		layoutDoPrintDoLivroCompleto(livrariaSlice[posicaoLivro])
	}
	return
}

func segundaOpcao(livrariaSlice []livroTipo) {
	for j := 0; j < len(livrariaSlice); j++ {
		layoutDoPrintDoLivroResumido(livrariaSlice[j])
	}
}

func quartaOpcao(livrariaSlice []livroTipo) {
	var posicaoLivro, opcaoEditar, valorGenericoInt int
	var valorGenericoString string
	fmt.Println("Selecione a posição do livro para editar")
	fmt.Scan(&posicaoLivro)
	fmt.Println("Digite qual o codigo correspondente ao elemento à editar")
	fmt.Println("1-Livro, 2-Autor, 3-Genero, 4-Publicacao, 5-quantidadeEstoque, 6-vendidos, 7-alugados, 8-paginas, 9-edicao")
	fmt.Scan(&opcaoEditar)
	if posicaoLivro > len(livrariaSlice) {
		fmt.Println("Opção inválida")
	} else {
		posicaoLivro = posicaoLivro - 1
		switch opcaoEditar {
		case 1:
			fmt.Println("Digite o novo titulo do livro ->")
			fmt.Scan(&valorGenericoString)
			livrariaSlice[posicaoLivro].livro = valorGenericoString
		case 2:
			fmt.Println("Digite o novo autor do livro ->")
			fmt.Scan(&valorGenericoString)
			livrariaSlice[posicaoLivro].autor = valorGenericoString
		case 3:
			fmt.Println("Digite o novo genero do livro ->")
			fmt.Scan(&valorGenericoString)
			livrariaSlice[posicaoLivro].genero = valorGenericoString
		case 4:
			fmt.Println("Digite o novo ano de publicação do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].anoPublicacao = valorGenericoInt
		case 5:
			fmt.Println("Digite o novo estoque do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].quantidadeEstoque = valorGenericoInt
		case 6:
			fmt.Println("Digite a nova quantidade de vendas do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].vendidos = valorGenericoInt
		case 7:
			fmt.Println("Digite a nova quantidade de aluguéis do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].alugados = valorGenericoInt
		case 8:
			fmt.Println("Digite a nova quantidade de páginas do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].paginas = valorGenericoInt
		case 9:
			fmt.Println("Digite a nova edição do livro ->")
			fmt.Scan(&valorGenericoInt)
			livrariaSlice[posicaoLivro].edicao = valorGenericoInt
		default:
			fmt.Println("Opção inválida!")
		}
		fmt.Println("Alterado com sucesso")
	}
	return
}

func terceiraOpcao() livroTipo {
	var livroCad, autorCad, generoCad string
	var anoPublicacaoCad, quantidadeEstoqueCad, paginasCad, edicaoCad int
	fmt.Println("Digite o nome do livro que deseja cadastrar")
	fmt.Scan(&livroCad)
	fmt.Println("Digite o autor do livro que deseja cadastrar")
	fmt.Scan(&autorCad)
	fmt.Println("Digite o genero do livro que deseja cadastrar")
	fmt.Scan(&generoCad)
	fmt.Println("Digite o ano de publicacao do livro que deseja cadastrar")
	fmt.Scan(&anoPublicacaoCad)
	fmt.Println("Digite a quantidade a cadastrar do livro que deseja cadastrar")
	fmt.Scan(&quantidadeEstoqueCad)
	fmt.Println("Digite a quantidade de paginas do livro que deseja cadastrar")
	fmt.Scan(&paginasCad)
	fmt.Println("Digite a edicao do livro que deseja cadastrar")
	fmt.Scan(&edicaoCad)
	var livro = livroTipo{
		livro:             livroCad,
		autor:             autorCad,
		genero:            generoCad,
		anoPublicacao:     anoPublicacaoCad,
		quantidadeEstoque: quantidadeEstoqueCad,
		paginas:           paginasCad,
		edicao:            edicaoCad,
	}
	return livro
}

func quintaOpcao(livrariaSlice []livroTipo) {
	fmt.Println("Pensando ainda como vou implementar")
	/* var posLivro int
	//	segundaOpcao(livrariaSlice)
	fmt.Println("Digite o numero do livro que deseja excluir")
	fmt.Scan(&posLivro)
	posLivro = posLivro - 1
	copy(livrariaSlice[posLivro:], livrariaSlice[posLivro+1:])
	livrariaSlice[len(livrariaSlice)-1] = ""
	livrariaSlice = livrariaSlice[:len(livrariaSlice)-1] */
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

func layoutDoPrintDoLivroCompleto(livrariaSlice livroTipo) {
	fmt.Printf("|Livro|: %s, |Autor|: %s, |Gênero|: %s, |Publicacao(Ano)|: %d, |Estoque|: %d, |Vendidos|: %d, |Alugados|: %d, |Páginas|: %d, |edicao|: %d \n",
		livrariaSlice.livro, livrariaSlice.autor, livrariaSlice.genero, livrariaSlice.anoPublicacao, livrariaSlice.quantidadeEstoque, livrariaSlice.vendidos,
		livrariaSlice.alugados, livrariaSlice.paginas, livrariaSlice.edicao)
}
func layoutDoPrintDoLivroResumido(livrariaSlice livroTipo) {
	fmt.Printf("|Livro|: %s, \n", livrariaSlice.livro)
}
