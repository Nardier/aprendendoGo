package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./aprendendoGo.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS livraria (id INTEGER PRIMARY KEY, livro TEXT, autor TEXT, genero TEXT)")
	statement.Exec()
	/* statement, _ = database.Prepare("INSERT INTO livraria (livro, autor, genero) VALUES (?, ?, ?)")
	statement.Exec("Reidran", "Nardier", "Dev") */
	rows, err := database.Query("SELECT livro, autor, genero FROM livraria")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows)
	var id int
	var livro string
	var autor string
	var genero string
	for rows.Next() {
		rows.Scan(&id, &livro, &autor, &genero)
		fmt.Println(strconv.Itoa(id) + ": " + livro + " " + autor + " " + genero)
	}
}
