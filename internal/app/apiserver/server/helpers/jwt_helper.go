package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	
	"github.com/dgrijalva/jwt-go"
)

type AccessDetails struct {
	UserId int64
}

func CreateJWT(userid int64) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") // this should be in an env file
	claims := jwt.MapClaims{}
	claims["userID"] = userid
	claims["IssuedAt"] = time.Now().Unix()
	claims["ExpiresAt"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["userID"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			UserId: userId,
		}, nil
	}
	return nil, err
}
