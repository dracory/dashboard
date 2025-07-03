package dashboard

import (
	"context"
	"net/http"

	"github.com/dracory/dashboard/shared"
)

func ThemeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		themeName := ThemeNameRetrieveFromCookie(r)

		ctx := context.WithValue(r.Context(), shared.ThemeNameContextKey{}, themeName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
