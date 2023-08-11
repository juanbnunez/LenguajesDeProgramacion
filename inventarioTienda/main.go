package main

import "fmt"

type calzado struct {
	modelo   string
	precio   int
	talla    int8
	cantidad int
}

func NuevoCalzado(modelo string, precio int, talla int8, cantidad int) calzado {
	var zapato calzado
	if talla >= 35 && talla <= 44 {
		zapato = calzado{
			modelo:   modelo,
			precio:   precio,
			talla:    talla,
			cantidad: cantidad,
		}
	} else {
		println("No se pudo agregar el modelo, la talla tiene que estar entre 35 y 44")
	}
	return zapato
}

func Venta(listaCalzado []calzado, modelo string, cantidad int) {
	for i, zapato := range listaCalzado {
		if zapato.modelo == modelo {
			if zapato.cantidad >= cantidad {
				listaCalzado[i].cantidad -= cantidad
				fmt.Println("Venta exitosa")
			} else {
				fmt.Println("No hay suficientes unidades para realizar la venta")
			}

		}
	}
}

func main() {
	var listaCalzado [15]calzado

	listaCalzado[0] = NuevoCalzado("Nike", 50000, 43, 22)
	listaCalzado[1] = NuevoCalzado("Adidas", 30000, 43, 0)
	listaCalzado[2] = NuevoCalzado("Reebok", 25900, 42, 22)
	listaCalzado[3] = NuevoCalzado("Puma", 100000, 35, 22)
	listaCalzado[4] = NuevoCalzado("Converse", 34500, 43, 22)
	listaCalzado[5] = NuevoCalzado("Vans", 27800, 43, 22)
	listaCalzado[6] = NuevoCalzado("New Balance", 62400, 43, 22)
	listaCalzado[7] = NuevoCalzado("Timberland", 44500, 43, 22)
	listaCalzado[8] = NuevoCalzado("Skechers", 68950, 43, 22)
	listaCalzado[9] = NuevoCalzado("Fila", 33500, 43, 22)
	listaCalzado[10] = NuevoCalzado("Under Armour", 200000, 43, 22)
	listaCalzado[11] = NuevoCalzado("Crocs", 45690, 43, 22)

	fmt.Print(listaCalzado[:], "\n")
	Venta(listaCalzado[:], "Nike", 22)
	Venta(listaCalzado[:], "Adidas", 3)
	Venta(listaCalzado[:], "Converse", 21)

	fmt.Println(listaCalzado[:])
}
