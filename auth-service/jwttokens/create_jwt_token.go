package jwttokens

import "github.com/golang-jwt/jwt/v5"

func CreateToken() (string, error) {
	var (
		key []byte
		t   *jwt.Token
	)

	key = []byte("temp_signing-key")
	t = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "my-auth-server",
			"sub": "john",
			"foo": 2,
		})

	return t.SignedString(key)
}
