package middleware

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/servergrpc"
	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: false,
			Error:  "Token Mismatch",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
		return
	} else {
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		tokenJwt, err := servergrpc.ReqAuthServiceToTokenGrpc(token)
		if err == nil {
			c.Keys["token_jwt"] = tokenJwt
			c.Next()
		} else {
			errorResponse := response.ErrorResponse{
				Code:   http.StatusUnauthorized,
				Status: false,
				Error:  err.Error(),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
			return
		}

	}
}
