package controller

import (
	"test.com/utils/e"

	"github.com/gin-gonic/gin"
	"test.com/model"
)

type Goods struct {
}

func (g *Goods) Index(c *gin.Context) {

	goodsModle := model.NewGoodsDao()

	r := e.NewResp()
	result, err := goodsModle.FindOneByGoodsID(37)
	if err != nil {
		r.Output(e.RESP_NOT_FOUND, "", nil)
		c.JSON(e.RESP_SUCC, r)
	}

	r.Output(e.RESP_SUCC, "", result)
	c.JSON(e.RESP_SUCC, r)
}
