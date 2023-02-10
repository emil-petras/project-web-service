package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	idemProto "github.com/emil-petras/project-proto/idempotency"
	"github.com/emil-petras/project-web-service/communication"
	"github.com/emil-petras/project-web-service/utils"
)

func Idempotency() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := idemProto.NewIdempotencyServiceClient(communication.IdemConn)
		body, err := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error reading body: %w", err))
			c.Abort()
			return
		}

		req := idemProto.Request{
			Value: string(body),
		}
		response, err := client.Check(context.Background(), &req)
		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error contacting idempotency service: %w", err))
			c.Abort()
			return
		}

		if response.Exists {
			utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("duplicate request"))
			c.Abort()
			return
		}

		c.Next()
	}
}
