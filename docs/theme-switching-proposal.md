# Theme Switching Implementation Proposal

This document outlines a comprehensive plan for implementing theme switching functionality in the Dracory Dashboard module. The goal is to support multiple themes that can be switched via configuration settings.

## Current State

Currently, the dashboard uses the Tabler theme with assets loaded from CDN. Theme-related code is primarily in:
- `tabler_cdn.go` - Contains CDN URLs and functions to generate HTML tags
- `render/page.go` - Contains theme rendering logic, including CSS/JS inclusion

## Proposed Themes

The implementation will support the following themes:
- [x] Tabler (current theme)
- [ ] Pure Bootstrap 5
- [ ] AdminLTE

## Implementation Plan

### 1. Create Theme Interface and Structure

- [ ] Create `render/theme/theme.go` with the Theme interface:
  ```go
  type Theme interface {
      GetName() string
      GetCSSLinks() []*hb.Tag
      GetJSScripts() []*hb.Tag
      GetCustomCSS() string
      GetCustomJS() string
      // Theme-specific rendering methods if needed
  }
  ```

- [ ] Create directory structure for themes:
  - [ ] `render/theme/tabler/`
  - [ ] `render/theme/bootstrap/`
  - [ ] `render/theme/adminlte/`

### 2. Refactor Current Tabler Implementation

- [ ] Move Tabler-specific code from `tabler_cdn.go` to `render/theme/tabler/assets.go`
- [ ] Create `render/theme/tabler/theme.go` implementing the Theme interface
- [ ] Update imports and references to use the new structure

### 3. Implement Bootstrap 5 Theme

- [ ] Create `render/theme/bootstrap/assets.go` with Bootstrap 5 CDN URLs
- [ ] Create `render/theme/bootstrap/theme.go` implementing the Theme interface
- [ ] Add any Bootstrap 5 specific components or styles

### 4. Implement AdminLTE Theme

- [ ] Create `render/theme/adminlte/assets.go` with AdminLTE CDN URLs
- [ ] Create `render/theme/adminlte/theme.go` implementing the Theme interface
- [ ] Add any AdminLTE specific components or styles

### 5. Create Theme Manager

- [ ] Create `render/theme/manager.go` with functions to:
  - [ ] Register available themes
  - [ ] Get theme by name
  - [ ] Get default theme

### 6. Update Dashboard Configuration

- [ ] Add theme selection to dashboard configuration:
  ```go
  type Config struct {
      // Existing fields
      ThemeName string // "tabler", "bootstrap", "adminlte"
  }
  ```

- [ ] Update `model.DashboardRenderer` interface to include theme selection methods:
  ```go
  type DashboardRenderer interface {
      // Existing methods
      GetThemeName() string
      // Other methods as needed
  }
  ```

### 7. Update Page Rendering

- [ ] Modify `render/page.go` to use the theme interface:
  - [ ] Replace direct Tabler references with theme interface calls
  - [ ] Update header, footer, and other components to be theme-agnostic
  - [ ] Use the theme manager to get the appropriate theme based on configuration

### 8. Add Theme Switching UI (Optional)

- [ ] Create a theme switcher component that can be included in the dashboard
- [ ] Implement client-side theme switching with localStorage persistence
- [ ] Add server-side theme switching via configuration update

### 9. Documentation and Examples

- [ ] Update README.md with theme switching documentation
- [ ] Create example code for each theme
- [ ] Document any theme-specific features or limitations

### 10. Testing

- [ ] Test each theme individually
- [ ] Test theme switching functionality
- [ ] Verify that all dashboard components render correctly in each theme
- [ ] Test with different browsers and screen sizes

## Implementation Details

### Theme Interface

The Theme interface will be the core of the implementation, providing a consistent way to access theme assets and rendering functions:

```go
// Theme defines the interface for dashboard themes
type Theme interface {
    // GetName returns the theme's name
    GetName() string
    
    // GetCSSLinks returns the CSS link tags for the theme
    GetCSSLinks(isDarkMode bool) []*hb.Tag
    
    // GetJSScripts returns the JavaScript script tags for the theme
    GetJSScripts() []*hb.Tag
    
    // GetCustomCSS returns any custom CSS for the theme
    GetCustomCSS() string
    
    // GetCustomJS returns any custom JavaScript for the theme
    GetCustomJS() string
}
```

### Configuration

Theme selection will be controlled via the dashboard configuration:

```go
// Example configuration
dashboard := NewDashboard(&Config{
    ThemeName: "bootstrap", // Use Bootstrap theme
    // Other configuration options
})
```

### Theme-Specific Components

Some components may need theme-specific implementations. These will be handled through the theme interface with methods like:

```go
// Optional theme-specific component methods
type Theme interface {
    // ... base methods
    
    // RenderNavbar renders the theme-specific navbar
    RenderNavbar(d model.DashboardRenderer) *hb.Tag
    
    // RenderSidebar renders the theme-specific sidebar
    RenderSidebar(d model.DashboardRenderer) *hb.Tag
}
```

## Benefits

1. **Modularity**: Each theme is self-contained in its own package
2. **Extensibility**: Easy to add new themes by implementing the Theme interface
3. **Configurability**: Simple theme switching via configuration
4. **Maintainability**: Clear separation of theme-specific code from core rendering logic

## Conclusion

This implementation plan provides a structured approach to adding theme switching functionality to the Dracory Dashboard. By following this plan, we'll create a flexible, maintainable system that supports multiple themes while keeping the core dashboard functionality intact.
