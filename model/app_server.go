package model

import (
	"math"
	"time"
)

type AppServer struct {
	Base
	App_id      int
	Server_id   int
	Server_name string
	Origin_id   int
	Start_time  string
}

func NewAppServer() *AppServer {
	return &AppServer{
		Base: Base{DB: UseDbConn()},
	}
}

func (a *AppServer) TableName() string {
	return "app_server"
}

func (a *AppServer) DiffDay() int {
	stime, _ := time.Parse("2006-01-02 15:16:05", a.Start_time)
	now := time.Now()

	return int(math.Ceil(now.Sub(stime).Hours()))
}
