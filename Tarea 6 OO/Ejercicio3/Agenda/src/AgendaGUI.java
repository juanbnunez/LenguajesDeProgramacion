/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */

import Data.Agenda;
import Data.ObjAgenda;
import GUI.AgendaMain;

/**
 *
 * @author oviquez
 */

//  FORMATO PARA FECHAS DD-MM-AAAA
public class AgendaGUI {
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // Obtén la instancia única de la Agenda
        Agenda miAgenda = Agenda.getInstancia();

        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }
}



