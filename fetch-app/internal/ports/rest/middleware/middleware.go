package middleware

import (
	"net/http"
	"strings"

	"github.com/hengkysuryaa/backend-service/fetch-app/pkg/jwt"
)

const ROLE_ADMIN_VALUE = "admin"

func AuthorizeAdmin(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// Token format: Bearer <token>
		if len(strings.Split(authorization, " ")) != 2 {
			http.Error(rw, "invalid authorization header format", http.StatusForbidden)
			return
		}

		token := strings.Split(authorization, " ")[1]
		ctx, err := jwt.ReadToken(r.Context(), token)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		user := jwt.GetTokenData(ctx.Value(jwt.TokenDataKey))
		if strings.ToLower(user["role"].(string)) != ROLE_ADMIN_VALUE {
			http.Error(rw, "the service only can be accessed by admin role", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func AuthorizeAll(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(rw, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		// Token format: Bearer <token>
		if len(strings.Split(authorization, " ")) != 2 {
			http.Error(rw, "invalid authorization header format", http.StatusForbidden)
			return
		}

		token := strings.Split(authorization, " ")[1]
		ctx, err := jwt.ReadToken(r.Context(), token)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(rw, r.WithContext(ctx))
	})
}
