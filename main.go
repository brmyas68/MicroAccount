package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"main.go/AccountGrpcServer"
	"main.go/account/pb"
)

func main() {

	fmt.Println("Server is Running ...")
	//grpclog.Println("Server is Running ...")
	port := ":8888"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(" Failed to listen ...")
	}
	//grpclog.Println("listening on 127.0.0.1:8888")
	fmt.Println("listening on 127.0.0.1:8888")

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	sqldb, err := sql.Open("mysql", "user66:1234@tcp(127.0.0.1:3306)/accounting?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	ACCGrpc := AccountGrpcServer.NewAccountGrpcServerStruct(sqldb)

	pb.RegisterAccountServiceServer(server, ACCGrpc)
	log.Println("register now...")

	err = server.Serve(listener)

	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err.Error())
	}
}
