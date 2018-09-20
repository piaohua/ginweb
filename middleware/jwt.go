package middleware

import (
	"errors"
	"time"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
            c.AbortWithStatusJSON(http.StatusOK, gin.H{
                "code": http.StatusBadRequest,
                "msg": "无访问权限",
            })
			return
		}
		//log.Print("get token: ", token)
		j := NewJWT()
		// parseToken
		claims, err := j.ParseToken(token)
        if err == TokenExpired {
            newToken, err := j.RefreshToken(token)
            if err != nil {
                c.AbortWithStatusJSON(http.StatusOK, gin.H{
                    "code": http.StatusBadRequest,
                    "msg": err.Error(),
                })
                return
            }
            //c.Header("Authorization", newToken)
            c.AbortWithStatusJSON(http.StatusOK, gin.H{
                "code": -1,
                "data": newToken,
            })
            return
        }
		if err != nil {
            c.AbortWithStatusJSON(http.StatusOK, gin.H{
                "code": http.StatusBadRequest,
                "msg": err.Error(),
            })
            return
        }

        if claims.OpenID == "" {
            c.AbortWithStatusJSON(http.StatusOK, gin.H{
                "code": http.StatusBadRequest,
                "msg": "filed missing",
            })
            return
        }

		//c.Set("claims", claims)
		c.Set("openid", claims.OpenID)
		c.Next()
	}
}

// JWT jwt签名
type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "vXHEhkGrVWAxCSOzrgwW2bKCGcAB2QZi0PlZEmZVqR4"
)

// CustomClaims 载荷
type CustomClaims struct {
	OpenID string `json:"openid"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(GetSignKey()),
	}
}

// GetSignKey get sign key
func GetSignKey() string {
	return SignKey
}

// SetSignKey set sign key
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken create token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken parse token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
    func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
        return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken refresh token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
    func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// GenToken gen token
func GenToken(openid string) (string, error) {
	j := NewJWT()
	claims := CustomClaims{
		OpenID: openid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(), // 过期时间
			Issuer:    "piaohua",                            //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	return token, err
}
