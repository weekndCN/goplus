package routers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // 路由
    "rwplus-backend/routers/api/v1"
    // 增加JWT认证
    "rwplus-backend/routers/api"
    // app.ini 初始化
    "rwplus-backend/pkg/setting"
    // 引入中间件
    "rwplus-backend/middleware/jwt"
    "rwplus-backend/pkg/upload"

)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    // 加载app.ini的RUN_MODE参数
    gin.SetMode(setting.ServerSetting.RunMode)

    r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
    r.GET("/auth", api.GetAuth)
    // add upload image
    r.POST("/upload", api.UploadImage)
    apiv1 := r.Group("/api/v1")
    apiv1.Use(jwt.JWT())
    {
        // 获取用户
        apiv1.GET("/users", v1.GetUser)
        // 新建用户
        apiv1.POST("/users", v1.DingMsg)
        // qrcode
        apiv1.POST("/qrcode", v1.GenerateArticlePoster)
    }

    return r
}