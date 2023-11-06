package Data;

import java.util.LinkedList;
import java.util.Date;

public interface ObjAgendaFactory {
    ObjAgenda createContacto(String parentezco, String nombre, String telefono);

    ObjAgenda createEventoReunion(int cantidadAsistentes, Date fecha, String nombre);
}