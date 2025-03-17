package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Isabelle", "Veiga", "Alencar"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

}
