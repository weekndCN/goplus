package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/boombuler/barcode/qr"
    //引入验证包
    //"github.com/astaxie/beego/validation"

    //"fmt"
    //"strconv"
    "net/http"
    "rwplus-backend/pkg/qrcode"

    //"rwplus-backend/pkg/setting"
    "rwplus-backend/pkg/e"
)
const (
    QRCODE_URL = "https://github.com/weekndchina/mkdocs"
)

func GenerateArticlePoster(c *gin.Context) {
    qrc := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
    path := qrcode.GetQrCodeFullPath()
    code := e.SUCCESS
    _, _, err := qrc.Encode(path)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "code" : code,
            "msg" : e.GetMsg(code),
            "data" : nil,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : nil,
    })
}