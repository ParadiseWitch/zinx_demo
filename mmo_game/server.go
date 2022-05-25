package main

import "zinx_demo/znet"

func main() {
	//创建服务器句柄
	s := znet.NewServer("mmo_game_server")

	//启动服务
	s.Serve()
}
