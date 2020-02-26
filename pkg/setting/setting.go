package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)

var (

    // 加载ini文件
    Cfg *ini.File

    // 设置运行模式。默认Debug
    RunMode string

    // 设置运行服务的端口
    HTTPPort int
    //设置读超时时间
    ReadTimeout time.Duration
    //设置写超时时间
    WriteTimeout time.Duration

    // 设置token key
    RWSecret string
)


func init() {
    var err error
    Cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("无法解析app.ini文件: %v", err)
    }
    LoadBase()
    LoadServer()
    LoadApp()
}

func LoadBase() {
    // 由 Must 开头的方法名允许接收一个相同类型的参数来作为默认值，
    // 当键不存在或者转换失败时，则会直接返回该默认值。
    // 但是，MustString 方法必须传递一个默认值。
    sec, err := Cfg.GetSection("mode")
    if err != nil {
        log.Fatalf("无法获取到'mode'分区: %v", err)
    }
    RunMode = sec.Key("RUN_MODE").MustString("debug")
}


func LoadServer() {
    // 获取server分区下的所有key
    sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("无法获取到'server'分区: %v", err)
    }

    // 获取key对应的value
    HTTPPort = sec.Key("HTTP_PORT").MustInt(9000)
    // Duration类型代表两个时间点之间经过的时间，以纳秒为单位
    // 将整数个某时间单元表示为Duration类型值,用乘法 * time.Second
    // 所以这里为60s
    ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
    WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
    // 获取app分区下的所有key
    sec, err := Cfg.GetSection("app")
    if err != nil {
        log.Fatalf("无法获取'app'分区: %v", err)
    }

    // 获取key对应的value
    RWSecret = sec.Key("RW_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
