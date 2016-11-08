package main

import (
	"log"
	"net"
	"net/http"
	"time"

	pb "../pointpb"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Point(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	log.Println("Enter in Point")
	q := GetPosition("asset1")
	return &pb.Response{Point: &q}, nil
}

func (s *server) PointStream(request *pb.Request, stream pb.Trackersvc_PointStreamServer) error {
	for {
		log.Println("Enter in PointStream")
		q := GetPosition("asset1")
		err := stream.Send(&pb.Response{Point: &q})
		if err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	log.Println("Starting Tracker service...")
	ln, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTrackersvcServer(s, &server{})
	log.Println("Tracker service started successfully.")
	go s.Serve(ln)

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}
	log.Fatal(http.ListenAndServe(":10001", nil))
}
