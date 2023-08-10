package main

import (
	"fmt"
)

func rotacion(arreglo *[8]string, movimientos int, direccion int) {
	if direccion == 0 {
		for movimientos > 0 {
			temp := arreglo[0] //Primer elemento
			for i := 0; i <= len(arreglo)-2; i++ {
				arreglo[i] = arreglo[i+1]
			}
			arreglo[len(arreglo)-1] = temp
			movimientos--
		}
	} else {
		for movimientos > 0 {
			temp := arreglo[0] //Primer elemento
			for i := len(arreglo) - 1; i == 1; i++ {
				arreglo[i] = arreglo[i-1]
			}
			arreglo[0] = temp
			movimientos--
		}
	}

}

func main() {
	arreglo := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println(arreglo)
	rotacion(&arreglo, 4, 0)
	fmt.Println(arreglo)
}
