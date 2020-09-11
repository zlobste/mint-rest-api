package apiserver

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateJWT(userid int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userid
	claims["IssuedAt"] = time.Now().Unix()
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
