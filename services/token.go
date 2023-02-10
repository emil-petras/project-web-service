package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	tokenProto "github.com/emil-petras/project-proto/token"
)

func Validate(conn *grpc.ClientConn, token string) (bool, string, error) {
	request := tokenProto.ReadToken{
		Value: token,
	}

	client := tokenProto.NewTokenServiceClient(conn)
	response, err := client.Read(context.Background(), &request)
	if err != nil {
		return false, "", fmt.Errorf("cannot read token: %w", err)
	}

	if response.ValidUntil.AsTime().Local().After(time.Now()) {
		return true, response.Username, nil
	}

	return false, "", nil
}

func Generate(conn *grpc.ClientConn, userID uint32) (string, error) {
	uuid := uuid.New()
	validUntil := time.Now().Local().Add(time.Hour * time.Duration(3))
	request := tokenProto.CreateToken{
		Value:      uuid.String(),
		UserID:     uint32(userID),
		ValidUntil: timestamppb.New(validUntil),
	}

	client := tokenProto.NewTokenServiceClient(conn)
	response, err := client.Create(context.Background(), &request)
	if err != nil {
		return "", fmt.Errorf("cannot create token: %w", err)
	}

	return response.Value, nil
}
