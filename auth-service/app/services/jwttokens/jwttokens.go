package jwttokens

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type NewTokenData struct {
	SessionID     int
	MinutesTilExp int
}
type TokenData struct {
	Expiration time.Time
	SessionID  int
}

// Create a JWT token with the given data
func CreateToken(data *NewTokenData) (string, error) {

	//Get environment configurations
	SigningSecret := []byte(os.Getenv("JWTSecretToken"))
	ServerName := os.Getenv("ServerName")

	//Construct token expiration
	minutes := data.MinutesTilExp
	if minutes == 0 {
		minutes = 60
	}
	exp := time.Now().Add(time.Duration(minutes) * time.Minute)

	//Construct JWT claims
	claims := jwt.MapClaims{
		"iss": ServerName,
		"sub": fmt.Sprintf("%d", data.SessionID),
		"exp": exp.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//return token as signed string
	return token.SignedString(SigningSecret)
}

// Decode a JWT token string into TokenData
func DecodeToken(tokenString string) (*TokenData, error) {
	//Get environment configurations
	SigningSecret := []byte(os.Getenv("JWTSecretToken"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SigningSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	//Extract JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println(err)
		return nil, err
	}

	//Construct Token data from claims
	exp, _ := claims.GetExpirationTime()
	sub, _ := claims.GetSubject()

	SessionID, err := strconv.Atoi(sub)
	if err != nil {
		return nil, err
	}
	data := &TokenData{
		Expiration: exp.Time,
		SessionID:  SessionID,
	}
	return data, nil
}
