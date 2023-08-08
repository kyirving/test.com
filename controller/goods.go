package controller

import (
	"fmt"

	"test.com/model"
	"test.com/utils/e"

	"github.com/gin-gonic/gin"
)

type Goods struct {
}

type QueryParams struct {
	AppId       int    `form:"app_id" binding:"required"`
	ServerId    int    `form:"server_id" binding:"required"`
	Category_id int    `form:"category_id"`
	IsHot       int    `form:"is_hot"`
	IsNew       int    `form:"is_new"`
	Keyword     string `form:"keyword"`
	Page        int    `form:"page" default:"1"`
	Pagesize    int    `form:"pagesize" default:"20"`
}

func (g *Goods) Index(c *gin.Context) {

	var queryParams QueryParams
	r := e.NewResp()
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		r.Output(e.RESP_PARAMS_ERR, "参数解析失败", nil)
		c.JSON(e.RESP_SUCC, r)
		return
	}

	fmt.Println(queryParams)

	propDao := model.NewPropDao()
	appServerDao := model.NewAppServer()
	propInfoDao := model.NewPropInfoDao()

	appServerDao.Table(appServerDao.TableName()).Where("app_id = ? AND server_id = ? ", queryParams.AppId, queryParams.ServerId).Find(&appServerDao)
	if appServerDao.Error != nil {
		r.Output(e.RESP_PARAMS_ERR, "区服不存在", nil)
		c.JSON(e.RESP_SUCC, r)
	}

	diffDay := appServerDao.DiffDay()
	subQuery := propDao.Table(propDao.TableName()).Select("prop_id , prop_name").Where("app_id = ? AND open_day <= ?", queryParams.AppId, diffDay)
	if queryParams.Keyword != "" {
		subQuery = propDao.Where("prop_name LIKE ?", queryParams.Keyword)
	}

	query := propInfoDao.Table("(?) as p1", subQuery).Select("p1.prop_id, p1.prop_name, MIN(p2.min_price) as min_price, MAX(p2.max_price) as max_price, SUM(p2.stock) as stock, MAX(p2.new_time) as new_time").
		Joins("left join prop_info p2 on p1.prop_id = p2.prop_id").
		Where("app_id = 1 AND server_id in ?", []int{1010, 1007, 1001}).
		Group("p1.prop_id")
		// Offset((queryParams.Page - 1) * queryParams.Pagesize).
		// Limit(queryParams.Pagesize).
	if queryParams.IsHot == 1 {
		query.Order("stock desc")
	}

	if queryParams.IsNew == 1 {
		query.Order("new_time desc")
	}

	var count int64
	query.Count(&count)

	props := []model.Props{}
	query.Offset((queryParams.Page - 1) * queryParams.Pagesize).Limit(queryParams.Pagesize).Find(&props)

	data := map[string]interface{}{
		"pageinfo": map[string]int{
			"page":     queryParams.Page,
			"pagesize": queryParams.Pagesize,
			"total":    int(count),
		},
		"list": props,
	}

	r.Output(e.RESP_SUCC, "", data)
	c.JSON(e.RESP_SUCC, r)
}
