package main

import (
    "client/globals"
    "client/utils"
    "log"
    "os"
    "bufio" 
)

func leerLineas() []string {
	reader := bufio.NewReader(os.Stdin)
	var lineas []string

	log.Println("Ingrese líneas. Enter vacío para finalizar:")

	for {
		text, _ := reader.ReadString('\n')

		if text == "" {
			break
		}

		log.Println("Leído:", text)
		lineas = append(lineas, text)
	}

	return lineas
}

func main() {
    utils.ConfigurarLogger()

    // loggear "Soy un Log" usando la biblioteca log
    globals.ClientConfig = utils.IniciarConfiguracion("config.json")

    log.Println("Soy un Log")

    // validar que la config este cargada correctamente
    if globals.ClientConfig == nil { //TODO cuando ocurre?
        log.Fatalf("No se pudo cargar la configuración")
    }
    
    // loggeamos el valor de la config
    log.Println("Mensaje desde config.json:", globals.ClientConfig.Mensaje)

    //lea de consola todas las líneas que se ingresen, las loguee y, si se ingresa una línea vacía, termine con el programa.
    
    /* reader := bufio.NewReader(os.Stdin)
    for{ 
        text, _ := reader.ReadString('\n')
        log.Println(text)
        if text == "\n"{
            log.Println("Salida del programa.")
            return
        }
    } */
    
    utils.LeerConsola()

    // ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

    // enviar un mensaje al servidor con el valor de la config
    log.Println("Enviando mensaje al servidor...")
    utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

    // leer de la consola el mensaje
    utils.LeerConsola()

    // generamos un paquete y lo enviamos al servidor
    utils.GenerarYEnviarPaquete()
}