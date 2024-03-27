package main
import (
	"context"
	"dice/pb"
	"time"
	"math/rand"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel"
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
