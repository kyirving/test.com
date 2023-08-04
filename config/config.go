package config

type WebConf struct {
	Host   string
	Port   string
	Debug  bool   //是否开发
	Dbtype string //是否开发
}

type MysqlConf struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

type RedisConf struct {
	Host      string
	Password  string
	DataBase  string
	MaxIdle   int
	MaxActive int
}

var (
	WebConfig   *WebConf
	MysqlConfig *MysqlConf
	RedisConfig *RedisConf
)

//todo 不能动态化配置
// func init() {
// 	viper.SetConfigName("config")         // name of config file (without extension)
// 	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
// 	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
// 	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
// 	viper.AddConfigPath(".")              // optionally look for config in the working directory
// 	err := viper.ReadInConfig()           // Find and read the config file
// 	if err != nil {                       // Handle errors reading the config file
// 		panic(fmt.Errorf("fatal error config file: %w", err))
// 	}
// }
