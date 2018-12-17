package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/envoyproxy/envoy/examples/grpc-bridge/service/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	action := flag.String("action", "", "action")
	key := flag.String("key", "", "key")
	val := flag.String("val", "", "value")
	addr := "localhost:9001"

	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithAuthority("grpc"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKVClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	switch *action {
	case "get":
		r, err := c.Get(ctx, &pb.GetRequest{Key: *key})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("got %v", r)
		}
	case "set":
		r, err := c.Set(ctx, &pb.SetRequest{Key: *key, Value: *val})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("got %v", r)
		}
	}
}
