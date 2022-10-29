package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func Comunicacion(DataTipo string) {
	connS, err := grpc.Dial("dist041:50051", grpc.WithInsecure())

	if err != nil {
		panic("No se pudo conectar con el servidor" + err.Error())
	}

	defer connS.Close()

	serviceCliente := pb.NewMessageServiceClient(connS)

	//envia el mensaje al laboratorio
	res, err := serviceCliente.Intercambio(context.Background(),
		&pb.Message{
			Body: "1:" + DataTipo,
		})

	if err != nil {
		panic("No se puede crear el mensaje " + err.Error())
	}

	fmt.Println(res.Body)       //respuesta del laboratorio
	time.Sleep(2 * time.Second) //espera de 5 segundos
}

func main() {
	flag := true
	for flag {
		MenuInicio := "Rebeldes\n	[ 1 ] Consultar Datos\n	[ 2 ] Cerrar programa\nIngrese su opción:"
		MenuConsulta := "Consulta\n	[ 1 ] Consultar Datos de Logística\n	[ 2 ] Consultar Datos Financieros\n	[ 3 ] Consultar Datos Militares\n	[ 4 ] Volver\nIngrese su opción:"
		fmt.Print(MenuInicio)

		var eleccion string
		fmt.Scanln(&eleccion)

		switch eleccion {
		case "1":

			fmt.Print(MenuConsulta)
			var eleccionC string
			fmt.Scanln(&eleccionC)
			switch eleccionC {
			case "1":
				fmt.Println("Entregando Datos de Logistica:")
				Comunicacion("LOGISTICA")

			case "2":
				fmt.Println("Entregando Datos Financieron")
				Comunicacion("FINANCIERA")

			case "3":
				fmt.Println("Entregando Datos Militares")
				Comunicacion("MILITAR")
			case "4":
				fmt.Println("Volviendo")
			default:
				fmt.Println("Opcion no valida, volviendo al menu de inicio")
			}

		case "2":

			fmt.Println("Realizar ACA cierre de las otras cosas")
			flag = false

		default:
			fmt.Println("Opcion no valida")
		}
	}
	fmt.Println("Se acabo este programa")

}
