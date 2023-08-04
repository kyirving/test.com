package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"test.com/model"
)

type Goods struct {
}

func (g *Goods) Index(c *gin.Context) {

	goodsModle := model.NewGoodsDao()

	result, err := goodsModle.FindOneByGoodsID(37)

	fmt.Println(result)
	fmt.Println(err)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
