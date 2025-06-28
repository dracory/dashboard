package components

import (
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
	"strconv"
)

// ShadowBoxConfig holds configuration for the ShadowBox component
type ShadowBoxConfig struct {
	// Content is the HTML content of the shadow box
	Content string
	
	// Class is the CSS class for the shadow box
	Class string
	
	// Style is the inline CSS style for the shadow box
	Style string
	
	// Padding is the padding in pixels
	Padding int
	
	// Margin is the margin in pixels
	Margin int
	
	// ShadowLevel is the level of shadow (1-5)
	ShadowLevel int
}

// NewShadowBox creates a new shadow box component and returns it as hb.TagInterface
func NewShadowBox(config ShadowBoxConfig) hb.TagInterface {
	// Determine shadow class based on level
	shadowClass := "shadow"
	if config.ShadowLevel > 0 && config.ShadowLevel <= 5 {
		switch config.ShadowLevel {
		case 1:
			shadowClass = "shadow-sm"
		case 2:
			shadowClass = "shadow"
		case 3:
			shadowClass = "shadow-lg"
		case 4:
			shadowClass = "shadow-xl" // Custom class
		case 5:
			shadowClass = "shadow-2xl" // Custom class
		}
	}
	
	// Create shadow box
	shadowBox := hb.Div().
		Class(shadowClass).
		ClassIf(!lo.IsEmpty(config.Class), config.Class).
		StyleIf(!lo.IsEmpty(config.Style), config.Style).
		StyleIf(config.Padding > 0, "padding: "+strconv.Itoa(config.Padding)+"px;").
		StyleIf(config.Margin > 0, "margin: "+strconv.Itoa(config.Margin)+"px;").
		HTML(config.Content)

	return shadowBox
}
