package main

func main() {
	// TODO: pull ports from config

	httpServer := NewHttpServer(":8002")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":9002")
	gRPCServer.Run()
}
