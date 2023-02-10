package services

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	userProto "github.com/emil-petras/project-proto/user"
	"github.com/emil-petras/project-web-service/models"
)

func GetUser(conn *grpc.ClientConn, username string) (*userProto.User, error) {
	request := userProto.ReadUser{
		Username: username,
	}

	client := userProto.NewUserServiceClient(conn)
	response, err := client.Read(context.Background(), &request)
	if err != nil {
		return nil, fmt.Errorf("cannot read user: %w", err)
	}

	return response, nil
}

func WithdrawBalance(conn *grpc.ClientConn, withdraw models.DepositWithdraw, user *userProto.User) (bool, error) {
	updateUser := &userProto.UpdateUser{
		Username: user.Username,
		Balance:  user.Balance - uint64(withdraw.Amount),
	}

	if uint(user.Balance) < withdraw.Amount {
		return false, nil
	}

	userClient := userProto.NewUserServiceClient(conn)
	user, err := userClient.Update(context.Background(), updateUser)
	if err != nil {
		return false, fmt.Errorf("cannot update user: %w", err)
	}

	return true, nil
}

func DepositBalance(conn *grpc.ClientConn, deposit models.DepositWithdraw, user *userProto.User) error {
	updateUser := &userProto.UpdateUser{
		Username: user.Username,
		Balance:  user.Balance + uint64(deposit.Amount),
		UserID:   user.Id,
	}

	client := userProto.NewUserServiceClient(conn)
	user, err := client.Update(context.Background(), updateUser)
	if err != nil {
		return fmt.Errorf("cannot update user: %w", err)
	}

	return nil
}
