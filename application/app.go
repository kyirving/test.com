package application

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"test.com/config"
	"test.com/routes"

	"github.com/spf13/viper"
)

//服务
type InitFnObserver func() (deferFunc func(), err error)

//属性
type NewAppOptions func(*app)

type app struct {
	appName         string
	configDir       string
	configName      string
	configExt       string
	InitFnObservers []InitFnObserver
	err             error
	deferFuncs      []func()
}

func NewApp(appName string, fns ...NewAppOptions) *app {
	app := &app{
		appName:    appName,
		configDir:  "config",
		configName: "config",
		configExt:  "ini",
	}

	for _, fn := range fns {
		//执行函数 WithConfigDir、WithConfigName、WithConfigExt
		fn(app)
	}

	return app
}

//服务注册
func RegisterInitFnObserver(fn InitFnObserver) NewAppOptions {
	return func(a *app) {
		a.InitFnObservers = append(a.InitFnObservers, fn)
	}
}

func (a *app) NotifyInitObserver() *app {
	a.deferFuncs = []func(){}

	for _, fnObserver := range a.InitFnObservers {
		var fn func()
		fn, a.err = fnObserver()

		if a.err != nil {
			return a
		}
		a.deferFuncs = append(a.deferFuncs, fn)
	}
	return a
}

//关闭app
func (a *app) Close() {
	for _, fn := range a.deferFuncs {
		fn()
	}
}

// 是否有异常
func (a *app) Error() (err error) {
	return a.err
}

func WithConfigDir(configDir string) NewAppOptions {
	return func(a *app) {
		a.configDir = configDir
	}
}

func WithConfigName(configName string) NewAppOptions {
	return func(a *app) {
		a.configName = configName
	}
}

func WithConfigExt(configExt string) NewAppOptions {
	return func(a *app) {
		a.configExt = configExt
	}
}

// 初始化配置
func (a *app) InitConfig() *app {
	viper := viper.New()
	viper.AddConfigPath(a.configDir)
	viper.SetConfigName(a.configName)
	viper.SetConfigType(a.configExt)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("GlobConfig err", err)
		a.err = err
		return a
	}

	var (
		webconf   config.WebConf
		mysqlconf config.MysqlConf
		redisconf config.RedisConf
	)

	if err := viper.UnmarshalKey("web", &webconf); err != nil {
		a.err = err
		return a
	}

	if err := viper.UnmarshalKey("mysql", &mysqlconf); err != nil {
		a.err = err
		return a
	}

	if err := viper.UnmarshalKey("redis", &redisconf); err != nil {
		a.err = err
		return a
	}

	config.WebConfig = &webconf
	config.MysqlConfig = &mysqlconf
	config.RedisConfig = &redisconf
	return a
}

func (a *app) RunServer() {
	router := routes.InitRouter()

	go func() {
		if err := router.Run(fmt.Sprintf("%s:%s", config.WebConfig.Host, config.WebConfig.Port)); err != nil {
			log.Panic(err)
		}
	}()
}

//监听是否退出
func (a *app) WaitForExitSign(exitFunc ...func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-c
	for index := range exitFunc {
		exitFunc[index]()
	}
}
