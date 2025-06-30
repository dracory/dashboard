# Theme Switching Example

This example demonstrates how to use the theme switching functionality in the Dracory Dashboard package.

## Features

- Switch between Tabler, Bootstrap 5, and AdminLTE themes
- Theme selection persists via URL parameters
- Responsive design that works with all included themes

## How to Run

1. Make sure you have Go installed on your system
2. Navigate to this directory in your terminal
3. Run the example:
   ```bash
   go run main.go
   ```
4. Open your browser and visit `http://localhost:8080`
5. Use the theme buttons to switch between different themes

## How It Works

The example creates a simple web server that serves a dashboard with theme switching capabilities. The theme is selected via URL parameters:

- `http://localhost:8080/?theme=tabler` - Tabler theme (default)
- `http://localhost:8080/?theme=bootstrap` - Bootstrap 5 theme
- `http://localhost:8080/?theme=adminlte` - AdminLTE theme

## Code Overview

The main components of this example are:

1. **Dashboard Setup**:
   ```go
   d := dashboard.New()
   d.SetContent(`...`)
   ```

2. **Theme Selection**:
   ```go
   theme := r.URL.Query().Get("theme")
   if theme != "" {
       d.SetThemeName(theme)
   }
   ```

3. **Rendering**:
   ```go
   html := dashboard.Render(d)
   w.Write([]byte(html))
   ```

## Customization

You can customize the dashboard by:

1. Adding more content using `d.SetContent()`
2. Adding menu items with `d.SetMenuItems()`
3. Setting a user with `d.SetUser()`
4. Customizing the navbar with `SetNavbarBackgroundColor()` and related methods

## Dependencies

- Go 1.16 or higher
- Dracory Dashboard package
- Go HTTP server (included in standard library)
