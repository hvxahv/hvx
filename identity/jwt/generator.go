package jwt

import "github.com/golang-jwt/jwt/v4"

// JWTTokenGenerator
// Signer Example: viper.GetString("authentication.token.signed")
func (c *Claims) JWTTokenGenerator(signer string) (token string, err error) {
	g, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(signer))
	if err != nil {
		return "", err
	}
	return g, nil
}
