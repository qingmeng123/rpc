/*******
* @Author:qingmeng
* @Description:
* @File:main
* @Date2022/5/17
 */

//server（被调用rpc的一方）
package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcStudy/proto"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() //获取新服务示例
	proto.RegisterBiliServer(s, &server{})

	// 开始处理
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	proto.UnimplementedBiliServer // 用于实现proto包里BiliServer接口
}

func (s *server) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp, error) {
	resp := &proto.LoginResp{}
	log.Println("recv:", req.UserName, req.PassWord)
	if req.PassWord != GetPassWord(req.UserName) {
		resp.OK = false
		return resp, nil
	}
	resp.OK = true
	return resp, nil
}

func GetPassWord(userName string) (password string) {
	return userName + "123456"
}
