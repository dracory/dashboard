package dashboard_test

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/shared"
)

func TestThemeHandler(t *testing.T) {
	tests := []struct {
		name           string
		theme          string
		redirect       string
		secureRequest  bool
		expectedStatus int
		expectedTheme  string
		expectedPath   string
		expectedSecure bool
	}{
		{
			name:           "secure request with theme and redirect",
			theme:          "dark",
			redirect:       "/dashboard",
			secureRequest:  true,
			expectedStatus: http.StatusFound,
			expectedTheme:  "dark",
			expectedPath:   "/dashboard",
			expectedSecure: true,
		},
		{
			name:           "non-secure request with theme and redirect",
			theme:          "light",
			redirect:       "/home",
			secureRequest:  false,
			expectedStatus: http.StatusFound,
			expectedTheme:  "light",
			expectedPath:   "/home",
			expectedSecure: false,
		},
		{
			name:           "empty theme with default redirect",
			theme:          "",
			redirect:       "",
			secureRequest:  true,
			expectedStatus: http.StatusFound,
			expectedTheme:  "",
			expectedPath:   "/",
			expectedSecure: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with query parameters
			req, err := http.NewRequest("GET", "/theme?theme="+tt.theme+"&redirect="+tt.redirect, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Set TLS if this is a secure request
			if tt.secureRequest {
				req.TLS = &tls.ConnectionState{}
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(dashboard.ThemeHandler)
			handler.ServeHTTP(rr, req)

			// Check the status code
			if gotStatus := rr.Code; gotStatus != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", gotStatus, tt.expectedStatus)
			}

			// Check the Location header for redirect
			location, err := rr.Result().Location()
			if err != nil {
				t.Fatal(err)
			}
			if location.Path != tt.expectedPath {
				t.Errorf("handler returned wrong redirect path: got %v want %v", location.Path, tt.expectedPath)
			}

			// Check the cookie
			cookies := rr.Result().Cookies()
			var themeCookie *http.Cookie
			for _, cookie := range cookies {
				if cookie.Name == shared.THEME_COOKIE_KEY {
					themeCookie = cookie
					break
				}
			}

			if tt.expectedTheme == "" {
				if themeCookie != nil {
					t.Error("Expected no theme cookie to be set")
				}
			} else {
				if themeCookie == nil {
					t.Fatal("Expected theme cookie to be set")
				}
				if themeCookie.Value != tt.expectedTheme {
					t.Errorf("Unexpected theme cookie value: got %v want %v", themeCookie.Value, tt.expectedTheme)
				}
				if themeCookie.Path != "/" {
					t.Errorf("Cookie path should be /, got %v", themeCookie.Path)
				}
				if !themeCookie.Expires.After(now) {
					t.Error("Cookie should have future expiration")
				}
				if themeCookie.Secure != tt.expectedSecure {
					t.Errorf("Cookie secure flag mismatch: got %v want %v", themeCookie.Secure, tt.expectedSecure)
				}
			}
		})
	}
}

// now is a helper function to get the current time for cookie expiration checks
var now = time.Now()
