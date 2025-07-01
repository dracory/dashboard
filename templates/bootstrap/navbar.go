package bootstrap

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
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
	user := dashboard.GetUser()

	dropdownUser := navbarDropdownUser(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, *user, dashboard.GetMenuUserItems())
	dropdownQuickAccess := navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode, dashboard.GetMenuQuickAccessItems())
	dropdownThemeSwitch := navbarDropdownThemeSwitch(navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode)

	buttonTheme := navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode)
	buttonMenuToggle := buttonMenuToggle(buttonTheme, hasNavbarTextColor, dashboard, iconStyle)
	buttonOffcanvasToggle := buttonOffcanvasToggle(buttonTheme, hasNavbarTextColor, navbarTextColor, iconStyle, dashboard)

	buttonMainMenu := buttonOffcanvasToggle
	if dashboard.GetMenuType() == types.MENU_TYPE_MODAL {
		buttonMainMenu = buttonMenuToggle
	}

	logo := lo.
		If(hasLogoRawHTML, hb.Raw(dashboard.GetLogoRawHtml())).
		ElseIf(hasLogoImage, hb.Image(dashboard.GetLogoImageURL()).Style("max-height:35px;")).
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

	toolbar := hb.Nav().
		ID("Toolbar").
		Class("navbar").
		ClassIf(navbarHasBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode), navbarThemeBackgroundClass).
		Style("z-index: 3;box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);transition: all .2s ease;padding-left: 20px;padding-right: 20px; display:block;").
		StyleIf(hasNavbarBackgroundColor, `background-color: `+navbarBackgroundColor+`;`).
		StyleIf(hasNavbarTextColor, `color: `+navbarTextColor+`;`).
		ChildIf(hasLogo, logoLink).
		Child(buttonMainMenu).
		ChildIf(user != nil,
			hb.Div().Class("float-end").
				Style("margin-left:10px;").
				Child(dropdownUser),
		).
		ChildIf(lo.IsEmpty(user) && dashboard.GetRegisterURL() != "", registerLink).
		ChildIf(lo.IsEmpty(user) && dashboard.GetLoginURL() != "", loginLink).
		ChildIf(dashboard.GetThemeHandlerUrl() != "", hb.Div().Class("float-end").Style("margin-left:10px;").Child(dropdownThemeSwitch)).
		ChildIf(len(dashboard.GetMenuQuickAccessItems()) > 0, hb.Div().
			Class("float-end").
			Style("margin-left:10px;").
			Child(dropdownQuickAccess))

	// Add modal menu to the DOM if menu type is modal
	if dashboard.GetMenuType() == types.MENU_TYPE_MODAL {
		modalMenuTag := menuModal(dashboard)
		toolbar.Child(modalMenuTag)
	} else {
		offcanvasMenuTag := menuOffcanvas(dashboard)
		toolbar.Child(offcanvasMenuTag)
	}

	mainMenu := lo.TernaryF(dashboard.GetMenuType() == types.MENU_TYPE_MODAL, func() *hb.Tag {
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
		Child(icons.Icon("bi-list", 24, 24, "").Style(iconStyle)).
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
			icons.Icon("bi-list", 24, 24, "").Style(iconStyle),
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
func navbarDropdownQuickAccess(iconStyle, navbarTextColor, navbarBackgroundColor, navbarBackgroundColorMode string, quickAccessItems []types.MenuItem) *hb.Tag {
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

	// Add quick access items from configuration
	if len(quickAccessItems) > 0 {
		var menuItems []hb.TagInterface

		// Group items into rows of 3
		for i := 0; i < len(quickAccessItems); i += 3 {
			end := i + 3
			if end > len(quickAccessItems) {
				end = len(quickAccessItems)
			}
			rowItems := quickAccessItems[i:end]

			row := hb.Div().Class("row g-0")

			for _, item := range rowItems {
				icon := item.Icon
				if icon == "" {
					icon = "bi bi-app"
				}

				col := hb.Div().Class("col-4 text-center").
					Child(hb.Hyperlink().
						Class("dropdown-item d-flex flex-column align-items-center").
						Href(item.URL).
						Child(icons.Icon(icon, 24, 24, "")).
						Child(hb.Span().Text(item.Title).Class("mt-1")),
					)
				row.Child(col)
			}

			menuItems = append(menuItems, row)
		}

		dropdownMenu.Children(menuItems)
	}

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
