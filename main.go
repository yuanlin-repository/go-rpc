package main

import server "go-rpc/server"

func main() {
	//
	serverOptions := []server.ServerOption{
		server.WithAddress("127.0.0.1:8080"),
		server.WithProtocol("proto"),
		server.WithNetWork("tcp"),
		server.WithSerialization("protobuff"),
	}
	svr := server.NewServer(serverOptions...)
	svr.Start()
}
