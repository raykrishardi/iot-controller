package main

import "os"

var grpcPort = os.Getenv("GRPC_PORT")

func main() {
	grpcListen()
}
