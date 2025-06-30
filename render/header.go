package render

// import (
// 	"github.com/dracory/dashboard/model"
// 	"github.com/gouniverse/hb"
// 	"github.com/samber/lo"
// )

// // RenderHeader generates the header HTML for the dashboard
// func RenderHeader(d model.DashboardRenderer) *hb.Tag {
// 	// Top navbar with logo and user menu
// 	topNavbar := renderTopNavbar(d)

// 	// Main menu navbar
// 	mainMenuNavbar := renderMainMenuNavbar(d)

// 	// Combine both navbars into a header
// 	headerWrap := hb.Wrap().
// 		Child(topNavbar).
// 		Child(mainMenuNavbar)

// 	return headerWrap
// }

// // renderTopNavbar generates the top navbar with logo and user menu
// func renderTopNavbar(d model.DashboardRenderer) *hb.Tag {
// 	// Logo
// 	logo := renderLogo(d)

// 	// User dropdown
// 	userDropdown := RenderUserDropdown(d)

// 	// Theme switcher
// 	themeSwitcher := RenderThemeSwitcher(d)

// 	// Quick access menu
// 	quickAccess := RenderQuickAccessMenu(d)
// 	if quickAccess == nil {
// 		quickAccess = hb.NewHTML("")
// 	}

// 	// Create navbar toggle button
// 	navbarToggler := hb.Button().Class("navbar-toggler").
// 		Attr("type", "button").
// 		Attr("data-bs-toggle", "collapse").
// 		Attr("data-bs-target", "#navbar-menu").
// 		Child(hb.Span().Class("navbar-toggler-icon"))

// 	// Create brand logo container
// 	brandContainer := hb.H1().
// 		Class("navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3").
// 		Child(logo)

// 	// Create user menu container
// 	userMenuContainer := hb.Div().
// 		Class("navbar-nav flex-row order-md-last").
// 		Child(quickAccess).
// 		Child(themeSwitcher).
// 		Child(hb.Div().Class("dropdown-divider")).
// 		Child(userDropdown)

// 	// Create container
// 	container := hb.Div().
// 		Class("container-xl").
// 		Child(navbarToggler).
// 		Child(brandContainer).
// 		Child(userMenuContainer)

// 	// Create header
// 	header := hb.Header().
// 		Class("navbar navbar-expand-md navbar-light d-print-none").
// 		Child(container)

// 	return header
// }

// // renderMainMenuNavbar generates the main menu navbar
// func renderMainMenuNavbar(d model.DashboardRenderer) *hb.Tag {
// 	// Main menu items
// 	menuItems := RenderMainMenuItems(d)

// 	return hb.Div().
// 		Class("navbar-expand-md").
// 		Child(
// 			hb.Div().
// 				Class("collapse navbar-collapse").
// 				ID("navbar-menu").
// 				Child(
// 					hb.Div().Class("navbar navbar-light").Child(
// 						hb.Div().Class("container-xl").Child(
// 							hb.Ul().Class("navbar-nav").Child(menuItems),
// 						),
// 					),
// 				),
// 		)
// }

// // renderLogo generates the logo HTML
// func renderLogo(d model.DashboardRenderer) *hb.Tag {
// 	if !lo.IsEmpty(d.GetLogoRawHtml()) {
// 		return hb.NewHTML(d.GetLogoRawHtml())
// 	}

// 	if !lo.IsEmpty(d.GetLogoImageURL()) {
// 		logoLink := d.GetLogoRedirectURL()
// 		if lo.IsEmpty(logoLink) {
// 			logoLink = "#"
// 		}

// 		imageTag := hb.NewTag("img")
// 		imageTag = imageTag.Attr("src", d.GetLogoImageURL())
// 		imageTag = imageTag.Attr("height", "36")
// 		imageTag = imageTag.Attr("alt", "Logo")

// 		logoAnchor := hb.A().Href(logoLink).Class("navbar-brand navbar-brand-autodark")
// 		logoAnchor = logoAnchor.Child(imageTag)

// 		return logoAnchor
// 	}

// 	return hb.NewHTML("")
// }

// // RenderMainMenuItems generates the main menu items HTML
// func RenderMainMenuItems(d model.DashboardRenderer) *hb.Tag {
// 	if len(d.GetMenuItems()) == 0 {
// 		return hb.NewHTML("")
// 	}

// 	menuItemsWrap := hb.Wrap()

// 	for _, menuItem := range d.GetMenuItems() {
// 		// Create icon element
// 		var icon *hb.Tag
// 		if !lo.IsEmpty(menuItem.Icon) {
// 			iconContainer := hb.Span().Class("nav-link-icon d-md-none d-lg-inline-block")
// 			iconElement := hb.I().Class(menuItem.Icon)
// 			iconContainer = iconContainer.Child(iconElement)
// 			icon = iconContainer
// 		} else {
// 			icon = hb.NewHTML("")
// 		}

// 		// Determine active class
// 		activeClass := ""
// 		if menuItem.Active {
// 			activeClass = " active"
// 		}

// 		// Create title element
// 		titleElement := hb.Span().Class("nav-link-title").Text(menuItem.Text)

// 		if len(menuItem.SubMenu) > 0 {
// 			// Dropdown menu
// 			subMenuItems := hb.Wrap()

// 			// Create submenu items
// 			for _, subItem := range menuItem.SubMenu {
// 				// Create submenu icon
// 				var subIcon *hb.Tag
// 				if !lo.IsEmpty(subItem.Icon) {
// 					subIconContainer := hb.Span().Class("nav-link-icon d-md-none d-lg-inline-block")
// 					subIconElement := hb.I().Class(subItem.Icon)
// 					subIconContainer = subIconContainer.Child(subIconElement)
// 					subIcon = subIconContainer
// 				} else {
// 					subIcon = hb.NewHTML("")
// 				}

// 				// Determine submenu active class
// 				subActiveClass := ""
// 				if subItem.Active {
// 					subActiveClass = " active"
// 				}

// 				// Create submenu link
// 				subLink := hb.A().Class("dropdown-item" + subActiveClass).Href(subItem.URL)
// 				subLink = subLink.Child(subIcon)
// 				subLink = subLink.Child(hb.Text(subItem.Text))

// 				// Add to submenu wrapper directly without wrapping in Li
// 				subMenuItems = subMenuItems.Child(subLink)
// 			}

// 			// Create dropdown toggle link
// 			dropdownToggle := hb.A().
// 				Class("nav-link dropdown-toggle").
// 				Href("#navbar-base").
// 				Attr("data-bs-toggle", "dropdown").
// 				Attr("data-bs-auto-close", "outside").
// 				Attr("role", "button").
// 				Attr("aria-expanded", "false").
// 				Child(icon).
// 				Child(titleElement)

// 			// Create dropdown menu column
// 			dropdownMenuColumn := hb.Div().Class("dropdown-menu-column")
// 			dropdownMenuColumn = dropdownMenuColumn.Child(subMenuItems)

// 			// Create dropdown menu columns container
// 			dropdownMenuColumns := hb.Div().Class("dropdown-menu-columns")
// 			dropdownMenuColumns = dropdownMenuColumns.Child(dropdownMenuColumn)

// 			// Create dropdown menu
// 			dropdownMenu := hb.Div().Class("dropdown-menu")
// 			dropdownMenu = dropdownMenu.Child(dropdownMenuColumns)

// 			// Create dropdown list item
// 			dropdownItem := hb.Li().Class("nav-item dropdown" + activeClass)
// 			dropdownItem = dropdownItem.Child(dropdownToggle)
// 			dropdownItem = dropdownItem.Child(dropdownMenu)

// 			// Add to menu wrapper
// 			menuItemsWrap = menuItemsWrap.Child(dropdownItem)
// 		} else {
// 			// Regular menu item

// 			// Create menu link
// 			menuLink := hb.A().Class("nav-link").Href(menuItem.URL)
// 			menuLink = menuLink.Child(icon)
// 			menuLink = menuLink.Child(titleElement)

// 			// Create menu item
// 			menuItem := hb.Li().Class("nav-item" + activeClass)
// 			menuItem = menuItem.Child(menuLink)

// 			// Add to menu wrapper
// 			menuItemsWrap = menuItemsWrap.Child(menuItem)
// 		}
// 	}

// 	return menuItemsWrap
// }
