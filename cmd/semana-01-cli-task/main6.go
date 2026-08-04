package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func produtor(entrada chan<- int) {

	for range 100 {
		entrada <- rand.IntN(61)
		time.Sleep(10 * time.Millisecond)
	}

}

func consumidor(entrada <-chan int, saida chan string) {

	maior := 0
	menor := 60
	sum := 0
	quantidade := 0

	for range 200 {
		temperatura := <-entrada
		sum += temperatura
		quantidade++

		if temperatura < menor {
			menor = temperatura
		}

		if temperatura > maior {
			maior = temperatura
		}
	}

	media := sum / quantidade

	saida <- fmt.Sprintf("Quantidade: %d, Maior: %d, Menor: %d, Media: %d", quantidade, maior, menor, media)

}
func main() {
	entrada := make(chan int, 100)
	saida := make(chan string)

	go produtor(entrada)
	go produtor(entrada)
	go consumidor(entrada, saida)

	fmt.Println(<-saida)
}
