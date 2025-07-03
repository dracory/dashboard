package dashboard_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/shared"
)

func TestThemeMiddleware(t *testing.T) {
	tests := []struct {
		name        string
		themeCookie string
		expected    string
	}{
		{
			name:        "no theme cookie",
			themeCookie: "",
			expected:    "",
		},
		{
			name:        "with theme cookie",
			themeCookie: "dark",
			expected:    "dark",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.themeCookie != "" {
				req.AddCookie(&http.Cookie{
					Name:  "theme",
					Value: tt.themeCookie,
				})
			}

			handlerCalled := false
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
				theme, ok := r.Context().Value(shared.ThemeNameContextKey{}).(string)
				if !ok {
					t.Fatal("theme not found in context")
				}
				if theme != tt.expected {
					t.Errorf("expected theme %q, got %q", tt.expected, theme)
				}
			})

			middleware := dashboard.ThemeMiddleware(handler)
			middleware.ServeHTTP(httptest.NewRecorder(), req)

			if !handlerCalled {
				t.Fatal("next handler was not called")
			}
		})
	}
}

func TestThemeMiddleware_ContextPropagation(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	originalCtx := context.WithValue(req.Context(), "testKey", "testValue")
	req = req.WithContext(originalCtx)

	handlerCalled := false
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		// Check that original context values are preserved
		val, ok := r.Context().Value("testKey").(string)
		if !ok {
			t.Fatal("original context value not found")
		}
		if val != "testValue" {
			t.Errorf("expected context value 'testValue', got %q", val)
		}

		// Check that theme value is set
		theme, ok := r.Context().Value(shared.ThemeNameContextKey{}).(string)
		if !ok {
			t.Fatal("theme not found in context")
		}
		if theme != "" {
			t.Errorf("expected empty theme, got %q", theme)
		}
	})

	middleware := dashboard.ThemeMiddleware(handler)
	middleware.ServeHTTP(httptest.NewRecorder(), req)

	if !handlerCalled {
		t.Fatal("next handler was not called")
	}
}
