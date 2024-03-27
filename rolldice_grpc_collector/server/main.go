package main

import (
	"math/rand"
	"strconv"
	"log"
	"dice/pb"
	"net"
	"context"
	"time"
	"os/signal"
	"os"

	"google.golang.org/grpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

var tracer = otel.Tracer("rolldice")

type server struct {
	pb.UnimplementedRollDicerServer
}

func (s *server) RollDice(ctx context.Context, req *pb.RollDiceRequest) (*pb.RollDiceReply, error) {
	ctx, span := tracer.Start(ctx, "roll")
	defer span.End()
	time.Sleep(1 * time.Second)
	roll := 1 + rand.Intn(6)
	resp := strconv.Itoa(roll) + "\n"
	rollValueAttr:=attribute.Int("roll.value",roll)
	span.SetAttributes(rollValueAttr)
	return &pb.RollDiceReply{DiceNumber: resp}, nil
}


func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	tp, err := newTraceProvider(ctx)
	if err != nil {
		return
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	lis, err := net.Listen("tcp", "server:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverHandler:=otelgrpc.NewServerHandler(
		otelgrpc.WithTracerProvider(tp),
	)
	s:=grpc.NewServer(grpc.StatsHandler(serverHandler))
	pb.RegisterRollDicerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
