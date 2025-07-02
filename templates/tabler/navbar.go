package tabler

import (
	"strings"

	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

// navbarHeader creates the main navbar container
func navbarHeader() *hb.Tag {
	return hb.NewHeader().Class("navbar navbar-expand-md d-print-none")
}

// navbarContainer creates the main container for navbar content
func navbarContainer() *hb.Tag {
	return hb.NewDiv().Class("container-xl")
}

// navbarBrand creates the brand/logo section
func navbarBrand(dashboard types.DashboardInterface) *hb.Tag {
	brand := hb.NewDiv().Class("navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3")
	brandLink := hb.NewA().Href(".").Attr("aria-label", "Tabler")

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

// topNavigation generates the top navigation bar
func topNavigation(dashboard types.DashboardInterface) string {
	header := navbarHeader()
	container := navbarContainer()
	rightNav := navbarRightContainer()

	// Navigation menu variables
	var (
		navMenu *hb.Tag
		nav     *hb.Tag
	)

	// Initialize main navigation
	navMenu = hb.Div().Class("collapse navbar-collapse").Attr("id", "navbar-menu")
	nav = hb.Ul().Class("navbar-nav")

	// Add main menu items
	for _, item := range dashboard.GetMenuMainItems() {
		if item.URL == "" && item.Title == "" {
			// Divider
			nav.Child(hb.Li().Class("nav-item").
				Child(hb.Div().Class("dropdown-divider")))
			continue
		}

		li := hb.Li().Class("nav-item")
		hasChildren := len(item.Children) > 0

		// Create link
		link := hb.NewA().Class("nav-link")
		if hasChildren {
			li.AddClass("dropdown")
			link.AddClass("dropdown-toggle")
			link.Attr("data-bs-toggle", "dropdown")
			link.Attr("role", "button")
			link.Attr("aria-expanded", "false")
		}

		link.Attr("href", lo.Ternary(item.URL != "", item.URL, "#"))
		if item.Target != "" {
			link.Attr("target", item.Target)
		}

		if item.Icon != "" {
			link.Child(hb.Raw(item.Icon)).AddClass("me-2")
		}

		link.Child(hb.Text(item.Title))

		// Add dropdown menu if there are children
		if hasChildren {
			dropdownMenu := hb.NewDiv().Class("dropdown-menu")
			for _, child := range item.Children {
				if child.URL == "" && child.Title == "" {
					dropdownMenu.Child(hb.NewDiv().Class("dropdown-divider"))
					continue
				}

				childLink := hb.NewA().Class("dropdown-item").Href(child.URL)
				if child.Target != "" {
					childLink.Attr("target", child.Target)
				}

				if child.Icon != "" {
					childLink.Child(hb.Raw(child.Icon)).AddClass("me-2")
				}

				childLink.Child(hb.Text(child.Title))
				dropdownMenu.Child(childLink)
			}
			li.Child(link).Child(dropdownMenu)
		} else {
			li.Child(link)
		}

		nav.Child(li)
	}

	navMenu.Child(hb.NewDiv().Class("d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center").
		Child(nav))

	// Add theme toggles to right nav
	if themeToggle := themetoggle(dashboard); themeToggle != nil {
		themeToggle.AddClass("nav-link px-0")
		rightNav.Child(hb.NewDiv().Class("nav-item").Child(themeToggle))
	}

	if themeDropdown := themedropdown(dashboard); themeDropdown != nil {
		themeDropdown.AddClass("nav-link px-0")
		rightNav.Child(hb.NewDiv().Class("nav-item dropdown d-none d-md-flex me-3").Child(themeDropdown))
	}

	// Add user menu to right nav
	if user := dashboard.GetUser(); user != nil {
		userMenu := hb.NewDiv().Class("nav-item dropdown")
		userLink := hb.NewA().Class("nav-link d-flex lh-1 text-reset p-0")
		userLink.Attr("data-bs-toggle", "dropdown")

		// User avatar with initials
		avatar := hb.NewSpan().Class("avatar avatar-sm")
		if user.AvatarURL != "" {
			avatar.Child(hb.Img(user.AvatarURL).Class("avatar"))
		} else {
			initials := ""
			if len(user.FirstName) > 0 {
				initials = string(strings.ToUpper(user.FirstName)[0])
			}
			avatar.Child(hb.NewSpan().Class("avatar-initials").Text(initials))
		}

		// User name
		userInfo := hb.NewDiv().Class("d-none d-xl-block ps-2")
		userInfo.Child(hb.NewDiv().Text(user.FirstName + " " + user.LastName))
		if user.Email != "" {
			userInfo.Child(hb.NewDiv().Class("mt-1 small text-muted").Text(user.Email))
		}

		userLink.Child(avatar).Child(userInfo)

		// Dropdown menu
		dropdownMenu := hb.NewDiv().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")
		for _, item := range dashboard.GetMenuUserItems() {
			if item.URL == "" && item.Title == "" {
				dropdownMenu.Child(hb.NewDiv().Class("dropdown-divider"))
				continue
			}
			itemLink := hb.NewA().Class("dropdown-item").Href(item.URL)
			if item.Target != "" {
				itemLink.Attr("target", item.Target)
			}
			if item.Icon != "" {
				itemLink.Child(hb.Raw(item.Icon)).AddClass("me-2")
			}
			itemLink.Child(hb.Text(item.Title))
			dropdownMenu.Child(itemLink)
		}

		userMenu.Child(userLink).Child(dropdownMenu)
		rightNav.Child(userMenu)
	} else {
		// Login button if no user
		loginBtn := hb.NewA().Class("btn btn-primary")
		loginBtn.Href("/login").Child(hb.Text("Sign in"))
		rightNav.Child(loginBtn)
	}

	// Initialize main navigation
	navMenu = hb.Div().Class("collapse navbar-collapse").Attr("id", "navbar-menu")
	nav = hb.Ul().Class("navbar-nav")

	// Add main menu items
	for _, item := range dashboard.GetMenuMainItems() {
		if item.URL == "" && item.Title == "" {
			// Divider
			nav.Child(hb.Li().Class("nav-item").
				Child(hb.Div().Class("dropdown-divider")))
			continue
		}

		li := hb.Li().Class("nav-item")
		hasChildren := len(item.Children) > 0

		// Create link
		link := hb.NewA().Class("nav-link")
		if hasChildren {
			li.AddClass("dropdown")
			link.AddClass("dropdown-toggle")
			link.Attr("data-bs-toggle", "dropdown")
			link.Attr("role", "button")
			link.Attr("aria-expanded", "false")
		}

		link.Attr("href", lo.Ternary(item.URL != "", item.URL, "#"))
		if item.Target != "" {
			link.Attr("target", item.Target)
		}

		// Add icon if exists
		if item.Icon != "" {
			icon := hb.NewSpan().Class("nav-link-icon d-none d-md-inline-block me-1")
			icon.Child(hb.Raw(item.Icon))
			link.Child(icon)
		}

		// Add title
		link.Child(hb.NewSpan().Class("nav-link-title").Text(item.Title))

		li.Child(link)

		// Add dropdown menu if has children
		if hasChildren {
			dropdown := hb.NewDiv().Class("dropdown-menu dropdown-menu-arrow")
			for _, child := range item.Children {
				if child.URL == "" && child.Title == "" {
					dropdown.Child(hb.NewDiv().Class("dropdown-divider"))
					continue
				}
				childLink := hb.NewA().Class("dropdown-item").Href(lo.Ternary(child.URL != "", child.URL, "#"))
				if child.Target != "" {
					childLink.Attr("target", child.Target)
				}
				if child.Icon != "" {
					childLink.Child(hb.Raw(child.Icon)).AddClass("me-2")
				}
				childLink.Child(hb.Text(child.Title))
				dropdown.Child(childLink)
			}
			li.Child(dropdown)
		}

		nav.Child(li)
	}

	navMenu.Child(hb.NewDiv().Class("d-flex flex-column flex-md-row flex-fill align-items-stretch align-items-md-center").
		Child(nav))

	container.Children(
		navbarButtonToggler(),
		navbarBrand(dashboard),
		navMenu,
		rightNav,
	)

	header.Child(container)

	return header.ToHTML()
}

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

// themetoggle generates the theme toggle button
func themetoggle(dashboard types.DashboardInterface) *hb.Tag {
	button := hb.NewButton().Class("btn btn-link px-2")
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
