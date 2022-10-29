package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/Kendovvul/Ejemplo/Proto"
	"google.golang.org/grpc"
)

func RetornarData(Tipo string) string {
	file, err := os.Open("DATA.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	StringRetorno := ""

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		Split_Msj := strings.Split(scanner.Text(), ":")

		if Split_Msj[0] == Tipo {

			StringRetorno = StringRetorno + Split_Msj[1] + ":" + Split_Msj[2] + "\n"

		}
	}

	file.Close()
	return StringRetorno
}

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Intercambio(ctx context.Context, msg *pb.Message) (*pb.Message, error) {
	file, err := os.Create("DATA.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	msn := ""

	Split_Msj := strings.Split(msg.Body, ":")

	if Split_Msj[0] == "1" {
		msn = RetornarData(Split_Msj[1])

	} else {

		file.WriteString(Split_Msj[1] + ":" + Split_Msj[2] + ":" + Split_Msj[3] + "\n")
		msn = "Guardado"

	}

	return &pb.Message{Body: msn}, nil

}

//"DateNode Synth"
func main() {

	listener, err := net.Listen("tcp", ":50051") //conexion sincrona
	if err != nil {
		panic("La conexion no se pudo crear" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err = serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}

}
