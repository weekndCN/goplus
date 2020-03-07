package v1

import (
    "github.com/gin-gonic/gin"
    //引入验证包
    //"github.com/astaxie/beego/validation"

    //"fmt"
    //"strconv"
    "net/http"

    "rwplus-backend/models"
    //"rwplus-backend/pkg/setting"
    "rwplus-backend/pkg/e"
)
const (
    QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

func GenerateArticlePoster(c *gin.Context) {
    appG := app.Gin{c}
    qrc := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
    path := qrcode.GetQrCodeFullPath()
    _, _, err := qrc.Encode(path)
    if err != nil {
        appG.Response(http.StatusOK, e.ERROR, nil)
        return
    }

    appG.Response(http.StatusOK, e.SUCCESS, nil)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}