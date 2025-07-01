package tabler

import (
	"fmt"
	"strings"

	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

// topNavigation generates the top navigation bar
func topNavigation(dashboard types.DashboardInterface) string {
	// Create the main navbar container
	header := hb.NewHeader().Class("navbar navbar-expand-md navbar-light d-print-none")
	header.Attr("data-bs-theme", "light")

	// Container for navbar content
	container := hb.NewDiv().Class("container-xl")

	// Toggle sidebar button (hamburger)
	toggleBtn := hb.NewButton().Class("navbar-toggler")
	toggleBtn.Attr("type", "button")
	toggleBtn.Attr("data-bs-toggle", "offcanvas")
	toggleBtn.Attr("data-bs-target", "#sidebar")
	toggleBtn.Attr("aria-controls", "sidebar")
	toggleBtn.Attr("aria-label", "Toggle navigation")
	toggleBtn.Child(hb.Span().Class("navbar-toggler-icon"))

	// Logo/Brand
	brand := hb.NewDiv().Class("navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3")
	brandLink := hb.NewA()
	if logoURL := dashboard.GetLogoImageURL(); logoURL != "" {
		brandLink.Child(hb.Img(logoURL).Class("navbar-brand-image"))
	} else if logoHTML := dashboard.GetLogoRawHtml(); logoHTML != "" {
		brandLink.Child(hb.Raw(logoHTML))
	} else {
		brandLink.Child(hb.Span().Text(dashboard.GetTitle()))
	}
	if redirectURL := dashboard.GetLogoRedirectURL(); redirectURL != "" {
		brandLink.Attr("href", redirectURL)
	} else {
		brandLink.Attr("href", "#")
	}
	brand.Child(brandLink)

	// Navbar collapse wrapper
	collapse := hb.NewDiv().Class("navbar-nav flex-row order-md-last")

	// Theme toggle
	themeToggle := hb.NewDiv().Class("nav-item dropdown d-none d-md-flex me-3")
	themeLink := hb.NewA().Class("nav-link px-0 hide-theme-dark")
	themeLink.Attr("data-bs-toggle", "tooltip")
	themeLink.Attr("data-bs-placement", "bottom")
	themeLink.Attr("title", "Enable dark mode")
	themeLink.Attr("data-bs-theme", "light")
	themeLink.Attr("data-bs-theme-value", "dark")
	themeLink.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z" /></svg>`))
	themeToggle.Child(themeLink)

	themeToggleDark := hb.NewDiv().Class("nav-item dropdown d-none d-md-flex me-3")
	themeLinkDark := hb.NewA().Class("nav-link px-0 hide-theme-light")
	themeLinkDark.Attr("data-bs-toggle", "tooltip")
	themeLinkDark.Attr("data-bs-placement", "bottom")
	themeLinkDark.Attr("title", "Enable light mode")
	themeLinkDark.Attr("data-bs-theme", "dark")
	themeLinkDark.Attr("data-bs-theme-value", "light")
	themeLinkDark.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-4 0a4 4 0 1 0 8 0a4 4 0 1 0 -8 0" /><path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" /></svg>`))
	themeToggleDark.Child(themeLinkDark)

	collapse.Child(themedropdown(dashboard))
	collapse.Child(themetoggle(dashboard))

	// User menu
	if user := dashboard.GetUser(); user != nil {
		userMenu := hb.NewDiv().Class("nav-item dropdown")
		userLink := hb.NewA().Class("nav-link d-flex lh-1 text-reset p-0")
		userLink.Attr("data-bs-toggle", "dropdown")
		userLink.Attr("aria-label", "Open user menu")

		// User avatar with initials
		initials := strings.ToUpper(string(user.FirstName[0])) + strings.ToUpper(string(user.LastName[0]))
		avatar := hb.NewSpan().Class("avatar avatar-sm rounded-circle")
		avatar.Child(hb.NewSpan().Class("avatar-initials").Text(initials))

		// User info
		userInfo := hb.NewDiv().Class("d-none d-xl-block ps-2")
		userInfo.Child(hb.NewDiv().Text(fmt.Sprintf("%s %s", user.FirstName, user.LastName)))
		userInfo.Child(hb.NewDiv().Class("mt-1 small text-muted").Text("Administrator"))

		userLink.Child(avatar)
		userLink.Child(userInfo)

		// Dropdown menu
		dropdownMenu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")

		// Add user menu items
		for _, item := range dashboard.GetMenuUserItems() {
			if item.URL == "" && item.Title == "" {
				dropdownMenu.Child(hb.NewDiv().Class("dropdown-divider"))
				continue
			}

			link := hb.NewA().Class("dropdown-item")
			if item.URL != "" {
				link.Attr("href", item.URL)
			} else {
				link.Attr("href", "#")
			}

			if item.Target != "" {
				link.Attr("target", item.Target)
			}

			if item.Icon != "" {
				link.Child(hb.Raw(item.Icon)).AddClass("me-2")
			}

			link.Child(hb.Text(item.Title))
			dropdownMenu.Child(link)
		}

		userMenu.Child(userLink)
		userMenu.Child(dropdownMenu)
		collapse.Child(userMenu)
	}

	// Add all elements to the container
	container.Child(toggleBtn)
	container.Child(brand)

	// Navbar navigation
	nav := hb.NewDiv().Class("navbar-nav")

	// Add main menu items
	for _, item := range dashboard.GetMenuMainItems() {
		if item.URL == "" && item.Title == "" {
			// Divider
			nav.Child(hb.NewDiv().Class("nav-item").Child(hb.NewDiv().Class("nav-link disabled").Text("|")))
			continue
		}

		link := hb.NewA().Class("nav-link")
		if item.URL != "" {
			link.Attr("href", item.URL)
		} else {
			link.Attr("href", "#")
		}

		if item.Target != "" {
			link.Attr("target", item.Target)
		}

		if item.Icon != "" {
			link.Child(hb.Raw(item.Icon)).AddClass("me-1")
		}

		link.Child(hb.Text(item.Title))

		nav.Child(hb.NewDiv().Class("nav-item").Child(link))
	}

	container.Child(nav)

	// Add the right side to the container
	container.Child(collapse)

	// Add the container to the header
	header.Child(container)

	return header.ToHTML()
}

// themetoggle generates the theme toggle dropdown
func themetoggle(dashboard types.DashboardInterface) *hb.Tag {
	dropdown := hb.NewDiv().Class("nav-item dropdown d-none d-md-flex me-3")
	link := hb.NewA().Class("nav-link px-0")
	link.Attr("data-bs-toggle", "dropdown")
	link.Attr("tabindex", "-1")
	link.Attr("aria-label", "Show theme menu")

	// Theme icon
	link.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" /></svg>`))

	// Dropdown menu
	menu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end")

	// Theme options
	themes := map[string]string{
		"light":  "Light",
		"dark":   "Dark",
		"system": "System",
	}

	for value, label := range themes {
		item := hb.NewButton().Class("dropdown-item")
		item.Attr("data-bs-theme", value)
		item.Child(hb.Text(label))
		menu.Child(item)
	}

	dropdown.Child(link)
	dropdown.Child(menu)

	return dropdown
}

// themedropdown generates the theme dropdown
func themedropdown(dashboard types.DashboardInterface) *hb.Tag {
	dropdown := hb.NewDiv().Class("nav-item dropdown d-none d-md-flex me-3")
	link := hb.NewA().Class("nav-link px-0")
	link.Attr("data-bs-toggle", "dropdown")
	link.Attr("tabindex", "-1")
	link.Attr("aria-label", "Show theme menu")

	// Theme icon
	link.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" /></svg>`))

	// Dropdown menu
	menu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end")

	// Theme options
	themes := map[string]string{
		"light":  "Light",
		"dark":   "Dark",
		"system": "System",
	}

	for value, label := range themes {
		item := hb.NewButton().Class("dropdown-item")
		item.Attr("data-bs-theme", value)
		item.Child(hb.Text(label))
		menu.Child(item)
	}

	dropdown.Child(link)
	dropdown.Child(menu)

	return dropdown
}

// buildNavItem creates a navigation item
func buildNavItem(item types.MenuItem) *hb.Tag {
	hasChildren := len(item.Children) > 0

	// Create list item
	li := hb.NewLI().Class("nav-item")
	if hasChildren {
		li.Class("nav-item dropdown")
	}

	// Create link
	link := hb.NewHyperlink().
		Class("nav-link")

	if hasChildren {
		link.Class("dropdown-toggle")
		link.Attr("data-bs-toggle", "dropdown")
		link.Attr("role", "button")
		link.Attr("aria-expanded", "false")
	}

	if item.URL != "" {
		link.Href(item.URL)
	} else {
		link.Href("#")
	}

	// Add icon if exists
	if item.Icon != "" {
		icon := hb.NewSpan().Class("nav-link-icon d-md-none d-lg-inline-block me-1")
		icon.Child(hb.Raw(item.Icon))
		link.Child(icon)
	}

	// Add title
	title := hb.NewSpan().Class("nav-link-title").HTML(item.Title)
	link.Child(title)

	li.Child(link)

	// Add dropdown menu if has children
	if hasChildren {
		dropdown := hb.NewDiv().Class("dropdown-menu")
		for _, child := range item.Children {
			dropdown.Child(buildDropdownItem(child))
		}
		li.Child(dropdown)
	}

	return li
}

// buildDropdownItem creates a dropdown item
func buildDropdownItem(item types.MenuItem) *hb.Tag {

	hasChildren := len(item.Children) > 0

	link := hb.NewHyperlink().
		Class("dropdown-item")

	if hasChildren {
		link.Class("dropdown-item")
		link.Attr("data-bs-toggle", "dropdown")
		link.Attr("role", "button")
		link.Attr("aria-expanded", "false")
	}

	if item.URL != "" {
		link.Href(item.URL)
	} else {
		link.Href("#")
	}

	// Add icon if exists
	if item.Icon != "" {
		icon := hb.NewSpan().Class("me-2").HTML(item.Icon)
		link.Child(icon)
	}

	// Add title
	link.Child(hb.NewRaw(item.Title))



	if hasChildren {
		dropdown := hb.NewDiv().Class("dropdown-menu")
		for _, child := range item.Children {
			dropdown.Child(buildDropdownItem(child))
		}
		return hb.NewDiv().Class("dropend").Child(link).Child(dropdown)
	}

	return link
}

// navbarDropdownUser creates the user dropdown menu
func navbarDropdownUser(dashboard types.DashboardInterface, user types.User) *hb.Tag {
	dropdown := hb.NewDiv().Class("dropdown")
	toggle := hb.NewA().
		Class("nav-link" + lo.Ternary(dashboard.GetNavbarTextColor() != "", " text-white", "")).
		Href("#")

	// User avatar with initials
	avatar := hb.NewDiv().Class("avatar avatar-sm")
	initials := "U"
	if len(user.FirstName) > 0 {
		initials = string(user.FirstName[0])
	}
	avatar.Child(hb.NewSpan().Class("avatar-initials").HTML(initials))
	toggle.Child(avatar)

	// User name
	userName := user.FirstName
	if user.LastName != "" {
		userName += " " + user.LastName
	}

	if userName != "" {
		toggle.Child(hb.NewSpan().Class("d-none d-md-inline ms-2").HTML(userName))
	}

	dropdown.Child(toggle)

	// Dropdown menu
	menu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end")

	// Add user menu items
	for _, item := range dashboard.GetMenuUserItems() {
		menu.Child(buildDropdownItem(item))
	}

	// Add quick access items if any
	if quickAccessItems := dashboard.GetMenuQuickAccessItems(); len(quickAccessItems) > 0 {
		menu.Child(hb.NewDiv().Class("dropdown-divider"))
		menu.Child(hb.NewDiv().Class("dropdown-header").HTML("Quick Access"))
		for _, item := range quickAccessItems {
			menu.Child(buildDropdownItem(item))
		}
	}

	dropdown.Child(menu)

	return dropdown
}
