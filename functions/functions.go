package functions

import ( //importação consignada
	"fmt"
	"math"
	"math/rand" //estamos importando arquivos que comecem com package rand
	"time"
)

/*Funções que começam com letra maiúscula podem ser exportados em um pacote para outro código,
  já funções que começam com letra minúscula não podem ser exportados são funções privadas do arquivo*/

func PrintRand() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("Meu número favorito é", r.Int31())
}

func PrintProblems() {
	fmt.Printf("Agora você tem %g problemas.\n", math.Sqrt(7)) //Printf para tratar as diretivas
}

func ExportedNames() {
	fmt.Println("Pi:", math.Pi) //nomes exportados começãm com letra maiúscula como Pi e E que são exportados do pacote math
	fmt.Println("E: ", math.E)
}

func Add(x, y int) int { //Os tipos vem após o nome da variável, isso torna mais legível quanto o que a função faz
	return x + y //Quando os tipos consecutivos da função possuem o mesmo tipo, você pode omitir todos os tipos menos o do último
} //lê-se da esquerda para direita

func Swap(x, y string) (string, string) { //uma função pode retornar qualquer número de resultados
	return y, x
}

func Split(sum int) (x, y int) { //os valores de retorno podem ser nomeados e agirem como variáveis,
	x = sum * 4 / 9 //o return sem arg é conhecido como retorno "limpo", retorna os valores atuais dos resultados
	y = sum - x     //deve-se ser usado apenas em funções curtas, pois elas podem prejudicar a legibilidade em funções longas
	return
}

func TypeConversion(x, y int) (float64, uint) {
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	return f, z
}

func Sum(elements []int) int {
	sum := 0
	for i := 0; i < len(elements); i++ { //o for é a única estrutura de laço em Go, e também não há parênteses nos componentes e sempre será utilizado as chaves
		sum += elements[i] //o for possui instrução inicial, expressão de condição e pós-instrução
	} //variáveis declaradas são apenas visíveis no escopo da declaração do for
	return sum
}
