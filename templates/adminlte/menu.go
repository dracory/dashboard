package adminlte

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
)

// buildMenuItem creates a menu item for the dashboard menu
func buildMenuItem(menuItem types.MenuItem, index int) *hb.Tag {
	title := menuItem.Title
	if title == "" {
		title = "n/a"
	}

	url := menuItem.URL
	if url == "" {
		url = "#"
	}

	icon := menuItem.Icon
	children := menuItem.Children
	hasChildren := len(children) > 0

	// Create the main list item
	li := hb.Li().Class("nav-item")
	if hasChildren {
		li.Class("has-treeview")
	}

	// Create the link
	link := hb.Hyperlink().Class("nav-link")

	// Add icon if present
	if icon != "" {
		iconHTML := "<i class=\"" + icon + " nav-icon\"></i>"
		link.Child(hb.Raw(iconHTML))
	} else {
		link.Child(hb.Raw("<i class=\"far fa-circle nav-icon\"></i>"))
	}

	// Add title
	link.Child(hb.P().HTML(title))

	// Add caret for dropdown if has children
	if hasChildren {
		link.Child(hb.Raw("<i class=\"right fas fa-angle-left\"></i>"))
	}

	// Set URL
	link.Href(url)

	// Add the link to the list item
	li.Child(link)

	// Add children if present
	if hasChildren {
		ul := hb.Ul().Class("nav nav-treeview")
		for i, child := range children {
			childItem := buildMenuItem(child, i)
			ul.Child(childItem)
		}
		li.Child(ul)
	}

	return li
}

// dashboardMenuNavbar generates the HTML for the dashboard menu navbar
func dashboardMenuNavbar(dashboard types.DashboardInterface) string {
	menuItems := dashboard.GetMenuMainItems()
	menu := hb.Ul().Class("navbar-nav")

	for i, item := range menuItems {
		menuItem := buildMenuItem(item, i)
		menu.Child(menuItem)
	}

	return menu.ToHTML()
}

// menuOffcanvas generates the offcanvas menu HTML
func menuOffcanvas(dashboard types.DashboardInterface) *hb.Tag {
	// Main sidebar container
	sidebar := hb.Div().Class("main-sidebar sidebar-dark-primary elevation-4")

	// Brand logo
	brandLink := hb.Div().Class("brand-link text-center")
	if dashboard.GetLogoImageURL() != "" {
		brandLink.Child(hb.Img(dashboard.GetLogoImageURL()).Class("brand-image img-circle elevation-3"))
	}
	brandLink.Child(hb.Span().Class("brand-text font-weight-light").HTML(dashboard.GetTitle()))
	sidebar.Child(brandLink)

	// Sidebar
	sidebarInner := hb.Div().Class("sidebar")

	// User panel
	user := dashboard.GetUser()
	if user != nil && user.FirstName != "" {
		userPanel := hb.Div().Class("user-panel mt-3 pb-3 mb-3 d-flex")
		userPanel.Child(hb.Div().Class("image").Child(
			hb.Img("https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y").Class("img-circle elevation-2").Style("width: 2.1rem"),
		))
		userPanel.Child(hb.Div().Class("info").Child(
			hb.A().Href("#").Class("d-block").HTML(user.FirstName + " " + user.LastName),
		))
		sidebarInner.Child(userPanel)
	}

	// Sidebar menu
	nav := hb.Nav().Class("mt-2")
	nav.Child(hb.Ul().Class("nav nav-pills nav-sidebar flex-column").Data("widget", "treeview").Role("menu").Data("accordion", "false").Child(BuildSidebarMenu(dashboard)))

	sidebarInner.Child(nav)
	sidebar.Child(sidebarInner)

	return sidebar
}

// BuildSidebarMenu builds the sidebar menu structure
func BuildSidebarMenu(dashboard types.DashboardInterface) *hb.Tag {
	menuItems := dashboard.GetMenuMainItems()
	menu := hb.Ul().Class("nav nav-pills nav-sidebar flex-column").Data("widget", "treeview").Role("menu")

	for i, item := range menuItems {
		menuItem := buildSidebarMenuItem(item, i)
		if menuItem != nil {
			menu.Child(menuItem)
		}
	}

	return menu
}

// buildSidebarMenuItem builds a single sidebar menu item
func buildSidebarMenuItem(item types.MenuItem, index int) *hb.Tag {
	hasChildren := len(item.Children) > 0
	icon := item.Icon
	if icon == "" {
		icon = "far fa-circle"
	}

	// Create list item
	li := hb.Li().Class("nav-item")
	if hasChildren {
		li.Class("has-treeview")
	}

	// Create link
	link := hb.A().Href(item.URL).Class("nav-link")
	if hasChildren {
		link.Attr("onclick", "return false;")
	}

	// Add icon
	iconHTML := "<i class=\"nav-icon " + icon + "\"></i>"
	link.Child(hb.Raw(iconHTML))

	// Add title
	title := hb.P().HTML(item.Title)
	link.Child(title)

	// Add dropdown arrow for items with children
	if hasChildren {
		arrowHTML := "<i class=\"right fas fa-angle-left\"></i>"
		title.Child(hb.Raw(arrowHTML))
	}

	li.Child(link)

	// Add child items if they exist
	if hasChildren {
		childList := hb.Ul().Class("nav nav-treeview")
		for _, child := range item.Children {
			childList.Child(buildSidebarMenuItem(child, index+1))
		}
		li.Child(childList)
	}

	return li
}
