package shared

import (
	"fmt"

	"github.com/dracory/dashboard/render/atomizer"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// DashboardRenderer defines the interface for dashboard renderers
type DashboardRenderer = model.DashboardRenderer

// Template defines the interface for dashboard templates
type Template interface {
	// RenderPage renders a complete page with the given content
	// and dashboard renderer
	RenderPage(content string, d DashboardRenderer) (*hb.Tag, error)
	// GetName returns the theme's name
	GetName() string

	// GetCSSLinks returns the CSS link tags for the theme
	GetCSSLinks(isDarkMode bool) []*hb.Tag

	// GetJSScripts returns the JavaScript script tags for the theme
	GetJSScripts() []*hb.Tag

	// GetCustomCSS returns any custom CSS for the theme
	GetCustomCSS() string

	// GetCustomJS returns any custom JavaScript for the theme
	GetCustomJS() string

	// RenderHeader renders the theme-specific header
	// Deprecated: Use RenderAtom with Omni atoms instead
	RenderHeader(d DashboardRenderer) *hb.Tag

	// RenderFooter renders the theme-specific footer
	// Deprecated: Use RenderAtom with Omni atoms instead
	RenderFooter(d DashboardRenderer) *hb.Tag

	// RenderAtom renders an Omni atom using the theme's styling
	RenderAtom(atom *omni.Atom) (*hb.Tag, error)

	// RenderDashboard renders a complete dashboard from Omni atoms
	RenderDashboard(dashboard *omni.Atom) (string, error)
}

// DefaultTemplate is a basic template implementation that can be used as a fallback
type DefaultTemplate struct{}

// RenderPage renders a complete page with the given content and dashboard renderer
func (t *DefaultTemplate) RenderPage(content string, d DashboardRenderer) (*hb.Tag, error) {
	header := t.RenderHeader(d)
	footer := t.RenderFooter(d)

	// Create the head section
	head := hb.NewTag("head").
		Child(hb.Meta().Attr("charset", "utf-8")).
		Child(hb.Meta().Name("viewport").Attr("content", "width=device-width, initial-scale=1, viewport-fit=cover")).
		Child(hb.Meta().Attr("http-equiv", "X-UA-Compatible").Attr("content", "ie=edge")).
		Child(hb.Title().Text("Dashboard"))

	// Add favicon if available
	if d.GetFaviconURL() != "" {
		head.Child(hb.Link().Rel("icon").Href(d.GetFaviconURL()))
	}

	// Add template CSS
	cssLinks := t.GetCSSLinks(t.isDarkColorScheme(d))
	for _, link := range cssLinks {
		head.Child(link)
	}

	// Create the body section
	bodyAttrs := map[string]string{}
	if t.isDarkColorScheme(d) {
		bodyAttrs["data-bs-theme"] = "dark"
	}

	body := hb.NewTag("body").
		Attrs(bodyAttrs).
		ChildIf(header != nil, header).
		Child(hb.NewHTML(content)).
		ChildIf(footer != nil, footer)

	// Add JavaScript
	for _, script := range t.GetJSScripts() {
		body.Child(script)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		body.Child(hb.Script(customJS))
	}

	// Create HTML document
	html := hb.NewTag("html").
		Attr("lang", "en").
		Child(head).
		Child(body)

	return hb.Wrap().
		Child(hb.NewHTML("<!DOCTYPE html>")).
		Child(html), nil
}

// isDarkColorScheme checks if the color scheme should be dark
func (t *DefaultTemplate) isDarkColorScheme(d DashboardRenderer) bool {
	return d.GetNavbarBackgroundColorMode() == "dark"
}

// GetName returns the name of the default template
func (t *DefaultTemplate) GetName() string {
	return "default"
}

// GetCSSLinks returns no CSS links for the default template
func (t *DefaultTemplate) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	// Return basic CSS reset as a fallback
	return []*hb.Tag{
		hb.Style(`
			* { box-sizing: border-box; margin: 0; padding: 0; }
			body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; line-height: 1.6; }
		`),
	}
}

// GetJSScripts returns no JavaScript for the default template
func (t *DefaultTemplate) GetJSScripts() []*hb.Tag {
	return nil
}

// GetCustomCSS returns empty string as there's no custom CSS for the default template
func (t *DefaultTemplate) GetCustomCSS() string {
	return ""
}

// GetCustomJS returns empty string as there's no custom JavaScript for the default template
func (t *DefaultTemplate) GetCustomJS() string {
	return ""
}

// RenderAtom renders an Omni atom using default HTML
func (t *DefaultTemplate) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	if atom == nil {
		return nil, nil
	}

	// Convert *omni.Atom to omni.AtomInterface
	atomInterface, ok := interface{}(atom).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Helper function to get string property with default
	getProp := func(key, def string) string {
		if atomInterface.Has(key) {
			return atomInterface.Get(key)
		}
		return def
	}

	switch atom.GetType() {
	case "container":
		tag := hb.NewDiv()
		for _, child := range atomInterface.ChildrenGet() {
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			tag.AddChild(childTag)
		}
		return tag, nil

	case "text":
		return hb.NewTag("span").Text(getProp("text", "")), nil

	case "image":
		img := hb.NewTag("img")
		if src := getProp("src", ""); src != "" {
			img.Attr("src", src)
		}
		if alt := getProp("alt", getProp("text", "")); alt != "" {
			img.Attr("alt", alt)
		}
		return img, nil

	case "button":
		text := getProp("text", "")
		button := hb.NewButton().Text(text)
		if href := getProp("href", ""); href != "" {
			button = hb.NewTag("a").Attr("href", href).Text(text)
		}
		return button, nil

	case "link":
		href := getProp("href", "#")
		text := getProp("text", "")
		link := hb.NewTag("a").Attr("href", href).Text(text)

		// Add children if any
		for _, child := range atomInterface.ChildrenGet() {
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			link.AddChild(childTag)
		}
		return link, nil

	case "menu", "menuItem":
		list := hb.NewTag("ul")
		for _, child := range atomInterface.ChildrenGet() {
			item := hb.NewTag("li")
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			item.AddChild(childTag)
			list.AddChild(item)
		}
		return list, nil

	case "icon":
		icon := hb.NewTag("i")
		switch getProp("name", "") {
		case "menu":
			icon.Class("fas fa-bars")
		default:
			icon.Class("fas fa-circle")
		}
		return icon, nil

	default:
		return hb.NewTag("div").Text(getProp("text", "")), nil
	}
}

// RenderDashboard renders a complete dashboard from Omni atoms
func (t *DefaultTemplate) RenderDashboard(dashboard *omni.Atom) (string, error) {
	templateName := dashboard.Get("template")
	if templateName == "" {
		templateName = "default"
	}

	// Add template class to body
	bodyAttrs := map[string]string{
		"class": fmt.Sprintf("template-%s", templateName),
	}

	// Only try to get color scheme if dashboard implements DashboardRenderer
	if d, ok := interface{}(dashboard).(DashboardRenderer); ok {
		if t.isDarkColorScheme(d) {
			bodyAttrs["data-bs-theme"] = "dark"
		}
	}

	// Convert *omni.Atom to omni.AtomInterface
	dashboardInterface, ok := interface{}(dashboard).(omni.AtomInterface)
	if !ok {
		return "", fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Get children
	children := dashboardInterface.ChildrenGet()
	if len(children) < 3 {
		return "", fmt.Errorf("dashboard must have at least 3 children (header, content, footer)")
	}

	// Render header if present
	headerTag := hb.NewTag("header")
	headerAtom, err := toAtom(children[0])
	if err == nil {
		header, err := t.RenderAtom(headerAtom)
		if err == nil && header != nil {
			headerTag.AddChild(header)
		}
	}

	// Render content if present
	contentTag := hb.NewTag("main")
	contentAtom, err := toAtom(children[1])
	if err == nil {
		content, err := t.RenderAtom(contentAtom)
		if err == nil && content != nil {
			contentTag.AddChild(content)
		}
	}

	// Render footer if present
	footerTag := hb.NewTag("footer")
	if len(children) > 2 {
		footerAtom, err := toAtom(children[2])
		if err == nil {
			footer, err := t.RenderAtom(footerAtom)
			if err == nil && footer != nil {
				footerTag.AddChild(footer)
			}
		}
	}

	// Create HTML document structure
	html := hb.NewTag("html")
	headTag := hb.NewTag("head")
	html.AddChild(headTag)

	// Add title
	title := "Dashboard"
	if dashboardInterface.Has("title") {
		title = dashboardInterface.Get("title")
	}
	headTag.AddChild(hb.NewTag("title").Text(title))

	// Add meta charset and viewport
	headTag.AddChild(hb.NewTag("meta").Attr("charset", "UTF-8"))
	headTag.AddChild(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1.0"))

	body := hb.NewTag("body")
	html.AddChild(body)

	// Add header, content, and footer to body
	body.AddChild(headerTag)
	body.AddChild(contentTag)
	body.AddChild(footerTag)

	// Add CSS
	for _, css := range t.GetCSSLinks(false) {
		headTag.AddChild(css)
	}

	// Add JavaScript
	for _, js := range t.GetJSScripts() {
		body.AddChild(js)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		body.AddChild(hb.NewTag("script").Text(customJS))
	}

	// Create HTML document with doctype
	return "<!DOCTYPE html>\n" + html.ToHTML(), nil
}

// RenderHeader renders a basic header for the default template
func (t *DefaultTemplate) RenderHeader(d DashboardRenderer) *hb.Tag {
	header := hb.NewHeader().Class("p-3 bg-light border-bottom")
	container := hb.NewDiv().Class("container-fluid")
	headerInner := hb.NewDiv().Class("d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start")
	logoLink := hb.NewA().Href("/").Class("d-flex align-items-center mb-2 mb-lg-0 text-dark text-decoration-none")
	logoLink.AddChild(hb.NewSpan().Text("Dashboard").Class("fs-4"))
	headerInner.AddChild(logoLink)
	container.AddChild(headerInner)
	header.AddChild(container)
	return header
}

// toAtom converts an interface to *omni.Atom
func toAtom(atom interface{}) (*omni.Atom, error) {
	switch a := atom.(type) {
	case *omni.Atom:
		return a, nil
	case omni.AtomInterface:
		// Try to convert AtomInterface to *omni.Atom if possible
		if atomPtr, ok := a.(*omni.Atom); ok {
			return atomPtr, nil
		}
		// Create a new atom using NewAtom with the same type
		newAtom := atomizer.NewAtom(a.GetType())
		// Copy properties
		allProps := a.GetAll()
		for key, value := range allProps {
			newAtom.Set(key, value)
		}
		// Copy children
		for _, child := range a.ChildrenGet() {
			if child != nil {
				newAtom.ChildAdd(child)
			}
		}
		// Convert back to *omni.Atom if possible
		if atomPtr, ok := newAtom.(*omni.Atom); ok {
			return atomPtr, nil
		}
		return nil, fmt.Errorf("failed to convert to *omni.Atom")
	default:
		return nil, fmt.Errorf("unsupported atom type: %T", atom)
	}
}

// RenderFooter renders a basic footer for the default template
func (t *DefaultTemplate) RenderFooter(d DashboardRenderer) *hb.Tag {
	leftCol := hb.NewDiv().
		Class("col-12 col-md-6").
		AddChild(hb.NewTag("small").Text("© 2025 Dashboard").Class("text-muted"))

	rightCol := hb.NewDiv().
		Class("col-12 col-md-6 text-end").
		AddChild(hb.NewTag("small").Text("Powered by Dashboard").Class("text-muted"))

	row := hb.NewDiv().
		Class("row").
		AddChild(leftCol).
		AddChild(rightCol)

	container := hb.NewDiv().Class("container").AddChild(row)

	footer := hb.NewFooter().Class("mt-auto py-3 bg-light").AddChild(container)
	return footer
}
