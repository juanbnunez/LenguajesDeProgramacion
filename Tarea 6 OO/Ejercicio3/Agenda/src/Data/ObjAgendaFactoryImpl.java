package Data;
import Data.ObjAgendaFactory;
import java.util.Date;
public class ObjAgendaFactoryImpl implements ObjAgendaFactory{

    @Override
    public ObjAgenda createContacto(String parentezco, String nombre, String telefono) {
        return new ContactoFamiliar(parentezco, nombre, telefono);
    }

    @Override
    public ObjAgenda createEventoReunion(int cantidadAsistentes, Date fecha, String nombre) {
        return new EventoReunion(cantidadAsistentes, fecha, nombre);
    }
}
