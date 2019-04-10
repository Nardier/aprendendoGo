/* package main

//não completo

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	conjunto1 := make([]int, 0)
	conjunto2 := make([]int, 0)
	conjunto3 := make([]int, 0)
	conjunto1 = preencheConjunto(conjunto1, 1)
	limpaTela()
	conjunto2 = preencheConjunto(conjunto2, 2)
	limpaTela()
	conjunto3 = preencheConjunto(conjunto3, 3)
	limpaTela()

	var optionMenu int
	fmt.Println("Qual item deseja ver ?")
	fmt.Println("1- !(Conjunto1)")
	fmt.Println("2- !(Conjunto1 ^ Conjunto2) = !(Conjunto1) v !(Conjunto2)")
	fmt.Println("3- !(!(Conjunto1 v Conjunto2) ^ (Conjunto3))")

	fmt.Scan(&optionMenu)

	switch optionMenu {
	case 1:
		primeiroItem(conjunto1, conjunto2, conjunto3)
	case 2:
	case 3:

	}
}

func preencheConjunto(conjunto []int, numConj int) []int {
	var tamanhoConjunto int
	fmt.Println("Escreva qual o tamanho do conjunto ", numConj)
	fmt.Scan(&tamanhoConjunto)
	for i := 1; i <= tamanhoConjunto; i++ {
		var valor int
		fmt.Println("Escreva o ", i, "valor do conjunto")
		fmt.Scan(&valor)
		conjunto = append(conjunto, valor)
	}
	return conjunto
}

func limpaTela() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func primeiroItem(conj1 []int, conj2 []int, conj3 []int) {
	var complementar = make([]int, 0)
	for i := 0; i < len(conj1); i++ {
		for j := 0; j < len(conj2); j++ {
			if conj1[i] != conj2[j] {
				if !procuraValorVetor(complementar, conj2[j]) {
					complementar = append(complementar, conj2[j])
				}
			}
		}
	}

	for i := 0; i < len(conj1); i++ {
		for j := 0; j < len(conj3); j++ {
			if conj1[i] != conj3[j] {
				if !procuraValorVetor(complementar, conj3[j]) {
					complementar = append(complementar, conj3[j])
				}
			}
		}
	}
	fmt.Println("Complementar do conjunto 1 é")
	printaSlice(complementar)
}

func procuraValorVetor(conj []int, valor int) bool {
	for i := 0; i < len(conj); i++ {
		if conj[i] == valor {
			return true
		}
	}
	return false
}

func printaSlice(slice []int) {
	fmt.Print("{")
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Print("}")
}

func debug(conj []int) {
	fmt.Println("CHegou")
	var any int
	fmt.Println("Itens dentro", len(conj))
	fmt.Println("Capacidade", cap(conj))
	fmt.Scan(&any)
}
*/