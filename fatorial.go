#file.go2 fazer fatorial

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Função para calcular o fatorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Semente para números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Gerando número aleatório de 0 a 10
	num := rand.Intn(10)
	fmt.Printf("Número gerado: %d\n", num)

	// Calculando o fatorial
	fat := factorial(num)
	fmt.Printf("Fatorial de %d é: %d\n", num, fat)
}
