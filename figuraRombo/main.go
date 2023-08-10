package main

func Rombo(lineaCentroSize int) {
	if lineaCentroSize%2 == 0 {
		println("La linea central debe tener una cantidad impar de asteriscos")
	} else {
		for i := 0; i <= lineaCentroSize; i++ {
			for j := lineaCentroSize; j >= 0; j-- {
				print("*")
			}
			println()
		}
	}
}

func main() {
	Rombo(5)
}
