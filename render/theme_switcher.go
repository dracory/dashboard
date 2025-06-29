package render

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
)

// RenderThemeSwitcher generates the theme switcher dropdown HTML
func RenderThemeSwitcher(d model.DashboardRenderer) *hb.Tag {
	// Create the trigger button
	triggerButton := hb.A().Href("#").Class("nav-link px-0")
	triggerButton = triggerButton.Attr("data-bs-toggle", "dropdown")
	triggerButton = triggerButton.Attr("tabindex", "-1")
	triggerButton = triggerButton.Attr("aria-label", "Show theme options")
	triggerButton = triggerButton.Child(hb.I().Class("ti ti-sun"))
	
	// Create light theme option
	lightOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "light")
	lightOption = lightOption.Child(hb.I().Class("ti ti-sun me-2"))
	lightOption = lightOption.Child(hb.Text("Light"))
	
	// Create dark theme option
	darkOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "dark")
	darkOption = darkOption.Child(hb.I().Class("ti ti-moon me-2"))
	darkOption = darkOption.Child(hb.Text("Dark"))
	
	// Create auto theme option
	autoOption := hb.A().Href("#").Class("dropdown-item").Attr("data-bs-theme-value", "auto")
	autoOption = autoOption.Child(hb.I().Class("ti ti-device-desktop me-2"))
	autoOption = autoOption.Child(hb.Text("Auto"))
	
	// Create dropdown menu
	dropdownMenu := hb.Div().Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow")
	dropdownMenu = dropdownMenu.Child(lightOption)
	dropdownMenu = dropdownMenu.Child(darkOption)
	dropdownMenu = dropdownMenu.Child(autoOption)
	
	// Create container
	container := hb.Div().Class("nav-item dropdown d-none d-md-flex me-3")
	container = container.Child(triggerButton)
	container = container.Child(dropdownMenu)
	
	return container
}
