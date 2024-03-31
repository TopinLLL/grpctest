package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/TopinLLL/grpctest/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:9005", "the address to connect to")
)

func main() {
	dial, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer dial.Close()
	c := pb.NewTestClient(dial)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := c.GetTestResult(ctx, &pb.Hello{TestNumber: 1})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
}
