package authz

import (
	"context"
	"net/http"

	"github.com/blue-health/blue-go-toolbox/authn"
	"github.com/google/uuid"
)

type (
	Guard interface {
		// Check authorizes the Identity ID stored in the context with the given permissions.
		// Therefore, it should be used after the Enforce() middleware.
		Check(...string) Middleware
	}

	// GuardImpl implements Guard. It stores a IdentityService to use to authorize an identity.
	GuardImpl struct{ authorizer Authorizer }

	Authorizer interface {
		Authorize(context.Context, uuid.UUID, ...string) (bool, error)
	}

	Middleware func(next http.Handler) http.Handler
)

func NewGuard(authorizer Authorizer) *GuardImpl {
	return &GuardImpl{authorizer: authorizer}
}

func (g *GuardImpl) Check(permissions ...string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, ok := authn.GetIdentityID(r.Context())
			if !ok {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if ok, err := g.authorizer.Authorize(r.Context(), id, permissions...); err != nil || !ok {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
