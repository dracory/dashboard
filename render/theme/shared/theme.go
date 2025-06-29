package shared

import (
	"fmt"

	"github.com/dracory/dashboard/model"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// DashboardRenderer defines the interface for dashboard renderers
type DashboardRenderer = model.DashboardRenderer

// Theme defines the interface for dashboard themes
type Theme interface {
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

// DefaultTheme is a basic theme implementation that can be used as a fallback
type DefaultTheme struct{}

// GetName returns the name of the default theme
func (t *DefaultTheme) GetName() string {
	return "default"
}

// GetCSSLinks returns no CSS links for the default theme
func (t *DefaultTheme) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	// Return basic CSS reset as a fallback
	return []*hb.Tag{
		hb.Style(`
			* { box-sizing: border-box; margin: 0; padding: 0; }
			body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; line-height: 1.6; }
		`),
	}
}

// GetJSScripts returns no JavaScript for the default theme
func (t *DefaultTheme) GetJSScripts() []*hb.Tag {
	return nil
}

// GetCustomCSS returns empty string as there's no custom CSS for the default theme
func (t *DefaultTheme) GetCustomCSS() string {
	return ""
}

// GetCustomJS returns empty string as there's no custom JavaScript for the default theme
func (t *DefaultTheme) GetCustomJS() string {
	return ""
}

// RenderAtom renders an Omni atom using default HTML
func (t *DefaultTheme) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
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
		if prop := atomInterface.GetProperty(key); prop != nil {
			return prop.GetValue()
		}
		return def
	}

	switch atom.GetType() {
	case "container":
		tag := hb.NewDiv()
		for _, child := range atom.GetChildren() {
			childTag, err := t.RenderAtom(child.(*omni.Atom))
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
		for _, child := range atom.GetChildren() {
			childTag, err := t.RenderAtom(child.(*omni.Atom))
			if err != nil {
				return nil, err
			}
			link.AddChild(childTag)
		}
		return link, nil

	case "menu", "menuItem":
		list := hb.NewTag("ul")
		for _, child := range atom.GetChildren() {
			item := hb.NewTag("li")
			childTag, err := t.RenderAtom(child.(*omni.Atom))
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
func (t *DefaultTheme) RenderDashboard(dashboard *omni.Atom) (string, error) {
	if dashboard == nil {
		return "", fmt.Errorf("dashboard cannot be nil")
	}

	// Convert *omni.Atom to omni.AtomInterface
	dashboardInterface, ok := interface{}(dashboard).(omni.AtomInterface)
	if !ok {
		return "", fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Get children
	children := dashboardInterface.GetChildren()
	if len(children) < 3 {
		return "", fmt.Errorf("dashboard must have at least 3 children (header, content, footer)")
	}

	// Render header if present
	headerTag := hb.NewTag("header")
	header, err := t.RenderAtom(children[0].(*omni.Atom))
	if err == nil && header != nil {
		headerTag.AddChild(header)
	}

	// Render content if present
	contentTag := hb.NewTag("main")
	content, err := t.RenderAtom(children[1].(*omni.Atom))
	if err == nil && content != nil {
		contentTag.AddChild(content)
	}

	// Render footer if present
	footerTag := hb.NewTag("footer")
	if len(children) > 2 {
		footer, err := t.RenderAtom(children[2].(*omni.Atom))
		if err == nil && footer != nil {
			footerTag.AddChild(footer)
		}
	}

	// Create HTML document structure
	html := hb.NewTag("html")
	headTag := hb.NewTag("head")
	html.AddChild(headTag)
	
	// Add title
	title := "Dashboard"
	if prop := dashboardInterface.GetProperty("title"); prop != nil {
		title = prop.GetValue()
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

// RenderHeader renders a basic header for the default theme
func (t *DefaultTheme) RenderHeader(d DashboardRenderer) *hb.Tag {
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

// RenderFooter renders a basic footer for the default theme
func (t *DefaultTheme) RenderFooter(d DashboardRenderer) *hb.Tag {
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

	container := hb.NewDiv().
		Class("container-fluid").
		AddChild(row)

	footer := hb.NewFooter().
		Class("d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top").
		AddChild(container)

	return footer
}
