package middleware

import (
	"net/http"
	"time"

	"test.com/utils/e"

	"github.com/gin-gonic/gin"

	"test.com/config"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Uid                  int64
	Mobile               int64
	jwt.RegisteredClaims // v5版本新加的方法
}

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		var resp = e.NewResp()
		access_token := c.Request.Header.Get("Access-Token")
		if access_token == "" {
			resp.Output(e.RESP_TOKEN_MISSING, "", nil)
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}

		myClaims, err := ParseJwt(access_token)
		if err != nil {
			resp.Output(e.RESP_TOKEN_INVALID, "", nil)
			c.JSON(http.StatusOK, resp)
			c.Abort()
			return
		}

		c.Set("user_id", myClaims.Uid)
		c.Set("mobile", myClaims.Mobile)
		c.Next()
	}
}

//生成Jwt
func GenerateJWT(uid int64, mobile int64) (string, error) {
	claims := MyClaims{
		uid,
		mobile,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}

	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(config.WebConfig.Secret))

	return s, err
}

// 解析JWT
func ParseJwt(token string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.WebConfig.Secret), nil
	})

	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
