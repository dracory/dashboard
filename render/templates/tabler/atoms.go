package tabler

import (
	"fmt"

	"github.com/dracory/dashboard/render/atomizer"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// getChildByType finds the first child of the given type
func getChildByType(atom omni.AtomInterface, childType string) omni.AtomInterface {
	for _, child := range atom.ChildrenGet() {
		if child.GetType() == childType {
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

// RenderAtom renders an Omni atom using Tabler classes and components
// Implements shared.Template interface
func (t *TablerTemplate) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
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

// toAtom converts an interface to *omni.Atom
func toAtom(atom interface{}) (*omni.Atom, error) {
	switch a := atom.(type) {
	case *omni.Atom:
		return a, nil
	case omni.AtomInterface:
		// Create a new atom using NewAtom with the same type
		newAtomInterface := atomizer.NewAtom(a.GetType())
		// Try to convert to *omni.Atom
		newAtom, ok := newAtomInterface.(*omni.Atom)
		if !ok {
			return nil, fmt.Errorf("failed to convert atom to *omni.Atom")
		}
		// Copy properties
		allProps := a.GetAll()
		for key, value := range allProps {
			newAtom.Set(key, value)
		}
		// Copy children
		for _, child := range a.ChildrenGet() {
			if childAtom, ok := child.(*omni.Atom); ok {
				newAtom.ChildAdd(childAtom)
			}
		}
		return newAtom, nil
	}
	return nil, fmt.Errorf("unsupported atom type: %T, cannot convert to *omni.Atom", atom)
}

// renderAtom is the internal implementation that works with both *omni.Atom and omni.AtomInterface
func (t *TablerTemplate) renderAtom(atom interface{}) (*hb.Tag, error) {
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
		var atomType string
		if atomInterface, ok := atom.(omni.AtomInterface); ok {
			atomType = atomInterface.GetType()
		} else if atomPtr, ok := atom.(*omni.Atom); ok {
			if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
				atomType = atomInterface.GetType()
			}
		}

		tag := hb.NewTag("div").Class(atomType)

		// Get children based on the atom type
		switch a := atom.(type) {
		case omni.AtomInterface:
			for _, child := range a.ChildrenGet() {
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
		case *omni.Atom:
			if atomInterface, ok := interface{}(a).(omni.AtomInterface); ok {
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
			}
		}
		return tag, nil
	}
}

func (t *TablerTemplate) renderContainer(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	container := hb.NewDiv().Class("container")

	// Get children through the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.ChildrenGet() {
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

func (t *TablerTemplate) renderHeader(atom interface{}) (*hb.Tag, error) {
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
		for _, child := range atomInterface.ChildrenGet() {
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

func (t *TablerTemplate) renderFooter(atom interface{}) (*hb.Tag, error) {
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
		for _, child := range atomInterface.ChildrenGet() {
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

func (t *TablerTemplate) renderMenu(atom interface{}) (*hb.Tag, error) {
	// Convert to *omni.Atom first
	atomPtr, err := toAtom(atom)
	if err != nil {
		return nil, err
	}

	menu := hb.NewTag("ul").Class("navbar-nav")

	// Get children through the interface
	if atomInterface, ok := interface{}(atomPtr).(omni.AtomInterface); ok {
		for _, child := range atomInterface.ChildrenGet() {
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

func (t *TablerTemplate) renderMenuItem(atom interface{}) (*hb.Tag, error) {
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
	href := atomInterface.Get("href")
	link := hb.NewTag("a").Attr("href", href).Class("nav-link")

	// Add active class if needed
	if activeProp := atomInterface.Get("active"); activeProp != "" && activeProp == "true" {
		link.Class("active")
	}

	// Add icon if available
	if icon := atomInterface.Get("icon"); icon != "" {
		link.Child(hb.I().Class(icon + " me-2"))
	}

	// Add text
	if text := atomInterface.Get("text"); text != "" {
		link.Child(hb.Text(text))
	}

	item.AddChild(link)

	// Add children (submenu)
	for _, child := range atomInterface.ChildrenGet() {
		if child == nil {
			continue
		}
		if child.GetType() == "menu" {
			submenu, err := t.renderAtom(child)
			if err != nil {
				return nil, fmt.Errorf("failed to render submenu: %w", err)
			}
			submenu.Class("nav nav-treeview")
			item.AddChild(submenu)
		}
	}

	return item, nil
}

func (t *TablerTemplate) renderLink(atom interface{}) (*hb.Tag, error) {
	// Get the atom interface for property access
	var atomInterface omni.AtomInterface
	switch a := atom.(type) {
	case *omni.Atom:
		var ok bool
		atomInterface, ok = interface{}(a).(omni.AtomInterface)
		if !ok {
			return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
		}
	case omni.AtomInterface:
		atomInterface = a
	default:
		return nil, fmt.Errorf("unsupported atom type: %T", atom)
	}

	href := getPropertyString(atomInterface, "href", "#")
	title := getPropertyString(atomInterface, "title", "")
	text := getPropertyString(atomInterface, "text", href)

	a := hb.NewTag("a").
		Attr("href", href).
		Child(hb.NewText(text))

	// Add title if provided
	if title != "" {
		a.Attr("title", title)
	}

	// Add any children (e.g., icons, badges)
	for _, child := range atomInterface.ChildrenGet() {
		if child == nil {
			continue
		}

		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, fmt.Errorf("failed to render link child: %w", err)
		}
		a.AddChild(childTag)
	}

	return a, nil
}

func (t *TablerTemplate) renderButton(atom interface{}) (*hb.Tag, error) {
	// Convert to omni.AtomInterface
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}

	button := hb.NewButton()

	// Add button variant class (default to primary)
	variant := getPropertyString(atomInterface, "variant", "primary")
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
	for _, child := range atomInterface.ChildrenGet() {
		if child == nil {
			continue
		}
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, fmt.Errorf("failed to render button child: %w", err)
		}
		button.AddChild(childTag)
	}

	return button, nil
}

func (t *TablerTemplate) renderImage(atom interface{}) (*hb.Tag, error) {
	// Convert to omni.AtomInterface
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}

	img := hb.NewTag("img")

	// Add src attribute
	src := getPropertyString(atomInterface, "src", "")
	if src != "" {
		img.Attr("src", src)
	}

	// Add alt text (fallback to text property if alt is not set)
	alt := getPropertyString(atomInterface, "alt", "")
	if alt == "" {
		alt = getPropertyString(atomInterface, "text", "")
	}
	if alt != "" {
		img.Attr("alt", alt)
	}

	// Add responsive image class
	img.AddClass("img-fluid")

	return img, nil
}

func (t *TablerTemplate) renderText(atom interface{}) (*hb.Tag, error) {
	// Convert to omni.AtomInterface
	atomInterface, ok := atom.(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("expected omni.AtomInterface, got %T", atom)
	}
	// Get the text property using the helper function
	text := getPropertyString(atomInterface, "text", "")
	return hb.NewTag("span").Text(text), nil
}

// RenderDashboard renders a complete dashboard using the Tabler template
func (t *TablerTemplate) RenderDashboard(d shared.DashboardRenderer) (*hb.Tag, error) {
	// Create the main content container
	content := hb.Div().Class("page")
	
	// Add the header
	header := t.RenderHeader(d)
	content.Child(header)
	
	// Add the main content
	mainContent := hb.Div().Class("page-wrapper")
	
	// Add a container for the content
	container := hb.Div().Class("container-xl")
	
	// Add the dashboard content
	rowDiv := hb.NewDiv().Class("row row-cards")
	colDiv := hb.NewDiv().Class("col-12")
	cardDiv := hb.NewDiv().Class("card")
	cardBody := hb.NewDiv().Class("card-body")
	
	// Add the content to the card body
	contentHTML := hb.NewHTML(d.GetContent())
	cardBody.Child(contentHTML)
	
	// Build the hierarchy
	cardDiv.Child(cardBody)
	colDiv.Child(cardDiv)
	rowDiv.Child(colDiv)
	container.Child(rowDiv)
	
	mainContent.Child(container)
	content.Child(mainContent)
	
	// Add the footer
	footer := t.RenderFooter(d)
	content.Child(footer)
	
	return content, nil
}
