package common

import (
	"github.com/dgrijalva/jwt-go"
	"go_gin+vue/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint // uint：无符号整型
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 有效期7天
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 有效期
			IssuedAt:  time.Now().Unix(),     //token发放时间
			Issuer:    "oceanlearn.tech",     // 发放人
			Subject:   "user token",          // 主题
		},
	}

	// 使用jwt密钥来生成的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	//生成的token如下全部 由三部分组成 由 . 拼接

	//加密协议
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
	// 终端通过base64来解密
	//echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -D
	//得到：alg：算法HS256 jwt类型
	//{"alg":"HS256","typ":"JWT"}%

	// claims 的储存信息（line18～line25）
	//eyJVc2VySWQiOjE0LCJleHAiOjE2NzU3NDk1MzksImlhdCI6MTY3NTE0NDczOSwiaXNzIjoib2NlYW5sZWFybi50ZWNoIiwic3ViIjoidXNlciB0b2tlbiJ9.
	// 同样的base64来解密 得：
	// {"UserId":14,"exp":1675749539,"iat":1675144739,"iss":"oceanlearn.tech","sub":"user token"}%

	// 前面两部分加上 jwtKey 值的哈希值
	//ByF_Wy5zQkpLSd27edy1_94hBMn9-EBstemtpq2xMS
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) { // 从tokenString 里面解析出claims 并且返回
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
