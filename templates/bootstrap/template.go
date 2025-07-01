package bootstrap

import (
	"strings"

	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/bs"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/icons"
	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type Template struct {
}

var _ types.TemplateInterface = (*Template)(nil)

func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	pageContent := dashboard.GetContent()

	styleURLs := []string{
		// Icons
		cdn.BootstrapIconsCss_1_11_3(),
	}

	// Bootstrap Css
	// if d.uncdnHandlerEndpoint != "" {
	// 	styleURLs = append(styleURLs, uncdnThemeStyleURL(d.uncdnHandlerEndpoint, d.theme))
	// } else {
	// 	styleURLs = append(styleURLs, cdnThemeStyleUrl(d.theme))
	// }
	styleURLs = append(styleURLs, cdn.BootstrapCss_5_3_3())

	scriptURLs := []string{}

	// // Bootstrap JS
	// if d.uncdnHandlerEndpoint != "" {
	// 	scriptURLs = append(scriptURLs, uncdn.BootstrapJs523())
	// } else {
	scriptURLs = append(scriptURLs, cdn.BootstrapJs_5_3_3())
	// }

	faviconURL := dashboard.GetFaviconURL()
	// if faviconURL == "" {
	// 	faviconURL = favicon()
	// }

	webpage := hb.Webpage()
	webpage.SetTitle(dashboard.GetTitle())

	// Required Style URLs
	webpage.AddStyleURLs(styleURLs)

	// Custom Style URLs
	webpage.AddStyleURLs(dashboard.GetStyleURLs())

	// Template Styles
	// webpage.AddStyle(d.templateStyle())

	// Custom Styles
	webpage.AddStyles(dashboard.GetStyles())

	// Required Script URLs
	webpage.AddScriptURLs(scriptURLs)

	// Custom Script URLs
	webpage.AddScriptURLs(dashboard.GetScriptURLs())

	// Template Scripts
	// webpage.AddScript(templateStyles())

	// Custom Scripts
	webpage.AddScripts(dashboard.GetScripts())

	// webpage.AddScript(scripts(d.scripts))
	webpage.SetFavicon(faviconURL)
	if dashboard.GetRedirectUrl() != "" && dashboard.GetRedirectTime() != "" {
		webpage.Meta(hb.Meta().
			Attr("http-equiv", "refresh").
			Attr("content", dashboard.GetRedirectTime()+"; url = "+dashboard.GetRedirectUrl()))
	}

	// menu := d.menuOffcanvas().ToHTML()

	// if d.menuType == MENU_TYPE_MODAL {
	// 	menu += d.menuModal().ToHTML()
	// }

	// webpage.AddChild(hb.Raw(d.layout() + menu))

	webpage.Child(hb.Raw(pageContent))

	return webpage.ToHTML()
}

func buildSubmenuItem(menuItem types.MenuItem, index int) *hb.Tag {
	title := menuItem.Title
	if title == "" {
		title = "n/a"
	}
	url := menuItem.URL
	if url == "" {
		url = "#"
	}
	icon := menuItem.Icon
	target := menuItem.Target
	if target == "" {
		target = "_self"
	}

	children := menuItem.Children
	hasChildren := len(children) > 0
	submenuId := "submenu_" + utils.ToString(index)
	if hasChildren {
		url = "#" + submenuId
	}

	link := hb.Hyperlink().Class("nav-link px-0")

	if icon != "" {
		link.Child(hb.Span().
			Class("icon").
			Style("margin-right: 5px;").
			HTML(icon))
	} else {
		link.Child(hb.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" fill="currentColor" class="bi bi-caret-right-fill" viewBox="0 0 16 16">
	<path d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"/>
</svg>`))
	}
	link.Child(hb.Span().Class("d-inline").HTML(title))
	link.Href(url)
	if hasChildren {
		link.Data("bs-toggle", "collapse")
	}

	return hb.LI().
		Class("w-100").
		Child(link)
}

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
	submenuId := "submenu_" + utils.ToString(index)
	if hasChildren {
		url = "#" + submenuId
	}

	link := hb.Hyperlink().Class("nav-link align-middle px-0")
	if icon != "" {
		link.Child(hb.Span().Class("icon").Style("margin-right: 5px;").HTML(icon))
	}
	link.HTML(title)
	link.Href(url)
	if hasChildren {
		link.Data("bs-toggle", "collapse")
	}

	if hasChildren {
		html := `<b class="caret">
			<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" fill="currentColor" class="bi bi-caret-down-fill" viewBox="0 0 16 16">
			<path d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"/>
			</svg>
		</b>`
		link.Child(hb.Raw(html))
	}

	li := hb.LI().Class("nav-item").Child(link)

	if hasChildren {
		ul := hb.UL().
			ID(submenuId).
			Class("collapse hide nav flex-column ms-1").
			Data("bs-parent", "#DashboardMenu")
		for childIndex, childMenuItem := range children {
			childItem := buildSubmenuItem(childMenuItem, childIndex)
			ul.Child(childItem)
		}
		li.Child(ul)
	}

	return li
}

func dashboardMenuNavbar(menuItems []types.MenuItem) string {
	items := []hb.TagInterface{}

	for index, menuItem := range menuItems {
		li := buildMenuItem(menuItem, index)
		items = append(items, li)
	}

	ul := hb.UL().
		ID("DashboardMenu").
		Class("navbar-nav justify-content-end flex-grow-1 pe-3").
		Children(items)

	return ul.ToHTML()
}

// topNavigation returns the HTML code for the top navigation toolbar in the Dashboard.
//
// No parameters.
// Returns a string.
func topNavigation(menuItems []types.MenuItem, logoImageURL string, logoRawHtml string, logoRedirectURL string, navbarBackgroundColor string, navbarTextColor string) string {
	hasNavbarBackgroundColor := lo.Ternary(navbarBackgroundColor == "", false, true)
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)

	hasLogoImage := lo.Ternary(logoImageURL != "", true, false)
	hasLogoRawHTML := lo.Ternary(logoRawHtml != "", true, false)
	hasLogo := hasLogoImage || hasLogoRawHTML
	logoRedirectURL := lo.Ternary(logoRedirectURL != "", logoRedirectURL, "#")

	navbarThemeBackgroundClass := navbarBackgroundThemeClass(navbarBackgroundColor)

	iconStyle := "margin-top:-4px;margin-right:5px;"

	dropdownUser := navbarDropdownUser(iconStyle)
	dropdownQuickAccess := navbarDropdownQuickAccess(iconStyle)
	dropdownThemeSwitch := navbarDropdownThemeSwitch()

	buttonTheme := navbarButtonThemeClass()

	buttonMenuToggle := hb.Button().
		Class("btn "+buttonTheme).
		Style("background: none; border:none;").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Data("bs-toggle", "modal").
		Data("bs-target", "#ModalDashboardMenu").
		Children([]hb.TagInterface{
			icons.Icon("bi-list", 24, 24, "").Style(iconStyle),
			hb.Span().
				Class("d-none d-md-inline-block").
				HTML("Menu"),
		})

	buttonOffcanvasToggle := hb.Button().
		Class("btn "+buttonTheme).
		Style("background: none; border:none;").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Data("bs-toggle", "offcanvas").
		Data("bs-target", "#OffcanvasMenu").
		Child(icons.Icon("bi-list", 24, 24, "").Style(iconStyle)).
		ChildIf(d.menuShowText, hb.Span().
			Class("d-none d-md-inline-block").
			HTML("Menu"))

	mainMenu := buttonOffcanvasToggle
	if d.menuType == MENU_TYPE_MODAL {
		mainMenu = buttonMenuToggle
	}

	logo := lo.
		If(hasLogoRawHTML, hb.Raw(d.logoRawHtml)).
		ElseIf(hasLogoImage, hb.Image(d.logoImageURL).Style("max-height:35px;")).
		Else(nil)

	logoLink := hb.Hyperlink().
		Href(logoRedirectURL).
		Class("navbar-brand").
		Child(logo)

	loginLink := hb.Hyperlink().
		Text("Login").
		Href(d.loginURL).
		//Class("btn "+buttonTheme+" float-end").
		Class("btn btn-outline-info float-end").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Style("margin-left:10px;")

	registerLink := hb.Hyperlink().
		Text("Register").
		Href(d.registerURL).
		Class("btn "+buttonTheme+" float-end").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Style("margin-left:10px;  border:none;")

	toolbar := hb.Nav().
		ID("Toolbar").
		Class("navbar").
		ClassIf(d.navbarHasBackgroundThemeClass(), navbarThemeBackgroundClass).
		Style("z-index: 3;box-shadow: 0 5px 20px rgba(0, 0, 0, 0.1);transition: all .2s ease;padding-left: 20px;padding-right: 20px; display:block;").
		StyleIf(hasNavbarBackgroundColor, `background-color: `+d.navbarBackgroundColor+`;`).
		StyleIf(hasNavbarTextColor, `color: `+d.navbarTextColor+`;`).
		ChildIf(hasLogo, logoLink).
		Children([]hb.TagInterface{
			mainMenu,

			// User Menu
			hb.If(!lo.IsEmpty(d.user) && (d.user.FirstName != "" || d.user.LastName != ""),
				hb.Div().Class("float-end").
					Style("margin-left:10px;").
					Child(dropdownUser),
			),

			// Register Link
			hb.If(lo.IsEmpty(d.user) && d.registerURL != "",
				registerLink,
			),

			// Login Link
			hb.If(lo.IsEmpty(d.user) && d.loginURL != "",
				loginLink,
			),

			// Theme Switcher
			hb.If(d.themeHandlerUrl != "",
				hb.Div().Class("float-end").
					Style("margin-left:10px;").
					Child(dropdownThemeSwitch),
			),

			// Quick Menu (if provided)
			hb.If(len(d.quickAccessMenu) > 0, hb.Div().
				Class("float-end").
				Style("margin-left:10px;").
				Child(dropdownQuickAccess)),
		})

	return toolbar.ToHTML()
}

func center(content string) string {
	return content
}

func menuOffcanvas() *hb.Tag {
	backgroundClass := navbarBackgroundThemeClass()

	offcanvasMenu := hb.Div().
		ID("OffcanvasMenu").
		Class("offcanvas offcanvas-start").
		Class(backgroundClass).
		ClassIfElse(backgroundClass == "bg-light", "text-bg-light", "text-bg-dark").
		Attr("tabindex", "-1").
		Children([]hb.TagInterface{
			hb.Div().Class("offcanvas-header").
				Children([]hb.TagInterface{
					hb.Heading5().
						Class("offcanvas-title").
						Text("Menu"),
					hb.Button().
						Class("btn-close btn-close-white").
						ClassIf(backgroundClass == "bg-light", "text-bg-light").
						Type(hb.TYPE_BUTTON).
						Data("bs-dismiss", "offcanvas").
						Attr("aria-label", "Close"),
				}),
			hb.Div().Class("offcanvas-body").
				Children([]hb.TagInterface{
					hb.Raw(d.DashboardLayoutMenu()),
				}),
		})

	return offcanvasMenu
}

func menuModal() *hb.Tag {
	modalHeader := hb.Div().Class("modal-header").
		Children([]hb.TagInterface{
			hb.Heading5().HTML("Menu").Class("modal-title"),
			hb.Button().Attrs(map[string]string{
				"type":            "button",
				"class":           "btn-close",
				"data-bs-dismiss": "modal",
				"aria-label":      "Close",
			}),
		})

	modalBody := hb.Div().Class("modal-body").Children([]hb.TagInterface{
		hb.Raw(dashboardMenuNavbar()),
	})

	modalFooter := hb.Div().Class("modal-footer").Children([]hb.TagInterface{
		hb.Button().
			HTML("Close").
			Class("btn btn-secondary w-100").
			Data("bs-dismiss", "modal"),
	})

	modal := hb.Div().
		ID("ModalDashboardMenu").
		Class("modal fade").
		Children([]hb.TagInterface{
			hb.Div().Class("modal-dialog modal-lg").
				Children([]hb.TagInterface{
					hb.Div().Class("modal-content").
						Children([]hb.TagInterface{
							modalHeader,
							modalBody,
							modalFooter,
						}),
				}),
		})

	return modal
}

func templateStyle() string {
	fullHeightSupport := `html, body{ height: 100%; }`

	css := fullHeightSupport
	return css
}

// scripts returns the JavaScript code for the Dashboard.
//
// No parameters.
// Returns a string.
func templateScript() string {
	js := ``
	return js
}

// favicon returns the data URI for a website favicon.
//
// No parameters.
// Returns a string.
func favicon() string {
	favicon := "data:image/x-icon;base64,AAABAAEAEBAQAAAAAAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAgAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAzMzMAAAAmQBmZpkA////AJmZzAAzM5kAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAzMzMzMxQAA1YiIiIiUQADViIiIiZERANTIiIiJBVRRTMiIiJBNmJFNSIiJEMmZlRlIiJlYiJmVDUiImIiIiIUMzImMiIiJRFGUiZiImJkQEMzIiImFlEABDMyIiZiVAAEFTNiI2ZEAABBFTJmJRAAAABBRjQjQAAAAABFEBFACABwAAAAcAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAEAAMABAADAAQAA4AMAAPgDAAD+IwAA"
	return favicon
}

// isThemeDark checks if the theme of the Dashboard is dark.
//
// It does so by checking if the Dashboard's theme name is contained
// in the list of dark themes.
//
// Returns a boolean indicating whether the theme is dark.
func isThemeDark(theme string) bool {
	isDark := lo.Contains(lo.Keys(themesDark), theme)
	return isDark
}

func navbarHasBackgroundThemeClass() bool {
	hasNavbarBackgroundColor := lo.Ternary(navbarBackgroundColor == "", false, true)
	hasNavbarBackgroundTheme := lo.Ternary(!hasNavbarBackgroundColor && navbarBackgroundColorMode != "", true, false)
	return hasNavbarBackgroundTheme
}

func navbarBackgroundThemeClass(navbarBackgroundColorMode string) string {
	navbarThemeBackgroundClass := lo.
		If(navbarHasBackgroundThemeClass(navbarBackgroundColorMode), "bg-"+navbarBackgroundColorMode).
		// ElseIf(!hasNavbarBackgroundTheme && !hasNavbarBackgroundColor, "bg-dark").
		Else("")

	return navbarThemeBackgroundClass
}

func navbarButtonThemeClass(navbarBackgroundColorMode string) string {
	buttonTheme := lo.
		If(navbarHasBackgroundThemeClass(navbarBackgroundColorMode), "btn-"+navbarBackgroundColorMode).
		Else("")
	return buttonTheme
}

func navbarDropdownQuickAccess(iconStyle string) *hb.Tag {
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)
	buttonTheme := navbarButtonThemeClass(navbarBackgroundColorMode)

	button := hb.Button().
		ID("ButtonQuickAccess").
		Class("btn "+buttonTheme+" dropdown-toggle").
		Style("background:none;border:0px;").
		StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
		Type(hb.TYPE_BUTTON).
		Data("bs-toggle", "dropdown").
		Children([]hb.TagInterface{
			icons.Icon("bi-microsoft", 24, 24, "").
				Style(iconStyle).
				Style("margin-top:-4px;margin-right:8px;"),
			hb.Span().
				Class("d-none d-md-inline-block").
				Text("Quick Access").
				Style("margin-right:10px;"),
		})

	dropdownQuickAccess := hb.Div().
		Class("dropdown").
		Style(`margin:0px;`).
		Child(button).
		Child(hb.UL().
			Class("dropdown-menu").
			Children(lo.Map(d.quickAccessMenu, func(item MenuItem, _ int) hb.TagInterface {
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
			})))

	return dropdownQuickAccess
}

// themeButton generates a dropdown menu with light and dark themes.
//
// It checks if the current theme is dark and creates dropdown items for both light and dark themes.
// The dropdown items are created dynamically based on the themesLight and themesDark maps.
// The function returns a *hb.Tag that represents the generated dropdown menu.
func navbarDropdownThemeSwitch(navbarTextColor string) *hb.Tag {
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)
	buttonTheme := navbarButtonThemeClass(navbarTextColor)

	isDark := isThemeDark(navbarTextColor)

	// Light Themes
	lightDropdownItems := lo.Map(lo.Keys(themesLight), func(theme string, index int) hb.TagInterface {
		name := themesLight[theme]
		active := lo.Ternary(d.theme == theme, " active", "")
		url := lo.Ternary(strings.Contains(d.themeHandlerUrl, "?"), d.themeHandlerUrl+"&theme="+theme, d.themeHandlerUrl+"?theme="+theme)

		if len(d.themesRestrict) > 0 {
			if customName, exists := d.themesRestrict[theme]; exists {
				name = customName
			} else {
				return nil
			}
		}

		return hb.LI().Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("dropdown-item"+active).
				Child(hb.I().Class("bi bi-sun").Style("margin-right:5px;")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	// Dark Themes
	darkDropdownItems := lo.Map(lo.Keys(themesDark), func(theme string, index int) hb.TagInterface {
		name := themesDark[theme]
		active := lo.Ternary(d.theme == theme, " active", "")
		url := lo.Ternary(strings.Contains(d.themeHandlerUrl, "?"), d.themeHandlerUrl+"&theme="+theme, d.themeHandlerUrl+"?theme="+theme)

		if len(d.themesRestrict) > 0 {
			if customName, exists := d.themesRestrict[theme]; exists {
				name = customName
			} else {
				return nil
			}
		}

		return hb.LI().Children([]hb.TagInterface{
			hb.Hyperlink().
				Class("dropdown-item"+active).
				Child(hb.I().Class("bi bi-moon-stars-fill").Style("margin-right:5px;")).
				HTML(name).
				Href(url).
				Attr("ref", "nofollow"),
		})
	})

	button := bs.Button().
		ID("buttonTheme").
		Class(buttonTheme+" dropdown-toggle").
		Style("background:none;border:0px;").
		StyleIf(hasNavbarTextColor, "color:"+d.navbarTextColor).
		Data("bs-toggle", "dropdown").
		Children([]hb.TagInterface{
			lo.Ternary(isDark, hb.I().Class("bi bi-sun"), hb.I().Class("bi bi-moon-stars-fill")),
		})

	return hb.Div().
		Class("dropdown").
		// Style(`margin:0px;`).
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

func navbarDropdownUser(iconStyle string, navbarTextColor string, user types.User) *hb.Tag {
	hasNavbarTextColor := lo.Ternary(navbarTextColor == "", false, true)
	buttonTheme := navbarButtonThemeClass(navbarTextColor)
	userName := user.FirstName + " " + user.LastName

	dropdownUser := hb.Div().
		Class("dropdown").
		Children([]hb.TagInterface{
			hb.Button().
				ID("ButtonUser").
				Class("btn "+buttonTheme+" dropdown-toggle").
				Style("background:none;border:0px;").
				StyleIf(hasNavbarTextColor, "color: "+d.navbarTextColor+";").
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
				Children(lo.Map(d.userMenu, func(item MenuItem, _ int) hb.TagInterface {
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
