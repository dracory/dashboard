package model

// MenuItem represents a single item in a navigation menu
// This is a consolidated version that includes all fields needed by the dashboard templates
type MenuItem struct {
	// The unique identifier for the menu item
	ID string `json:"id"`

	// The display text of the menu item
	Title string `json:"title"`

	// The URL or path the menu item links to
	URL string `json:"url"`

	// An optional icon to display with the menu item
	Icon string `json:"icon,omitempty"`

	// Optional child menu items for dropdown/submenus
	Children []MenuItem `json:"children,omitempty"`

	// Alternative field name for submenu items (for backward compatibility)
	SubMenu []MenuItem `json:"submenu,omitempty"`

	// Whether the menu item is currently active/selected
	Active bool `json:"active,omitempty"`

	// Alternative field for title (for backward compatibility)
	Text string `json:"text,omitempty"`

	// Badge text to display next to the menu item
	BadgeText string `json:"badge_text,omitempty"`

	// CSS class for the badge
	BadgeClass string `json:"badge_class,omitempty"`

	// Whether to open the link in a new window
	NewWindow bool `json:"new_window,omitempty"`

	// JavaScript to execute when the item is clicked
	OnClick string `json:"on_click,omitempty"`
}

// Text returns the display text of the menu item
// This provides backward compatibility with code that uses the Text field
func (m MenuItem) GetText() string {
	if m.Text != "" {
		return m.Text
	}
	return m.Title
}

// SetText sets the display text of the menu item
// This provides backward compatibility with code that uses the Text field
func (m *MenuItem) SetText(text string) {
	m.Text = text
	if m.Title == "" {
		m.Title = text
	}
}

// AddChild adds a child menu item
func (m *MenuItem) AddChild(child MenuItem) *MenuItem {
	m.Children = append(m.Children, child)
	return m
}

// SetActive sets the active state of the menu item
func (m *MenuItem) SetActive(active bool) *MenuItem {
	m.Active = active
	return m
}
