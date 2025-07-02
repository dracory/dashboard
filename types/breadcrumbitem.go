package types

// BreadcrumbItem represents a single item in the breadcrumb navigation
type BreadcrumbItem struct {
	Title string // Display text for the breadcrumb item
	URL   string // URL for the breadcrumb link (empty for current/last item)
}
