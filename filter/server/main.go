package main

import (
	"log"
	"dice/pb"
	"net"
	"context"
	"os/signal"
	"os"

	"google.golang.org/grpc"
)


func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	statsHandler,err:=setupOTelSDK(ctx)
	if err != nil {
		return
	}

	lis, err := net.Listen("tcp", "server:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:=grpc.NewServer(grpc.StatsHandler(statsHandler))
	pb.RegisterRollDicerServer(s, &server{})


	go func(){
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
