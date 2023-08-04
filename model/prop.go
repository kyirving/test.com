package model

type Prop struct {
	Base
	App_id      int
	Prop_id     int
	Prop_name   string
	Category_id int
	Open_day    int
	Icon_num    int
}

func NewPropDao() *Prop {
	return &Prop{
		Base: Base{DB: UseDbConn()},
	}
}

func (p *Prop) TableName() string {
	return "prop"
}
