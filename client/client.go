/*******
* @Author:qingmeng
* @Description:
* @File:client
* @Date2022/6/1
 */



package main

import (

"context"
"fmt"
"google.golang.org/grpc"
"google.golang.org/grpc/credentials/insecure"
	"grpcStudy/proto"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewBiliClient(conn)

	for {
		//这段不重要
		fmt.Println("input username&password:")
		iptName := ""
		_, _ = fmt.Scanln(&iptName)
		iptPassword := ""
		_, _ = fmt.Scanln(&iptPassword)

		loginResp, _ := c.Login(context.Background(), &proto.LoginReq{
			UserName: iptName,
			PassWord: iptPassword,
		})

		if loginResp.OK {
			fmt.Println("success")
			break
		}
		fmt.Println("retry")
	}
}
