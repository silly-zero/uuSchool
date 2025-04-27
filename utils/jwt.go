package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CustomClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type CustomClaims struct {
	// 自定义字段
	Id                   string `json:"id"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(id string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		id, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //定义过期时间
			Issuer:    "ken	",                                             // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	return token.SignedString([]byte("secret"))
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	//对token对象中的Claim进行类型断言
	if token.Valid { // 校验token
		return mc, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
