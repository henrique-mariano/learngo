package main

import (
	"fmt"
	"time"
)

var c, python, java bool //pode estar num pacote ou a nível de função.

var k, j int = 1, 2 //a declaração pode incluir inicializadores, uma por variável

func main() {
	fmt.Println("Vamos aprenderrr!!")

	fmt.Println("A hora agora é:", time.Now())

	printRand()

	printProblems()

	exportedNames()

	fmt.Println("Soma:", add(25, 12))

	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b)

	x, y := split(17)
	fmt.Println("Split:", x, y) //é possível só chamar Printl(split(17))

	var i int //var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.
	fmt.Println("Variáveis:", i, c, python, java)

	var c, python, java = true, false, "no!" //se tiver um inicializador, o tipo pode ser omitido
	fmt.Println(k, j, c, python, java)
}
