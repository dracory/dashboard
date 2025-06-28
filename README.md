# Dracory Dashboard

A modern, responsive dashboard implementation using the [Tabler](https://github.com/tabler/tabler) template. This dashboard is designed to be easy to use, customizable, and feature-rich.

## Features

- **Modern UI**: Built on top of the Tabler template for a clean, modern look
- **Responsive Design**: Works on all screen sizes and devices
- **Theme Switching**: Support for light and dark themes
- **Customizable Navigation**: Configure main menu, user menu, and quick access menu
- **Flexible Layout**: Various layout components (cards, grids, tabs, shadow boxes)
- **User Management**: Built-in user authentication UI components

## Installation

```bash
go get github.com/dracory/dashboard
```

## Usage

### Basic Example

```go
package main

import (
    "net/http"
    
    "github.com/dracory/dashboard"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Create dashboard configuration
        config := dashboard.Config{
            Content: "<h1>Welcome to Dashboard</h1>",
            MenuItems: []dashboard.MenuItem{
                {
                    ID:   "home",
                    Icon: "ti ti-home",
                    Text: "Home",
                    URL:  "/",
                    Active: true,
                },
                {
                    ID:   "settings",
                    Icon: "ti ti-settings",
                    Text: "Settings",
                    URL:  "/settings",
                },
            },
        }
        
        // Create dashboard from config
        dash := dashboard.NewFromConfig(config)
        
        // Render dashboard
        html := dash.ToHTML()
        
        // Write response
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte(html))
    })
    
    http.ListenAndServe(":8080", nil)
}
```

### Components

The dashboard includes several UI components:

#### Card Component

```go
card := components.NewCard(components.CardConfig{
    Title:   "Card Title",
    Content: "<p>Card content</p>",
    Margin:  15,
})
```

#### Grid Component

```go
grid := components.NewGrid(components.GridConfig{
    Columns: []components.GridColumn{
        {Content: "Column 1 content", Width: 4},
        {Content: "Column 2 content", Width: 4},
        {Content: "Column 3 content", Width: 4},
    },
})
```

#### Tab Component

```go
tabs := components.NewTab(components.TabConfig{
    Tabs: []components.Tab{
        {
            ID:      "tab1",
            Title:   "Tab 1",
            Content: "<p>Tab 1 content</p>",
            Active:  true,
        },
        {
            ID:      "tab2",
            Title:   "Tab 2",
            Content: "<p>Tab 2 content</p>",
        },
    },
})
```

#### Shadow Box Component

```go
shadowBox := components.NewShadowBox(components.ShadowBoxConfig{
    Content:     "<p>Content in a shadow box</p>",
    ShadowLevel: 2,
    Padding:     15,
})
```

## Configuration Options

The dashboard can be configured with the following options:

- **Content**: The main content to display in the dashboard
- **FaviconURL**: URL for the favicon
- **LogoImageURL**: URL for the logo image
- **LogoRawHtml**: Raw HTML for the logo (overrides LogoImageURL if set)
- **LogoRedirectURL**: URL to redirect to when the logo is clicked
- **MenuItems**: Array of menu items for the main menu
- **MenuShowText**: Whether to show text for menu items
- **MenuType**: Type of menu (MENU_TYPE_OFFCANVAS or MENU_TYPE_MODAL)
- **NavbarBackgroundColorMode**: Background color mode for the navbar (light or dark)
- **NavbarBackgroundColor**: Background color for the navbar
- **NavbarTextColor**: Text color for the navbar
- **LoginURL**: URL for the login page
- **RegisterURL**: URL for the registration page
- **QuickAccessMenu**: Array of menu items for the quick access menu
- **User**: User information for the user dropdown
- **UserMenu**: Array of menu items for the user dropdown
- **ThemeName**: Name of the theme to use (light or dark)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
