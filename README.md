# Dracory Dashboard

A modern, responsive dashboard implementation using the [Tabler](https://github.com/tabler/tabler) template. This dashboard is designed to be easy to use, customizable, and feature-rich.

## Features

- **Modern UI**: Built on top of the Tabler template for a clean, modern look
- **Responsive Design**: Works on all screen sizes and devices
- **Theme Switching**: Support for light and dark themes
- **Customizable Navigation**: Configure main menu, user menu, and quick access menu
- **Component-Based Rendering**: Modular rendering system using the gouniverse/hb HTML builder
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

### Architecture

The dashboard uses a component-based rendering architecture with these key packages:

- **dashboard**: Core dashboard configuration and implementation
- **model**: Shared interfaces and types to avoid import cycles
- **render**: Component-based rendering logic using gouniverse/hb

### Components

The dashboard includes several UI components rendered using the gouniverse/hb HTML builder:

#### Header Components

- Top navbar with logo and user menu
- Main menu with customizable items
- User dropdown with profile information
- Theme switcher for light/dark mode
- Quick access menu for frequently used actions

#### Content Components

```go
// Example of creating content with cards
content := `
<div class="row">
  <div class="col-md-4">
    <div class="card" style="margin: 15px;">
      <div class="card-header">Users</div>
      <div class="card-body">
        <h3>1,234</h3>
        <p>Total users</p>
      </div>
    </div>
  </div>
</div>
`

// Set the content in the dashboard config
config := dashboard.Config{
    Content: content,
    // other config options...
}
```

#### Tab Component

```go
// Example of creating tabs in your content
tabsHTML := `
<div class="tab-container mt-3 mb-3" id="tab-example">
  <ul class="nav nav-tabs" role="tablist">
    <li class="nav-item" role="presentation">
      <a class="nav-link active" data-bs-toggle="tab" href="#overview" role="tab">
        <i class="ti ti-chart-bar me-2"></i>Overview
      </a>
    </li>
    <li class="nav-item" role="presentation">
      <a class="nav-link" data-bs-toggle="tab" href="#details" role="tab">
        <i class="ti ti-list me-2"></i>Details
      </a>
    </li>
  </ul>
  <div class="tab-content">
    <div class="tab-pane fade show active" id="overview" role="tabpanel">
      <p>Overview content</p>
    </div>
    <div class="tab-pane fade" id="details" role="tabpanel">
      <p>Details content</p>
    </div>
  </div>
</div>
`
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
