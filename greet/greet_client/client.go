package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/VJ-Vijay77/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

//main function

func main() {
	fmt.Println("Hello im a client!")

	cc,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect ",err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created Client:%f",c)
	//doUnary(c)
	doServerStreaming(c)
	
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a server streaming rpc...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vijay",
			LastName: "Dinesh",
		},
	}
	res,err := c.GreetManyTimes(context.Background(),req)
	if err != nil {
		log.Fatalln(err)
	}
	for{
	msg,err := res.Recv()
	if err == io.EOF{
		break
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(msg.GetResult())
	}
}



func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vijay",
			LastName: "Dinesh",
		},
	}

	res,err := c.Greet(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC : %v", err)
	}
	log.Printf("Response from Greet : %v",res.Result)
}