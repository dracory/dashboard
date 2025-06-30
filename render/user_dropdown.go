package render

// import (
// 	"github.com/dracory/dashboard/model"
// 	"github.com/gouniverse/hb"
// 	"github.com/samber/lo"
// )

// // RenderUserDropdown generates the user dropdown HTML
// func RenderUserDropdown(d model.DashboardRenderer) *hb.Tag {
// 	user := d.GetUser()

// 	if user.ID == "" {
// 		// If no user, show login/register links if available
// 		if !lo.IsEmpty(d.GetLoginURL()) || !lo.IsEmpty(d.GetRegisterURL()) {
// 			loginLink := hb.NewHTML("")
// 			if !lo.IsEmpty(d.GetLoginURL()) {
// 				loginLink = hb.A().Href(d.GetLoginURL()).Class("dropdown-item").Text("Login")
// 			}

// 			registerLink := hb.NewHTML("")
// 			if !lo.IsEmpty(d.GetRegisterURL()) {
// 				registerLink = hb.A().Href(d.GetRegisterURL()).Class("dropdown-item").Text("Register")
// 			}

// 			// Create avatar icon
// 			avatarIcon := hb.Span().Class("avatar avatar-sm").Child(
// 				hb.I().Class("ti ti-user"),
// 			)

// 			// Create user menu trigger
// 			userMenuTrigger := hb.A().Href("#").Class("nav-link d-flex lh-1 text-reset p-0")
// 			userMenuTrigger = userMenuTrigger.Attr("data-bs-toggle", "dropdown")
// 			userMenuTrigger = userMenuTrigger.Attr("aria-label", "Open user menu")
// 			userMenuTrigger = userMenuTrigger.Child(avatarIcon)

// 			// Create dropdown menu
// 			dropdownMenu := hb.Div().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")
// 			dropdownMenu = dropdownMenu.Child(loginLink)
// 			dropdownMenu = dropdownMenu.Child(registerLink)

// 			// Create dropdown container
// 			dropdownContainer := hb.Div().Class("nav-item dropdown")
// 			dropdownContainer = dropdownContainer.Child(userMenuTrigger)
// 			dropdownContainer = dropdownContainer.Child(dropdownMenu)
// 			return dropdownContainer
// 		}

// 		return hb.NewHTML("")
// 	}

// 	// User is logged in
// 	var avatar *hb.Tag
// 	if !lo.IsEmpty(user.AvatarURL) {
// 		avatar = hb.Span().Class("avatar avatar-sm").Style("background-image: url(" + user.AvatarURL + ")")
// 	} else {
// 		avatar = hb.Span().Class("avatar avatar-sm").Text(user.Name[:1])
// 	}

// 	// User menu items
// 	userMenuItems := hb.Wrap()
// 	for _, menuItem := range d.GetUserMenu() {
// 		// Check if this is a divider
// 		if menuItem.Icon == "dropdown-divider" {
// 			userMenuItems.Child(hb.Div().Class("dropdown-divider"))
// 			continue
// 		}

// 		var icon *hb.Tag
// 		if !lo.IsEmpty(menuItem.Icon) {
// 			icon = hb.I().Class(menuItem.Icon + " me-2")
// 		} else {
// 			icon = hb.NewHTML("")
// 		}

// 		var badge *hb.Tag
// 		if !lo.IsEmpty(menuItem.BadgeText) {
// 			badgeClass := "badge bg-primary ms-auto"
// 			if !lo.IsEmpty(menuItem.BadgeClass) {
// 				badgeClass = menuItem.BadgeClass
// 			}
// 			badge = hb.Span().Class(badgeClass).Text(menuItem.BadgeText)
// 		} else {
// 			badge = hb.NewHTML("")
// 		}

// 		itemTag := hb.A().Href(menuItem.URL).Class("dropdown-item")
// 		itemTag = itemTag.Child(icon)
// 		itemTag = itemTag.Child(hb.Text(menuItem.Text))
// 		itemTag = itemTag.Child(badge)

// 		if menuItem.NewWindow {
// 			itemTag.Target("_blank")
// 		}

// 		if !lo.IsEmpty(menuItem.OnClick) {
// 			itemTag.Attr("onclick", menuItem.OnClick)
// 		}

// 		userMenuItems.Child(itemTag)
// 	}

// 	// Create user menu trigger with avatar
// 	userMenuTrigger := hb.A().
// 		Href("#").
// 		Class("nav-link d-flex lh-1 text-reset p-0").
// 		Attr("data-bs-toggle", "dropdown").
// 		Attr("aria-label", "Open user menu")

// 	// Create user info section
// 	userInfo := hb.Div().
// 		Class("d-none d-xl-block ps-2").
// 		Child(hb.Div().Text(user.Name)).
// 		Child(hb.Div().Class("mt-1 small text-muted").Text(user.Email))

// 	// Add avatar and user info to the trigger
// 	userMenuTrigger = userMenuTrigger.Child(avatar)
// 	userMenuTrigger = userMenuTrigger.Child(userInfo)

// 	// Create dropdown container
// 	dropdownContainer := hb.Div().
// 		Class("nav-item dropdown").
// 		Child(userMenuTrigger).
// 		Child(hb.Div().
// 			Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow").
// 			Child(userMenuItems))

// 	return dropdownContainer
// }
