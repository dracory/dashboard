package components

import (
	"github.com/gouniverse/hb"
	"github.com/samber/lo"
	"github.com/spf13/cast"
)

// GridConfig holds configuration for the Grid component
type GridConfig struct {
	// Columns is the list of columns in the grid
	Columns []GridColumn

	// Class is the CSS class for the grid
	Class string

	// Style is the inline CSS style for the grid
	Style string
}

// GridColumn represents a column in the grid
type GridColumn struct {
	// Content is the HTML content of the column
	Content string

	// Width is the width of the column (1-12)
	Width int

	// Class is the CSS class for the column
	Class string

	// Style is the inline CSS style for the column
	Style string
}

// NewGrid creates a new grid component and returns it as hb.TagInterface
func NewGrid(config GridConfig) hb.TagInterface {
	// Create row
	row := hb.Div().
		Class("row").
		ClassIf(!lo.IsEmpty(config.Class), config.Class).
		StyleIf(!lo.IsEmpty(config.Style), config.Style)

	// Add columns
	for _, column := range config.Columns {
		// Default width is 12 (full width) if not specified
		width := column.Width
		if width <= 0 || width > 12 {
			width = 12
		}

		// Create column
		col := hb.Div().
			Class("col-md-"+cast.ToString(width)).
			ClassIf(!lo.IsEmpty(column.Class), column.Class).
			StyleIf(!lo.IsEmpty(column.Style), column.Style).
			HTML(column.Content)

		row.AddChild(col)
	}

	return row
}

// NewGridWithColumns creates a new grid with the specified number of equal-width columns
func NewGridWithColumns(numColumns int, contents []string) hb.TagInterface {
	if numColumns <= 0 {
		numColumns = 1
	}

	// Calculate column width
	width := 12 / numColumns

	// Create columns
	columns := []GridColumn{}
	for i, content := range contents {
		if i >= numColumns {
			break
		}

		columns = append(columns, GridColumn{
			Content: content,
			Width:   width,
		})
	}

	// Create grid
	return NewGrid(GridConfig{
		Columns: columns,
	})
}
