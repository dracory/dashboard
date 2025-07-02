package types

// Modal represents a Bootstrap/Tabler modal dialog
type Modal struct {
	ID          string // Unique ID for the modal
	Title       string // Modal title
	Content     string // Modal content (HTML)
	Size        string // Modal size (sm, lg, xl, or empty for default)
	Footer      string // Modal footer content (HTML)
	CloseButton bool   // Whether to show the close button (default: true)
}
