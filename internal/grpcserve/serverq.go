package grpcserve

import (
	context "context"
	qserve "go-queue/internal/qsys"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type servergrpc struct {
	UnimplementedQServiceServer
}

func (s *servergrpc) NewQ(ctx context.Context, in *QueueMSG) (*NewQResponse, error) {
	var q qserve.ConfigQueue
	q.Name = in.Name
	q.Persistent = true
	for _, i := range in.Variables {
		q.Variable = append(q.Variable, qserve.Variable{
			Key:   i.Key,
			Value: i.Value,
		})

	}
	q.MaxSize = 0
	q.TTL = 0
	err := qserve.CreateQ(q)

	if err != nil {
		return nil, err
	}
	out := new(NewQResponse)
	out.Msg = "Created"
	return out, nil
}

// InitServer -grpc
func InitServer() {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	RegisterQServiceServer(s, &servergrpc{})
	reflection.Register(s)
	s.Serve(listener)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
