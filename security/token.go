package security

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

var (
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtSigningMethod = jwt.SigningMethodHS256.Name
)

func GenerateToken(userID string) (string, error) {
	claims := jwt.StandardClaims{
		Id: userID,
		Issuer: userID,
		IssuedAt: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}

func validateSignMethod(token *jwt.Token) (interface{}, error) {
	if _, ok  := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return JwtSecretKey, nil
}


func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	claims := new(jwt.StandardClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignMethod)	
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.StandardClaims)
	if ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}