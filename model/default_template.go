package model

import (
	"fmt"

	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// DefaultTemplate is a basic template implementation that can be used as a fallback
type DefaultTemplate struct{}

// RenderPage renders a complete page with the given content and dashboard renderer
func (t *DefaultTemplate) RenderPage(content string, d DashboardRenderer) (*hb.Tag, error) {
	doc := hb.NewTag("html")

	// Head section
	head := hb.NewTag("head")
	head.Child(hb.NewTag("meta").Attr("charset", "UTF-8"))
	head.Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1.0"))
	head.Child(hb.NewTag("title").Text("Dashboard"))

	// Add favicon if available
	if faviconURL := d.GetFaviconURL(); faviconURL != "" {
		head.Child(hb.NewTag("link").Attr("rel", "icon").Attr("href", faviconURL))
	}

	// Add basic styling
	style := hb.NewTag("style")
	style.Text(`
		body {
			font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
			line-height: 1.5;
			color: #333;
			margin: 0;
			padding: 0;
		}
		.container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 1rem;
		}
		.header {
			padding: 1rem 0;
			border-bottom: 1px solid #eee;
			margin-bottom: 2rem;
		}
		.content {
			min-height: 60vh;
		}
		.footer {
			margin-top: 2rem;
			padding: 1rem 0;
			border-top: 1px solid #eee;
			text-align: center;
			color: #666;
		}
	`)
	head.Child(style)
	doc.Child(head)

	// Body section
	body := hb.NewTag("body")

	// Header
	header := hb.NewTag("header").Class("header")
	headerContainer := hb.NewTag("div").Class("container")

	// Add logo if available
	if logo := d.GetLogoRawHtml(); logo != "" {
		logoLink := hb.NewTag("a").Attr("href", "/")
		logoLink.HTML(logo)
		headerContainer.Child(logoLink)
	} else {
		headerContainer.Child(hb.NewTag("h1").Text("Dashboard"))
	}

	header.Child(headerContainer)
	body.Child(header)

	// Main content
	main := hb.NewTag("main").Class("content")
	mainContainer := hb.NewTag("div").Class("container")
	mainContainer.HTML(content)
	main.Child(mainContainer)
	body.Child(main)

	// Footer
	footer := hb.NewTag("footer").Class("footer")
	footerContainer := hb.NewTag("div").Class("container")
	footerContainer.Text("© 2023 Dashboard. All rights reserved.")
	footer.Child(footerContainer)
	body.Child(footer)

	doc.Child(body)

	return doc, nil
}

// GetName returns the name of the default template
func (t *DefaultTemplate) GetName() string {
	return "default"
}

// GetCSSLinks returns no CSS links for the default template
func (t *DefaultTemplate) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	return []*hb.Tag{}
}

// GetJSScripts returns no JavaScript for the default template
func (t *DefaultTemplate) GetJSScripts() []*hb.Tag {
	return []*hb.Tag{}
}

// GetCustomCSS returns empty string as there's no custom CSS for the default template
func (t *DefaultTemplate) GetCustomCSS() string {
	return ""
}

// GetCustomJS returns empty string as there's no custom JavaScript for the default template
func (t *DefaultTemplate) GetCustomJS() string {
	return ""
}

// RenderHeader renders a basic header for the default template
func (t *DefaultTemplate) RenderHeader(d DashboardRenderer) *hb.Tag {
	header := hb.NewTag("header")
	header.Child(hb.NewTag("h1").Text("Dashboard"))
	return header
}

// RenderFooter renders a basic footer for the default template
func (t *DefaultTemplate) RenderFooter(d DashboardRenderer) *hb.Tag {
	footer := hb.NewTag("footer")
	footer.Child(hb.NewTag("p").Text("© 2023 Dashboard. All rights reserved."))
	return footer
}

// RenderAtom renders an Omni atom using default HTML
func (t *DefaultTemplate) RenderAtom(a *omni.Atom) (*hb.Tag, error) {
	if a == nil {
		return nil, fmt.Errorf("cannot render nil atom")
	}

	// Convert *omni.Atom to omni.AtomInterface
	aInterface, ok := interface{}(a).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Create a container for the atom
	container := hb.NewTag("div").Class("omni-atom")

	// Add the atom's type as a class
	container.Class("omni-" + a.GetType())

	// Handle different atom types
	switch a.GetType() {
	case "container":
		// For containers, render children
		for _, child := range aInterface.ChildrenGet() {
			childTag, err := t.RenderAtom(child.(*omni.Atom))
			if err != nil {
				return nil, fmt.Errorf("error rendering child: %v", err)
			}
			container.Child(childTag)
		}

	case "text":
		// For text atoms, use the content property
		content := aInterface.Get("content")
		container.Text(content)

	case "heading":
		// For headings, use the level and text properties
		level := aInterface.Get("level")
		if level == "" {
			level = "h1"
		}
		text := aInterface.Get("text")
		heading := hb.NewTag(level).Text(text)
		container.Child(heading)

	default:
		// For unknown types, just show the type and properties
		typeSpan := hb.NewTag("div").Style("font-weight:bold").Text("Type: " + a.GetType())
		container.Child(typeSpan)

		// Show properties
		props := hb.NewTag("div").Style("margin-left:1em")
		// Get known properties that might exist on the atom
		knownProps := []string{"id", "class", "style", "content", "text", "level", "href", "src", "alt"}
		for _, key := range knownProps {
			if val := aInterface.Get(key); val != "" {
				propSpan := hb.NewTag("div").Text(fmt.Sprintf("%s: %v", key, val))
				props.Child(propSpan)
			}
		}
		container.Child(props)

		// Show children if any
		if len(aInterface.ChildrenGet()) > 0 {
			childrenDiv := hb.NewTag("div").Style("margin-left:1em;border-left:1px solid #ccc;padding-left:0.5em")
			for _, child := range aInterface.ChildrenGet() {
				childTag, err := t.RenderAtom(child.(*omni.Atom))
				if err != nil {
					return nil, fmt.Errorf("error rendering child: %v", err)
				}
				childrenDiv.Child(childTag)
			}
			container.Child(childrenDiv)
		}
	}

	// Add basic styling for the atom
	style := hb.NewTag("style")
	style.Text(`
		.omni-atom {
			margin: 1rem 0;
			padding: 1rem;
			border: 1px solid #eee;
			border-radius: 4px;
		}
		.omni-atom h1, .omni-atom h2, .omni-atom h3, .omni-atom h4, .omni-atom h5, .omni-atom h6 {
			margin-top: 0;
		}
	`)
	container.Child(style)

	return container, nil
}

// RenderDashboard renders a dashboard using the template's layout
func (t *DefaultTemplate) RenderDashboard(d DashboardRenderer) (*hb.Tag, error) {
	doc := hb.NewTag("html")

	// Head section
	head := hb.NewTag("head")
	head.Child(hb.NewTag("meta").Attr("charset", "UTF-8"))
	head.Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1.0"))
	head.Child(hb.NewTag("title").Text("Dashboard"))

	// Add favicon if available
	if faviconURL := d.GetFaviconURL(); faviconURL != "" {
		head.Child(hb.NewTag("link").Attr("rel", "icon").Attr("href", faviconURL))
	}

	// Add basic styling
	style := hb.NewTag("style")
	style.Text(`
		body {
			font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
			line-height: 1.5;
			color: #333;
			margin: 0;
			padding: 0;
		}
		.container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 1rem;
		}
	`)
	head.Child(style)
	doc.Child(head)

	// Body section
	body := hb.NewTag("body")

	// Add header
	header := t.RenderHeader(d)
	if header != nil {
		body.Child(header)
	}

	// Add main content
	main := hb.NewTag("main")
	main.HTML(d.GetContent())
	body.Child(main)

	// Add footer
	footer := t.RenderFooter(d)
	if footer != nil {
		body.Child(footer)
	}

	doc.Child(body)

	return doc, nil
}

// isDarkColorScheme checks if the color scheme should be dark
func (t *DefaultTemplate) isDarkColorScheme(d DashboardRenderer) bool {
	// Default to light theme
	return false
}
