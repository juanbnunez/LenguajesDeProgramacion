#Juan Bautista Nunez Parrales Ejercicio 1

class Object:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def move(self, dx, dy):
        self.x += dx
        self.y += dy

class TextObject(Object):
    def __init__(self, x, y, text):
        super().__init__(x, y)
        self.text = text

    def edit_text(self, new_text):
        self.text = new_text

class GeometricObject(Object):
    def __init__(self, x, y):
        super().__init__(x, y)

class Circle(GeometricObject):
    def __init__(self, x, y, radius):
        super().__init__(x, y)
        self.radius = radius

class Ellipse(GeometricObject):
    def __init__(self, x, y, width, height):
        super().__init__(x, y)
        self.width = width
        self.height = height

class Rectangle(GeometricObject):
    def __init__(self, x, y, width, height):
        super().__init__(x, y)
        self.width = width
        self.height = height

class Line(GeometricObject):
    def __init__(self, x1, y1, x2, y2):
        super().__init__(x1, y1)
        self.x2 = x2
        self.y2 = y2

class Square(Rectangle):
    def __init__(self, x, y, side_length):
        super().__init__(x, y, side_length, side_length)

class Group(Object):
    def __init__(self, x, y):
        super().__init__(x, y)
        self.objects = []

    def add_object(self, obj):
        self.objects.append(obj)

    def move(self, dx, dy):
        super().move(dx, dy)
        for obj in self.objects:
            obj.move(dx, dy)

# Ejemplo de uso:
if __name__ == "__main__":
    # Crear objetos
    circle = Circle(10, 10, 5)
    ellipse = Ellipse(20, 20, 10, 15)
    rectangle = Rectangle(30, 30, 12, 8)
    line = Line(5, 5, 15, 15)
    square = Square(40, 40, 6)
    text_obj = TextObject(50, 50, "Hello, World!")

    # Crear un grupo y agregar objetos
    group1 = Group(0, 0)
    group1.add_object(circle)
    group1.add_object(ellipse)

    group2 = Group(100, 100)
    group2.add_object(rectangle)
    group2.add_object(line)

    # Agregar un grupo dentro de otro grupo
    nested_group = Group(200, 200)
    nested_group.add_object(square)
    nested_group.add_object(group1)

    # Mover objetos y grupos
    group2.move(10, 10)
    text_obj.move(5, 5)

    # Editar texto
    text_obj.edit_text("New Text")

    # Mostrar la posición de los objetos
    print(f"Circle Position: ({circle.x}, {circle.y})")
    print(f"Text Object Position: ({text_obj.x}, {text_obj.y})")

    # Imprimir los objetos dentro del grupo anidado
    for obj in nested_group.objects:
        if isinstance(obj, Circle):
            print("Circle in Nested Group")
        elif isinstance(obj, Group):
            print("Group in Nested Group")
