package main

import (
	"fmt"
	"math/cmplx"

	"time"

	"github.com/henrique-mariano/learngo/functions"
)

var c, python, java bool //pode estar num pacote ou a nível de função.

var k, j int = 1, 2 //a declaração pode incluir inicializadores, uma por variável

// eoq := 12 //não é possível utilizar a atribuição curta fora do escopo de funções

/*Tipos báiscos em Go
bool, string

int, int8, int16, int32, int64
uint, uint8, uint16, uint32, uint64, uintptr

byte //pseudônimo para uint8
rune //pseudônimo para int32
	 //representa um ponto de código Unicode

float32, float64

complex64, complex128
*/

//variáveis sem um valor inicial explicitado terão seu valor como zero
//0 para números, false para booleanos e "" para strings

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("Vamos aprenderrr!!")

	fmt.Println("A hora agora é:", time.Now())

	functions.PrintRand()

	functions.PrintProblems()

	functions.ExportedNames()

	fmt.Println("Soma:", functions.Add(25, 12)) //Println -> imprime um valor com uma quebra de linha no final

	a, b := functions.Swap("hello", "world") //atribuição multipla
	fmt.Println("Swap:", a, b)

	x, y := functions.Split(17)
	fmt.Println("Split:", x, y) //é possível só chamar Printl(split(17))

	var i int //var declara uma lista de variáveis, como em listas de argumentos de função, o tipo é o último passado.
	fmt.Println("Variáveis:", i, c, python, java)

	var c, python, java = true, false, "no!" //se tiver um inicializador, o tipo pode ser omitido
	fmt.Println(k, j, c, python, java)

	complexodmais := 1 + 2i
	fmt.Println(complexodmais)

	fmt.Printf("Tipo: %T Valor: %v\n", ToBe, ToBe)
	fmt.Printf("Tipo: %T Valor: %v\n", MaxInt, MaxInt)
	fmt.Printf("Tipo: %T Valor: %v\n", z, z)

	w, z := functions.TypeConversion(3, 4)
	fmt.Println("Type Conversion:", w, z) //conversões de tipo em Go precisam ser explícitas

	const Valor = "Oii"                    //constantes precisam ser declaradas utilizando a palavra-chave const, não é possível utilizar o operador := para definir uma constante
	fmt.Println("Valor constante:", Valor) //uma constante sem tipo pega o tipo pelo seu contexto.

	array := []int{2, 4, 6, 8, 10} //array de inteiros instanciados com os valores 2, 4, 6, 8, 10
	fmt.Println("Somatório de um array:", array, "Soma:", functions.Sum(array))
}
