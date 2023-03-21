package main

import "fmt"

func main() {
	var indices []float32
	indices = []float32{3.0, 2.0, 1.0}
	fmt.Print("Inserte coeficiente/s: \n")
	fmt.Scanln(&indices)

	for indice := 0; indice < len(indices); indice++ {
		if indice == 0 {
			fmt.Printf("El polinomio completo es: %v ", indices[0])
		} else if indice == 1 {
			fmt.Printf("+ %v X ", indices[1])
		} else if indice >= 2 {
			fmt.Printf("+ %v X**%v ", indices[indice], indice)
		}
	}

	fmt.Print("\n")
}
