package components

import (
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
	"strconv"
)

// CardConfig holds configuration for the Card component
type CardConfig struct {
	// Content is the HTML content of the card body
	Content string
	
	// Title is the title of the card header
	Title string
	
	// HeaderClass is the CSS class for the card header
	HeaderClass string
	
	// BodyClass is the CSS class for the card body
	BodyClass string
	
	// CardClass is the CSS class for the card
	CardClass string
	
	// Style is the inline CSS style for the card
	Style string
	
	// Padding is the padding in pixels
	Padding int
	
	// Margin is the margin in pixels
	Margin int
}

// NewCard creates a new card component and returns it as hb.TagInterface
func NewCard(config CardConfig) hb.TagInterface {
	// Create card
	card := hb.Div().
		Class("card").
		ClassIf(!lo.IsEmpty(config.CardClass), config.CardClass).
		StyleIf(!lo.IsEmpty(config.Style), config.Style).
		StyleIf(config.Padding > 0, "padding: "+strconv.Itoa(config.Padding)+"px;").
		StyleIf(config.Margin > 0, "margin: "+strconv.Itoa(config.Margin)+"px;")

	// Add header if title is provided
	if !lo.IsEmpty(config.Title) {
		headerClass := "card-header"
		if !lo.IsEmpty(config.HeaderClass) {
			headerClass = config.HeaderClass
		}
		
		header := hb.Div().
			Class(headerClass).
			HTML(config.Title)
		
		card.AddChild(header)
	}

	// Add body with content
	bodyClass := "card-body"
	if !lo.IsEmpty(config.BodyClass) {
		bodyClass = config.BodyClass
	}
	
	body := hb.Div().
		Class(bodyClass).
		HTML(config.Content)
	
	card.AddChild(body)

	return card
}
