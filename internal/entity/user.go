package entity

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserID         int64  `json:"user_id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	StandardClaims jwt.RegisteredClaims
}
