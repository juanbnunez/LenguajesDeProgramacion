#Juan Bautista Nunez Parrales Ejercicio 2

class Libro:
    def __init__(self, codigo, titulo, autor, disponible=True):
        self.codigo = codigo
        self.titulo = titulo
        self.autor = autor
        self.disponible = disponible

    def __str__(self):
        estado = "Disponible" if self.disponible else "No disponible"
        return f"Código: {self.codigo}, Título: {self.titulo}, Autor: {self.autor}, Estado: {estado}"

class Socio:
    def __init__(self, numero, nombre, direccion):
        self.numero = numero
        self.nombre = nombre
        self.direccion = direccion
        self.prestamos = []

    def prestar_libro(self, libro, fecha):
        if libro.disponible:
            libro.disponible = False
            self.prestamos.append((libro, fecha))
            print(f"{self.nombre} ha prestado el libro '{libro.titulo}' en la fecha {fecha}")
        else:
            print(f"El libro '{libro.titulo}' no está disponible para préstamo.")

    def libros_prestados(self):
        return len(self.prestamos)

    def __str__(self):
        return f"Número de socio: {self.numero}, Nombre: {self.nombre}, Dirección: {self.direccion}, Libros Prestados: {self.libros_prestados()}"

def encontrar_socios_con_mas_de_tres_prestamos(socios):
    return list(filter(lambda socio: socio.libros_prestados() > 3, socios))

# Crear algunos libros y socios de prueba
libro1 = Libro(codigo="001", titulo="El Gran Gatsby", autor="F. Scott Fitzgerald")
libro2 = Libro(codigo="002", titulo="1984", autor="George Orwell")
libro3 = Libro(codigo="003", titulo="Cien años de soledad", autor="Gabriel García Márquez")

socio1 = Socio(numero=101, nombre="Juan", direccion="Calle A")
socio2 = Socio(numero=102, nombre="Maria", direccion="Calle B")
socio3 = Socio(numero=103, nombre="Pedro", direccion="Calle C")

# Realizar algunos préstamos
socio1.prestar_libro(libro1, fecha="2023-11-06")
socio2.prestar_libro(libro2, fecha="2023-11-06")
socio3.prestar_libro(libro3, fecha="2023-11-06")

# Mostrar el estado de los libros
print(libro1)
print(libro2)
print(libro3)

# Mostrar socios con más de 3 libros prestados
socios = [socio1, socio2, socio3]
socios_con_mas_de_tres_libros = encontrar_socios_con_mas_de_tres_prestamos(socios)
print("\nSocios con más de 3 libros prestados:")
for socio in socios_con_mas_de_tres_libros:
    print(socio)
