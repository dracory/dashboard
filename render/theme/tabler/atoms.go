package tabler

import (
	"fmt"

	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// getChildByType finds the first child of the given type
func getChildByType(atom omni.AtomInterface, childType string) omni.AtomInterface {
	for _, child := range atom.GetChildren() {
		if child.GetType() == childType {
			return child
		}
	}
	return nil
}

// getPropertyString safely gets a string property with a default value
func getPropertyString(atom omni.AtomInterface, key, defaultValue string) string {
	if prop := atom.GetProperty(key); prop != nil {
		return prop.GetValue()
	}
	return defaultValue
}

// RenderAtom renders an Omni atom using Tabler classes and components
// Implements shared.Theme interface
func (t *TablerTheme) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	if atom == nil {
		return nil, fmt.Errorf("atom cannot be nil")
	}
	
	// Convert *omni.Atom to omni.AtomInterface
	atomInterface, ok := interface{}(atom).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}
	
	return t.renderAtom(atomInterface)
}

// toAtom converts an interface{} to *omni.Atom
func toAtom(atom interface{}) (*omni.Atom, error) {
	switch a := atom.(type) {
	case *omni.Atom:
		return a, nil
	case omni.AtomInterface:
		// Try to convert AtomInterface to *omni.Atom if possible
		if atomPtr, ok := a.(*omni.Atom); ok {
			return atomPtr, nil
		}
		// If we can't convert directly, we'll need to create a new Atom
		// using the factory function if available
		return omni.NewAtom(
			a.GetType(),
			a.GetProperties(),
			a.GetChildren()...,
		), nil
	}
	return nil, fmt.Errorf("unsupported atom type: %T, cannot convert to *omni.Atom", atom)
}

// renderAtom is the internal implementation that works with both *omni.Atom and omni.AtomInterface
func (t *TablerTheme) renderAtom(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first since that's what RenderAtom expects
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	// Get the type from the atom
	var atomType string
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		atomType = atomInterface.GetType()
	} else {
		return nil, fmt.Errorf("atom does not implement omni.AtomInterface")
	}

	switch atomType {
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

func (t *TablerTheme) renderContainer(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	container := hb.NewDiv().Class("container")

	// Get children through the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.GetChildren() {
			// Convert child to *omni.Atom for RenderAtom
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			container.AddChild(childTag)
		}
	}

	return container, nil
}

func (t *TablerTheme) renderHeader(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	header := hb.NewTag("header")
	header.Class("navbar navbar-expand-md navbar-light d-print-none")

	// Add container
	container := hb.NewDiv().Class("container")
	header.AddChild(container)

	// Add children to container if atom implements the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.GetChildren() {
			// Convert child to *omni.Atom for RenderAtom
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			container.AddChild(childTag)
		}
	}

	return header, nil
}

func (t *TablerTheme) renderFooter(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	footer := hb.NewTag("footer").Class("footer footer-transparent")
	container := hb.NewDiv().Class("container")
	footer.AddChild(container)

	// Get children through the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.GetChildren() {
			// Convert child to *omni.Atom for RenderAtom
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			container.AddChild(childTag)
		}
	}

	return footer, nil
}

func (t *TablerTheme) renderMenu(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	menu := hb.NewTag("ul").Class("navbar-nav")

	// Get children through the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.GetChildren() {
			// Convert child to *omni.Atom for RenderAtom
			childPtr, err := toAtom(child)
			if err != nil {
				return nil, err
			}
			childTag, err := t.RenderAtom(childPtr)
			if err != nil {
				return nil, err
			}
			menu.AddChild(hb.NewTag("li").Class("nav-item").AddChild(childTag))
		}
	}

	return menu, nil
}

func (t *TablerTheme) renderMenuItem(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	// Get the atom interface for property access
	var atomInterface omni.AtomInterface
	var ok bool
	if atomInterface, ok = interface{}(atomPtr).(omni.AtomInterface); !ok {
		return nil, fmt.Errorf("expected *omni.Atom to implement omni.AtomInterface")
	}

	// Create the menu item
	item := hb.NewTag("li").Class("nav-item")

	// Create link
	href := getPropertyString(atomInterface, "href", "#")
	link := hb.NewTag("a").Attr("href", href).Class("nav-link")

	// Add active class if needed
	if activeProp := atomInterface.GetProperty("active"); activeProp != nil && activeProp.GetValue() == "true" {
		link.Class("active")
	}

	// Add icon if available
	if icon := getPropertyString(atomInterface, "icon", ""); icon != "" {
		link.Child(hb.I().Class(icon + " me-2"))
	}

	// Add text
	if text := getPropertyString(atomInterface, "text", ""); text != "" {
		link.Child(hb.Text(text))
	}

	item.AddChild(link)

	// Add children (submenu)
	for _, child := range atomInterface.GetChildren() {
		if child.GetType() == "menu" {
			submenu, err := t.RenderAtom(child)
			if err != nil {
				return nil, err
			}
			submenu.Class("nav nav-treeview")
			item.AddChild(submenu)
		}
	}

	return item, nil
}

func (t *TablerTheme) renderLink(atom interface{}) (*hb.Tag, error) {
	atomInterface, err := toAtomInterface(atom)
	if err != nil {
		return nil, err
	}
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}
	href := getPropertyString(atomInterface, "href", "#")
	title := getPropertyString(atomInterface, "title", "")
	text := getPropertyString(atomInterface, "text", href)

	a := hb.NewTag("a").Attr("href", href)
	if title != "" {
		a.Attr("title", title)
	}
	a.Child(hb.NewText(text))

	// Add any children (e.g., icons, badges)
	for _, child := range atomInterface.GetChildren() {
		childTag, err := t.RenderAtom(child)
		if err != nil {
			return nil, err
		}
		a.AddChild(childTag)
	}

	return a, nil
}

func (t *TablerTheme) renderButton(atom interface{}) (*hb.Tag, error) {
	atomInterface, err := toAtomInterface(atom)
	if err != nil {
		return nil, err
	}
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}
	button := hb.NewButton()

	// Add button variant class (default to primary)
	variant := "primary"
	if variantProp := atomInterface.GetProperty("variant"); variantProp != nil {
		variant = variantProp.GetValue()
	}
	button.AddClass(fmt.Sprintf("btn btn-%s", variant))

	// Add icon if specified
	if icon := getPropertyString(atomInterface, "icon", ""); icon != "" {
		iconTag := hb.NewTag("i").Class(fmt.Sprintf("ti ti-%s me-2", icon))
		button.AddChild(iconTag)
	}

	// Add text if present
	if text := getPropertyString(atomInterface, "text", ""); text != "" {
		button.Text(text)
	}

	// Add children
	for _, child := range atomInterface.GetChildren() {
		childTag, err := t.RenderAtom(child)
		if err != nil {
			return nil, err
		}
		button.AddChild(childTag)
	}

	return button, nil
}

func (t *TablerTheme) renderImage(atom interface{}) (*hb.Tag, error) {
	atomInterface, err := toAtomInterface(atom)
	if err != nil {
		return nil, err
	}
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}
	img := hb.NewTag("img")

	// Add src attribute
	if srcProp := atomInterface.GetProperty("src"); srcProp != nil {
		img.Attr("src", srcProp.GetValue())
	}

	// Add alt text
	if alt := getPropertyString(atomInterface, "alt", getPropertyString(atomInterface, "text", "")); alt != "" {
		img.Attr("alt", alt)
	}

	// Add responsive image class
	img.AddClass("img-fluid")

	return img, nil
}

func (t *TablerTheme) renderText(atom interface{}) (*hb.Tag, error) {
	atomInterface, err := toAtomInterface(atom)
	if err != nil {
		return nil, err
	}
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}
	text := getPropertyString(atomInterface, "text", "")
	return hb.NewTag("span").Text(text), nil
}

// RenderDashboard renders a complete dashboard from Omni atoms
func (t *TablerTheme) RenderDashboard(dashboard *omni.Atom) (string, error) {
	// Convert *omni.Atom to omni.AtomInterface
	if dashboard == nil {
		return "", fmt.Errorf("dashboard cannot be nil")
	}

	atomInterface, ok := interface{}(dashboard).(omni.AtomInterface)
	if !ok {
		return "", fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}

	return t.renderDashboard(atomInterface)
}

// renderDashboard is the internal implementation that works with both *omni.Atom and omni.AtomInterface
func (t *TablerTheme) renderDashboard(dashboard interface{}) (string, error) {
	// Convert to *omni.Atom first
	dashboardPtr, err := toAtom(dashboard)
	if err != nil {
		return "", err
	}

	// Get the dashboard interface
	var dashboardInterface omni.AtomInterface
	var ok bool
	if dashboardInterface, ok = interface{}(dashboardPtr).(omni.AtomInterface); !ok {
		return "", fmt.Errorf("dashboard does not implement omni.AtomInterface")
	}

	// Start with a page wrapper
	page := hb.NewTag("div").Class("page")

	// Get header and content children
	header := getChildByType(dashboardInterface, "header")
	contentAtom := getChildByType(dashboardInterface, "content")

	// Render header if exists
	if header != nil {
		headerPtr, err := toAtom(header)
		if err != nil {
			return "", fmt.Errorf("error converting header: %v", err)
		}
		headerTag, err := t.RenderAtom(headerPtr)
		if err != nil {
			return "", fmt.Errorf("failed to render header: %w", err)
		}
		page.AddChild(headerTag)

		// Add main content
		content := hb.NewTag("div").Class("page-wrapper")
		page.AddChild(content)

		// Render content if exists
		if contentAtom != nil {
			contentPtr, err := toAtom(contentAtom)
			if err != nil {
				return "", fmt.Errorf("error converting content: %v", err)
			}
			contentTag, err := t.RenderAtom(contentPtr)
			if err != nil {
				return "", fmt.Errorf("failed to render content: %w", err)
			}
			contentTag.Class("content")
			page.AddChild(contentTag)
		}
	}

	return page.ToHTML(), nil
}
