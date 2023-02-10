package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/emil-petras/project-web-service/communication"
	"github.com/emil-petras/project-web-service/models"
	"github.com/emil-petras/project-web-service/services"
	"github.com/emil-petras/project-web-service/utils"
)

func Login(c *gin.Context) {
	user := models.Login{}
	err := utils.UnmarshalJSON(c.Request, &user)
	defer c.Request.Body.Close()
	if err != nil {
		err = fmt.Errorf("issue with reading JSON: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	loggedIn, err := services.Login(communication.DBConn, user)
	if err != nil {
		err = fmt.Errorf("issue with logging in: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if loggedIn != nil {
		token, err := services.Generate(communication.DBConn, uint32(loggedIn.Id))
		if err != nil {
			err = fmt.Errorf("issue with token generation: %w", err)
			utils.WriteError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

	c.JSON(http.StatusUnauthorized, "wrong username and/or password")
}

func Deposit(c *gin.Context) {
	deposit := models.DepositWithdraw{}
	err := utils.UnmarshalJSON(c.Request, &deposit)
	defer c.Request.Body.Close()
	if err != nil {
		err = fmt.Errorf("issue with reading JSON: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	valid, username, err := services.Validate(communication.DBConn, deposit.Token)
	if err != nil {
		err = fmt.Errorf("issue with token validation: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if valid {
		user, err := services.GetUser(communication.DBConn, username)
		if err != nil {
			err = fmt.Errorf("issue with getting user: %w", err)
			utils.WriteError(c, http.StatusBadRequest, err)
			return
		}

		err = services.DepositBalance(communication.DBConn, deposit, user)
		if err != nil {
			err = fmt.Errorf("issue with updating amount: %w", err)
			utils.WriteError(c, http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusAccepted, "success")
		return
	}

	c.JSON(http.StatusUnauthorized, "token invalid")
}

func Withdraw(c *gin.Context) {
	withdraw := models.DepositWithdraw{}
	err := utils.UnmarshalJSON(c.Request, &withdraw)
	defer c.Request.Body.Close()
	if err != nil {
		err = fmt.Errorf("issue with reading JSON: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	valid, username, err := services.Validate(communication.DBConn, withdraw.Token)
	if err != nil {
		err = fmt.Errorf("issue with token validation: %w", err)
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if valid {
		user, err := services.GetUser(communication.DBConn, username)
		if err != nil {
			err = fmt.Errorf("issue with getting user: %w", err)
			utils.WriteError(c, http.StatusBadRequest, err)
			return
		}

		updated, err := services.WithdrawBalance(communication.DBConn, withdraw, user)
		if err != nil {
			err = fmt.Errorf("issue with updating amount: %w", err)
			utils.WriteError(c, http.StatusBadRequest, err)
			return
		}

		if !updated {
			c.JSON(http.StatusAccepted, "balance can't be less than zero")
			return
		}

		c.JSON(http.StatusAccepted, "success")
		return
	}

	c.JSON(http.StatusUnauthorized, "token invalid")
}
