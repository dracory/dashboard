package types

// Alert represents a notification message to display to the user
type Alert struct {
	Type    string // Alert type (e.g., "success", "danger", "warning", "info")
	Message string // The message to display
}
