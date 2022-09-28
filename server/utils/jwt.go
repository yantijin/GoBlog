package utils

import (
	"GoLog/commen"
	"GoLog/model"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	TokenExpired     = errors.New("Token过期")
	TokenNotValidYet = errors.New("Token目前无效")
	TokenMalformed   = errors.New("这并不是token,请检查")
	TokenInvalid     = errors.New("无法处理此token")
)

type JWT struct {
	SigningKey []byte
}

func parseDuration(d string) (time.Duration, error) {
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.HasSuffix(d, "d") { // 如果后缀有d,即表示多少天
		h := strings.TrimSuffix(d, "d")
		day, _ := strconv.Atoi(h)
		return time.Hour * 24 * time.Duration(day), nil
	}
	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func (j *JWT) SetClaims(claims model.BaseClaims) model.CustomClaims {

	ep, _ := parseDuration(commen.GVA_CONFIG.JWT.ExpiresTime)

	resClaims := model.CustomClaims{
		BaseClaims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ep).Unix(), // 过期时间
			NotBefore: time.Now().Unix() - 1000,  // 不早于某个时间
			Issuer:    commen.GVA_CONFIG.JWT.Issuer,
		},
	}
	return resClaims
}

// 根据claim来生成token
func (j *JWT) GenToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 对token进行解析
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
