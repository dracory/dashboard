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
	link := hb.A().Href(url).Class("nav-link")

	// Add icon if present
	iconClass := "far fa-circle"
	if icon != "" {
		iconClass = icon
	}

	iconEl := hb.I().Class("nav-icon " + iconClass)
	link.Child(iconEl)

	// Add title and caret for dropdown if has children
	titleEl := hb.P()
	titleEl.Child(hb.Span().HTML(title))

	if hasChildren {
		caret := hb.I().Class("right fas fa-angle-left")
		titleEl.Child(caret)
	}

	link.Child(titleEl)

	// Add link to list item
	li.Child(link)

	// Add child items if they exist
	if hasChildren {
		ul := hb.Ul().Class("nav nav-treeview")
		for _, child := range children {
			childItem := buildMenuItem(child, index+1)
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
	brandLink := hb.A().Href("#").Class("brand-link")
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

	// Sidebar search
	searchDiv := hb.Div().Class("form-inline mt-2")
	inputGroup := hb.Div().Class("input-group").Data("widget", "sidebar-search")
	input := hb.Input().Class("form-control form-control-sidebar").Type("search").Placeholder("Search...").Attr("aria-label", "Search")
	inputGroup.Child(input)
	inputGroupAppend := hb.Div().Class("input-group-append")
	button := hb.Button().Class("btn btn-sidebar").Type("submit")
	iconHTML := "<i class=\"fas fa-search fa-fw\"></i>"
	button.Child(hb.Raw(iconHTML))
	inputGroupAppend.Child(button)
	inputGroup.Child(inputGroupAppend)
	searchDiv.Child(inputGroup)
	sidebarInner.Child(searchDiv)

	nav := hb.Nav().Class("mt-2")
	ul := hb.Ul().Class("nav nav-pills nav-sidebar flex-column").Data("widget", "treeview").Role("menu").Data("accordion", "false")
	ul.Child(BuildSidebarMenu(dashboard))
	nav.Child(ul)

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
	link := hb.A()
	if item.URL != "" {
		link.Href(item.URL)
	} else {
		link.Href("#")
	}
	link.Class("nav-link")

	// Add icon
	iconEl := hb.I().Class("nav-icon " + icon)
	link.Child(iconEl)

	// Add title and caret
	titleContainer := hb.P()
	titleContainer.Child(hb.Span().HTML(item.Title))

	if hasChildren {
		caret := hb.I().Class("right fas fa-angle-left")
		titleContainer.Child(caret)
	}

	link.Child(titleContainer)
	li.Child(link)

	// Add child items if they exist
	if hasChildren {
		childList := hb.Ul().Class("nav nav-treeview")
		for _, child := range item.Children {
			childItem := buildSidebarMenuItem(child, index+1)
			childList.Child(childItem)
		}
		li.Child(childList)
	}

	return li
}
