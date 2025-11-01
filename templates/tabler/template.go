package tabler

import (
	"fmt"
	"strings"

	"github.com/dracory/dashboard/types"
	"github.com/dracory/hb"
	"github.com/samber/lo"
)

type Template struct{}

func New() *Template {
	return &Template{}
}

func (t *Template) getStylesAndScripts(dashboard types.DashboardInterface) (
	styleURLs []string,
	scriptURLs []string,
	styles []string,
	scripts []string,
) {
	styleURLs = make([]string, 0)
	scriptURLs = make([]string, 0)
	styles = make([]string, 0)
	scripts = make([]string, 0)

	// Add CSS
	if style := templateStyle(); style != "" {
		styles = append(styles, style)
	}

	// Add JavaScript
	if script := templateScript(); script != "" {
		scripts = append(scripts, script)
	}

	styleURLs = append(styleURLs, "https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/css/tabler.min.css")
	styleURLs = append(styleURLs, "https://cdn.jsdelivr.net/npm/@tabler/icons-webfont@2.47.0/tabler-icons.min.css")
	scriptURLs = append(scriptURLs, "https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/dist/js/tabler.min.js")

	styles = append(styles, "@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap');")

	styleURLs = append(styleURLs, dashboard.GetStyleURLs()...)
	scriptURLs = append(scriptURLs, dashboard.GetScriptURLs()...)

	styles = append(styles, dashboard.GetStyles()...)
	scripts = append(scripts, dashboard.GetScripts()...)

	return lo.Uniq(styleURLs), lo.Uniq(scriptURLs), lo.Uniq(styles), lo.Uniq(scripts)
}

// ToHTML generates the complete HTML page
func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	styleURLs, scriptURLs, styles, scripts := t.getStylesAndScripts(dashboard)

	// Create a new webpage
	webpage := hb.Webpage()

	// Set the page title
	webpage.SetTitle(dashboard.GetTitle())

	// Add favicon
	if favicon := favicon(); favicon != "" {
		webpage.SetFavicon(favicon)
	}

	// Add meta tags
	webpage.Meta(hb.Meta().Attr("charset", "utf-8"))
	webpage.Meta(hb.Meta().Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1, viewport-fit=cover"))

	// Add styles URLs
	for _, styleURL := range styleURLs {
		if styleURL != "" {
			webpage.AddStyleURL(styleURL)
		}
	}

	// Add styles
	for _, style := range styles {
		if style != "" {
			webpage.AddStyle(style)
		}
	}

	// Add scripts URLs
	for _, scriptURL := range scriptURLs {
		if scriptURL != "" {
			webpage.AddScriptURL(scriptURL)
		}
	}

	// Add scripts
	for _, script := range scripts {
		if script != "" {
			webpage.AddScript(script)
		}
	}

	// Set HTML attributes
	webpage.Attr("lang", "en")
	webpage.Attr("data-bs-theme", lo.Ternary(dashboard.GetTheme() == "dark", "dark", "light"))

	// Set body classes
	bodyClasses := []string{
		"antialiased",
		"theme-" + dashboard.GetTheme(),
	}

	// Add sidebar collapse class if sidebar is collapsed
	if dashboard.GetSidebarCollapsed() {
		bodyClasses = append(bodyClasses, "sidebar-collapse")
	}

	webpage.Body().AddClass(strings.Join(bodyClasses, " "))

	// Generate the layout
	layoutHTML := t.layout(dashboard)

	// Add the layout to the webpage
	webpage.Body().Child(hb.Raw(layoutHTML))

	// Add back to top button
	webpage.Body().Child(hb.NewDiv().Class("back-to-top").Child(
		hb.NewA().Href("#").Class("btn btn-primary btn-icon").Child(
			hb.NewI().Class("ti ti-arrow-up"),
		),
	))

	// Handle redirect if needed
	if redirectURL := dashboard.GetRedirectUrl(); redirectURL != "" {
		redirectTime := "0"
		if dashboard.GetRedirectTime() != "" {
			redirectTime = dashboard.GetRedirectTime()
		}
		webpage.Meta(hb.Meta().
			Attr("http-equiv", "refresh").
			Attr("content", redirectTime+"; url="+redirectURL))
	}

	// Generate the final HTML
	return webpage.ToHTML()
}

// layout generates the main layout structure
func (t *Template) layout(dashboard types.DashboardInterface) string {
	content := dashboard.GetContent()

	// Create page container
	container := hb.NewDiv().Class("page")

	// Add navbar
	container.Child(hb.Raw(topNavigation(dashboard)))

	// Create page wrapper
	contentWrapper := hb.NewDiv().Class("page-wrapper")

	// Add page header if title exists
	if title := dashboard.GetTitle(); title != "" {
		header := hb.NewDiv().Class("page-header d-print-none")
		headerContent := hb.NewDiv().Class("container-xl")
		headerRow := hb.NewDiv().Class("row g-2 align-items-center")

		// Left side (title and breadcrumb)
		headerCol := hb.NewDiv().Class("col")
		headerCol.Child(hb.NewH1().Class("page-title").Text(title))

		// Add breadcrumb if available
		if breadcrumb := dashboard.GetBreadcrumb(); len(breadcrumb) > 0 {
			// Add page pretitle (subtitle) if available
			if subtitle := dashboard.GetSubtitle(); subtitle != "" {
				headerCol.Child(hb.NewDiv().Class("page-pretitle").Text(subtitle))
			}

			// Create breadcrumb navigation
			navEl := hb.NewNav().Class("breadcrumb")
			for i, item := range breadcrumb {
				itemEl := hb.NewSpan().Class("breadcrumb-item")
				if i == len(breadcrumb)-1 {
					// Last item is active (not clickable)
					itemEl.Class("active").AddChild(hb.NewSpan().Text(item.Title))
				} else if item.URL != "" {
					// Clickable breadcrumb item
					itemEl.AddChild(hb.NewA().Href(item.URL).Text(item.Title))
				} else {
					// Non-clickable breadcrumb item
					itemEl.AddChild(hb.NewSpan().Text(item.Title))
				}
				navEl.AddChild(itemEl)
			}
			headerCol.Child(navEl)
		}

		headerRow.Child(headerCol)

		// Right side (actions)
		actionsCol := hb.NewDiv().Class("col-auto ms-auto d-print-none")
		actionsRow := hb.NewDiv().Class("btn-list")

		// Add action buttons
		actions := dashboard.GetActions()
		for _, action := range actions {
			btn := hb.NewButton().
				Class(fmt.Sprintf("btn %s", lo.Ternary(action.Primary, "btn-primary", "btn-outline-secondary"))).
				Attr("onclick", action.OnClick).
				Child(hb.NewI().Class(fmt.Sprintf("ti ti-%s me-2", action.Icon))).
				Child(hb.NewSpan().Text(action.Title))

			if action.ID != "" {
				btn.ID(action.ID)
			}

			actionsRow.Child(btn)
		}

		actionsCol.Child(actionsRow)
		headerRow.Child(actionsCol)
		headerContent.Child(headerRow)
		header.Child(headerContent)
		contentWrapper.Child(header)
	}

	// Add page body
	pageBody := hb.NewDiv().Class("page-body")
	pageBodyContent := hb.NewDiv().Class("container-xl")

	// Add alerts if any
	alerts := dashboard.GetAlerts()
	for _, alert := range alerts {
		alertEl := hb.NewDiv().
			Class(fmt.Sprintf("alert alert-%s alert-dismissible fade show", alert.Type)).
			Attr("role", "alert")

		// Add message
		alertEl.AddChild(hb.NewDiv().Text(alert.Message))

		// Add close button
		closeBtn := hb.NewButton().
			Type("button").
			Class("btn-close").
			Attr("data-bs-dismiss", "alert").
			Attr("aria-label", "Close")

		alertEl.AddChild(closeBtn)
		pageBodyContent.AddChild(alertEl)
	}

	// Add main content
	pageBodyContent.Child(hb.Raw(content))
	pageBody.Child(pageBodyContent)

	contentWrapper.Child(pageBody)
	container.Child(contentWrapper)

	// Add modals if any
	modals := dashboard.GetModals()
	for _, modal := range modals {
		modalEl := hb.NewDiv()
		modalEl.Class("modal fade")
		modalEl.Attr("id", modal.ID)
		modalEl.Attr("tabindex", "-1")
		modalEl.Attr("role", "dialog")
		modalEl.Attr("aria-hidden", "true")

		dialogEl := hb.NewDiv().Class("modal-dialog")
		switch modal.Size {
		case "sm":
			dialogEl.Class("modal-sm")
		case "lg":
			dialogEl.Class("modal-lg")
		case "xl":
			dialogEl.Class("modal-xl")
		}

		contentEl := hb.NewDiv().Class("modal-content")

		// Header
		headerEl := hb.NewDiv().Class("modal-header")
		headerEl.Child(hb.NewH5().Class("modal-title").Text(modal.Title))
		if modal.CloseButton {
			closeBtn := hb.NewButton()
			closeBtn.Type("button")
			closeBtn.Class("btn-close")
			closeBtn.Attr("data-bs-dismiss", "modal")
			closeBtn.Attr("aria-label", "Close")
			headerEl.Child(closeBtn)
		}
		contentEl.Child(headerEl)

		// Body
		bodyEl := hb.NewDiv().Class("modal-body").HTML(modal.Content)
		contentEl.Child(bodyEl)

		// Footer
		if modal.Footer != "" {
			footerEl := hb.NewDiv().Class("modal-footer").HTML(modal.Footer)
			contentEl.Child(footerEl)
		} else if modal.CloseButton {
			// Add default close button if no footer is provided
			footerEl := hb.NewDiv().Class("modal-footer")
			closeBtn := hb.NewButton().
				Type("button").
				Class("btn btn-link link-secondary").
				Attr("data-bs-dismiss", "modal").
				Text("Close")
			footerEl.Child(closeBtn)
			contentEl.Child(footerEl)
		}

		dialogEl.Child(contentEl)
		modalEl.Child(dialogEl)
		pageBodyContent.Child(modalEl)
	}

	return container.ToHTML()
}
