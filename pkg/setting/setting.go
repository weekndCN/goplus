package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)


// 设置app配置

/*
RW_SECRET = 'callmeweeknd'

# Setting path
ImagePrefixUrl = 'http://localhost:9000'
ImageSavePath = upload/images/
# MB
ImageMaxSize = 5
# Allow Exts
ImageAllowExts = .jpg,.jpeg,.png
*/

type App struct {
    RWSecret string
    RuntimeRootPath string
    ImagePrefixUrl string
    ImageSavePath string
    ImageMaxSize int
    ImageAllowExts []string
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

func Setup() {
    Cfg, err := ini.Load("conf/app.ini")

    if err != nil {
        log.Fatalf("fail to parse 'conf/app.ini': %v", err)
    }

    err = Cfg.Section("app").MapTo(AppSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }
    // trans to MB from byte
    AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

    err = Cfg.Section("server").MapTo(ServerSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }
    // trans to second
    ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
    ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

    err = Cfg.Section("database").MapTo(DatabaseSetting)

    if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
    }
}

