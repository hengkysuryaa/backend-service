package jwt

import (
	"context"
	"errors"
	"log"

	jwt "github.com/golang-jwt/jwt/v4"
)

const JWT_SIGNATURE_KEY = "H31kR4l3J4JMsLDpJsMlpcY6kqAqNBKQ8A4u91lhuJsVOOVGkfZYXaP2R3MjskWn"

var (
	TokenDataKey = &contextKey{"token-data"}
)

type contextKey struct {
	name string
}

func ReadToken(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		}

		return []byte(JWT_SIGNATURE_KEY), nil
	})

	if err != nil {
		log.Println(err)
		return ctx, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New("failed to decode token")
		log.Println(err)
		return ctx, err
	}

	ctx = context.WithValue(ctx, TokenDataKey, claims)

	return ctx, nil
}

func GetTokenData(claims interface{}) map[string]interface{} {
	return claims.(jwt.MapClaims)
}
