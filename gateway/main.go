package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "e13n/proto"

	"google.golang.org/grpc"
)

const (
	default_address = "localhost:50051"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := "World"
	p := r.URL.Query().Get("name")
	if p != "" {
		name = p
	}
	log.Printf("Get request " + p)
	address := default_address
	if os.Getenv("GRPC_ADDRESS") != "" {
		address = os.Getenv("GRPC_ADDRESS")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Get gRPC response " + res.GetMessage())

	fmt.Fprintf(w, "Greeting: %s", res.GetMessage())
}

func main() {
	http.HandleFunc("/", handler)
	log.Print("Listen on 8484")
	http.ListenAndServe(":8484", nil)
}
