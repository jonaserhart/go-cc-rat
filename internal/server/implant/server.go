package implant

import (
	"context"
	"errors"

	"github.com/jonaserhart/go_c2_rat/pkg/pb"
)

type implantServer struct {
	pb.UnimplementedImplantServer
	work, output chan *pb.Command
}

func NewImplantserver(work, output chan *pb.Command) *implantServer {
	return &implantServer{work: work, output: output}
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *pb.Empty) (*pb.Command, error) {
	cmd := new(pb.Command)
	select {
	case cmd, ok := <-s.work:
		if ok {
			return cmd, nil
		}
		return cmd, errors.New("channel closed")
	default:
		return cmd, nil
	}
}

func (s *implantServer) SendOutput(ctx context.Context, result *pb.Command) (*pb.Empty, error) {
	s.output <- result
	return &pb.Empty{}, nil
}
