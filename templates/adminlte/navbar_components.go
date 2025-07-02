package adminlte

import (
	"strconv"

	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
)

// navbarMessagesDropdown creates a messages dropdown menu
func navbarMessagesDropdown(navbarTextColor string, messages []types.MenuItem) *hb.Tag {
	// Only show first 4 messages in the dropdown
	visibleMessages := messages
	if len(visibleMessages) > 4 {
		visibleMessages = visibleMessages[:4]
	}

	dropdown := hb.Li().Class("nav-item dropdown")
	a := hb.A().Href("#").Class("nav-link").Data("toggle", "dropdown")

	// Message icon with badge
	icon := hb.I().Class("far fa-comments")
	if navbarTextColor != "" {
		icon.Style("color: " + navbarTextColor + " !important")
	}
	a.Child(icon)
	a.Child(hb.Span().Class("badge badge-danger navbar-badge").HTML(strconv.Itoa(len(messages))))

	// Dropdown menu
	menu := hb.Div().Class("dropdown-menu dropdown-menu-lg dropdown-menu-right")

	// Header
	header := hb.Span().Class("dropdown-item dropdown-header").HTML(strconv.Itoa(len(messages)) + " Messages")
	menu.Child(header)

	// Messages
	for _, msg := range visibleMessages {
		menu.Child(createMessageItem(msg, navbarTextColor))
	}

	// Footer
	footer := hb.A().Href("#").Class("dropdown-item dropdown-footer").HTML("See All Messages")
	menu.Child(footer)

	dropdown.Child(a)
	dropdown.Child(menu)

	return dropdown
}

// createMessageItem creates a single message item for the messages dropdown
func createMessageItem(msg types.MenuItem, textColor string) *hb.Tag {
	item := hb.Div().Class("dropdown-item")

	// Message header
	header := hb.Div().Class("media")

	// User image
	userImg := hb.Img("https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y")
	userImg.Class("img-size-50 mr-3 img-circle")
	header.Child(userImg)

	mediaBody := hb.Div().Class("media-body")
	mediaBody.Child(hb.H3().Class("dropdown-item-title").HTML(msg.Title))

	// Only add icon if it's not empty
	if msg.Icon != "" {
		mediaBody.Child(hb.P().Class("text-sm text-muted").Child(hb.I().Class(msg.Icon + " mr-2")).HTML(msg.URL))
	} else {
		mediaBody.Child(hb.P().Class("text-sm text-muted").HTML(msg.URL))
	}

	header.Child(mediaBody)
	item.Child(header)

	// Make the entire item clickable if URL is provided
	if msg.URL != "" {
		link := hb.A().Href(msg.URL).Class("dropdown-item")
		if msg.Target != "" {
			link.Attr("target", msg.Target)
		}
		link.Child(item)
		return link
	}

	return item
}

// navbarNotificationsDropdown creates a notifications dropdown menu
func navbarNotificationsDropdown(navbarTextColor string, notifications []types.Alert) *hb.Tag {
	// Only show first 5 notifications in the dropdown
	visibleNotifications := notifications
	if len(visibleNotifications) > 5 {
		visibleNotifications = visibleNotifications[:5]
	}

	dropdown := hb.Li().Class("nav-item dropdown")
	a := hb.A().Href("#").Class("nav-link").Data("toggle", "dropdown")

	// Notification icon with badge
	icon := hb.I().Class("far fa-bell")
	if navbarTextColor != "" {
		icon.Style("color: " + navbarTextColor + " !important")
	}
	a.Child(icon)
	a.Child(hb.Span().Class("badge badge-warning navbar-badge").HTML(strconv.Itoa(len(notifications))))

	// Dropdown menu
	menu := hb.Div().Class("dropdown-menu dropdown-menu-lg dropdown-menu-right")

	// Header
	header := hb.Span().Class("dropdown-item dropdown-header").HTML(strconv.Itoa(len(notifications)) + " Notifications")
	menu.Child(header)

	// Notifications
	for _, notif := range visibleNotifications {
		menu.Child(createNotificationItem(notif, navbarTextColor))
	}

	// Footer
	footer := hb.A().Href("#").Class("dropdown-item dropdown-footer").HTML("View All Notifications")
	menu.Child(footer)

	dropdown.Child(a)
	dropdown.Child(menu)

	return dropdown
}

// createNotificationItem creates a single notification item for the notifications dropdown
func createNotificationItem(alert types.Alert, textColor string) *hb.Tag {
	item := hb.Div().Class("dropdown-item")

	// Notification icon based on alert type
	iconClass := "fas fa-info-circle"
	switch alert.Type {
	case "success":
		iconClass = "fas fa-check-circle"
	case "danger":
		iconClass = "fas fa-exclamation-circle"
	case "warning":
		iconClass = "fas fa-exclamation-triangle"
	}

	icon := hb.I().Class(iconClass + " mr-2")
	if textColor != "" {
		icon.Style("color: " + textColor + " !important")
	}

	// Notification content
	content := hb.Div().Class("d-flex align-items-center")
	content.Child(icon)
	content.Child(hb.Div().HTML(alert.Message))

	item.Child(content)
	return item
}

// navbarUserMenu creates the user menu dropdown
func navbarUserMenu(navbarTextColor string, user types.User, userMenuItems []types.MenuItem) *hb.Tag {
	dropdown := hb.Li().Class("nav-item dropdown user-menu")

	// User menu toggle
	toggle := hb.A().Href("#").Class("nav-link").Data("toggle", "dropdown")

	// User image
	userImage := hb.Img("https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y")
	userImage.Class("user-image img-circle elevation-2")

	toggle.Child(userImage)

	// User name
	if user.FirstName != "" {
		toggle.Child(hb.Span().Class("d-none d-md-inline").HTML(user.FirstName))
	}

	dropdown.Child(toggle)

	// Dropdown menu
	dropdownMenu := hb.Div().Class("dropdown-menu dropdown-menu-lg dropdown-menu-right")

	// User info section
	userInfo := hb.Div().Class("dropdown-item dropdown-header")

	// User image in menu
	userInfo.Child(userImage.Class("img-size-50 mr-3 img-circle"))

	// User details
	userDetails := hb.Div()
	userDetails.Child(hb.P().HTML(user.FirstName))

	if user.Email != "" {
		emailSpan := hb.Span().Class("text-muted").Style("font-size: 0.875em;").HTML(user.Email)
		userDetails.Child(emailSpan)
	}

	userInfo.Child(userDetails)
	dropdownMenu.Child(userInfo)

	// Menu items
	if len(userMenuItems) > 0 {
		dropdownMenu.Child(hb.Div().Class("dropdown-divider"))

		for _, item := range userMenuItems {
			// Handle dividers (empty title indicates a divider)
			if item.Title == "" && item.URL == "" && item.Icon == "" {
				dropdownMenu.Child(hb.Div().Class("dropdown-divider"))
				continue
			}

			link := hb.A().Href(item.URL).Class("dropdown-item")

			// Add icon if available
			if item.Icon != "" {
				iconHTML := "<i class=\"" + item.Icon + " mr-2"
				if navbarTextColor != "" {
					iconHTML += " style=\"color: " + navbarTextColor + " !important\""
				}
				iconHTML += "></i>"
				link.Child(hb.Raw(iconHTML))
			}

			// Add title
			link.Child(hb.Span().HTML(item.Title))

			// Show sequence number as a badge if it's greater than 0
			if item.Sequence > 0 {
				badge := hb.Span().Class("float-right text-muted text-sm badge badge-secondary")
				badge.HTML(strconv.Itoa(item.Sequence))
				link.Child(badge)
			}

			dropdownMenu.Child(link)
		}
	}

	dropdown.Child(dropdownMenu)

	return dropdown
}

// navbarThemeSwitcher creates a theme switcher dropdown
func navbarThemeSwitcher(navbarTextColor, currentTheme, themeHandlerUrl string) *hb.Tag {
	dropdown := hb.Li().Class("nav-item dropdown")
	a := hb.A().Href("#").Class("nav-link").Data("toggle", "dropdown")

	// Theme icon
	iconHTML := "<i class=\"fas fa-moon"
	if navbarTextColor != "" {
		iconHTML += " style=\"color: " + navbarTextColor + " !important\""
	}
	iconHTML += "></i>"
	a.Child(hb.Raw(iconHTML))

	// Dropdown menu
	menu := hb.Div().Class("dropdown-menu dropdown-menu-right")

	// Header
	header := hb.H6().Class("dropdown-header").HTML("Theme")
	menu.Child(header)

	// Theme options
	themes := []string{"default", "light", "dark"}
	themeTitles := map[string]string{
		"default": "Default",
		"light":   "Light",
		"dark":    "Dark",
	}

	themeIcons := map[string]string{
		"default": "fas fa-adjust",
		"light":   "fas fa-sun",
		"dark":    "fas fa-moon",
	}

	for _, theme := range themes {
		link := hb.A().Href("#").Class("dropdown-item theme-switch")
		link.Data("theme", theme)

		iconHTML := "<i class=\"" + themeIcons[theme] + " mr-2\"></i>"
		link.Child(hb.Raw(iconHTML))

		link.Child(hb.Span().HTML(themeTitles[theme]))

		if theme == currentTheme {
			checkIcon := "<i class=\"fas fa-check float-right mt-1\"></i>"
			link.Child(hb.Raw(checkIcon))
		}

		menu.Child(link)
	}

	dropdown.Child(a)
	dropdown.Child(menu)

	return dropdown
}
