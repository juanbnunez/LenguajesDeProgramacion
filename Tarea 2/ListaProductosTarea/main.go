package ListaProductosTarea

import "fmt"

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

// Hacer una función adicional que reciba una cantidad potencialmente infinita de productos previamente creados y los agregue a la lista

func (l *listaProductos) buscarProducto(nombre string) (producto, int) { //el retorno es el índice del producto encontrado y -1 si no existe
	var productoEncontrado producto
	var err int = 0
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			productoEncontrado = (*l)[i]
			return productoEncontrado, err
		}
	}
	err = -1
	return productoEncontrado, err
}

//REALIZADO
//para la funcion buscarProducto
//modifique la función para que esta vez solo retorne el producto como tal y que retorne otro valor adicional "err"
//conteniendo un 0 si no hay error y números mayores si existe algún
//tipo de error como por ejemplo que el producto no exista. Una vez implementada la nueva función, cambie los usos de la
//anterior función en funciones posteriores, realizando los cambios
//que el uso de la nueva función ameriten

func (l *listaProductos) venderProducto(nombre string, cantVenta int) {
	prod, index := l.buscarProducto(nombre)
	if prod != nil {
		if prod.cantidad >= cantVenta {
			prod.cantidad -= cantVenta

			if prod.cantidad == 0 {
				// Eliminar el producto de la lista
				copy(l.productos[index:], l.productos[index+1:])
				l.productos = l.productos[:len(l.productos)-1]
				fmt.Printf("Producto '%s' se agotó y fue eliminado de la lista.\n", prod.nombre)
			} else {
				fmt.Printf("Venta de %d unidades de '%s' realizada con éxito.\n", cantVenta, prod.nombre)
			}
		} else {
			fmt.Println("Cantidad insuficiente para realizar la venta")
		}
	} else {
		fmt.Println("Producto no encontrado")
	}
}
En este código, he implementado la eliminación del producto de la lista utilizando la función copy para eliminar el producto en la posición index y reorganizar los elementos restantes en el slice. Además, he agregado mensajes de notificación para cuando se realiza una venta exitosa o cuando un producto se agota y se elimina de la lista.

Recuerda que este código es un ejemplo simplificado y que podría necesitar ajustes según la estructura y requerimientos específicos de tu programa.







//haga una función para a partir del nombre del producto, se pueda modificar el precio del mismo Pero utilizando la función buscarProducto MODIFICADA PREVIAMENTE

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)

}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	return nil
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	//venda productos
}
