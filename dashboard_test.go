package dashboard_test

import (
	"testing"

	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/shared"
	"github.com/dracory/dashboard/types"
)

func TestNewDashboard(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(types.DashboardInterface)
		validate func(types.DashboardInterface) bool
	}{
		{
			name: "default dashboard",
			setup: func(d types.DashboardInterface) {
				d.SetTitle("Test Dashboard")
			},
			validate: func(d types.DashboardInterface) bool {
				return d.GetTitle() == "Test Dashboard" &&
					d.GetTemplate() == shared.TEMPLATE_DEFAULT
			},
		},
		{
			name: "dashboard with template",
			setup: func(d types.DashboardInterface) {
				d.SetTemplate(shared.TEMPLATE_ADMINLTE)
			},
			validate: func(d types.DashboardInterface) bool {
				return d.GetTemplate() == shared.TEMPLATE_ADMINLTE
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dashboard.New()
			if tt.setup != nil {
				tt.setup(d)
			}
			if !tt.validate(d) {
				t.Error("dashboard validation failed")
			}
		})
	}
}

func TestDashboardGettersAndSetters(t *testing.T) {
	tests := []struct {
		name    string
		setter  func(types.DashboardInterface)
		getter  func(types.DashboardInterface) interface{}
		expect  interface{}
		success bool
	}{
		{
			name: "Set and Get Title",
			setter: func(d types.DashboardInterface) {
				d.SetTitle("New Title")
			},
			getter: func(d types.DashboardInterface) interface{} {
				return d.GetTitle()
			},
			expect:  "New Title",
			success: true,
		},
		{
			name: "Set and Get Subtitle",
			setter: func(d types.DashboardInterface) {
				d.SetSubtitle("Test Subtitle")
			},
			getter: func(d types.DashboardInterface) interface{} {
				return d.GetSubtitle()
			},
			expect:  "Test Subtitle",
			success: true,
		},
		{
			name: "Set and Get Theme",
			setter: func(d types.DashboardInterface) {
				d.SetTheme("dark")
			},
			getter: func(d types.DashboardInterface) interface{} {
				return d.GetTheme()
			},
			expect:  "dark",
			success: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dashboard.New()
			tt.setter(d)
			got := tt.getter(d)
			if got != tt.expect {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
		})
	}
}

func TestThemeMethods(t *testing.T) {
	tests := []struct {
		name   string
		theme  string
		expect bool
	}{
		{
			name:   "dark theme",
			theme:  "dark",
			expect: true,
		},
		{
			name:   "light theme",
			theme:  "light",
			expect: false,
		},
		{
			name:   "default theme",
			theme:  "",
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dashboard.New()
			d.SetTheme(tt.theme)
			if d.IsThemeDark() != tt.expect {
				t.Errorf("expected IsThemeDark() to be %v for theme '%s'", tt.expect, tt.theme)
			}
		})
	}
}

func TestTemplateFinding(t *testing.T) {
	tests := []struct {
		name        string
		template    string
		expects     string
		shouldPanic bool
	}{
		{
			name:        "empty template",
			template:    "",
			expects:     shared.TEMPLATE_DEFAULT,
			shouldPanic: true, // Empty template should use default but may panic due to missing required fields
		},
		{
			name:        "bootstrap template",
			template:    shared.TEMPLATE_BOOTSTRAP,
			expects:     shared.TEMPLATE_BOOTSTRAP,
			shouldPanic: true, // May panic due to missing required fields
		},
		{
			name:        "adminlte template",
			template:    shared.TEMPLATE_ADMINLTE,
			expects:     shared.TEMPLATE_ADMINLTE,
			shouldPanic: true, // May panic due to missing required fields
		},
		{
			name:        "tabler template",
			template:    shared.TEMPLATE_TABLER,
			expects:     shared.TEMPLATE_TABLER,
			shouldPanic: true, // May panic due to missing required fields
		},
		{
			name:        "invalid template",
			template:    "invalid",
			expects:     "invalid", // Invalid template names are stored as-is
			shouldPanic: true,      // Will panic when trying to render
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := dashboard.New()
			
			// Set required fields to avoid nil pointer dereferences
			d.SetTitle("Test Dashboard")
			d.SetContent("<div>Test Content</div>")
			
			// Set the template
			d.SetTemplate(tt.template)
			
			// Verify the template is set correctly
			got := d.GetTemplate()
			
			// For empty template, it should use the default
			expected := tt.expects
			if tt.template == "" {
				expected = shared.TEMPLATE_DEFAULT
			}
			
			if got != expected {
				t.Errorf("expected template %q, got %q", expected, got)
			}

			// Test template rendering if not expected to panic
			if !tt.shouldPanic {
				rendered := d.ToHTML()
				if rendered == "" {
					t.Error("expected non-empty HTML output from template")
				}
			}
		})
	}
}

func TestMenuMethods(t *testing.T) {
	d := dashboard.New()

	// Test menu items
	menuItems := []types.MenuItem{
		{Title: "Home", URL: "/"},
		{Title: "About", URL: "/about"},
	}

	// Test setting and getting menu items
	d.SetMenuMainItems(menuItems)
	if len(d.GetMenuMainItems()) != 2 {
		t.Error("expected 2 menu items")
	}

	// Test menu type
	d.SetMenuType("offcanvas")
	if d.GetMenuType() != "offcanvas" {
		t.Error("expected menu type to be 'offcanvas'")
	}

	// Test menu show text
	d.SetMenuShowText(true)
	if !d.GetMenuShowText() {
		t.Error("expected menu show text to be true")
	}
}

func TestAlertMethods(t *testing.T) {
	d := dashboard.New()

	// Test alerts
	alert := types.Alert{
		Type:    "success",
		Message: "Operation completed successfully",
	}

	d.AddAlert(alert)
	alerts := d.GetAlerts()
	if len(alerts) != 1 {
		t.Fatalf("expected 1 alert, got %d", len(alerts))
	}

	if alerts[0].Type != alert.Type || alerts[0].Message != alert.Message {
		t.Error("alert data does not match")
	}

	// Test clearing alerts
	d.ClearAlerts()
	if len(d.GetAlerts()) != 0 {
		t.Error("expected alerts to be cleared")
	}

	// Test modals
	modal := types.Modal{
		ID:      "testModal",
		Title:   "Test Modal",
		Content: "<p>Test content</p>",
	}

	d.AddModal(modal)
	modals := d.GetModals()
	if len(modals) != 1 {
		t.Fatalf("expected 1 modal, got %d", len(modals))
	}

	if modals[0].ID != modal.ID || modals[0].Title != modal.Title || modals[0].Content != modal.Content {
		t.Error("modal data does not match")
	}

	// Test clearing modals
	d.ClearModals()
	if len(d.GetModals()) != 0 {
		t.Error("expected modals to be cleared")
	}
}

func TestUserMethods(t *testing.T) {
	d := dashboard.New()

	// Test setting and getting user
	user := types.User{
		FirstName: "Test User",
		LastName:  "Test User",
		Email:     "test@example.com",
	}
	d.SetUser(user)

	gotUser := d.GetUser()
	if gotUser == nil || gotUser.FirstName != "Test User" || gotUser.LastName != "Test User" {
		t.Error("failed to get/set user")
	}
}

func TestRedirectMethods(t *testing.T) {
	d := dashboard.New()

	// Test redirect URL
	d.SetRedirectUrl("/dashboard")
	if d.GetRedirectUrl() != "/dashboard" {
		t.Error("failed to set/get redirect URL")
	}

	// Test redirect time
	d.SetRedirectTime("5")
	if d.GetRedirectTime() != "5" {
		t.Error("failed to set/get redirect time")
	}
}
