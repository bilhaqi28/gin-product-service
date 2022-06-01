package controller

import "github.com/gin-gonic/gin"

type ControllerProduct interface {
	FindAll(c *gin.Context)
	Store(c *gin.Context)
	FindById(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}
