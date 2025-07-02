package types

// Action represents an action button in the header
type Action struct {
	ID      string // Optional ID for the button
	Title   string // Button text
	Icon    string // Icon name (without the 'ti ti-' prefix)
	Primary bool   // Whether this is a primary action
	OnClick string // JavaScript to execute on click
}
