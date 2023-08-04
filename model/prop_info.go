package model

type ProdInfo struct {
	Base
	App_id    int
	Server_id int
	Prop_id   int
	Stock     int
	Price     float64
	Min_price float64
	Max_price float64
	New_time  string
}

func NewPropInfoDao() *ProdInfo {
	return &ProdInfo{
		Base: Base{DB: UseDbConn()},
	}
}
