package render

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
)

// RenderFooter generates the footer HTML for the dashboard
func RenderFooter(d model.DashboardRenderer) *hb.Tag {
	// Create the "Powered by Tabler" link
	tablerLink := hb.A().
		Href("https://github.com/tabler/tabler").
		Target("_blank").
		Class("link-secondary").
		Rel("noopener").
		Text("Powered by Tabler")

	// Create the list item containing the link
	tablerListItem := hb.Li().
		Class("list-inline-item").
		Child(tablerLink)

	// Create the list for the right side
	rightList := hb.Ul().
		Class("list-inline list-inline-dots mb-0").
		Child(tablerListItem)

	// Create the right column
	rightColumn := hb.Div().
		Class("col-lg-auto ms-lg-auto").
		Child(rightList)

	// Create the copyright list item
	copyrightListItem := hb.Li().
		Class("list-inline-item").
		Text("Copyright © 2025")

	// Create the list for the left side
	leftList := hb.Ul().
		Class("list-inline list-inline-dots mb-0").
		Child(copyrightListItem)

	// Create the left column
	leftColumn := hb.Div().
		Class("col-12 col-lg-auto mt-3 mt-lg-0").
		Child(leftList)

	// Create the row containing both columns
	row := hb.Div().
		Class("row text-center align-items-center flex-row-reverse").
		Child(rightColumn).
		Child(leftColumn)

	// Create the container
	container := hb.Div().
		Class("container-xl").
		Child(row)

	// Create the footer
	footer := hb.Footer().
		Class("footer footer-transparent d-print-none").
		Child(container)

	return footer
}
