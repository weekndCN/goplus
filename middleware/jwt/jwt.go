package jwt

import (
    "time"
    "net/http"

    "github.com/gin-gonic/gin"

    "rwplus-backend/pkg/e"
    "rwplus-backend/pkg/util"
)

func JWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        var code int
        var data interface{}

        code = e.SUCCESS
        token := c.Query("token")
        if token == "" {
            code = e.INVALID_PARAMS
        } else {
            claims, err := util.ParseToken(token)
            if err != nil {
                code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }
        // http.StatusUnauthorized=401
        if code != e.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : e.GetMsg(code),
                "data" : data,
            })
            // Abort prevents pending handlers from being called. Note that this will not stop the current handler.
            // Let's say you have an authorization middleware that validates that the current request is authorized.
            // If the authorization fails (ex: the password does not match),
            // call Abort to ensure the remaining handlers for this request are not called
            c.Abort()
            return
        }
        // Next should be used only inside middleware.
        // It executes the pending handlers in the chain inside the calling handler.
        c.Next()
    }
}