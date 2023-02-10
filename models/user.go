package models

import "github.com/emil-petras/project-web-service/utils"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DepositWithdraw struct {
	Token     string         `json:"token"`
	Amount    uint           `json:"amount"`
	Timestamp utils.JSONTime `json:"timestamp"`
}
