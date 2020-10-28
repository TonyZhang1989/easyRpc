package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "easyRpc/protos"
	"log"
	"os"
	"time"
)

const (
	address     = "127.0.0.1:8888"
	rpcTimeout = 5
	defaultCmd = "echo Hello!"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCmdExecutorClient(conn)

	// Contact the server and print out its response.
	cmd := defaultCmd
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(rpcTimeout * time.Second)))
	defer cancel()
	r, err := c.SendCmd(ctx, &pb.RpcRequest{Command: cmd})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	respMap := map[string]interface{}{}
	respMap["code"] = r.Code
	respMap["msg"] = r.Msg
	respMap["cmd"] = cmd

	respJson, err := json.Marshal(respMap)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	log.Printf(string(respJson))
}