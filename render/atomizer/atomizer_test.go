package atomizer

import (
	"testing"

	"github.com/dracory/dashboard/config"
	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
)



// Helper functions to work with omni.AtomInterface
type atomHelper struct{}

func (h *atomHelper) getAtomType(a omni.AtomInterface) string {
	if a == nil {
		return ""
	}
	return a.GetType()
}

func (h *atomHelper) getChildren(a omni.AtomInterface) []omni.AtomInterface {
	if a == nil {
		return nil
	}
	return a.ChildrenGet()
}

func (h *atomHelper) getProperty(a omni.AtomInterface, key string) string {
	if a == nil {
		return ""
	}
	return a.Get(key)
}

var helper = &atomHelper{}

// mockDashboard implements the DashboardRenderer interface for testing
type mockDashboard struct {
	content      string
	menuItems    []model.MenuItem
	user         *model.User
	userMenu     []model.MenuItem
	logoImageURL string
	logoURL      string
}

func (m *mockDashboard) GetContent() string                   { return m.content }
func (m *mockDashboard) GetFaviconURL() string                { return "" }
func (m *mockDashboard) GetLogoImageURL() string              { return m.logoImageURL }
func (m *mockDashboard) GetLogoRawHtml() string               { return "" }
func (m *mockDashboard) GetTemplateName() string              { return config.TEMPLATE_TABLER }
func (m *mockDashboard) GetLogoRedirectURL() string           { return m.logoURL }
func (m *mockDashboard) GetMenuItems() []model.MenuItem       { return m.menuItems }
func (m *mockDashboard) GetMenuShowText() bool                { return true }
func (m *mockDashboard) GetQuickAccessMenu() []model.MenuItem { return nil }
func (m *mockDashboard) GetUser() model.User {
	if m.user == nil {
		return model.User{}
	}
	return *m.user
}
func (m *mockDashboard) GetUserMenu() []model.MenuItem        { return m.userMenu }
func (m *mockDashboard) GetLoginURL() string                  { return "/login" }
func (m *mockDashboard) GetRegisterURL() string               { return "/register" }
func (m *mockDashboard) GetNavbarBackgroundColorMode() string { return "light" }
func (m *mockDashboard) GetNavbarBackgroundColor() string     { return "#ffffff" }
func (m *mockDashboard) GetNavbarTextColor() string           { return "#000000" }

func TestTransformDashboard(t *testing.T) {
	tests := []struct {
		name      string
		dashboard *mockDashboard
		validate  func(t *testing.T, atom omni.AtomInterface, err error)
	}{
		{
			name: "empty dashboard",
			dashboard: &mockDashboard{
				content: "Test content",
			},
			validate: func(t *testing.T, atom omni.AtomInterface, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if atom == nil {
					t.Fatal("expected atom to not be nil")
				}
				if gotType := helper.getAtomType(atom); gotType != atomizer.AtomTypeDashboard {
					t.Errorf("expected type %q, got %q", atomizer.AtomTypeDashboard, gotType)
				}
				// Should have header, content, and footer children
				children := helper.getChildren(atom)
				if len(children) < 3 {
					t.Errorf("expected at least 3 children, got %d", len(children))
				}
			},
		},
		{
			name: "dashboard with menu",
			dashboard: &mockDashboard{
				content: "Test content",
				menuItems: []model.MenuItem{
					{ID: "1", Text: "Home", URL: "/"},
					{ID: "2", Text: "About", URL: "/about"},
				},
			},
			validate: func(t *testing.T, atom omni.AtomInterface, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				// Verify menu items were transformed
				var menu omni.AtomInterface
				children := helper.getChildren(atom)
				for _, child := range children {
					if helper.getAtomType(child) == atomizer.AtomTypeHeader {
						headerChildren := helper.getChildren(child)
						for _, headerChild := range headerChildren {
							if helper.getAtomType(headerChild) == atomizer.AtomTypeMenu {
								menu = headerChild
								break
							}
						}
					}
				}
				if menu == nil {
					t.Fatal("Menu should exist in header")
				}
				menuChildren := helper.getChildren(menu)
				if len(menuChildren) < 2 {
					t.Errorf("expected at least 2 menu items, got %d", len(menuChildren))
				}
			},
		},
		{
			name: "dashboard with user menu",
			dashboard: &mockDashboard{
				content: "Test content",
				user: &model.User{
					ID:   "123",
					Name: "Test User",
				},
				userMenu: []model.MenuItem{
					{ID: "profile", Text: "Profile", URL: "/profile"},
					{ID: "settings", Text: "Settings", URL: "/settings"},
				},
			},
			validate: func(t *testing.T, atom omni.AtomInterface, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				// Verify user menu was transformed
				var userMenu omni.AtomInterface
				children := helper.getChildren(atom)
				header := children[0] // First child should be header
				headerChildren := helper.getChildren(header)
				for _, child := range headerChildren {
					if helper.getAtomType(child) == "user_menu" {
						userMenu = child
						break
					}
				}
				if userMenu == nil {
					t.Fatal("User menu should exist in header")
				}
			},
		},
	}

	transformer := atomizer.NewTransformer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			atom, err := transformer.TransformDashboard(tt.dashboard)
			tt.validate(t, atom, err)
		})
	}
}

func TestTransformMenu(t *testing.T) {
	transformer := atomizer.NewTransformer()

	t.Run("empty menu", func(t *testing.T) {
		menu, err := transformer.TransformMenu(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if menu == nil {
			t.Fatal("expected menu to not be nil")
		}
		children := helper.getChildren(menu)
		if len(children) != 0 {
			t.Errorf("expected 0 children, got %d", len(children))
		}
	})

	t.Run("menu with items", func(t *testing.T) {
		menuItems := []model.MenuItem{
			{
				ID:   "1",
				Text: "Home",
				URL:  "/",
			},
			{
				ID:     "2",
				Text:   "About",
				URL:    "/about",
				Active: true,
			},
			{
				ID:   "3",
				Text: "Services",
				Children: []model.MenuItem{
					{
						ID:   "3.1",
						Text: "Web",
						URL:  "/services/web",
					},
					{
						ID:   "3.2",
						Text: "Mobile",
						URL:  "/services/mobile",
					},
				},
			},
		}

		menu, err := transformer.TransformMenu(menuItems)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		children := helper.getChildren(menu)
		if len(children) != 3 {
			t.Fatalf("expected 3 children, got %d", len(children))
		}

		// Verify submenu
		services := children[2]
		servicesChildren := helper.getChildren(services)
		if len(servicesChildren) == 0 {
			t.Fatal("expected services to have children")
		}
		submenu := servicesChildren[0]
		if helper.getAtomType(submenu) != AtomTypeMenu {
			t.Errorf("expected type %q, got %q", AtomTypeMenu, helper.getAtomType(submenu))
		}
		submenuChildren := helper.getChildren(submenu)
		if len(submenuChildren) != 2 {
			t.Errorf("expected 2 submenu children, got %d", len(submenuChildren))
		}
	})
}

func TestTransformHeader_WithLogo(t *testing.T) {
	transformer := atomizer.NewTransformer()
	h := &atomHelper{}
	dashboard := &mockDashboard{
		logoImageURL: "/logo.png",
		logoURL:      "/",
	}

	header, err := transformer.TransformHeader(dashboard)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Find logo link
	var logoLink omni.AtomInterface
	headerChildren := h.getChildren(header)
	for _, child := range headerChildren {
		children := h.getChildren(child)
		if len(children) > 0 {
			firstChild := children[0]
			if h.getAtomType(child) == atomizer.AtomTypeLink && h.getAtomType(firstChild) == atomizer.AtomTypeImage {
				logoLink = child
				break
			}
		}
	}

	if logoLink == nil {
		t.Fatal("Logo link should exist")
	}

	prop := h.getProperty(logoLink, atomizer.PropHref)
	if prop == "" {
		t.Fatal("expected href property to exist")
	}
	href := prop
	if href != "/" {
		t.Errorf("expected href to be %q, got %q", "/", href)
	}

	logoChildren := h.getChildren(logoLink)
	if len(logoChildren) == 0 {
		t.Fatal("logo link should have an image child")
	}
	imgProp := h.getProperty(logoChildren[0], atomizer.PropSrc)
	if imgProp == "" {
		t.Fatal("expected src property to exist on image")
	}
	imgSrc := imgProp
	if imgSrc != "/logo.png" {
		t.Errorf("expected image src to be %q, got %q", "/logo.png", imgSrc)
	}
}
