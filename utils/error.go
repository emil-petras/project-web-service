package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func WriteError(c *gin.Context, code int, err error) {
	c.Error(err)
	c.JSON(code, gin.H{"error": err.Error()})
	logrus.Println(err.Error())
}
