package main

import "strings"

func contadorCaracteres(texto string) {

	caracteres := len(texto)
	palabras := strings.Count(texto, " ") + 1

	lineas := strings.Count(texto, "\n")

	println("El texto tiene: ", caracteres, " caracteres")
	println("El texto tiene: ", palabras, " palabras")
	println("El texto tiene: ", lineas, " lineas")
}

func main() {
	var texto = "Hola me llamo Juan\n Tengo 20 años"
	println(texto)
	contadorCaracteres(texto)

}
