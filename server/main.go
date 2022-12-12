package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/user/session"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	session.RegisterAuthCheckerServer(server, NewSessionManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
