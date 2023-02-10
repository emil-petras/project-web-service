package services

import (
	"context"
	"crypto/sha256"
	"fmt"

	"google.golang.org/grpc"

	userProto "github.com/emil-petras/project-proto/user"
	"github.com/emil-petras/project-web-service/models"
)

func Login(conn *grpc.ClientConn, user models.Login) (*userProto.User, error) {
	request := userProto.ReadUser{
		Username: user.Username,
	}

	client := userProto.NewUserServiceClient(conn)
	response, err := client.Read(context.Background(), &request)
	if err != nil {
		return nil, fmt.Errorf("cannot read user: %w", err)
	}

	hash := sha256.New()
	hash.Write([]byte(user.Password))
	sum := hash.Sum(nil)
	user.Password = fmt.Sprintf("%x", sum)

	if response.Password == user.Password {
		return response, nil
	}

	// username and/or password is invalid
	return nil, nil
}
