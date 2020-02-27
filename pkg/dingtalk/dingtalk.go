package dingtalk

import (
    "net/http"
    "fmt"
    "encoding/json"
    "strings"
)

// 定义markdown格式结构体
type MsgContent struct {
    Title string `json:"title"`
    Text  string `json:"text"`
}

type MDMsg struct {
    MsgType string `json:"msgtype"`
    Msgcontent MsgContent `json:"markdown"`
}

func MsgPush(dtoken, dtitle, dtext string) {
    webHook := "https://oapi.dingtalk.com/robot/send?access_token="+dtoken
    //fmt.Println(webHook)
    s := MDMsg{"markdown", MsgContent{dtitle, dtext}}
    buf, _ := json.Marshal(s)
    req, _ := http.NewRequest("POST", webHook, strings.NewReader(string(buf)))
    client := &http.Client{}
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    return
}

