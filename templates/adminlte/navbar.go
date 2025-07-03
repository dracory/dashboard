package adminlte

import (
	"fmt"

	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
)

func topNavigation(dashboard types.DashboardInterface) *hb.Tag {
	hasLogoRawHTML := dashboard.GetLogoRawHtml() != ""
	hasLogoImage := dashboard.GetLogoImageURL() != ""
	hasNavbarBackgroundColor := dashboard.GetNavbarBackgroundColor() != ""
	hasNavbarTextColor := dashboard.GetNavbarTextColor() != ""

	// Navbar with proper text color classes
	navbar := hb.Nav().Class("main-header navbar navbar-expand navbar-white navbar-light")

	// Left navbar container
	leftContainer := hb.Div().Class("container-fluid")
	navbar.Child(leftContainer)

	// Left navbar links
	leftNavbar := hb.Ul().Class("navbar-nav")
	leftContainer.Child(leftNavbar)

	// Navbar right icons
	rightNavbar := hb.Ul().Class("navbar-nav ms-auto")
	leftContainer.Child(hb.Div().Class("d-flex").Child(rightNavbar))

	// Left navbar menu toggle
	leftNavbar.Child(
		hb.Li().Class("nav-item").Child(
			hb.A().Class("nav-link").Attr("data-lte-toggle", "sidebar").Attr("href", "#").Child(
				hb.I().Class("fas fa-bars"),
			),
		),
	)

	// Logo
	logo := hb.Div()
	switch {
	case hasLogoRawHTML:
		logo.Child(hb.Raw(dashboard.GetLogoRawHtml()))
	case hasLogoImage:
		logo.Child(hb.Img(dashboard.GetLogoImageURL()).Style("max-height:35px;height:35px;width:auto;"))
	default:
		logo.Child(hb.Div().Text(dashboard.GetTitle()))
	}

	leftNavbar.Child(
		hb.Li().Class("nav-item d-none d-sm-inline-block").Child(
			hb.A().Class("nav-link").Href("/").Child(logo),
		),
	)

	// User menu
	if user := dashboard.GetUser(); user != nil {
		navbarTextColor := dashboard.GetNavbarTextColor()
		userMenu := navbarUserMenu(navbarTextColor, *user, dashboard.GetMenuUserItems())
		rightNavbar.Child(userMenu)
	}

	// Theme switcher
	navbarTextColor := dashboard.GetNavbarTextColor()
	themeSwitcher := navbarThemeSwitcher(navbarTextColor, "default", "/theme")
	rightNavbar.Child(themeSwitcher)

	// Add navbar background color if set
	if hasNavbarBackgroundColor {
		navbar.Style(fmt.Sprintf("background-color: %s !important;", dashboard.GetNavbarBackgroundColor()))
	}

	// Add navbar text color if set
	if hasNavbarTextColor {
		navbar.Style(fmt.Sprintf("color: %s !important;", dashboard.GetNavbarTextColor()))
	}

	return navbar
}
