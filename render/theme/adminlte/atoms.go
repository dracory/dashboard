package adminlte

import (
	"fmt"

	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// getPropertyString safely gets a string property with a default value
func getPropertyString(atom omni.AtomInterface, key, defaultValue string) string {
	if prop := atom.GetProperty(key); prop != nil {
		if str, ok := prop.(interface{ String() string }); ok {
			return str.String()
		}
	}
	return defaultValue
}

// getChildByType finds the first child atom of the specified type
func getChildByType(atom omni.AtomInterface, atomType string) omni.AtomInterface {
	if atom == nil {
		return nil
	}
	for _, child := range atom.GetChildren() {
		if child.GetType() == atomType {
			return child
		}
	}
	return nil
}

package adminlte

import (
	"github.com/dracory/dashboard/render/theme/shared"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// renderAtom is the internal implementation that works with omni.AtomInterface
func (t *AdminLTETheme) renderAtom(atom omni.AtomInterface) (*hb.Tag, error) {
	if atom == nil {
		return nil, fmt.Errorf("atom cannot be nil")
	}

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
	case "card":
		return t.renderCard(atom)
	default:
		// For unknown atom types, fall back to a div with the atom type as a class
		tag := hb.NewTag("div").Class(atom.GetType())
		for _, child := range atom.GetChildren() {
			childTag, err := t.RenderAtom(child)
			if err != nil {
				return nil, err
			}
			tag.AddChild(childTag)
		}
		return tag, nil
	}
}

func (t *AdminLTETheme) renderContainer(atom omni.AtomInterface) (*hb.Tag, error) {
	container := hb.NewDiv().Class("container")
	for _, child := range atom.GetChildren() {
		childTag, err := t.RenderAtom(child)
		if err != nil {
			return nil, err
		}
		container.AddChild(childTag)
	}
	return container, nil
}

func (t *AdminLTETheme) renderHeader(atom omni.AtomInterface) (*hb.Tag, error) {
	header := hb.NewTag("header").Class("main-header")
	
	// Navbar
	navbar := hb.NewTag("nav").Class("navbar navbar-expand navbar-white navbar-light")
	
	// Left navbar links
	leftNav := hb.NewTag("ul").Class("navbar-nav")
	leftNav.Child(hb.NewTag("li").Class("nav-item").
		Child(hb.NewTag("a").Class("nav-link").
			Attr("data-widget", "pushmenu").
			Attr("href", "#").
			Child(hb.I().Class("fas fa-bars"))))
	
	navbar.Child(leftNav)
	
	// Add children to header
	for _, child := range atom.GetChildren() {
		childTag, err := t.RenderAtom(child)
		if err != nil {
			return nil, err
		}
		header.AddChild(childTag)
	}
	
	header.AddChild(navbar)
	return header, nil
}

func (t *AdminLTETheme) renderFooter(atom omni.AtomInterface) (*hb.Tag, error) {
	footer := hb.NewFooter().Class("main-footer")
	
	// Add children to footer
	for _, child := range atom.GetChildren() {
		childTag, err := t.RenderAtom(child)
		if err != nil {
			return nil, err
		}
		footer.AddChild(childTag)
	}
	
	return footer, nil
}

func (t *AdminLTETheme) renderMenu(atom omni.AtomInterface) (*hb.Tag, error) {
	menu := hb.NewNav().Class("nav nav-pills nav-sidebar flex-column")
	menu.Attr("data-widget", "treeview")
	menu.Attr("role", "menu")
	menu.Attr("data-accordion", "false")

	// Add menu items
	for _, child := range atom.GetChildren() {
		item, err := t.renderMenuItem(child)
		if err != nil {
			return nil, err
		}
		menu.AddChild(item)
	}
	
	return menu, nil
}

func (t *AdminLTETheme) renderMenuItem(atom omni.AtomInterface) (*hb.Tag, error) {
	item := hb.NewTag("li").Class("nav-item")
	link := hb.NewTag("a").Class("nav-link")
	
	// Set href if available
	if href := getPropertyString(atom, "href", "#"); href != "" {
		link.Attr("href", href)
	}
	
	// Set active class if needed
	if active, _ := atom.GetProperty("active").Bool(); active {
		link.Class("active")
	}
	
	// Add icon if available
	if icon := getPropertyString(atom, "icon", ""); icon != "" {
		link.Child(hb.I().Class(icon))
	}
	
	// Add text if available
	if text := getPropertyString(atom, "text", ""); text != "" {
		link.Child(hb.Text(" " + text))
	}
	
	item.Child(link)
	
	// Handle submenu
	var submenu *hb.Tag
	for _, child := range atom.GetChildren() {
		if child.GetType() == "menu" {
			submenu, _ = t.RenderAtom(child)
			submenu.Class("nav nav-treeview")
			item.AddChild(submenu)
		} else {
			childTag, err := t.RenderAtom(child)
			if err != nil {
				return nil, err
			}
			item.AddChild(childTag)
		}
	}
	
	return item, nil
}

func (t *AdminLTETheme) renderLink(atom omni.AtomInterface) (*hb.Tag, error) {
	href := getPropertyString(atom, "href", "#")
	link := hb.NewLink().Href(href)

	// Add text if exists
	if text := getPropertyString(atom, "text", ""); text != "" {
		link.Child(hb.Text(text))
	}

	// Add classes if any
	if class := getPropertyString(atom, "class", ""); class != "" {
		link.Class(class)
	}

	return link, nil
}

func (t *AdminLTETheme) renderButton(atom omni.AtomInterface) (*hb.Tag, error) {
	button := hb.NewButton()

	// Add text if exists
	if text := getPropertyString(atom, "text", ""); text != "" {
		button.Child(hb.Text(text))
	}

	// Add classes if any
	if class := getPropertyString(atom, "class", ""); class != "" {
		button.Class(class)
	}

	// Add type if specified, default to button
	buttonType := getPropertyString(atom, "type", "button")
	button.Attr("type", buttonType)

	return button, nil
}

func (t *AdminLTETheme) renderImage(atom omni.AtomInterface) (*hb.Tag, error) {
	src := getPropertyString(atom, "src", "")
	img := hb.NewImage().Src(src)

	// Add alt text if exists
	if alt := getPropertyString(atom, "alt", ""); alt != "" {
		img.Attr("alt", alt)
	}

	// Add classes if any
	if class := getPropertyString(atom, "class", ""); class != "" {
		img.Class(class)
	}

	return img, nil
}

func (t *AdminLTETheme) renderText(atom omni.AtomInterface) (*hb.Tag, error) {
	text := getPropertyString(atom, "text", "")
	tag := hb.NewTag("span").Child(hb.Text(text))

	// Add classes if any
	if class := getPropertyString(atom, "class", ""); class != "" {
		tag.Class(class)
	}

	return tag, nil
}

func (t *AdminLTETheme) renderCard(atom omni.AtomInterface) (*hb.Tag, error) {
	card := hb.NewTag("div").Class("card")
	
	// Add header if exists
	header := getChildByType(atom, "header")
	if header != nil {
		headerTag, err := t.renderAtom(header)
		if err != nil {
			return nil, err
		}
		card.AddChild(headerTag)
	}
	
	// Add body if exists
	body := getChildByType(atom, "body")
	if body != nil {
		bodyTag, err := t.renderAtom(body)
		if err != nil {
			return nil, err
		}
		card.AddChild(bodyTag)
	}
	
	// Add footer if exists
	footer := getChildByType(atom, "footer")
func (t *AdminLTETheme) renderDashboard(dashboard omni.AtomInterface) (string, error) {
	// Create HTML document
	doc := hb.NewHTML()
	doc.DocType()

	// Add HTML head
	head := hb.NewTag("head")
	doc.AddChild(head)

	// Add meta charset
	head.Child(hb.NewTag("meta").Attr("charset", "UTF-8"))

	// Add viewport meta tag
	head.Child(hb.NewTag("meta").
		Attr("name", "viewport").
		Attr("content", "width=device-width, initial-scale=1.0"))

	// Add title
	title := getPropertyString(dashboard, "title", "AdminLTE Dashboard")
	head.Child(hb.NewTag("title").Text(title))

	// Add CSS links
	for _, css := range t.GetCSSLinks(false) {
		head.Child(css)
	}

	// Add custom CSS
	if customCSS := t.GetCustomCSS(); customCSS != "" {
		style := hb.NewTag("style").Text(customCSS)
		head.Child(style)
	}

	// Start body
	body := hb.NewTag("body").Class("hold-transition sidebar-mini")
	doc.AddChild(body)

	// Add header
	header := hb.NewTag("header").Class("main-header")
	body.AddChild(header)

	// Add content wrapper
	contentWrapper := hb.NewDiv().Class("content-wrapper")
	body.AddChild(contentWrapper)

	// Render dashboard content
	content, err := t.renderAtom(dashboard)
	if err != nil {
		return "", err
	}
	contentWrapper.AddChild(content)

	// Add footer
	footer := hb.NewTag("footer").Class("main-footer")
	body.AddChild(footer)

	// Add JavaScript
	for _, js := range t.GetJSScripts() {
		body.Child(js)
	}

	// Add custom JavaScript
	if customJS := t.GetCustomJS(); customJS != "" {
		script := hb.NewTag("script").Text(customJS)
		body.Child(script)
	}

	return doc.String(), nil
}
