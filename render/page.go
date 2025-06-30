package render

import (
	"fmt"

	"github.com/dracory/dashboard/config"
	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/model/interfaces"
	"github.com/dracory/dashboard/model/templateregistry"
	"github.com/gouniverse/hb"
)

// RenderPage generates the complete page HTML for the dashboard
func RenderPage(d interfaces.DashboardRenderer) *hb.Tag {
	// Get the template name from the dashboard configuration
	templateName := d.GetTemplateName()
	if templateName == "" {
		fmt.Printf("[DEBUG] No template specified in dashboard, using default: %s\n", config.TEMPLATE_DEFAULT)
		templateName = config.TEMPLATE_DEFAULT
	} else {
		fmt.Printf("[DEBUG] Dashboard requested template: %s\n", templateName)
	}

	// Get the template instance from the registry
	templateInstance := templateregistry.Get(templateName)
	if templateInstance == nil {
		fmt.Printf("[WARN] Requested template not found: %s, falling back to default: %s\n", templateName, config.TEMPLATE_DEFAULT)
		templateName = config.TEMPLATE_DEFAULT
		templateInstance = templateregistry.GetDefault()

		if templateInstance == nil {
			// If we still can't find a template, use a default one
			fmt.Printf("[ERROR] No templates registered, using fallback\n")
			// Create a new default template from the model package
			template := &model.DefaultTemplate{}
			page, err := template.RenderPage("", d)
			if err != nil {
				// Return a basic error page if rendering fails
				errPage := hb.NewTag("html")
				head := hb.NewTag("head")
				head.Child(hb.NewTag("title").Text("Error"))
				errPage.Child(head)
				body := hb.NewTag("body")
				body.Child(hb.NewTag("h1").Text("Error Rendering Page"))
				body.Child(hb.NewTag("p").Text(fmt.Sprintf("Failed to render page: %v", err)))
				errPage.Child(body)
				return errPage
			}
			return page
		}
	}
	fmt.Printf("[DEBUG] Using template: %s (type: %T)\n", templateInstance.GetName(), templateInstance)

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

	// Template CSS with appropriate color scheme
	cssLinks := templateInstance.GetCSSLinks(isDarkColorScheme)
	for _, link := range cssLinks {
		head = head.Child(link)
	}

	// Template custom styles
	head = head.Child(hb.Style(templateInstance.GetCustomCSS()))

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
		Child(templateInstance.RenderFooter(d))

	// Create page container
	pageContainer := hb.Div().
		Class("page").
		Child(templateInstance.RenderHeader(d)).
		Child(pageWrapper)

	// Create body with scripts
	body := hb.NewTag("body").
		Attrs(bodyAttrs).
		Child(pageContainer)

	// Add template JavaScript
	jsScripts := templateInstance.GetJSScripts()
	for _, script := range jsScripts {
		body = body.Child(script)
	}

	// Add template custom JavaScript
	body = body.Child(hb.Script(templateInstance.GetCustomJS()))

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
	// For now, default to light theme
	// TODO: Implement proper theme detection based on user preferences
	return false
}
