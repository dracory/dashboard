package bootstrap

import (
	"fmt"

	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// getChildByType finds the first child of the given type
func getChildByType(atom omni.AtomInterface, childType string) omni.AtomInterface {
	for _, child := range atom.ChildrenGet() {
		if child != nil && child.GetType() == childType {
			return child
		}
	}
	return nil
}

// getPropertyString safely gets a string property with a default value
func getPropertyString(atom omni.AtomInterface, key, defaultValue string) string {
	if atom.Has(key) {
		return atom.Get(key)
	}
	return defaultValue
}

// RenderAtom renders an Omni atom using Bootstrap 5 classes and components
func (t *BootstrapTemplate) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	// Convert *omni.Atom to omni.AtomInterface
	atomInterface, ok := interface{}(atom).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	return t.renderAtom(atomInterface)
}

func (t *BootstrapTemplate) renderAtom(atom omni.AtomInterface) (*hb.Tag, error) {
	switch atom.GetType() {
	case "container":
		return t.renderContainer(atom)
	case "header":
		return t.renderHeader(atom)
	case "footer":
		return t.renderFooter(atom)
	case "menu":
		return t.renderMenu(atom)
	case "menuItem":
		return t.renderMenuItem(atom)
	case "link":
		return t.renderLink(atom)
	case "button":
		return t.renderButton(atom)
	case "image":
		return t.renderImage(atom)
	case "text":
		return t.renderText(atom)
	default:
		// For unknown atom types, fall back to a div with the atom type as a class
		tag := hb.NewTag("div").Class(atom.GetType())
		for _, child := range atom.ChildrenGet() {
			if child == nil {
				continue
			}
			childTag, err := t.renderAtom(child)
			if err != nil {
				return nil, err
			}
			tag.AddChild(childTag)
		}
		return tag, nil
	}
}

func (t *BootstrapTemplate) renderContainer(atom omni.AtomInterface) (*hb.Tag, error) {
	container := hb.NewDiv().Class("container")
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		container.AddChild(childTag)
	}
	return container, nil
}

func (t *BootstrapTemplate) renderHeader(atom omni.AtomInterface) (*hb.Tag, error) {
	header := hb.NewTag("header").Class("navbar navbar-expand-lg navbar-light bg-light")

	// Add children to header
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		header.AddChild(childTag)
	}

	return header, nil
}

func (t *BootstrapTemplate) renderFooter(atom omni.AtomInterface) (*hb.Tag, error) {
	footer := hb.NewTag("footer").Class("bg-light py-3 mt-5")
	container := hb.NewDiv().Class("container")

	// Add children to footer
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		container.AddChild(childTag)
	}

	footer.AddChild(container)
	return footer, nil
}

func (t *BootstrapTemplate) renderMenu(atom omni.AtomInterface) (*hb.Tag, error) {
	menu := hb.NewTag("ul").Class("navbar-nav me-auto mb-2 mb-lg-0")

	// Add menu items
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		menu.AddChild(childTag)
	}

	return menu, nil
}

func (t *BootstrapTemplate) renderMenuItem(atom omni.AtomInterface) (*hb.Tag, error) {
	item := hb.NewTag("li").Class("nav-item")
	link := hb.NewTag("a").Class("nav-link")

	href := getPropertyString(atom, "href", "#")
	text := getPropertyString(atom, "text", "")

	link.Attr("href", href)

	if text != "" {
		link.Child(hb.NewText(text))
	}

	// Add any children (e.g., icons, badges)
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		link.AddChild(childTag)
	}

	item.AddChild(link)
	return item, nil
}

func (t *BootstrapTemplate) renderLink(atom omni.AtomInterface) (*hb.Tag, error) {
	href := getPropertyString(atom, "href", "#")
	text := getPropertyString(atom, "text", "")

	link := hb.NewTag("a").Attr("href", href).Text(text)

	// Add any children (e.g., icons, badges)
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		link.AddChild(childTag)
	}

	return link, nil
}

func (t *BootstrapTemplate) renderButton(atom omni.AtomInterface) (*hb.Tag, error) {
	button := hb.NewTag("button").Class("btn")

	// Set button style
	style := getPropertyString(atom, "style", "primary")
	button.Class("btn-" + style)

	// Set button size if specified
	size := getPropertyString(atom, "size", "")
	switch size {
	case "sm", "small":
		button.Class("btn-sm")
	case "lg", "large":
		button.Class("btn-lg")
	}

	// Set button text
	text := getPropertyString(atom, "text", "")
	if text != "" {
		button.Child(hb.NewText(text))
	}

	// Add any children (e.g., icons, badges)
	for _, child := range atom.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		button.AddChild(childTag)
	}

	return button, nil
}

func (t *BootstrapTemplate) renderImage(atom omni.AtomInterface) (*hb.Tag, error) {
	img := hb.NewTag("img")

	// Add src and alt attributes using helper functions
	src := getPropertyString(atom, "src", "")
	alt := getPropertyString(atom, "alt", "")
	class := getPropertyString(atom, "class", "")

	if src != "" {
		img.Attr("src", src)
	}

	if alt != "" {
		img.Attr("alt", alt)
	}

	// Add responsive image class by default
	img.Class("img-fluid")

	// Add any additional classes
	if class != "" {
		img.Class(class)
	}

	return img, nil
}

func (t *BootstrapTemplate) renderText(atom omni.AtomInterface) (*hb.Tag, error) {
	tag := hb.NewTag("div").Text(atom.Get("text"))
	return tag, nil
}

// RenderDashboard renders a complete dashboard from Omni atoms
func (t *BootstrapTemplate) RenderDashboard(dashboard *omni.Atom) (string, error) {
	// Convert *omni.Atom to omni.AtomInterface
	dashboardInterface, ok := interface{}(dashboard).(omni.AtomInterface)
	if !ok {
		return "", fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	// Create the HTML document
	html := hb.NewTag("html")

	// Create head section
	head := hb.NewTag("head")
	head.Child(hb.NewTag("meta").Attr("charset", "UTF-8"))
	head.Child(hb.NewTag("meta").Attr("name", "viewport").Attr("content", "width=device-width, initial-scale=1.0"))
	head.Child(hb.NewTag("title").Text("Dashboard"))

	// Add CSS links
	for _, cssLink := range t.GetCSSLinks(false) {
		head.Child(cssLink)
	}

	// Add custom CSS
	if customCSS := t.GetCustomCSS(); customCSS != "" {
		head.Child(hb.NewTag("style").HTML(customCSS))
	}

	html.Child(head)

	// Create body section
	body := hb.NewTag("body")

	// Create the main container
	container := hb.NewTag("div").Class("container-fluid").Style("min-height: 100vh; display: flex; flex-direction: column;")

	// Add header if exists
	header := getChildByType(dashboardInterface, "header")
	if header != nil {
		headerTag, err := t.renderAtom(header)
		if err != nil {
			return "", fmt.Errorf("error rendering header: %w", err)
		}
		container.Child(headerTag)
	}

	// Create main content area
	mainContent := hb.NewTag("main").Class("flex-grow-1 py-3")

	// Add main content children
	for _, child := range dashboard.ChildrenGet() {
		if child.GetType() != "header" && child.GetType() != "footer" {
			childTag, err := t.renderAtom(child)
			if err != nil {
				return "", fmt.Errorf("error rendering child %s: %w", child.GetType(), err)
			}
			mainContent.Child(childTag)
		}
	}

	container.Child(mainContent)

	// Add footer if exists
	footer := getChildByType(dashboardInterface, "footer")
	if footer != nil {
		footerTag, err := t.renderAtom(footer)
		if err != nil {
			return "", fmt.Errorf("error rendering footer: %w", err)
		}
		container.Child(footerTag)
	}

	body.Child(container)

	// Add JavaScript files
	for _, jsScript := range t.GetJSScripts() {
		body.Child(jsScript)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		body.Child(hb.NewTag("script").HTML(customJS))
	}

	html.Child(body)

	// Return the complete HTML document
	return "<!DOCTYPE html>\n" + html.ToHTML(), nil
}
