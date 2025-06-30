package render

import (
	"fmt"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render/templates"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/gouniverse/hb"
)

const THEME_ADMINLTE = "adminlte"
const THEME_BOOTSTRAP = "bootstrap"
const THEME_TABLER = "tabler"
const THEME_DEFAULT = THEME_TABLER

// RenderPage generates the complete page HTML for the dashboard
func RenderPage(d model.DashboardRenderer) *hb.Tag {
	// Get the theme manager instance
	themeManager := templates.Manager()

	// Get the theme from the dashboard configuration
	themeName := d.GetThemeName()
	if themeName == "" {
		fmt.Printf("[DEBUG] No theme specified in dashboard, using default: %s\n", THEME_DEFAULT)
		themeName = THEME_DEFAULT // Default theme
	} else {
		fmt.Printf("[DEBUG] Dashboard requested theme: %s\n", themeName)
	}

	// Get the theme instance
	themeInstance := themeManager.Get(themeName)
	if themeInstance == nil {
		fmt.Printf("[WARN] Requested theme not found: %s, falling back to default: %s\n", themeName, THEME_DEFAULT)
		themeInstance = themeManager.Get(THEME_DEFAULT)
	}

	// Ensure we have a valid theme instance
	if themeInstance == nil {
		fmt.Printf("[ERROR] No theme available, using default theme implementation\n")
		themeInstance = &shared.DefaultTemplate{}
	}

	fmt.Printf("[DEBUG] Using theme: %s (type: %T)\n", themeInstance.GetName(), themeInstance)

	// Check if we should use dark color scheme
	isDarkColorScheme := isDarkColorScheme(d)
	// Create the head section
	head := hb.NewTag("head").
		Child(hb.NewTag("meta").Attr("charset", "utf-8")).
		Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1, viewport-fit=cover")).
		Child(hb.NewTag("meta").Attr("http-equiv", "X-UA-Compatible").Attr("content", "ie=edge")).
		Child(hb.NewTag("title").Text("Dashboard"))

	// Favicon
	head = head.Child(renderFavicon(d))

	// Theme CSS with appropriate color scheme
	cssLinks := themeInstance.GetCSSLinks(isDarkColorScheme)
	for _, link := range cssLinks {
		head = head.Child(link)
	}

	// Theme custom styles
	head = head.Child(hb.Style(themeInstance.GetCustomCSS()))

	// Create the body section
	bodyAttrs := map[string]string{}
	if isDarkColorScheme {
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
		Child(themeInstance.RenderFooter(d))

	// Create page container
	pageContainer := hb.Div().
		Class("page").
		Child(themeInstance.RenderHeader(d)).
		Child(pageWrapper)

	// Create body with scripts
	body := hb.NewTag("body").
		Attrs(bodyAttrs).
		Child(pageContainer)

	// Add theme JavaScript
	jsScripts := themeInstance.GetJSScripts()
	for _, script := range jsScripts {
		body = body.Child(script)
	}

	// Add theme custom JavaScript
	body = body.Child(hb.Script(themeInstance.GetCustomJS()))

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

// isDarkColorScheme checks if the dashboard should use dark color scheme
func isDarkColorScheme(d model.DashboardRenderer) bool {
	// Check the navbar background color mode to determine if we're in dark mode
	return d.GetNavbarBackgroundColorMode() == "dark"
}
