package bootstrap

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/samber/lo"
)

// topNavigation returns the HTML code for the top navigation bar
func topNavigation(
	dashboard types.DashboardInterface,
) string {
	logoImageURL := dashboard.GetLogoImageURL()
	logoRawHtml := dashboard.GetLogoRawHtml()
	logoRedirectURL := dashboard.GetLogoRedirectURL()
	mainMenuItems := dashboard.GetMenuMainItems()
	navbarBackgroundColor := dashboard.GetNavbarBackgroundColor()
	navbarTextColor := dashboard.GetNavbarTextColor()
	navbarBackgroundColorMode := dashboard.GetNavbarBackgroundColorMode()
	user := dashboard.GetUser()

	iconStyle := lo.Ternary(navbarTextColor == "", "", "color: "+navbarTextColor)
	bgClass := navbarBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	dropdownQuickAccess := navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode)
	dropdownThemeSwitch := navbarDropdownThemeSwitch(navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode)

	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	navbar := hb.NewNavbar().
		Class("navbar navbar-expand-lg").
		ClassIf(bgClass != "", bgClass).
		StyleIf(navbarTextColor != "", "color: "+navbarTextColor)

	// Container
	container := hb.NewDiv().Class("container-fluid")

	// Brand/logo
	brand := hb.NewHyperlink().
		Class("navbar-brand").
		ChildIf(logoRawHtml == "", hb.Image(logoImageURL).Style("height:40px")).
		ChildIf(logoRawHtml != "", hb.Raw(logoRawHtml)).
		ChildIf(logoRawHtml == "" && logoImageURL == "", hb.Text("Logo"))

	if logoRedirectURL != "" {
		brand.Href(logoRedirectURL)
	}

	// Toggle button for mobile
	toggleButton := hb.NewButton().
		Class("navbar-toggler"+buttonTheme).
		Type("button").
		Data("bs-toggle", "collapse").
		Data("bs-target", "#navbarNav").
		Child(hb.NewSpan().Class("navbar-toggler-icon"))

	// Navbar content
	navbarContent := hb.NewDiv().
		Class("collapse navbar-collapse").
		ID("navbarNav")

	// Menu items
	nav := hb.NewUL().
		Class("navbar-nav me-auto mb-2 mb-lg-0").
		Children(lo.Map(mainMenuItems, func(item types.MenuItem, _ int) hb.TagInterface {
			return hb.NewLI().Class("nav-item").
				Child(hb.NewHyperlink().
					Class("nav-link").
					Href(item.URL).
					Text(item.Title))
		}))

	// Right-aligned items
	navbarNavRight := hb.NewUL().Class("navbar-nav ms-auto")

	// Add dropdowns if they exist
	if dropdownQuickAccess != nil {
		navbarNavRight.Child(hb.NewLI().Class("nav-item").Child(dropdownQuickAccess))
	}

	if dropdownThemeSwitch != nil {
		navbarNavRight.Child(hb.NewLI().Class("nav-item").Child(dropdownThemeSwitch))
	}

	if user != nil {
		dropdownUser := navbarDropdownUser(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, *user, dashboard.GetMenuUserItems())
		navbarNavRight.Child(hb.NewLI().Class("nav-item").Child(dropdownUser))
	}

	// Build the navbar
	navbarContent.Children([]hb.TagInterface{
		nav,
		navbarNavRight,
	})

	container.Children([]hb.TagInterface{
		brand,
		toggleButton,
		navbarContent,
	})

	navbar.Child(container)

	return navbar.ToHTML()
}

// navbarDropdownUser creates a user dropdown menu
func navbarDropdownUser(
	iconStyle,
	navbarTextColor,
	navbarBackgroundColor,
	navbarBackgroundColorMode string,
	user types.User,
	userMenuItems []types.MenuItem,
) *hb.Tag {
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)
	userName := user.FirstName + " " + user.LastName

	dropdownUser := hb.Div().
		Class("dropdown").
		Children([]hb.TagInterface{
			hb.Button().
				ID("ButtonUser").
				Class("btn "+buttonTheme+" dropdown-toggle").
				Style("background:none;border:0px;").
				StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
				Type(hb.TYPE_BUTTON).
				Data("bs-toggle", "dropdown").
				Children([]hb.TagInterface{
					icons.Icon("bi-person", 24, 24, "").Style(iconStyle),
					hb.Span().
						Class("d-none d-md-inline-block").
						Text(userName).
						Style("margin-right:10px;"),
				}),
			hb.UL().
				Class("dropdown-menu dropdown-menu-dark").
				Class(buttonTheme).
				Children(lo.Map(userMenuItems, func(item types.MenuItem, _ int) hb.TagInterface {
					target := lo.Ternary(item.Target == "", "_self", item.Target)
					url := lo.Ternary(item.URL == "", "#", item.URL)

					return hb.LI().Children([]hb.TagInterface{
						hb.If(item.Title == "",
							hb.HR().
								Class("dropdown-divider"),
						),

						hb.If(item.Title != "",
							hb.Hyperlink().
								Class("dropdown-item").
								ChildIf(item.Icon != "", hb.Span().Class("icon").Style("margin-right: 5px;").HTML(item.Icon)).
								Text(item.Title).
								Href(url).
								Target(target),
						),
					})
				})),
		})

	return dropdownUser
}

// navbarDropdownQuickAccess creates a quick access dropdown menu
func navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode string) *hb.Tag {
	hasNavbarTextColor := navbarTextColor != ""
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	button := hb.Button().
		ID("ButtonQuickAccess").
		Class("btn "+buttonTheme+" dropdown-toggle").
		Style("background:none;border:0px;").
		StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
		Type(hb.TYPE_BUTTON).
		Data("bs-toggle", "dropdown").
		Child(icons.Icon("bi bi-grid", 24, 24, ""))

	dropdownMenu := hb.Div().
		Class("dropdown-menu dropdown-menu-end").
		Style("min-width:300px;padding:0.5rem;")

	// Add quick access items here
	dropdownMenu.Children([]hb.TagInterface{
		hb.Div().
			Class("row g-0").
			Children([]hb.TagInterface{
				hb.Div().
					Class("col-4 text-center").
					Child(hb.Hyperlink().
						Class("dropdown-item d-flex flex-column align-items-center").
						Href("/dashboard").
						Child(icons.Icon("bi bi-speedometer2", 24, 24, "")).
						Child(hb.Span().Text("Dashboard").Class("mt-1")),
					),
				hb.Div().
					Class("col-4 text-center").
					Child(hb.Hyperlink().
						Class("dropdown-item d-flex flex-column align-items-center").
						Href("/profile").
						Child(icons.Icon("bi bi-person", 24, 24, "")).
						Child(hb.Span().Text("Profile").Class("mt-1")),
					),
				hb.Div().
					Class("col-4 text-center").
					Child(hb.Hyperlink().
						Class("dropdown-item d-flex flex-column align-items-center").
						Href("/settings").
						Child(icons.Icon("bi bi-gear", 24, 24, "")).
						Child(hb.Span().Text("Settings").Class("mt-1")),
					),
			}),
	})

	return hb.Div().
		Class("dropdown").
		Children([]hb.TagInterface{
			button,
			dropdownMenu,
		})
}

// navbarDropdownThemeSwitch creates a theme switcher dropdown
func navbarDropdownThemeSwitch(navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode string) *hb.Tag {
	hasNavbarTextColor := navbarTextColor != ""
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	isDark := isThemeDark(navbarBackgroundColorMode)

	darkDropdownItems := []hb.TagInterface{
		hb.Hyperlink().
			Class("dropdown-item").
			Href("/?theme=light").
			Child(hb.I().Class("bi bi-sun me-2")).
			HTML("Light"),
		hb.Hyperlink().
			Class("dropdown-item").
			Href("/?theme=dark").
			Child(hb.I().Class("bi bi-moon-stars-fill me-2")).
			HTML("Dark"),
	}

	lightDropdownItems := []hb.TagInterface{
		hb.Hyperlink().
			Class("dropdown-item").
			Href("/?theme=light").
			Child(hb.I().Class("bi bi-sun me-2")).
			HTML("Light"),
		hb.Hyperlink().
			Class("dropdown-item").
			Href("/?theme=dark").
			Child(hb.I().Class("bi bi-moon-stars-fill me-2")).
			HTML("Dark"),
	}

	button := hb.Button().
		ID("buttonTheme").
		Class(buttonTheme+" dropdown-toggle").
		Style("background:none;border:0px;").
		StyleIf(hasNavbarTextColor, "color:"+navbarTextColor).
		Data("bs-toggle", "dropdown").
		Children([]hb.TagInterface{
			lo.Ternary(isDark, hb.I().Class("bi bi-sun"), hb.I().Class("bi bi-moon-stars-fill")),
		})

	return hb.Div().
		Class("dropdown").
		Child(button).
		Child(hb.UL().
			Class(buttonTheme+" dropdown-menu dropdown-menu-dark").
			Children(lightDropdownItems).
			ChildIf(
				len(lo.Filter(darkDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0 && len(lo.Filter(lightDropdownItems, func(item hb.TagInterface, _ int) bool { return item != nil })) > 0,
				hb.LI().Children([]hb.TagInterface{
					hb.HR().Class("dropdown-divider"),
				}),
			).
			Children(darkDropdownItems))
}
