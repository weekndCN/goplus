package main

import (
    "fmt"
    "net/http"

    "rwplus-backend/models"
    "rwplus-backend/pkg/setting"

	"rwplus-backend/routers"
)

func init() {
    setting.Setup()
	models.Setup()
}

func main() {
	routersInit := routers.InitRouter()

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	maxHeaderBytes := 1 << 20

    s := &http.Server{
        Addr:           endPoint,
        Handler:        routersInit,
        ReadTimeout:    readTimeout,
        WriteTimeout:   writeTimeout,
        MaxHeaderBytes: maxHeaderBytes,
    }

    s.ListenAndServe()
}
