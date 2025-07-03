# Dracory Dashboard

A modern, responsive dashboard framework that supports multiple UI frameworks,
including [Bootstrap](https://getbootstrap.com/), [Tabler](https://tabler.io/),
and [AdminLTE](https://adminlte.io/). This dashboard is designed to be easy to use,
highly customizable, and feature-rich.

## Features

- **Multiple UI Frameworks**: Choose between Bootstrap, Tabler, or AdminLTE
- **Responsive Design**: Works on all screen sizes and devices
- **Theme Switching**: Support for light and dark themes with persistent user preferences
- **Customizable Navigation**: Configure main menu, user menu, and quick access menu
- **Component-Based Rendering**: Modular rendering system
- **Flexible Layout**: Various layout components (cards, grids, tabs, shadow boxes)
- **User Management**: Built-in user authentication UI components
- **Modern Icons**: Built-in support for Tabler Icons
- **Custom Styling**: Easy to customize with your own CSS

## Installation

```bash
go get github.com/dracory/dashboard
```

## Usage

### Basic Example with Bootstrap

```go
package main

import (
    "fmt"
    "net/http"
    
    "github.com/dracory/dashboard"
    "github.com/dracory/dashboard/shared"
    "github.com/dracory/dashboard/types"
    "github.com/samber/lo"
)

func main() {
    http.HandleFunc("/", handleHome)
    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    // Create a new dashboard instance
    d := dashboard.New()

    // Set the template to use Bootstrap
    d.SetTemplate(dashboard.TEMPLATE_BOOTSTRAP)
    
    // Set dashboard title and basic configuration
    d.SetTitle("My Dashboard")
    d.SetLogoImageURL("https://example.com/logo.png")
    d.SetLogoRedirectURL("/")
    
    // Set theme handling
    d.SetThemeHandlerUrl("/")
    
    // Set up main menu
    mainMenuItems := []types.MenuItem{
        {
            Title: "Dashboard",
            URL:   "/",
            Icon:  `<i class="bi bi-speedometer2"></i>`,
        },
        {
            Title: "Users",
            URL:   "/users",
            Icon:  `<i class="bi bi-people"></i>`,
        },
    }
    d.SetMenuMainItems(mainMenuItems)
    
    // Set user information
    user := types.User{
        FirstName: "John",
        LastName:  "Doe",
    }
    d.SetUser(user)
    
    // Set the dashboard content
    d.SetContent(`
        <div class="container-fluid">
            <h1>Welcome to Your Dashboard</h1>
            <p>This is a sample dashboard using Bootstrap.</p>
        </div>
    `)
    
    // Render and write the response
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(d.ToHTML()))
}
```

### Available Templates

The dashboard supports multiple UI frameworks:

1. **Bootstrap** - `dashboard.TEMPLATE_BOOTSTRAP`
   - Modern, responsive design
   - Built with Bootstrap 5
   - Clean and simple interface

2. **Tabler** - `dashboard.TEMPLATE_TABLER`
   - Open Source dashboard template
   - Responsive and accessible
   - Modern UI components

3. **AdminLTE** - `dashboard.TEMPLATE_ADMINLTE`
   - AdminLTE 3 based dashboard
   - Responsive admin template
   - Wide range of components

### Architecture

The dashboard uses a component-based rendering architecture with these key packages:

- **dashboard**: Core dashboard configuration and implementation
- **types**: Shared data structures and interfaces
- **shared**: Common constants and utilities

## Components

The dashboard includes several UI components that work across all supported frameworks:

### Navigation Components

- **Main Menu**: Customizable sidebar navigation
- **User Menu**: Dropdown menu for user actions
- **Breadcrumbs**: Navigation hierarchy
- **Quick Access**: Frequently used actions

### Content Components

#### Cards

```go
content := `
<div class="row">
  <div class="col-md-4">
    <div class="card">
      <div class="card-header">
        <h3 class="card-title">Users</h3>
      </div>
      <div class="card-body">
        <h2>1,234</h2>
        <p>Total registered users</p>
      </div>
    </div>
  </div>
</div>
`
```

#### Tabs

```go
tabs := `
<ul class="nav nav-tabs" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" data-bs-toggle="tab" href="#overview">
      <i class="ti ti-chart-bar me-1"></i> Overview
    </a>
  </li>
  <li class="nav-item">
    <a class="nav-link" data-bs-toggle="tab" href="#details">
      <i class="ti ti-list-details me-1"></i> Details
    </a>
  </li>
</ul>
<div class="tab-content p-3 border border-top-0 rounded-bottom">
  <div class="tab-pane fade show active" id="overview">
    <p>Overview content goes here</p>
  </div>
  <div class="tab-pane fade" id="details">
    <p>Detailed information here</p>
  </div>
</div>
`
```

## Configuration Options

### Dashboard Settings

| Method | Description |
|--------|-------------|
| `SetTemplate(template string)` | Set the UI framework (Bootstrap, Tabler, or AdminLTE) |
| `SetTitle(title string)` | Set the page title |
| `SetFaviconURL(url string)` | Set the favicon URL |
| `SetTheme(theme string)` | Set the color theme (i.e. default, dark, light) |
| `SetThemeHandlerUrl(url string)` | Set the URL for theme switching |

### Navigation

| Method | Description |
|--------|-------------|
| `SetMenuMainItems(items []MenuItem)` | Set main menu items |
| `SetMenuUserItems(items []MenuItem)` | Set user menu items |
| `SetMenuQuickAccessItems(items []MenuItem)` | Set quick access menu items |
| `SetBreadcrumb(items []BreadcrumbItem)` | Set breadcrumb navigation |

### User Interface

| Method | Description |
|--------|-------------|
| `SetLogoImageURL(url string)` | Set the logo image URL |
| `SetLogoRedirectURL(url string)` | Set the logo click URL |
| `SetNavbarBackgroundColor(color string)` | Set navbar background color |
| `SetNavbarTextColor(color string)` | Set navbar text color |
| `SetNavbarBackgroundColorMode(mode string)` | Set navbar color mode (light/dark) |
```

## Advanced Usage

### Theme Switching

The dashboard includes built-in theme switching support:

```go
// In your request handler
theme := r.URL.Query().Get("theme")
if theme != "" {
    // Set theme cookie
    cookie := &http.Cookie{
        Name:     "theme",
        Value:    theme,
        Path:     "/",
        MaxAge:   86400 * 30, // 30 days
        HttpOnly: false,
    }
    http.SetCookie(w, cookie)
}

// Apply theme to dashboard
d.SetTheme(theme)
```

### Custom Styling

Add custom CSS to your dashboard:

```go
d.SetStyles([]string{
    `/* Custom styles */
    .my-custom-class { 
        background-color: #f8f9fa;
        padding: 1rem;
    }`,
})
```

### Adding JavaScript

Add custom JavaScript:

```go
d.SetScripts([]string{
    `// Initialize tooltips
    document.addEventListener('DOMContentLoaded', function() {
        var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
        var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
            return new bootstrap.Tooltip(tooltipTriggerEl);
        });
    });`,
})
```

## Examples

Check the `examples/` directory for complete working examples:

- `bootstrap/` - Example using Bootstrap template
- `tabler/` - Example using Tabler template
- `adminlte/` - Example using AdminLTE template

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
