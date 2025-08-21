package main

import (
	"fmt"
	"sync" // Pacote de sincronização
	"time"
)

func fale(pessoa string, texto string, vezes int, wg *sync.WaitGroup) {
	// 3. Avisa ao WaitGroup que esta goroutine terminou
	defer wg.Done()

	for i := 0; i < vezes; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// 1. Cria um WaitGroup
	var wg sync.WaitGroup

	// 2. Avisa que estamos esperando 2 goroutines terminarem
	wg.Add(2)

	go fale("Maria", "Ei...", 10, &wg)
	go fale("João", "Opa!", 10, &wg)

	// 4. Espera o contador do WaitGroup chegar a zero
	wg.Wait()

	fmt.Println("\nFim da conversa!")
}
