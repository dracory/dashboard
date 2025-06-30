package render

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
)

// RenderThemeSwitcher generates the theme switcher dropdown HTML
func RenderThemeSwitcher(d model.DashboardRenderer) *hb.Tag {
	// Create the trigger button
	triggerButton := hb.A().Href("#").Class("nav-link px-0").
		Attr("data-bs-toggle", "dropdown").
		Attr("tabindex", "-1").
		Attr("aria-label", "Show theme options").
		Child(hb.I().Class("ti ti-sun"))

	// Create light theme option
	lightOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "light").
		Child(hb.I().Class("ti ti-sun me-2")).
		Child(hb.Text("Light"))

	// Create dark theme option
	darkOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "dark").
		Child(hb.I().Class("ti ti-moon me-2")).
		Child(hb.Text("Dark"))

	// Create auto theme option
	autoOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "auto").
		Child(hb.I().Class("ti ti-device-desktop me-2")).
		Child(hb.Text("Auto"))

	// Create dropdown menu
	dropdownMenu := hb.Div().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow").
		Child(lightOption).
		Child(darkOption).
		Child(autoOption)

	// Create container
	container := hb.Div().Class("nav-item dropdown d-none d-md-flex me-3").
		Child(triggerButton).
		Child(dropdownMenu)

	return container
}
