package main

import (
	"context"
	"fmt"
	pb "golang-learning/grpc_practice/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func printResult(client pb.UserServerClient, name string, age int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := []grpc.CallOption{}
	resp, err := client.Login(ctx, &pb.LoginRequest{Name: name, Age: age}, opt...)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("result", resp.Success)
}

func main() {
	opt := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("localhost:1234", opt...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserServerClient(conn)
	printResult(client, "andy", 18)
	printResult(client, "chris", 18)
}
