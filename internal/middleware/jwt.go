package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"tracker-server/internal/security"
)

type deviceIDKey struct{}

// TODO нормальный логгер
func JWTAuthMiddleware(manager *security.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				err := errors.New("token not found in header") //TODO: заменить на кастомную ошибку
				log.Println(err)
				return
			}
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			deviceClaims, err := manager.ParseToken(tokenStr)
			if err != nil {
				err = fmt.Errorf("parse token: %w", err)
				log.Println(err)
				return
			}
			ctx := ctxWithClaimsValues(r.Context(), deviceClaims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ctxWithClaimsValues(ctx context.Context, claims security.DeviceClaims) context.Context {
	return context.WithValue(ctx, deviceIDKey{}, claims.DeviceID)
}
