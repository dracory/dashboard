package tabler

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

// topNavigation generates the top navigation bar with two headers
func topNavigation(dashboard types.DashboardInterface) string {
	header1 := navbarHeader1(dashboard)
	header2 := navbarHeader2(dashboard)

	// Wrap both headers in a parent container
	return hb.Wrap(header1, header2).ToHTML()
}

// navbarHeader1 creates the first header with logo and user menu
func navbarHeader1(dashboard types.DashboardInterface) *hb.Tag {
	dropdownTheme := themedropdown(dashboard)

	rightNav := navbarRightContainer()
	if dropdownTheme != nil {
		dropdownTheme.AddClass("nav-link px-0")
		rightNav = rightNav.Child(hb.Div().
			Class("nav-item dropdown d-none d-md-flex me-3").
			Child(dropdownTheme))
	}

	rightNav = rightNav.Child(lo.TernaryF(dashboard.GetUser() != nil,
		func() *hb.Tag { return navbarUserMenu(dashboard, dashboard.GetUser()) },
		func() *hb.Tag { return navbarLoginButton() }))

	headerClass := "navbar-expand-md"
	if dashboard.GetNavbarBackgroundColorMode() == "dark" {
		headerClass += " navbar-dark"
	}

	return navbarHeader().
		Class(headerClass).
		Child(navbarContainer().
			Child(navbarButtonToggler()).
			Child(navbarBrand(dashboard)).
			Child(hb.Div().
				Class("d-flex ms-auto").
				Child(rightNav)))
}

// navbarHeader2 creates the second navbar with main navigation menu
func navbarHeader2(dashboard types.DashboardInterface) *hb.Tag {
	nav := buildNavigation(dashboard.GetMenuMainItems())

	headerClass := "navbar navbar-expand-md"
	if dashboard.GetNavbarBackgroundColorMode() == "dark" {
		headerClass += " navbar-dark"
	}

	container := hb.Div().Class("container-xl")
	navMenu := navMenuContainer()
	navContainer := hb.Div().Class("navbar-nav flex-row").Child(nav)

	return hb.Div().
		Class(headerClass).
		Child(container.Child(navMenu.Child(navContainer)))
}

// navMenuContainer creates the main navigation menu container
func navMenuContainer() *hb.Tag {
	return hb.Div().
		Class("collapse navbar-collapse").
		Attr("id", "navbar-menu")
}

// navList creates the main navigation list
func navList() *hb.Tag {
	return hb.Ul().Class("navbar-nav")
}

// navDivider creates a divider for the navigation menu
func navDivider() *hb.Tag {
	return hb.Li().
		Class("nav-item").
		Child(hb.Div().Class("dropdown-divider"))
}

// navLink creates a navigation link
func navLink(item types.MenuItem) *hb.Tag {
	link := hb.A().
		Class("nav-link").
		Href(lo.Ternary(item.URL != "", item.URL, "#")).
		Attr("target", item.Target)

	// Add active class if the item is active
	if item.IsActive {
		link.AddClass("active")
	}

	// Add icon if present
	if item.Icon != "" {
		icon := hb.Span().Class("nav-link-icon")
		icon.Child(hb.Raw(item.Icon))
		link.Child(icon)
	}

	// Add title
	title := hb.Span().Class("nav-link-title").Child(hb.Text(item.Title))
	return link.Child(title)
}

// navDropdown creates a dropdown menu for navigation items with children
func navDropdown(item types.MenuItem) *hb.Tag {
	li := hb.Li().Class("nav-item dropdown")

	// Add active class to dropdown parent if any child is active
	isChildActive := false
	for _, child := range item.Children {
		if child.IsActive {
			isChildActive = true
			break
		}
	}

	link := navLink(item).
		AddClass("dropdown-toggle").
		Attr("data-bs-toggle", "dropdown").
		Attr("role", "button").
		Attr("aria-expanded", "false")

	if isChildActive {
		link.AddClass("active")
	}

	dropdownMenu := hb.Div().Class("dropdown-menu")
	for _, child := range item.Children {
		if child.URL == "" && child.Title == "" {
			dropdownMenu.Child(navDivider())
			continue
		}

		link := navLink(child).AddClass("dropdown-item")
		if child.IsActive {
			link.AddClass("active")
		}
		dropdownMenu.Child(link)
	}

	return li.Child(link).Child(dropdownMenu)
}

// navItem creates a single navigation item
func navItem(item types.MenuItem) *hb.Tag {
	if item.URL == "" && item.Title == "" {
		return navDivider()
	}

	li := hb.Li().Class("nav-item")

	// Add active class to parent li if the item is active
	if item.IsActive {
		li.AddClass("active")
	}

	if len(item.Children) > 0 {
		return navDropdown(item)
	}

	return li.Child(navLink(item))
}

// buildNavigation builds the complete navigation menu
func buildNavigation(menuItems []types.MenuItem) *hb.Tag {
	nav := navList()
	for _, item := range menuItems {
		nav.Child(navItem(item))
	}
	return nav
}

// navbarHeader creates the main navbar container
func navbarHeader() *hb.Tag {
	return hb.Header().
		Class("navbar navbar-expand-md d-print-none")
}

// navbarContainer creates the main container for navbar content
func navbarContainer() *hb.Tag {
	return hb.Div().
		Class("container-xl")
}

// navbarBrand creates the brand/logo section
func navbarBrand(dashboard types.DashboardInterface) *hb.Tag {
	brand := hb.Div().
		Class("navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3")
	brandLink := hb.A().Href(".").Attr("aria-label", "Tabler")

	// Set logo or title
	switch {
	case dashboard.GetLogoImageURL() != "":
		brandLink.Child(hb.Img(dashboard.GetLogoImageURL()).Class("navbar-brand-image"))
	case dashboard.GetLogoRawHtml() != "":
		brandLink.Child(hb.Raw(dashboard.GetLogoRawHtml()))
	default:
		brandLink.Child(hb.Span().Text(dashboard.GetTitle()))
	}

	// Set redirect URL
	brandLink.Attr("href", lo.Ternary(dashboard.GetLogoRedirectURL() != "", dashboard.GetLogoRedirectURL(), "#"))
	return brand.Child(brandLink)
}

// navbarRightContainer creates the right-aligned navigation container
func navbarRightContainer() *hb.Tag {
	return hb.NewDiv().Class("navbar-nav flex-row order-md-last")
}

// navbarButtonToggler creates the mobile menu toggle button
func navbarButtonToggler() *hb.Tag {
	toggleBtn := hb.Button().
		Class("navbar-toggler").
		Type("button").
		Data("bs-toggle", "collapse").
		Data("bs-target", "#navbar-menu").
		Aria("controls", "navbar-menu").
		Aria("expanded", "false").
		Aria("label", "Toggle navigation").
		Child(hb.Span().
			Class("navbar-toggler-icon"))
	return toggleBtn
}

// navbarUserMenu creates the user menu dropdown
func navbarUserMenu(dashboard types.DashboardInterface, user *types.User) *hb.Tag {
	userMenu := hb.Div().Class("nav-item dropdown")
	userLink := hb.A().
		Class("nav-link d-flex lh-1 text-reset p-0").
		Attr("data-bs-toggle", "dropdown")

	// User avatar with initials
	initials := ""
	if len(user.FirstName) > 0 {
		initials = string(user.FirstName[0])
	}

	avatar := hb.Span().
		Class("avatar avatar-sm").
		Child(hb.Span().
			Class("avatar-initials").
			Text(initials))

	// User info
	userInfo := hb.Div().Class("d-none d-xl-block ps-2")
	userInfo.Child(hb.Div().Text(user.FirstName + " " + user.LastName))

	if user.Email != "" {
		userInfo.Child(hb.Div().
			Class("mt-1 small text-muted").
			Text(user.Email))
	}

	userLink.Children([]hb.TagInterface{avatar, userInfo})

	// Dropdown menu
	dropdownMenu := hb.Div().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")
	for _, item := range dashboard.GetMenuUserItems() {
		if item.URL == "" && item.Title == "" {
			dropdownMenu.Child(hb.Div().Class("dropdown-divider"))
			continue
		}

		link := hb.A().
			Class("dropdown-item").
			Href(item.URL).
			Attr("target", item.Target)

		if item.Icon != "" {
			link.Child(hb.Raw(item.Icon)).AddClass("me-2")
		}

		link.Child(hb.Text(item.Title))
		dropdownMenu.Child(link)
	}

	return userMenu.Children([]hb.TagInterface{userLink, dropdownMenu})
}

// navbarLoginButton creates the sign-in button
func navbarLoginButton() *hb.Tag {
	return hb.A().
		Class("btn btn-primary").
		Href("/login").
		Child(hb.Text("Sign in"))
}

// themetoggle generates the theme toggle button
func themetoggle(dashboard types.DashboardInterface) *hb.Tag {
	button := hb.Button().Class("btn btn-link px-2")
	button.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" class="icon" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 12m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0" /><path d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" /></svg>`))
	return button
}

// themedropdown generates the theme dropdown menu
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
		link := hb.NewA().Class("dropdown-item").Href(item.URL)
		if item.Icon != "" {
			link.Child(hb.Raw(item.Icon)).AddClass("me-2")
		}
		link.Child(hb.Text(item.Title))
		menu.Child(link)
	}

	// Add quick access items if any
	if quickAccessItems := dashboard.GetMenuQuickAccessItems(); len(quickAccessItems) > 0 {
		menu.Child(hb.NewDiv().Class("dropdown-divider"))
		menu.Child(hb.NewDiv().Class("dropdown-header").HTML("Quick Access"))
		for _, item := range quickAccessItems {
			link := hb.NewA().Class("dropdown-item").Href(item.URL)
			if item.Icon != "" {
				link.Child(hb.Raw(item.Icon)).AddClass("me-2")
			}
			link.Child(hb.Text(item.Title))
			menu.Child(link)
		}
	}

	dropdown.Child(menu)

	return dropdown
}
