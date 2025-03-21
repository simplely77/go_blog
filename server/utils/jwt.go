package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"server/global"
	"server/model/request"
	"time"
)

type JWT struct {
	AccessTokenSecret  []byte
	RefreshTokenSecret []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		AccessTokenSecret:  []byte(global.Config.Jwt.AccessTokenSecret),
		RefreshTokenSecret: []byte(global.Config.Jwt.RefreshTokenSecret),
	}
}

func (j *JWT) CreateAccessClaims(baseClaims request.BaseClaims) request.JwtCustomClaims {
	ep, _ := ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	claims := request.JwtCustomClaims{
		BaseClaims: baseClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},                //受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), //过期时间
			Issuer:    global.Config.Jwt.Issuer,               //签名的发行者
		},
	}
	return claims
}

func (j *JWT) CreateAccessToken(claims request.JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.AccessTokenSecret)
}

func (j *JWT) CreateRefreshClaims(baseClaims request.BaseClaims) request.JwtCustomRefreshClaims {
	ep, _ := ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime) // 获取过期时间
	claims := request.JwtCustomRefreshClaims{
		UserID: baseClaims.UserID, // 用户 ID
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"TAP"},                // 受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,               // 签名的发行者
		},
	}
	return claims
}

func (j *JWT) CreateRefreshToken(claims request.JwtCustomRefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 创建新的 JWT Token
	return token.SignedString(j.RefreshTokenSecret)            // 使用 RefreshToken 密钥签名并返回 Token 字符串
}
