package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	pb "../pointpb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var (
		server  = flag.String("server", "127.0.0.1:10000", "Server address.")
		assetID = flag.String("assetID", "1", "The assetID to use.")
	)
	flag.Parse()

	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewTrackersvcClient(conn)

	stream, err := c.PointStream(context.Background(), &pb.Request{AssetID: *assetID})
	if err != nil {
		log.Fatal("test")
		log.Fatal(err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(response.Point)
	}
}
