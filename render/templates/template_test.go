package templates_test

import (
	"testing"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render/templates/adminlte"
	"github.com/dracory/dashboard/render/templates/bootstrap"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/dracory/dashboard/render/templates/tabler"
)

func TestTemplateRendering(t *testing.T) {
	tests := []struct {
		name  string
		theme shared.Template
	}{
		{
			name:  "Bootstrap",
			theme: bootstrap.NewBootstrapTemplate(),
		},
		{
			name:  "AdminLTE",
			theme: adminlte.NewAdminLTETemplate(),
		},
		{
			name:  "Tabler",
			theme: tabler.NewTablerTemplate(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test dashboard rendering
			t.Run("dashboard", func(t *testing.T) {
				mockRenderer := &mockDashboardRenderer{
					content:      "<div>Test Dashboard Content</div>",
					templateName: tt.theme.GetName(),
				}
				
				// Test RenderDashboard
				_, err := tt.theme.RenderDashboard(mockRenderer)
				if err != nil {
					t.Fatalf("%s: Failed to render dashboard: %v", tt.name, err)
				}

				// Test RenderPage
				_, err = tt.theme.RenderPage("<div>Test Page Content</div>", mockRenderer)
				if err != nil {
					t.Fatalf("%s: Failed to render page: %v", tt.name, err)
				}
			})

			// Test template methods
			t.Run("template_methods", func(t *testing.T) {
				// Test GetName
				name := tt.theme.GetName()
				if name == "" {
					t.Errorf("%s: Expected template name to be non-empty", tt.name)
				}

				// Test GetCSSLinks
				cssLinks := tt.theme.GetCSSLinks(false)
				if len(cssLinks) == 0 {
					t.Logf("%s: No CSS links returned (this might be expected)", tt.name)
				}

				// Test GetJSScripts
				jsScripts := tt.theme.GetJSScripts()
				if len(jsScripts) == 0 {
					t.Logf("%s: No JS scripts returned (this might be expected)", tt.name)
				}

				// Test GetCustomCSS and GetCustomJS (just check they don't panic)
				_ = tt.theme.GetCustomCSS()
				_ = tt.theme.GetCustomJS()
			})
		})
	}
}

// mockDashboardRenderer is a test implementation of the DashboardRenderer interface
type mockDashboardRenderer struct {
	content      string
	templateName string
}

func (m *mockDashboardRenderer) GetContent() string {
	return m.content
}

func (m *mockDashboardRenderer) GetFaviconURL() string {
	return ""
}

func (m *mockDashboardRenderer) GetLogoImageURL() string {
	return ""
}

func (m *mockDashboardRenderer) GetLogoRawHtml() string {
	return ""
}

func (m *mockDashboardRenderer) GetLogoRedirectURL() string {
	return "/"
}

func (m *mockDashboardRenderer) GetTemplateName() string {
	return m.templateName
}

func (m *mockDashboardRenderer) GetMenuItems() []model.MenuItem {
	return []model.MenuItem{
		{Text: "Home", URL: "/"},
	}
}

func (m *mockDashboardRenderer) GetMenuShowText() bool {
	return true
}

func (m *mockDashboardRenderer) GetQuickAccessMenu() []model.MenuItem {
	return []model.MenuItem{}
}

func (m *mockDashboardRenderer) GetUser() model.User {
	return model.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
}

func (m *mockDashboardRenderer) GetUserMenu() []model.MenuItem {
	return []model.MenuItem{
		{Text: "Profile", URL: "/profile"},
		{Text: "Logout", URL: "/logout"},
	}
}

func (m *mockDashboardRenderer) GetLoginURL() string {
	return "/login"
}

func (m *mockDashboardRenderer) GetRegisterURL() string {
	return "/register"
}

func (m *mockDashboardRenderer) GetNavbarBackgroundColorMode() string {
	return "light"
}

func (m *mockDashboardRenderer) GetNavbarBackgroundColor() string {
	return "#fff"
}

func (m *mockDashboardRenderer) GetNavbarTextColor() string {
	return "#000"
}
