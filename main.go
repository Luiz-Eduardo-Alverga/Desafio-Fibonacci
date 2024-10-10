package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFibonacci(num int) bool {
	// Inicializando os dois primeiros números da sequência de Fibonacci
	a, b := 0, 1

	// Verifica se o número informado é 0 ou 1, que são parte da sequência
	if num == 0 || num == 1 {
		return true
	}

	// Calcula a sequência até encontrar um número maior ou igual ao informado
	for b < num {
		a, b = b, a+b
	}

	// Verifica se o número informado faz parte da sequência
	return b == num
}

func main() {
	// Criando um scanner para ler a entrada do terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Informe um número: ")
	input, _ := reader.ReadString('\n')

	// Removendo espaços e convertendo para int
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Por favor, insira um número válido.")
		return
	}

	// Verifica e imprime o resultado
	if isFibonacci(num) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", num)
	}
}
