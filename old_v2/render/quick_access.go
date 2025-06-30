package render

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

// RenderQuickAccessMenu generates the quick access dropdown HTML
func RenderQuickAccessMenu(d model.DashboardRenderer) *hb.Tag {
	if len(d.GetQuickAccessMenu()) == 0 {
		return hb.NewHTML("")
	}

	menuItems := hb.Wrap()
	for _, menuItem := range d.GetQuickAccessMenu() {
		var icon *hb.Tag
		if !lo.IsEmpty(menuItem.Icon) {
			icon = hb.I().Class(menuItem.Icon + " me-2")
		} else {
			icon = hb.NewHTML("")
		}

		var badge *hb.Tag
		if !lo.IsEmpty(menuItem.BadgeText) {
			badgeClass := "badge bg-primary ms-auto"
			if !lo.IsEmpty(menuItem.BadgeClass) {
				badgeClass = menuItem.BadgeClass
			}
			badge = hb.Span().Class(badgeClass).Text(menuItem.BadgeText)
		} else {
			badge = hb.NewHTML("")
		}

		itemTag := hb.A().Href(menuItem.URL).
			Class("dropdown-item").
			Child(icon).
			Child(hb.Text(menuItem.Text)).
			Child(badge)

		if menuItem.NewWindow {
			itemTag.Target("_blank")
		}

		if !lo.IsEmpty(menuItem.OnClick) {
			itemTag.Attr("onclick", menuItem.OnClick)
		}

		menuItems.Child(itemTag)
	}

	// Create the trigger button
	triggerButton := hb.A().
		Href("#").
		Class("nav-link px-0").
		Attr("data-bs-toggle", "dropdown").
		Attr("tabindex", "-1").
		Attr("aria-label", "Quick access").
		Child(hb.I().Class("ti ti-layout-grid-add"))

	// Create dropdown menu
	dropdownMenu := hb.Div().
		Class("dropdown-menu dropdown-menu-end dropdown-menu-arrow").
		Child(menuItems)

	// Create container
	container := hb.Div().
		Class("nav-item dropdown d-none d-md-flex me-3").
		Child(triggerButton).
		Child(dropdownMenu)

	return container
}
