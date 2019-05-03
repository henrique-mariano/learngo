package main

import (
	"fmt"

	"time"

	"github.com/henrique-mariano/learngo/functions"
)

var c, python, java bool //pode estar num pacote ou a nível de função.

var k, j int = 1, 2 //a declaração pode incluir inicializadores, uma por variável

// eoq := 12 //não é possível utilizar a atribuição curta fora do escopo de funções

func main() {
	fmt.Println("Vamos aprenderrr!!")

	fmt.Println("A hora agora é:", time.Now())

	functions.PrintRand()

	functions.PrintProblems()

	functions.ExportedNames()

	fmt.Println("Soma:", functions.Add(25, 12))

	a, b := functions.Swap("hello", "world")
	fmt.Println("Swap:", a, b)

	x, y := functions.Split(17)
	fmt.Println("Split:", x, y) //é possível só chamar Printl(split(17))

	var i int //var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.
	fmt.Println("Variáveis:", i, c, python, java)

	var c, python, java = true, false, "no!" //se tiver um inicializador, o tipo pode ser omitido
	fmt.Println(k, j, c, python, java)

	complexodmais := 1 + 2i
	fmt.Println(complexodmais)
}
