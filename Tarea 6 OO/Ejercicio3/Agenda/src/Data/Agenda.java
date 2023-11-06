package Data;

import java.util.LinkedList;

public class Agenda {
    private LinkedList<Object> listaObjetos;
    private LinkedList<String> contactos;
    private LinkedList<String> eventos;

    // Declaración de la instancia única
    private static Agenda instanciaUnica;

    private Agenda() {
        this.listaObjetos = new LinkedList<Object>();
        this.contactos = new LinkedList<String>();
        this.eventos = new LinkedList<String>();
    }

    // Método para obtener la instancia única (Lazy Initialization)
    public static Agenda getInstancia() {
        if (instanciaUnica == null) {
            instanciaUnica = new Agenda();
        }
        return instanciaUnica;
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }
}
