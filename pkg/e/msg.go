package e

// 设置错误返回信息Map
var MsgFlags = map[int]string {
    SUCCESS : "成功",
    ERROR : "失败",
    INVALID_PARAMS : "请求参数错误",
    ERROR_EXIST_USER :  "已存在该用户",
    ERROR_NOT_EXIST_USER : "该用户不存在",
    ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
    ERROR_AUTH_TOKEN : "Token生成失败",
    ERROR_AUTH : "Token错误",
}

// 返回Flags信息
func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}