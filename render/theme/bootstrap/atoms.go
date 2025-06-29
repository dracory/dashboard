package bootstrap

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

// RenderAtom renders an Omni atom using Bootstrap 5 classes and components
func (t *BootstrapTheme) RenderAtom(atom *omni.Atom) (*hb.Tag, error) {
	// Convert *omni.Atom to omni.AtomInterface
	if atom == nil {
		return nil, fmt.Errorf("atom cannot be nil")
	}
	
	atomInterface, ok := interface{}(atom).(omni.AtomInterface)
	if !ok {
		return nil, fmt.Errorf("failed to convert *omni.Atom to omni.AtomInterface")
	}
	
	// Call the implementation that works with omni.AtomInterface
	return t.renderAtom(atomInterface)
}

func (t *BootstrapTheme) renderAtom(atom omni.AtomInterface) (*hb.Tag, error) {
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
		for _, child := range atom.GetChildren() {
			childTag, err := t.renderAtom(child)
			if err != nil {
				return nil, err
			}
			tag.AddChild(childTag)
		}
		return tag, nil
	}
}

func (t *BootstrapTheme) renderContainer(atom omni.AtomInterface) (*hb.Tag, error) {
	container := hb.NewDiv().Class("container")
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		container.AddChild(childTag)
	}
	return container, nil
}

func (t *BootstrapTheme) renderHeader(atom omni.AtomInterface) (*hb.Tag, error) {
	header := hb.NewTag("header").Class("navbar navbar-expand-lg navbar-light bg-light")

	// Add children to header
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		header.AddChild(childTag)
	}

	return header, nil
}

func (t *BootstrapTheme) renderFooter(atom omni.AtomInterface) (*hb.Tag, error) {
	footer := hb.NewTag("footer").Class("bg-light py-3 mt-5")
	container := hb.NewDiv().Class("container")

	// Add children to footer
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		container.AddChild(childTag)
	}

	footer.AddChild(container)
	return footer, nil
}

func (t *BootstrapTheme) renderMenu(atom omni.AtomInterface) (*hb.Tag, error) {
	menu := hb.NewTag("ul").Class("navbar-nav me-auto mb-2 mb-lg-0")

	// Add menu items
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		menu.AddChild(childTag)
	}

	return menu, nil
}

func (t *BootstrapTheme) renderMenuItem(atom omni.AtomInterface) (*hb.Tag, error) {
	item := hb.NewTag("li").Class("nav-item")
	link := hb.NewTag("a").Class("nav-link")

	href := getPropertyString(atom, "href", "#")
	text := getPropertyString(atom, "text", "")

	link.Attr("href", href)

	if text != "" {
		link.Child(hb.NewText(text))
	}

	// Add any children (e.g., icons, badges)
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		link.AddChild(childTag)
	}

	item.AddChild(link)
	return item, nil
}

func (t *BootstrapTheme) renderLink(atom omni.AtomInterface) (*hb.Tag, error) {
	href := getPropertyString(atom, "href", "#")
	text := getPropertyString(atom, "text", "")

	link := hb.NewTag("a").Attr("href", href).Text(text)

	// Add any children (e.g., icons, badges)
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		link.AddChild(childTag)
	}

	return link, nil
}

func (t *BootstrapTheme) renderButton(atom omni.AtomInterface) (*hb.Tag, error) {
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
	for _, child := range atom.GetChildren() {
		childTag, err := t.renderAtom(child)
		if err != nil {
			return nil, err
		}
		button.AddChild(childTag)
	}

	return button, nil
}

func (t *BootstrapTheme) renderImage(atom omni.AtomInterface) (*hb.Tag, error) {
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

func (t *BootstrapTheme) renderText(atom *omni.Atom) (*hb.Tag, error) {
	return hb.NewTag("span").Text(atom.Text), nil
}
