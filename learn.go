package main

import (
	"fmt"
	"time"
)

var c, python, java bool //pode estar num pacote ou a nível de função.

func main() {
	fmt.Println("Vamos aprenderrr!!")

	fmt.Println("A hora agora é:", time.Now())

	print_rand()

	print_problems()

	exported_names()

	fmt.Println("Soma:", add(25, 12))

	a, b := swap("hello", "world")
	fmt.Println("Swap:", a, b)

	x, y := split(17)
	fmt.Println("Split:", x, y) //é possível só chamar Printl(split(17))

	var i int //var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.
	fmt.Println("Variáveis:", i, c, python, java)
}