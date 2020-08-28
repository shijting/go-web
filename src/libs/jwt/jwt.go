package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// access token过期时间
	ATokenExpireDuration = time.Second * 1
	// refresh token过期时间
	RTokenExpireDuration = time.Second * 3
)

// 密钥(盐)
var Secret = []byte("亚索主E，QEQ")

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// 生成jwt token
func GenToken(userId int64) (aToken string, rToken string, err error) {
	c := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ATokenExpireDuration).Unix(),
		},
	}
	// 生成access token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(Secret)
	if err != nil {
		return
	}
	// 生成refresh token
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RTokenExpireDuration).Unix(),
	}).SignedString(Secret)

	return
}

// 解析token
func ParseToken(tokenString string) (claims *Claims, err error) {
	var token *jwt.Token
	claims = new(Claims)
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
		return

	}
	return
}

// 解析refresh token 并生成新的access token
func ParseRefreshToken(aToken, rToken string) (newAccessToken string, err error) {
	// 解析refresh token， 无效或过期则之间返回
	var refreshToken *jwt.Token
	refreshToken, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return
	}
	if !refreshToken.Valid {
		err = errors.New("invalid refresh  token")
		return
	}
	// 从旧的access token 中解析出claims 数据
	claims := new(Claims)
	_, err = jwt.ParseWithClaims(aToken, claims, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if claims.UserId == 0 {
		err = errors.New("access token is invalid")
		return
	}
	// 判断access token是过期 并且refresh token没有过期则生成新的access token
	if err != nil {
		v := err.(*jwt.ValidationError)
		fmt.Println(v.Errors)
		//access token以过期
		if v.Errors == jwt.ValidationErrorExpired {
			newAccessToken, _, err = GenToken(claims.UserId)
		}
		return
	}
	newAccessToken = aToken
	return
}
