package domain

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	CustomerId string   `json:"customer_id"`
	Accounts   []string `json:"accounts"`
	Username   string   `json:"username"`
	jwt.Claims
}
