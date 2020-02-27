package v1

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "rwplus-backend/pkg/e"
    "rwplus-backend/pkg/dingtalk"
)

type Check struct {
    Token string `valid:"Required"`
    Title string `valid:"Required"`
    Content string `valid:"Required"`
}


type Dingmsg struct {
   Token    string `json:"token"`
   Title     string  `json:"title"`
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
            dtoken := json.Token
            dtitle := json.Title
            dcontent := json.Content
            go dingtalk.MsgPush(dtoken, dtitle, dcontent)
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



