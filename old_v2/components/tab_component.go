package components

import (
	"github.com/dracory/base/str"
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
)

// TabConfig holds configuration for the Tab component
type TabConfig struct {
	// Tabs is the list of tabs
	Tabs []Tab

	// ID is the unique ID for the tab component
	ID string

	// Class is the CSS class for the tab component
	Class string

	// Style is the inline CSS style for the tab component
	Style string

	// TabsClass is the CSS class for the tabs navigation
	TabsClass string

	// ContentClass is the CSS class for the tab content
	ContentClass string

	// Vertical indicates whether the tabs should be displayed vertically
	Vertical bool

	// Pills indicates whether the tabs should be displayed as pills
	Pills bool
}

// Tab represents a tab in the tab component
type Tab struct {
	// ID is the unique ID for the tab
	ID string

	// Title is the title of the tab
	Title string

	// Content is the HTML content of the tab
	Content string

	// Active indicates whether the tab is active
	Active bool

	// Disabled indicates whether the tab is disabled
	Disabled bool

	// Icon is the icon for the tab
	Icon string
}

// NewTab creates a new tab component and returns it as hb.TagInterface
func NewTab(config TabConfig) hb.TagInterface {
	// Generate unique ID if not provided
	id := config.ID
	if lo.IsEmpty(id) {
		id = "tab-" + str.Random(8)
	}

	// Create container
	container := hb.Div().
		ID(id).
		Class("tab-container").
		ClassIf(!lo.IsEmpty(config.Class), config.Class).
		StyleIf(!lo.IsEmpty(config.Style), config.Style)

	// Create tabs navigation
	navClass := "nav"
	if config.Pills {
		navClass += " nav-pills"
	} else {
		navClass += " nav-tabs"
	}

	if config.Vertical {
		navClass += " flex-column"
	}

	if !lo.IsEmpty(config.TabsClass) {
		navClass += " " + config.TabsClass
	}

	nav := hb.Ul().
		Class(navClass).
		Role("tablist")

	// Create tab content container
	contentClass := "tab-content"
	if !lo.IsEmpty(config.ContentClass) {
		contentClass += " " + config.ContentClass
	}

	content := hb.Div().
		Class(contentClass)

	// Add tabs
	for i, tab := range config.Tabs {
		tabID := tab.ID
		if lo.IsEmpty(tabID) {
			tabID = "tab-" + str.Random(8)
		}

		// Set first tab as active if none is active
		isActive := tab.Active

		// Check if any tab is active (replacing lo.Any)
		if i == 0 {
			anyTabActive := false
			for _, t := range config.Tabs {
				if t.Active {
					anyTabActive = true
					break
				}
			}

			if !anyTabActive {
				isActive = true
			}
		}

		// Create tab nav item
		navItemClass := "nav-item"
		navLinkClass := "nav-link"
		if isActive {
			navLinkClass += " active"
		}
		if tab.Disabled {
			navLinkClass += " disabled"
		}

		icon := ""
		if !lo.IsEmpty(tab.Icon) {
			icon = `<i class="` + tab.Icon + ` me-2"></i>`
		}

		navLink := hb.A().
			Class(navLinkClass).
			ID(tabID+"-tab").
			Data("bs-toggle", "tab").
			Data("bs-target", "#"+tabID).
			Role("tab").
			Aria("controls", tabID).
			Aria("selected", lo.If(isActive, "true").Else("false")).
			HTML(icon + tab.Title)

		navItem := hb.Li().
			Class(navItemClass).
			Role("presentation").
			AddChild(navLink)

		nav.AddChild(navItem)

		// Create tab content pane
		contentPaneClass := "tab-pane fade"
		if isActive {
			contentPaneClass += " show active"
		}

		contentPane := hb.Div().
			Class(contentPaneClass).
			ID(tabID).
			Role("tabpanel").
			Aria("labelledby", tabID+"-tab").
			HTML(tab.Content)

		content.AddChild(contentPane)
	}

	// Add tabs navigation and content to container
	if config.Vertical {
		row := hb.Div().Class("row")

		navCol := hb.Div().Class("col-md-3").AddChild(nav)
		contentCol := hb.Div().Class("col-md-9").AddChild(content)

		row.AddChild(navCol).AddChild(contentCol)
		container.AddChild(row)
	} else {
		container.AddChild(nav).AddChild(content)
	}

	return container
}
