package adminlte

import (
	"github.com/dracory/dashboard/render/theme/shared"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// AdminLTETheme implements the shared.Theme interface for AdminLTE
type AdminLTETheme struct {
	shared.DefaultTheme // Embed DefaultTheme to inherit default implementations
}

// newAdminLTETheme creates a new instance of the AdminLTE theme
// This is an internal helper function used by NewAdminLTETheme in export.go
func newAdminLTETheme() *AdminLTETheme {
	return &AdminLTETheme{}
}

// RenderAtom renders an Omni atom using AdminLTE classes and components
func (t *AdminLTETheme) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	if atom == nil {
		return nil, nil
	}
	return t.renderAtom(atom)
}

// RenderHeader renders the AdminLTE theme header
func (t *AdminLTETheme) RenderHeader(d shared.DashboardRenderer) *hb.Tag {
	header := hb.NewTag("header").Class("main-header")
	
	// Navbar
	navbar := hb.NewTag("nav").Class("navbar navbar-expand navbar-white navbar-light")
	
	// Left navbar links
	leftNav := hb.NewTag("ul").Class("navbar-nav")
	leftNav.Child(hb.NewTag("li").Class("nav-item").
		Child(hb.NewTag("a").Class("nav-link").
			Attr("data-widget", "pushmenu").
			Attr("href", "#").
			Child(hb.I().Class("fas fa-bars"))))
	
	navbar.Child(leftNav)
	
	// Right navbar links
	rightNav := hb.NewTag("ul").Class("navbar-nav ml-auto")
	
	// User menu
	user := d.GetUser()
	if user.Name != "" {
		// User dropdown menu
		userMenu := hb.NewTag("li").Class("nav-item dropdown")
		userLink := hb.NewTag("a").Class("nav-link").
			Attr("data-toggle", "dropdown").
			Attr("href", "#")
		
		// User avatar or initials
		userInfo := hb.NewTag("div").Class("user-panel d-flex align-items-center")
		if user.AvatarURL != "" {
			userInfo.Child(hb.NewTag("div").Class("image").
				Child(hb.Img(user.AvatarURL).Class("img-circle elevation-2").Attr("alt", "User Image")))
		} else {
			initials := ""
			if len(user.Name) > 0 {
				initials = string(user.Name[0])
			}
			userInfo.Child(hb.NewTag("div").Class("image").
				Child(hb.Span().Class("img-circle elevation-2 bg-primary text-white d-flex align-items-center justify-content-center").
					Style("width: 2.1rem; height: 2.1rem; display: inline-flex;").
					Text(initials)))
		}
		userInfo.Child(hb.Span().Class("d-none d-md-inline").Text(user.Name))
		userLink.Child(userInfo)
		
		dropdownMenu := hb.NewTag("div").Class("dropdown-menu dropdown-menu-lg dropdown-menu-right")
		dropdownMenu.Child(hb.Span().Class("dropdown-item dropdown-header").Text("Account"))
		dropdownMenu.Child(hb.NewTag("div").Class("dropdown-divider"))
		dropdownMenu.Child(hb.NewTag("a").Class("dropdown-item").
			Attr("href", "/profile").
			Child(hb.I().Class("fas fa-user mr-2")).
			Text("Profile"))
		dropdownMenu.Child(hb.NewTag("div").Class("dropdown-divider"))
		dropdownMenu.Child(hb.NewTag("a").Class("dropdown-item").
			Attr("href", "/logout").
			Child(hb.I().Class("fas fa-sign-out-alt mr-2")).
			Text("Logout"))
		
		userMenu.Child(userLink)
		userMenu.Child(dropdownMenu)
		rightNav.Child(userMenu)
	} else {
		// Login/Signup buttons for guests
		rightNav.Child(hb.NewTag("li").Class("nav-item").
			Child(hb.NewTag("a").Class("nav-link").
				Attr("href", "/login").
				Text("Login")))
		rightNav.Child(hb.NewTag("li").Class("nav-item").
			Child(hb.NewTag("a").Class("nav-link").
				Attr("href", "/register").
				Text("Register")))
	}
	
	navbar.Child(rightNav)
	header.Child(navbar)
	
	return header
}

// RenderFooter renders the AdminLTE theme footer
func (t *AdminLTETheme) RenderFooter(d shared.DashboardRenderer) *hb.Tag {
	footer := hb.NewTag("footer").Class("main-footer")
	
	// Footer content
	container := hb.NewTag("div").Class("container-fluid")
	
	// Left side (copyright)
	leftCol := hb.NewTag("div").Class("float-right d-none d-sm-block")
	leftCol.Child(hb.NewTag("b").Text("Version")).Text(" 1.0.0")
	
	// Right side (copyright)
	rightCol := hb.NewTag("div")
	rightCol.Child(hb.Text("© 2023 ")).
		Child(hb.NewTag("a").Attr("href", "#").Text("AdminLTE.io")).
		Child(hb.Text(". All rights reserved."))
	
	container.Child(leftCol)
	container.Child(rightCol)
	footer.Child(container)
	
	return footer
}

// RenderDashboard renders a complete dashboard from Omni atoms
func (t *AdminLTETheme) RenderDashboard(dashboard *omni.Atom) (string, error) {
	if dashboard == nil {
		return "", nil
	}
	return t.renderDashboard(dashboard)
}

// Ensure AdminLTETheme implements shared.Theme
var _ shared.Theme = (*AdminLTETheme)(nil)

// GetName returns the name of the theme
func (t *AdminLTETheme) GetName() string {
	return "adminlte"
}

// GetCSSLinks returns the CSS link tags for the theme
func (t *AdminLTETheme) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return GetAdminLTEAssets()
}

// GetJSScripts returns the JavaScript script tags for the theme
func (t *AdminLTETheme) GetJSScripts() []*hb.Tag {
	return GetAdminLTEScripts()
}

// GetCustomCSS returns any custom CSS for the theme
func (t *AdminLTETheme) GetCustomCSS() string {
	return `
		/* AdminLTE custom styles */
		.content-wrapper {
			min-height: calc(100vh - 56px);
		}
		.main-sidebar {
			padding-top: 56px;
		}
		.navbar-nav .user-menu {
			height: 50px;
		}
		.brand-link {
			height: 56px;
		}
		.brand-link img {
			height: 33px;
			width: auto;
		}
	`
}

// GetCustomJS returns any custom JavaScript for the theme
func (t *AdminLTETheme) GetCustomJS() string {
	return `
		// Initialize AdminLTE
		document.addEventListener('DOMContentLoaded', function() {
			// Enable tooltips
			$('[data-toggle="tooltip"]').tooltip();
			
			// Enable popovers
			$('[data-toggle="popover"]').popover();
			
			// Initialize sidebar menu
			$('[data-widget="treeview"]').each(function() {
				$(this).Treeview('init');
			});
			
			// Theme switcher
			$('[data-bs-theme-value]').on('click', function() {
				var theme = $(this).data('bs-theme-value');
				$('body')
					.removeClass('dark-mode')
					.removeClass('light-mode');
				
				if (theme === 'dark') {
					$('body').addClass('dark-mode');
				} else {
					$('body').addClass('light-mode');
				}
				
				localStorage.setItem('theme', theme);
			});
			
			// Set theme from localStorage
			var theme = localStorage.getItem('theme');
			if (theme) {
				$('body')
					.removeClass('dark-mode')
					.removeClass('light-mode')
					.addClass(theme + '-mode');
			}
		});
	`
}
