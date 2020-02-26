package v1

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "rwplus-backend/pkg/e"

)

type Check struct {
    Token string `valid:"Required"`
    Type string `valid:"Required"`
    Content string `valid:"Required"`
}


type Dingmsg struct {
   Token    string `json:"token"`
   Type     string  `json:"type"`
   Content  string  `json:"content"`
}


func DingMsg(c *gin.Context) {
    var json Dingmsg
    code := e.SUCCESS
    data := make(map[string]interface{})
    if err := c.BindJSON(&json); err!=nil {
        code = 30001
        data["result"]="只支持JSON传入"
    } else {
        valid := validation.Validation{}
        a := Check(json)
        ok, _ := valid.Valid(&a)

        if ok {
            data["token"] = json.Token
            data["type"] = json.Type
            data["content"] = json.Content
        } else {
            for _, err := range valid.Errors {
                code = e.ERROR_VALID_FAIL
                data["result"] = e.ValidMsg(err.Key)
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}



