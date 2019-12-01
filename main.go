package main

import (
	_ "go-snowflake-service/cache"
	"go-snowflake-service/conf"
	_ "go-snowflake-service/conf"
	"go-snowflake-service/pb"
	"go-snowflake-service/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", conf.ConfigObject.Addr)

	if err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	pb.RegisterIdGeneratorServiceServer(server, &services.IdGeneratorService{})

	reflection.Register(server)

	log.Printf("[Snowflake] Server started on %s", conf.ConfigObject.Addr)

	if err := server.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
