package atomizer

import (
	omni "github.com/dracory/omni"
)

// Common atom types used in the dashboard
const (
	// Layout atoms
	AtomTypeDashboard = "dashboard"
	AtomTypeHeader    = "header"
	AtomTypeFooter    = "footer"
	AtomTypeContent   = "content"
	AtomTypeMenu      = "menu"
	AtomTypeMenuItem  = "menu_item"

	// UI component atoms
	AtomTypeButton = "button"
	AtomTypeLink   = "link"
	AtomTypeImage  = "image"
	AtomTypeText   = "text"
)

// Common property keys
const (
	PropText     = "text"
	PropHref     = "href"
	PropSrc      = "src"
	PropAlt      = "alt"
	PropClass    = "class"
	PropStyle    = "style"
	PropVariant  = "variant"
	PropSize     = "size"
	PropDisabled = "disabled"
	PropActive   = "active"
	PropTarget   = "target"
	PropRel      = "rel"
)

// Common variants for UI components
const (
	VariantPrimary   = "primary"
	VariantSecondary = "secondary"
	VariantSuccess   = "success"
	VariantDanger    = "danger"
	VariantWarning   = "warning"
	VariantInfo      = "info"
	VariantLight     = "light"
	VariantDark      = "dark"
)

// Common sizes for UI components
const (
	SizeSmall  = "sm"
	SizeMedium = "md"
	SizeLarge  = "lg"
)

// NewAtom creates a new Omni atom with the given type and options
func NewAtom(atomType string, options ...omni.AtomOption) omni.AtomInterface {
	return omni.NewAtom(atomType, options...)
}

// WithText sets the text content of an atom
func WithText(text string) omni.AtomOption {
	return func(a *omni.Atom) {
		a.Set(PropText, text)
	}
}

// WithClass adds a CSS class to an atom
func WithClass(className string) omni.AtomOption {
	return func(a *omni.Atom) {
		a.Set(PropClass, className)
	}
}

// WithVariant sets the variant of a UI component
func WithVariant(variant string) omni.AtomOption {
	return func(a *omni.Atom) {
		a.Set(PropVariant, variant)
	}
}

// WithChildren adds child atoms to a parent atom
func WithChildren(children ...*omni.Atom) omni.AtomOption {
	return func(a *omni.Atom) {
		for _, child := range children {
			if child != nil {
				a.ChildAdd(child)
			}
		}
	}
}
