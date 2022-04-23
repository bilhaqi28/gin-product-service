package controller

import "github.com/gin-gonic/gin"

type ControllerProduct interface {
	FindAll(c *gin.Context)
}
