package controller

import "github.com/gin-gonic/gin"

type Goods struct {
}

func (g *Goods) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
