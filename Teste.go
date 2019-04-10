// go get github.com/mattn/go-sqlite3
package main

import (
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
)

type livrariaTipo struct {
	livro, autor, genero                                                  string
	anoPublicacao, quantidadeEstoque, vendidos, alugados, paginas, edicao int
}

func main() {
	isso := "ok"
	livro := livrariaTipo{}
	livro1 := livrariaTipo{livro: isso}
	livraria := []livrariaTipo{}
	livraria = append(livraria, livro)
	livraria = append(livraria, livro1)
	fmt.Println(livraria[1].livro)
}
