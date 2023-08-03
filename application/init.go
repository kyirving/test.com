package application

func InitMysql() {

}

func InitRedis() (fn func(), err error) {

	// db.RedisPool = db.NewRedisPool(config.Addr, config.Passwd, config.Db, config.MaxIdle, config.MaxActive)

	// log.Println(fmt.Sprintf("Redis组件初始化成功！连接：%v，DB：%v，密码:%v MaxIdle:%v MaxActive:%v",
	// 	config.Addr,
	// 	config.Db,
	// 	config.Passwd,
	// 	config.MaxIdle,
	// 	config.MaxActive,
	// ))
	fn = func() {}
	return

}
