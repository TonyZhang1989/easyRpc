package main

import (
	pb "easyRpc/protos"
	"easyRpc/src/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"log"
	"net"
	"runtime"
)


// server is used to implement myRpc.CmdExecutorServer.
type server struct{}

// SendCmd implements myRpc.CmdExecutorServer
func (s *server) SendCmd(ctx context.Context, in *pb.RpcRequest) (*pb.RpcReply, error) {
	var cmdoutUtf8 string
	var cmdoutGbk string
	var code int
	switch runtime.GOOS {
	case "linux":
		code, cmdoutUtf8 = common.RunInLinux(in.Command)
	case "windows":
		code, cmdoutGbk = common.RunInWindows(in.Command)
		cmdoutByteGbk := []byte(cmdoutGbk)
		var cmdoutByteUtf8 []byte
		cmdoutByteUtf8, _ = common.GbkToUtf8(cmdoutByteGbk)
		cmdoutUtf8 = string(cmdoutByteUtf8[:])
	}
	return &pb.RpcReply{Code: int32(code), Msg: cmdoutUtf8 }, nil
}

func main() {
	cfg, err := ini.Load("etc/server.ini")
	port := ":"+cfg.Section("default").Key("port").String()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCmdExecutorServer(s, &server{})
	s.Serve(lis)
}