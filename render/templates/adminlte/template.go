package adminlte

import (
	"strings"

	"github.com/dracory/dashboard/render"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// AdminLTETemplate implements the shared.Template interface for AdminLTE
type AdminLTETemplate struct {
	shared.DefaultTemplate // Embed DefaultTemplate to inherit default implementations
}

// newAdminLTETemplate creates a new instance of the AdminLTE template
// This is an internal helper function used by NewAdminLTETemplate in export.go
func newAdminLTETemplate() *AdminLTETemplate {
	return &AdminLTETemplate{}
}

// RenderAtom renders an Omni atom using AdminLTE classes and components
func (t *AdminLTETemplate) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	if atom == nil {
		return nil, nil
	}
	return t.renderAtom(atom)
}

// RenderHeader renders the AdminLTE template header
func (t *AdminLTETemplate) RenderHeader(d shared.DashboardRenderer) *hb.Tag {
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

// RenderFooter renders the AdminLTE template footer
func (t *AdminLTETemplate) RenderFooter(d shared.DashboardRenderer) *hb.Tag {
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
func (t *AdminLTETemplate) RenderDashboard(dashboard *omni.Atom) (string, error) {
	if dashboard == nil {
		return "", nil
	}
	return t.renderDashboard(dashboard)
}

// Ensure AdminLTETemplate implements shared.Template
var _ shared.Template = (*AdminLTETemplate)(nil)

// GetName returns the name of the template
func (t *AdminLTETemplate) GetName() string {
	return render.THEME_ADMINLTE
}

// GetCSSLinks returns the CSS link tags for the theme
func (t *AdminLTETemplate) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return GetAdminLTEAssets()
}

// GetJSScripts returns the JavaScript script tags for the theme
func (t *AdminLTETemplate) GetJSScripts() []*hb.Tag {
	return GetAdminLTEScripts()
}

// GetCustomCSS returns any custom CSS for the theme
func (t *AdminLTETemplate) GetCustomCSS() string {
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
func (t *AdminLTETemplate) GetCustomJS() string {
	return `
		// Enable sidebar toggle
		document.addEventListener('DOMContentLoaded', function() {
			// Enable push menu
			$('[data-widget="pushmenu"]').PushMenu('toggle');

			// Enable tooltips
			$('[data-toggle="tooltip"]').tooltip();

			// Enable popovers
			$('[data-toggle="popover"]').popover();

			// Enable sidebar tree view
			$("[data-widget='treeview']").each(function() {
				$(this).Treeview('init');
			});

			// Enable direct tree link
			$(".nav-link, .nav-item a").on('click', function() {
				var $this = $(this);
				if ($this.is('a') && $this.attr('href') === '#') {
					event.preventDefault();
					return false;
				}
				if (!$this.parent().hasClass('menu-open') && !$this.parent().find('> .nav-treeview').is(':visible')) {
					var menu = $this.next();
					if (menu.is('.nav-treeview')) {
						event.preventDefault();
						menu.slideToggle();
						$this.parent().toggleClass('menu-open');
					}
				}
			});

			// Add active class to current menu item
			var url = window.location.href;
			$("ul.nav-sidebar a").filter(function() {
				return this.href === url;
			}).addClass('active');

			// Add active class to parent menu item
			$("ul.nav-treeview a").filter(function() {
				return this.href === url;
			}).parentsUntil('.nav-sidebar > .nav-treeview').addClass('menu-open').prev('a').addClass('active');

			// Auto-expand active menu item
			$('ul.nav-treeview').each(function() {
				if ($(this).find('a.active').length) {
					$(this).parent().addClass('menu-open');
				}
			});
		});
	`
}

// RenderPage renders a complete page with the given content and dashboard renderer
func (t *AdminLTETemplate) RenderPage(content string, d shared.DashboardRenderer) (*hb.Tag, error) {
	// Create the head section
	head := hb.NewTag("head").
		Child(hb.NewTag("meta").Attr("charset", "utf-8")).
		Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1")).
		Child(hb.NewTag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "ie=edge")).
		Child(hb.NewTag("title").Text("Dashboard"))

	// Add favicon if available
	if d.GetFaviconURL() != "" {
		head.Child(hb.NewTag("link").Attr("rel", "icon").Attr("href", d.GetFaviconURL()))
	}

	// Add theme CSS
	cssLinks := t.GetCSSLinks(t.isDarkColorScheme(d))
	for _, link := range cssLinks {
		head.Child(link)
	}

	// Create the body section with AdminLTE classes
	bodyClasses := []string{
		"hold-transition",
		"sidebar-mini",
		"layout-fixed",
	}
	if t.isDarkColorScheme(d) {
		bodyClasses = append(bodyClasses, "dark-mode")
	}

	body := hb.NewTag("body").Class(strings.Join(bodyClasses, " "))

	// Create wrapper
	wrapper := hb.NewDiv().Class("wrapper")

	// Add header
	header := t.RenderHeader(d)
	if header != nil {
		wrapper.Child(header)
	}

	// Create main content area
	contentWrapper := hb.NewDiv().Class("content-wrapper")

	// Content header
	contentHeader := hb.NewDiv().Class("content-header")
	contentHeader.Child(hb.NewDiv().Class("container-fluid"))
	contentWrapper.Child(contentHeader)

	// Main content
	mainContent := hb.NewDiv().Class("content")
	container := hb.NewDiv().Class("container-fluid")
	container.Child(hb.NewHTML(content))
	mainContent.Child(container)
	contentWrapper.Child(mainContent)

	wrapper.Child(contentWrapper)

	// Add footer
	footer := t.RenderFooter(d)
	if footer != nil {
		wrapper.Child(footer)
	}

	// Add control sidebar (required for AdminLTE)
	controlSidebar := hb.NewDiv().Class("control-sidebar control-sidebar-dark")
	wrapper.Child(controlSidebar)

	// Add the wrapper to the body
	body.Child(wrapper)

	// Add JavaScript
	for _, script := range t.GetJSScripts() {
		body.Child(script)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		body.Child(hb.NewTag("script").Text(customJS))
	}

	// Create HTML document
	html := hb.NewTag("html").Attr("lang", "en").
		Child(head).
		Child(body)

	return hb.Wrap().
		Child(hb.NewHTML("<!DOCTYPE html>")).
		Child(html), nil
}

// isDarkColorScheme checks if the color scheme should be dark
func (t *AdminLTETemplate) isDarkColorScheme(d shared.DashboardRenderer) bool {
	return d.GetNavbarBackgroundColorMode() == "dark"
}
