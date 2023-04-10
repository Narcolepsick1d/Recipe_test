package models

import "errors"

var (
	ErrRecipeNotFound      = errors.New("recipe not found")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)
