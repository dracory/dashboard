package shared

import (
	"fmt"

	"github.com/dracory/dashboard/model/interfaces"
	"github.com/dracory/dashboard/render/atomizer"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// DashboardRenderer defines the interface for dashboard renderers
type DashboardRenderer = interfaces.DashboardRenderer

// Template defines the interface for dashboard templates
type Template = interfaces.Template

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
func (t *DefaultTemplate) RenderAtom(a *omni.Atom) (*hb.Tag, error) {
	if a == nil {
		return nil, nil
	}

	// Convert *omni.Atom to omni.AtomInterface
	aInterface, ok := interface{}(a).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Helper function to get string property with default
	getProp := func(key, def string) string {
		if aInterface.Has(key) {
			return aInterface.Get(key)
		}
		return def
	}

	switch a.GetType() {
	case "container":
		tag := hb.NewDiv()
		for _, child := range aInterface.ChildrenGet() {
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
		for _, child := range aInterface.ChildrenGet() {
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
		for _, child := range aInterface.ChildrenGet() {
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

// RenderDashboard renders a dashboard using the template's layout
func (t *DefaultTemplate) RenderDashboard(d DashboardRenderer) (*hb.Tag, error) {
	templateName := "default"

	// Add template class to body
	bodyAttrs := map[string]string{
		"class": fmt.Sprintf("template-%s", templateName),
	}

	// Only try to get color scheme if dashboard implements DashboardRenderer
	if t.isDarkColorScheme(d) {
		bodyAttrs["data-bs-theme"] = "dark"
	}

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
	cssLinks := t.GetCSSLinks(false)
	for _, link := range cssLinks {
		head.Child(link)
	}

	// Create the body section
	body := hb.NewTag("body").
		Attrs(bodyAttrs).
		Child(t.RenderHeader(d)).
		Child(hb.NewTag("main")).
		Child(t.RenderFooter(d))

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
		AddChild(hb.NewTag("small").Text(" 2025 Dashboard").Class("text-muted"))

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
