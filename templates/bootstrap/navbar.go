package bootstrap

import (
	"strings"

	"github.com/dracory/dashboard/shared"
	"github.com/dracory/dashboard/types"
	"github.com/dracory/hb"
	"github.com/samber/lo"
)

// topNavigation returns the HTML code for the top navigation toolbar in the Dashboard.
//
// No parameters.
// Returns a string.
func topNavigation(dashboard types.DashboardInterface) string {
	hasNavbarBackgroundColor := lo.Ternary(dashboard.GetNavbarBackgroundColor() == "", false, true)
	hasNavbarTextColor := lo.Ternary(dashboard.GetNavbarTextColor() == "", false, true)

	hasLogoImage := lo.Ternary(dashboard.GetLogoImageURL() != "", true, false)
	hasLogoRawHTML := lo.Ternary(dashboard.GetLogoRawHtml() != "", true, false)
	hasLogo := hasLogoImage || hasLogoRawHTML
	logoRedirectURL := lo.Ternary(dashboard.GetLogoRedirectURL() != "", dashboard.GetLogoRedirectURL(), "#")

	navbarThemeBackgroundClass := navbarBackgroundThemeClass(dashboard.GetNavbarBackgroundColor(), dashboard.GetNavbarBackgroundColorMode())

	iconStyle := "margin-top:-4px;margin-right:5px;"
	navbarTextColor := dashboard.GetNavbarTextColor()
	navbarBackgroundColor := dashboard.GetNavbarBackgroundColor()
	navbarBackgroundColorMode := dashboard.GetNavbarBackgroundColorMode()
	if navbarTextColor == "" {
		navbarTextColor = defaultNavbarTextColor(navbarBackgroundColor, navbarBackgroundColorMode, dashboard.GetTheme())
	}
	user := dashboard.GetUser()

	dropdownUser := navbarDropdownUser(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, *user, dashboard.GetMenuUserItems())
	dropdownQuickAccess := navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, dashboard.GetMenuQuickAccessItems())
	dropdownThemeSwitch := navbarDropdownThemeSwitch(navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, dashboard.GetTheme(), dashboard.GetThemeHandlerUrl())

	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)
	buttonMenuToggle := buttonMenuToggle(buttonTheme, hasNavbarTextColor, dashboard, iconStyle)
	buttonOffcanvasToggle := buttonOffcanvasToggle(buttonTheme, hasNavbarTextColor, navbarTextColor, iconStyle, dashboard)

	buttonMainMenu := buttonOffcanvasToggle
	if dashboard.GetMenuType() == shared.TEMPLATE_BOOTSTRAP_MENU_TYPE_MODAL {
		buttonMainMenu = buttonMenuToggle
	}

	logo := lo.
		If(hasLogoRawHTML, hb.Raw(dashboard.GetLogoRawHtml())).
		ElseIf(hasLogoImage, hb.Image(dashboard.GetLogoImageURL()).Style("max-height:35px;height:35px; width:auto;")).
		Else(nil)

	logoLink := hb.Hyperlink().
		Href(logoRedirectURL).
		Class("navbar-brand").
		Child(logo)

	loginLink := hb.Hyperlink().
		Text("Login").
		Href(dashboard.GetLoginURL()).
		//Class("btn "+buttonTheme+" float-end").
		Class("btn btn-outline-info float-end").
		StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
		Style("margin-left:10px;")

	registerLink := hb.Hyperlink().
		Text("Register").
		Href(dashboard.GetRegisterURL()).
		Class("btn "+buttonTheme+" float-end").
		StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
		Style("margin-left:10px;  border:none;")

	// Create a container for the right-aligned items
	rightItems := hb.Div().Class("d-flex align-items-center ms-auto")

	// Add quick access dropdown if items exist
	if len(dashboard.GetMenuQuickAccessItems()) > 0 {
		rightItems.Child(dropdownQuickAccess)
	}

	// Add theme switch dropdown if handler URL exists
	if dashboard.GetThemeHandlerUrl() != "" {
		rightItems.Child(hb.Div().Style("margin-left:10px;").Child(dropdownThemeSwitch))
	}

	// Add user dropdown or login/register links
	if user != nil {
		rightItems.Child(hb.Div().Style("margin-left:10px;").Child(dropdownUser))
	} else {
		if dashboard.GetLoginURL() != "" {
			rightItems.Child(loginLink)
		}
		if dashboard.GetRegisterURL() != "" {
			rightItems.Child(registerLink)
		}
	}

	// Create items array and add conditionally
	var items []hb.TagInterface

	// Add main menu button
	items = append(items, buttonMainMenu)

	// User Menu - add conditionally
	hasUser := user != nil && (user.FirstName != "" || user.LastName != "")
	if hasUser {
		userDiv := hb.Div().Class("float-end").Style("margin-left:10px;").Child(dropdownUser)
		items = append(items, userDiv)
	}

	// Theme Switcher - add conditionally
	hasThemeHandler := dashboard.GetThemeHandlerUrl() != ""
	if hasThemeHandler {
		themeDiv := hb.Div().Class("float-end").Style("margin-left:10px;").Child(dropdownThemeSwitch)
		items = append(items, themeDiv)
	}

	// Quick Access Menu - add conditionally
	hasQuickAccess := len(dashboard.GetMenuQuickAccessItems()) > 0
	if hasQuickAccess {
		quickAccessDiv := hb.Div().Class("float-end").Style("margin-left:10px;").Child(dropdownQuickAccess)
		items = append(items, quickAccessDiv)
	}

	toolbar := hb.Nav().
		ID("Toolbar").
		Class("navbar").
		ClassIf(navbarHasBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode), navbarThemeBackgroundClass).
		Style("z-index: 3;box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);transition: all .2s ease;padding-left: 20px;padding-right: 20px; display:block;").
		StyleIf(hasNavbarBackgroundColor, `background-color: `+navbarBackgroundColor+`;`).
		StyleIf(hasNavbarTextColor, `color: `+navbarTextColor+`;`).
		ChildIf(hasLogo, logoLink).
		Children(items)

	// Create the main menu based on menu type
	mainMenu := lo.TernaryF(dashboard.GetMenuType() == shared.TEMPLATE_BOOTSTRAP_MENU_TYPE_MODAL, func() *hb.Tag {
		return menuModal(dashboard)
	}, func() *hb.Tag {
		return menuOffcanvas(dashboard)
	})

	return hb.Wrap().
		Child(toolbar).
		Child(mainMenu).
		ToHTML()
}

func buttonOffcanvasToggle(buttonTheme string, hasNavbarTextColor bool, navbarTextColor string, iconStyle string, dashboard types.DashboardInterface) *hb.Tag {
	buttonOffcanvasToggle := hb.Button().
		Class("btn "+buttonTheme).
		Style("background: none; border:none;").
		StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
		Data("bs-toggle", "offcanvas").
		Data("bs-target", "#OffcanvasMenu").
		Child(hb.I().Class("bi bi-list").Style(iconStyle)).
		ChildIf(dashboard.GetMenuShowText(), hb.Span().
			Class("d-none d-md-inline-block").
			HTML("Menu"))
	return buttonOffcanvasToggle
}

func buttonMenuToggle(buttonTheme string, hasNavbarTextColor bool, dashboard types.DashboardInterface, iconStyle string) *hb.Tag {
	buttonMenuToggle := hb.Button().
		Class("btn "+buttonTheme).
		Style("background: none; border:none;").
		StyleIf(hasNavbarTextColor, "color: "+dashboard.GetNavbarTextColor()+";").
		Data("bs-toggle", "modal").
		Data("bs-target", "#ModalDashboardMenu").
		Children([]hb.TagInterface{
			hb.I().Class("bi bi-list").Style(iconStyle),
			hb.Span().
				Class("d-none d-md-inline-block").
				HTML("Menu"),
		})
	return buttonMenuToggle
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
					hb.I().Class("bi bi-person").Style(iconStyle),
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
func navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode string, quickAccessItems []types.MenuItem) *hb.Tag {
	hasNavbarTextColor := navbarTextColor != ""
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	button := hb.Button().
		ID("ButtonQuickAccess").
		Class("btn "+buttonTheme).
		Style("background:none;border:0px;padding:0.375rem;").
		StyleIf(hasNavbarTextColor, "color: "+navbarTextColor+";").
		Type(hb.TYPE_BUTTON).
		Data("bs-toggle", "dropdown").
		Aria("expanded", "false").
		Child(hb.Span().Class("d-flex align-items-center").Children([]hb.TagInterface{
			hb.I().Class("bi bi-grid-3x3-gap-fill").Style(iconStyle),
		}))

	dropdownMenu := hb.Div().
		Class("dropdown-menu dropdown-menu-end shadow").
		Style("min-width:300px;padding:0.5rem;").
		Aria("labelledby", "ButtonQuickAccess")

	// Add quick access items from configuration
	if len(quickAccessItems) > 0 {
		var menuItems []hb.TagInterface

		for _, item := range quickAccessItems {
			icon := item.Icon
			if icon == "" {
				icon = "bi-app"
			}

			menuItem := hb.Hyperlink().
				Class("dropdown-item d-flex align-items-center").
				Href(item.URL).
				Children([]hb.TagInterface{
					hb.Span().Class("me-2").HTML(icon),
					hb.Span().Text(item.Title),
				})

			menuItems = append(menuItems, menuItem)
		}

		dropdownMenu.Children(menuItems)
	}

	return hb.Div().
		Class("nav-item dropdown").
		Children([]hb.TagInterface{
			button,
			dropdownMenu,
		})
}

// navbarDropdownThemeSwitch creates a theme switcher dropdown
func navbarDropdownThemeSwitch(navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, currentTheme, themeHandlerUrl string) *hb.Tag {
	hasNavbarTextColor := navbarTextColor != ""
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)

	isDark := isThemeDark(navbarBackgroundColorMode)

	// Use default handler URL if none provided
	handlerUrl := lo.Ternary(themeHandlerUrl == "", "/", themeHandlerUrl)

	// Generate Light Theme dropdown items
	lightDropdownItems := lo.Map(lo.Keys(themesLight), func(theme string, index int) hb.TagInterface {
		name := themesLight[theme]
		// Mark current theme as active
		active := lo.Ternary(currentTheme == theme, " active", "")
		// Build URL with proper query parameter handling
		url := lo.Ternary(strings.Contains(handlerUrl, "?"),
			handlerUrl+"&theme="+theme,
			handlerUrl+"?theme="+theme)

		return hb.LI().Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("dropdown-item"+active).
				Child(hb.I().Class("bi bi-sun me-2")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	// Generate Dark Theme dropdown items
	darkDropdownItems := lo.Map(lo.Keys(themesDark), func(theme string, index int) hb.TagInterface {
		name := themesDark[theme]
		// Mark current theme as active
		active := lo.Ternary(currentTheme == theme, " active", "")
		// Build URL with proper query parameter handling
		url := lo.Ternary(strings.Contains(handlerUrl, "?"),
			handlerUrl+"&theme="+theme,
			handlerUrl+"?theme="+theme)

		return hb.LI().Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("dropdown-item"+active).
				Child(hb.I().Class("bi bi-moon-stars-fill me-2")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	button := hb.Button().
		ID("buttonTheme").
		Class(buttonTheme+" dropdown-toggle d-flex align-items-center justify-content-center").
		Style("background:none;border:0px;padding:0.375rem;").
		Style("gap:0.25rem;").
		StyleIf(hasNavbarTextColor, "color:"+navbarTextColor).
		Data("bs-toggle", "dropdown").
		Children([]hb.TagInterface{
			lo.Ternary(isDark, hb.I().Class("bi bi-sun"), hb.I().Class("bi bi-moon-stars-fill")),
		})

	return hb.Div().
		Class("dropdown d-flex align-items-center").
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
