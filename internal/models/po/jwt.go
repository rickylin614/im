package po

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	DeviceID string `json:"device_id"`
	Token    string `json:"token"`
	User     *Users `json:"user"`
	jwt.RegisteredClaims
}

func (k JWTClaims) Validate() error {
	return nil
}
