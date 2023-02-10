package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/emil-petras/project-web-service/communication"
	"github.com/emil-petras/project-web-service/controllers"
	"github.com/emil-petras/project-web-service/middleware"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	gin.SetMode(os.Getenv("MODE"))
	router := gin.New()
	router.SetTrustedProxies(nil)

	err := communication.CreateDbServiceConn()
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	defer communication.DBConn.Close()
	logrus.Println("db service connection created")

	err = communication.CreateIdemServiceConn()
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	defer communication.IdemConn.Close()
	logrus.Println("idempotency service connection created")

	router.Use(middleware.Idempotency())

	router.POST("/login", controllers.Login)
	router.POST("/deposit", controllers.Deposit)
	router.POST("/withdraw", controllers.Withdraw)

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}
