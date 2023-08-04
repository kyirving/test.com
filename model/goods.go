package model

type GoodsModel struct {
	Base
	Goods_name string
	Prop_id    int
	Prop_name  string
	App_id     int
	Item_id    string
}

type Props struct {
	PropId    int
	PropName  string
	Stock     int
	Min_price float64
	Max_price float64
	New_time  string
}

// 表名
func (u *GoodsModel) TableName() string {
	return "goods"
}

func NewGoodsDao() *GoodsModel {
	return &GoodsModel{Base: Base{DB: UseDbConn()}}
}

func (g *GoodsModel) FindOneByGoodsID(goods_id int) (*GoodsModel, error) {
	result := g.Table(g.TableName()).First(g, goods_id)
	return g, result.Error
}
