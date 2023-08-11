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
	} else if direccion == 1 {
		for movimientos > 0 {
			temp := arreglo[len(arreglo)-1] // Guardar el último elemento
			for i := len(arreglo) - 1; i > 0; i-- {
				arreglo[i] = arreglo[i-1] // Mover los elementos a la derecha
			}
			arreglo[0] = temp // Colocar el último elemento al principio
			movimientos--
		}
	}

}

func main() {
	arreglo := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println(arreglo)
	rotacion(&arreglo, 3, 0)
	fmt.Println(arreglo)
}
