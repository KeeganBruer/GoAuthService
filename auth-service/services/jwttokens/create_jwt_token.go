package jwttokens

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type NewTokenData struct {
	UserID string
}
type TokenData struct {
	Expiration time.Time
	UserID     string
}

func CreateToken(data *NewTokenData) (string, error) {
	var (
		t *jwt.Token
	)
	SigningSecret := []byte("temp_signing-key")

	ServerName := os.Getenv("ServerName")

	exp := time.Now().Add(1 * time.Hour)
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": ServerName,
			"sub": data.UserID,
			"exp": exp.Unix(),
		})

	return t.SignedString(SigningSecret)
}
func DecodeToken(tokenString string) (*TokenData, error) {
	SigningSecret := []byte("temp_signing-key")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SigningSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp, _ := claims.GetExpirationTime()
		sub, _ := claims.GetSubject()
		data := &TokenData{
			Expiration: exp.Time,
			UserID:     sub,
		}
		return data, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
