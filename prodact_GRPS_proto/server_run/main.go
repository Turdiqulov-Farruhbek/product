package main

import (
	"log"
	"net"
	us_proto "product/genproto"
	ser "product/service"
	ps "product/storagecon/postgresdb"

	"google.golang.org/grpc"
)

func main() {
	db, err := ps.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	us_proto.RegisterUserServiceServer(server, ser.NewUserService(db))

	log.Printf("server listening at %v", listenPort.Addr())
	if err := server.Serve(listenPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
