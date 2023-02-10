package communication

import (
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var DBConn *grpc.ClientConn

func CreateDbServiceConn() error {
	target := fmt.Sprintf("%v", os.Getenv("DB_SERVICE_TARGET"))
	port := fmt.Sprintf(":%v", os.Getenv("DB_SERVICE_PORT"))

	conn, err := grpc.Dial(target+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("cannot create connection to db service: %w", err)
	}

	DBConn = conn

	return nil
}
