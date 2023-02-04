package authn

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type (
	Key        int
	Middleware func(next http.Handler) http.Handler
)

const (
	// Keys
	IdentityID Key = iota
	// Headers
	tokenLength   = 2
	authorization = "Authorization"
)

var ErrTokenInvalid = errors.New("token_invalid")

func FromJWT(ctx context.Context, jwksURL string) (Middleware, error) {
	options := keyfunc.Options{
		Ctx:               ctx,
		Client:            &http.Client{Timeout: 30 * time.Second},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Minute,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWKs: %w", err)
	}

	parse := func(t string) (uuid.UUID, error) {
		token, err := jwt.ParseWithClaims(t, &jwt.RegisteredClaims{}, jwks.Keyfunc)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to parse token: %w", err)
		}

		claims, ok := token.Claims.(*jwt.RegisteredClaims)
		if !ok || !token.Valid {
			return uuid.Nil, ErrTokenInvalid
		}

		if err = claims.Valid(); err != nil {
			return uuid.Nil, fmt.Errorf("failed to validate claims: %w", err)
		}

		u, err := uuid.Parse(claims.Subject)
		if err != nil {
			return uuid.Nil, fmt.Errorf("failed to validate claims: %w", err)
		}

		return u, nil
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if q, ok := TokenFromHeader(r); ok {
				if t, err := parse(q); err == nil {
					ctx = context.WithValue(ctx, IdentityID, t)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}, nil
}

func FromUUID(id uuid.UUID) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, IdentityID, id)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func FromRandomUUID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, IdentityID, uuid.New())

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Enforce(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(IdentityID) == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func TokenFromHeader(r *http.Request) (value string, found bool) {
	if arr := strings.Split(r.Header.Get(authorization), " "); len(arr) == tokenLength {
		found = true
		value = arr[1]
	}

	return
}

func GetIdentityID(ctx context.Context) (uuid.UUID, bool) {
	if i, ok := ctx.Value(IdentityID).(uuid.UUID); ok && i != uuid.Nil {
		return i, true
	}

	return uuid.Nil, false
}
