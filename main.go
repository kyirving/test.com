package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"test.com/application"
)

var (
	configDir  string
	configName string
	configExt  string
	APP_ENV    string
)

func init() {
	//命令行配置
	flag.StringVar(&configDir, "configDir", "config", "配置文件路径")
	flag.StringVar(&configName, "configName", "config", "配置文件名称")
	flag.StringVar(&configExt, "configExt", "ini", "配置文件后缀")

	//解析命令行参数写入注册
	flag.Parse()

	//读取不通环境配置
	APP_ENV := os.Getenv("APP_ENV")
	if APP_ENV == "prod" {
		configName += ".prod"
	} else {
		configName += ".dev"
	}
}

func main() {
	app := application.NewApp(
		"test",
		application.WithConfigDir(configDir),
		application.WithConfigName(configName),
		application.WithConfigExt(configExt),
		application.RegisterInitFnObserver(application.InitRedis),
	)

	err := app.InitConfig().
		NotifyInitObserver().
		Error()

	if err != nil {
		log.Println(fmt.Sprintf("初始化失败%s", err.Error()))
		panic(err)
	}

	app.RunServer()

	app.WaitForExitSign(func() {
		fmt.Println("服务停止...")
	})

	defer app.Close()
}
