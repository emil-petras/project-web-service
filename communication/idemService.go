package communication

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var IdemConn *grpc.ClientConn

func CreateIdemServiceConn() error {
	target := fmt.Sprintf("%v", os.Getenv("IDEM_SERVICE_TARGET"))
	port := fmt.Sprintf(":%v", os.Getenv("IDEM_SERVICE_PORT"))

	conn, err := grpc.Dial(target+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("cannot create connection to idempotency service: %w", err)
	}

	IdemConn = conn

	return nil
}
