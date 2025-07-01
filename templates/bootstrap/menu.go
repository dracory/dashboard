package bootstrap

import (
	"strconv"

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
	submenuID := "submenu_" + strconv.Itoa(index)

	if hasChildren {
		url = "#" + submenuID
	}

	link := hb.NewHyperlink().Class("nav-link align-middle px-0")

	if icon != "" {
		link.Child(hb.NewSpan().
			Class("icon").
			Style("margin-right: 5px;").
			HTML(icon))
	} else if hasChildren {
		link.Child(hb.NewRaw(`<svg xmlns="http://www.w3.org/2000/svg" width="8" height="8" fill="currentColor" class="bi bi-caret-right-fill" viewBox="0 0 16 16">
		<path d="m12.14 8.753-5.482 4.796c-.646.566-1.658.106-1.658-.753V3.204a1 1 0 0 1 1.659-.753l5.48 4.796a1 1 0 0 1 0 1.506z"/>
		</svg>`))
	}

	link.Child(hb.NewSpan().Class("d-inline").HTML(title))
	link.Href(url)

	if hasChildren {
		link.Data("bs-toggle", "collapse")
	}

	li := hb.NewLI().Class("nav-item").Child(link)

	if hasChildren {
		ul := hb.NewUL().
			ID(submenuID).
			Class("collapse hide nav flex-column ms-1").
			Data("bs-parent", "#DashboardMenu")

		for childIndex, childMenuItem := range children {
			childItem := buildMenuItem(childMenuItem, childIndex)
			ul.Child(childItem)
		}

		li.Child(ul)
	}

	return li
}

// dashboardMenuNavbar generates the HTML for the dashboard menu navbar
func dashboardMenuNavbar(dashboard types.DashboardInterface) string {
	nav := hb.NewNav().Class("nav nav-pills flex-column mb-auto").ID("DashboardMenu")

	for i, item := range dashboard.GetMenuMainItems() {
		nav.Child(buildMenuItem(item, i))
	}

	return nav.ToHTML()
}

// menuOffcanvas generates the offcanvas menu HTML
func menuOffcanvas(dashboard types.DashboardInterface) *hb.Tag {
	// Get the background class based on the current theme
	var backgroundClass string
	if navbarBg, ok := dashboard.GetNavbarBackground(); ok && navbarBg != "" {
		backgroundClass = navbarBg
	} else {
		backgroundClass = "bg-dark"
	}

	offcanvas := hb.NewDiv().
		ID("OffcanvasMenu").
		Class("offcanvas offcanvas-start").
		Class(backgroundClass).
		ClassIfElse(backgroundClass == "bg-light", "text-bg-light", "text-bg-dark").
		Attr("tabindex", "-1").
		Attr("aria-labelledby", "OffcanvasMenuLabel").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("offcanvas-header").
				Children([]hb.TagInterface{
					hb.NewH5().
						ID("OffcanvasMenuLabel").
						Class("offcanvas-title").
						Text("Menu"),
					hb.NewButton().
						Class("btn-close btn-close-white").
						ClassIf(backgroundClass == "bg-light", "text-bg-light").
						Type(hb.TYPE_BUTTON).
						Data("bs-dismiss", "offcanvas").
						Attr("aria-label", "Close"),
				}),
			hb.NewDiv().Class("offcanvas-body").
				Children([]hb.TagInterface{
					hb.Raw(dashboardMenuNavbar(dashboard)),
				}),
		})

	return offcanvas
}

// menuModal generates the modal menu HTML
func menuModal(dashboard types.DashboardInterface) *hb.Tag {
	modalHeader := hb.NewDiv().Class("modal-header").
		Children([]hb.TagInterface{
			hb.NewH5().HTML("Menu").Class("modal-title").ID("ModalDashboardMenuLabel"),
			hb.NewButton().Attrs(map[string]string{
				"type":            "button",
				"class":           "btn-close",
				"data-bs-dismiss": "modal",
				"aria-label":      "Close",
			}),
		})

	modalBody := hb.NewDiv().Class("modal-body").
		Children([]hb.TagInterface{
			hb.Raw(dashboardMenuNavbar(dashboard)),
		})

	modalFooter := hb.NewDiv().Class("modal-footer").
		Children([]hb.TagInterface{
			hb.NewButton().
				HTML("Close").
				Class("btn btn-secondary w-100").
				Data("bs-dismiss", "modal"),
		})

	modal := hb.NewDiv().
		ID("ModalDashboardMenu").
		Class("modal fade").
		Attr("tabindex", "-1").
		Attr("aria-labelledby", "ModalDashboardMenuLabel").
		Attr("aria-hidden", "true").
		Children([]hb.TagInterface{
			hb.NewDiv().Class("modal-dialog modal-lg").
				Children([]hb.TagInterface{
					hb.NewDiv().Class("modal-content").
						Children([]hb.TagInterface{
							modalHeader,
							modalBody,
							modalFooter,
						}),
				}),
		})

	return modal
}
