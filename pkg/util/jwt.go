package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "rwplus-backend/pkg/setting"
)

// 将app.ini的RW_SECRET的值转换为byte类似，一个字节用8个”0”或”1”字符表示
var jwtSecret = []byte(setting.RWSecret)


// 构造一个所有权的结构体
type Claims struct {
    Name string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}


// 定以生成token函数
func GenerateToken(name, password string) (string, error) {
    // 获取当前时间
    nowTime := time.Now()
    // 设置过期时间。当前时间+1小时
    expireTime := nowTime.Add(1 * time.Hour)

    // 生成一个Claims变量
    claims := Claims{
        name,
        password,
        jwt.StandardClaims {
            ExpiresAt: expireTime.Unix(),
            Issuer: "rwplus",
        },
    }
    // 生成一个tokenClaims
    // NewWithClaims(method SigningMethod, claims Claims)，
    // method对应着SigningMethodHMAC struct{}，
    // 其包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    // 生成token
    // func (t *Token) SignedString(key interface{}) 该方法内部生成签名字符串，再用于获取完整、已签名的token
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}


func ParseToken(token string) (*Claims, error) {
    // func (p *Parser) ParseWithClaims 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error){
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        // func (m MapClaims) Valid() 验证基于时间的声明exp, iat, nbf，
        //注意如果没有任何声明在令牌中，仍然会被认为是有效的。并且对于时区偏差没有计算方法
        if claims, ok :=tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}
