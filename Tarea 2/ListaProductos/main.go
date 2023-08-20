package main

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

// -----------------------------método de agregar producto con puntero hacia la listaProductos-----------------------------
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	bandera := false
	for n := 0; n < len(lProductos); n++ {
		if nombre == (*l)[n].nombre {
			(*l)[n].cantidad += cantidad
			bandera = true
			if precio != (*l)[n].precio {
				(*l)[n].precio = precio
			}
		}
	}
	if !bandera {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista
func (l *listaProductos) incrementarProducto(nombre string, cantidad int) {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
		}
	}
}

/*
func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}*/

// modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err" conteniendo un 0 si no hay error y números mayores si existe algún
// tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la anterior función en funciones posteriores, realizando los cambios
// que el uso de la nueva función ameriten
func (l *listaProductos) buscarProducto(nombre string) (*producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	//no hacer un retorno de una copia, sino un retorno de un cursor que apunte hacia la posición del producto
	//así si se ocupa modificar se modifica la posición y no una copia
	var err = -1
	var p *producto = nil
	for i := 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			//& asigna la dirección
			p = &((*l)[i])
			err = 0
		}
	}
	//retorna la dirección del producto y si existe error o no
	return p, err
}

func (l *listaProductos) venderProducto(nombre string) {
	var prod, err = l.buscarProducto(nombre)
	if err != -1 {
		(*prod).cantidad = (*prod).cantidad - 1
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar dicha acción
		//modifique posteriormente para que se ingrese de parámetro la cantidad de productos a vender
	}
}

//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("café", 12, 5000)
	lProductos.incrementarProducto("arroz", 30)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	return nil
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	//venda productos
	//lProductos.venderProducto("arroz")
	//lProductos.venderProducto("frijoles")
	//fmt.Println("Lista actualizada: ", lProductos)
}
