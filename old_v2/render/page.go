package render

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
)

// RenderPage generates the complete page HTML for the dashboard
func RenderPage(d model.DashboardRenderer) *hb.Tag {
	isDarkTheme := isThemeDark(d)

	// Create the head section
	head := hb.NewTag("head").
		Child(hb.NewTag("meta").Attr("charset", "utf-8")).
		Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1, viewport-fit=cover")).
		Child(hb.NewTag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "ie=edge")).
		Child(hb.NewTag("title").Text("Dashboard"))

	// Favicon
	head = head.Child(renderFavicon(d))

	// Tabler Core CSS
	cssLinks := renderTablerCSSLinks(isDarkTheme)
	for _, link := range cssLinks {
		head = head.Child(link)
	}

	// Custom styles
	head = head.Child(hb.Style(dashboardStyle(d)))

	// Create the body section
	bodyAttrs := map[string]string{}
	if isDarkTheme {
		bodyAttrs["data-bs-theme"] = "dark"
	}

	// Create page content
	contentContainer := hb.Div().
		Class("container-xl").
		Child(hb.NewHTML(d.GetContent()))

	pageContent := hb.Div().
		Class("page-body").
		Child(contentContainer)

	// Create page wrapper
	pageWrapper := hb.Div().
		Class("page-wrapper").
		Child(pageContent).
		Child(RenderFooter(d))

	// Create page container
	pageContainer := hb.Div().
		Class("page").
		Child(RenderHeader(d)).
		Child(pageWrapper)

	// Create body with scripts
	body := hb.NewTag("body").
		Attrs(bodyAttrs).
		Child(pageContainer)

	// Add Tabler Core JS
	jsScripts := renderTablerJSScripts()
	for _, script := range jsScripts {
		body = body.Child(script)
	}

	// Add custom scripts
	body = body.Child(hb.Script(dashboardScript(d)))

	// Create the complete HTML document
	html := hb.NewTag("html").Attr("lang", "en").
		Child(head).
		Child(body)

	// Add DOCTYPE and render the HTML
	doctype := hb.NewHTML("<!DOCTYPE html>")

	// Create the final wrapper
	wrapper := hb.Wrap().
		Child(doctype).
		Child(html)

	return wrapper
}

// renderFavicon generates the favicon link tag
func renderFavicon(d model.DashboardRenderer) *hb.Tag {
	if d.GetFaviconURL() == "" {
		return hb.NewHTML("")
	}

	return hb.NewLink().Attr("rel", "icon").Attr("href", d.GetFaviconURL())
}

// renderTablerCSSLinks generates the Tabler CSS link tags
func renderTablerCSSLinks(isDarkTheme bool) []*hb.Tag {
	links := []*hb.Tag{
		hb.NewLink().Href("https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler.min.css").Rel("stylesheet"),
		hb.NewLink().Href("https://cdn.jsdelivr.net/npm/@tabler/icons@latest/iconfont/tabler-icons.min.css").Rel("stylesheet"),
	}

	if isDarkTheme {
		links = append(links,
			hb.NewLink().Href("https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler-dark.min.css").Rel("stylesheet"),
		)
	}

	return links
}

// renderTablerJSScripts generates the Tabler JS script tags
func renderTablerJSScripts() []*hb.Tag {
	return []*hb.Tag{
		hb.NewTag("script").Attr("src", "https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/js/tabler.min.js"),
	}
}

// isThemeDark returns whether the theme is dark
func isThemeDark(d model.DashboardRenderer) bool {
	return d.GetThemeName() == "dark"
}

// dashboardStyle returns the dashboard custom CSS
func dashboardStyle(d model.DashboardRenderer) string {
	return `
		.navbar-brand-image {
			height: 2rem;
		}
		.navbar-vertical.navbar-expand-lg {
			width: 15rem;
		}
		.navbar-vertical.navbar-expand-lg .navbar-collapse {
			margin: 0 -0.5rem;
		}
	`
}

// dashboardScript returns the dashboard custom JavaScript
func dashboardScript(d model.DashboardRenderer) string {
	return `
		// Theme switcher
		document.querySelectorAll('[data-bs-theme-value]').forEach(function(element) {
			element.addEventListener('click', function() {
				var theme = this.getAttribute('data-bs-theme-value');
				document.body.setAttribute('data-bs-theme', theme);
				localStorage.setItem('theme', theme);
			});
		});
		
		// Set theme from localStorage
		var theme = localStorage.getItem('theme');
		if (theme) {
			document.body.setAttribute('data-bs-theme', theme);
		}
	`
}
