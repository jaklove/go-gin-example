package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize int
	PrefixUrl string
	RuntimeRootPath string

	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string
	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
	RUN_MODE string
}

var AppSetting = &App{}


type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}


type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}


type Redis struct {
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func SetUp()  {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout  = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil{
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil{
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}

	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout *time.Second

}













//var (
//	Cfg          *ini.File
//	Runmode      string
//	HTTPPort     int
//	ReadTimeout  time.Duration
//	WriteTimeout time.Duration
//	PageSize     int
//	JwtSecret    string
//)
//
//func init() {
//	var err error
//	Cfg, err = ini.Load("conf/app.ini")
//	if err != nil {
//		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
//	}
//
//	//加载运行模式
//	LoadBase()
//	LoadServer()
//	LoadApp()
//}
//
//func LoadBase() {
//	Runmode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
//}
//
//func LoadServer() {
//	section, err := Cfg.GetSection("server")
//	if err != nil {
//		log.Fatalf("Fail to get section 'server': %v", err)
//	}
//
//	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)
//	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
//	WriteTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
//}
//
//func LoadApp()  {
//	section, err := Cfg.GetSection("app")
//	if err != nil{
//		log.Fatalf("Fail to get section 'app': %v", err)
//	}
//
//	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = section.Key("PAGE_SIZE").MustInt(10)
//}
