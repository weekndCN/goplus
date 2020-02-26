package api

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "rwplus-backend/pkg/e"
    "rwplus-backend/pkg/util"
    "rwplus-backend/models"
)

type auth struct {
    Name string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    // 获取url的name和password的参数
    name := c.Query("name")
    password := c.Query("password")

    valid := validation.Validation{}
    a := auth{Name: name, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS

    if ok {
        //  调用models下的checkauth函数检查用户是否存在
        isExist := models.CheckAuth(name, password)
        if isExist {
            // 生成token
            token, err := util.GenerateToken(name, password)
            if err != nil {
                // 生成失败
                code = e.ERROR_AUTH_TOKEN
            } else {
                // 生成成功
                data["token"] = token

                code = e.SUCCESS
            }
        } else {
            // 用户不存在，鉴权失败
            code = e.ERROR_AUTH
        }
    // beego验证不通过，返回错误
    } else {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}