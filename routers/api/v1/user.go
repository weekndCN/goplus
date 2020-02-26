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

func GetUser(c *gin.Context) {

    // &name传入的key/value
    name := c.Query("name")

    //maps := make(map[string]interface{})
    data := make(map[string]interface{})
    //if name != "" {
    //   maps["name"] = name
    //}

    code := e.SUCCESS

    // 根根url请求的参数值查询数据库，数据写入data保存
    data["name"] = models.ExistUserByName(name)
    //data["phone"] = models.ExistUserByPhone(maps)


    // 返回请求结果
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}