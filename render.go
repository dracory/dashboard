package dashboard

import (
	"github.com/samber/lo"
)

// ToHTML renders the dashboard as HTML
func (d *Dashboard) ToHTML() string {
	isDarkTheme := d.isThemeDark()

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, viewport-fit=cover">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Dashboard</title>
    <!-- Tabler Core CSS -->
    ` + GetTablerCDNLinks(isDarkTheme) + `
    <style>` + d.dashboardStyle() + `</style>
</head>
<body` + d.bodyAttributes() + `>
    <div class="page">
        <!-- Navbar -->
        ` + d.topNavigation() + `
        <div class="page-wrapper">
            <!-- Page content -->
            <div class="page-body">
                <div class="container-fluid">
                    ` + d.content + `
                </div>
            </div>
            <!-- Footer -->
            <footer class="footer footer-transparent d-print-none">
                <div class="container-fluid">
                    <div class="row text-center align-items-center flex-row-reverse">
                        <div class="col-lg-auto ms-lg-auto">
                            <ul class="list-inline list-inline-dots mb-0">
                                <li class="list-inline-item">
                                    <a href="https://github.com/tabler/tabler" target="_blank" class="link-secondary" rel="noopener">
                                        Powered by Tabler
                                    </a>
                                </li>
                            </ul>
                        </div>
                        <div class="col-12 col-lg-auto mt-3 mt-lg-0">
                            <ul class="list-inline list-inline-dots mb-0">
                                <li class="list-inline-item">
                                    Copyright &copy; 2023
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </footer>
        </div>
    </div>
    
    <!-- Menu modal/offcanvas -->
    ` + d.renderMenu() + `
    
    <!-- Tabler Core JS -->
    ` + GetTablerCDNScripts() + `
    <script>` + d.dashboardScript() + `</script>
</body>
</html>`

	return html
}

// bodyAttributes returns the body tag attributes
func (d *Dashboard) bodyAttributes() string {
	attrs := ""

	if d.isThemeDark() {
		attrs += ` data-bs-theme="dark"`
	}

	return attrs
}

// renderMenu returns the menu HTML based on the menu type
func (d *Dashboard) renderMenu() string {
	if d.menuType == MENU_TYPE_MODAL {
		return d.menuModal()
	}

	return d.menuOffcanvas()
}

// menuOffcanvas returns the offcanvas menu HTML
func (d *Dashboard) menuOffcanvas() string {
	if len(d.menuItems) == 0 {
		return ""
	}

	menuItems := ""
	for i, menuItem := range d.menuItems {
		menuItems += d.buildMenuItem(menuItem, i)
	}

	offcanvas := `<div class="offcanvas offcanvas-start" tabindex="-1" id="offcanvasMenu" aria-labelledby="offcanvasMenuLabel">
  <div class="offcanvas-header">
    <h5 class="offcanvas-title" id="offcanvasMenuLabel">Menu</h5>
    <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
  </div>
  <div class="offcanvas-body">
    <div class="list-group list-group-flush">
      ` + menuItems + `
    </div>
  </div>
</div>`

	return offcanvas
}

// menuModal returns the modal menu HTML
func (d *Dashboard) menuModal() string {
	if len(d.menuItems) == 0 {
		return ""
	}

	menuItems := ""
	for i, menuItem := range d.menuItems {
		menuItems += d.buildMenuItem(menuItem, i)
	}

	modal := `<div class="modal fade" id="menuModal" tabindex="-1" aria-labelledby="menuModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="menuModalLabel">Menu</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        <div class="list-group list-group-flush">
          ` + menuItems + `
        </div>
      </div>
    </div>
  </div>
</div>`

	return modal
}

// buildMenuItem builds a menu item HTML
func (d *Dashboard) buildMenuItem(menuItem MenuItem, index int) string {
	if len(menuItem.SubMenu) > 0 {
		return d.buildSubmenuItem(menuItem, index)
	}

	icon := ""
	if !lo.IsEmpty(menuItem.Icon) {
		icon = `<i class="` + menuItem.Icon + ` me-2"></i>`
	}

	badge := ""
	if !lo.IsEmpty(menuItem.BadgeText) {
		badgeClass := "badge bg-primary ms-auto"
		if !lo.IsEmpty(menuItem.BadgeClass) {
			badgeClass = menuItem.BadgeClass
		}
		badge = `<span class="` + badgeClass + `">` + menuItem.BadgeText + `</span>`
	}

	target := ""
	if menuItem.NewWindow {
		target = ` target="_blank"`
	}

	onClick := ""
	if !lo.IsEmpty(menuItem.OnClick) {
		onClick = ` onclick="` + menuItem.OnClick + `"`
	}

	activeClass := ""
	if menuItem.Active {
		activeClass = " active"
	}

	text := ""
	if d.menuShowText {
		text = menuItem.Text
	}

	return `<a href="` + menuItem.URL + `" class="list-group-item list-group-item-action` + activeClass + `"` + target + onClick + `>
  ` + icon + text + badge + `
</a>`
}

// buildSubmenuItem builds a submenu item HTML
func (d *Dashboard) buildSubmenuItem(menuItem MenuItem, index int) string {
	icon := ""
	if !lo.IsEmpty(menuItem.Icon) {
		icon = `<i class="` + menuItem.Icon + ` me-2"></i>`
	}

	text := ""
	if d.menuShowText {
		text = menuItem.Text
	}

	submenuItems := ""
	for i, subMenuItem := range menuItem.SubMenu {
		submenuItems += d.buildMenuItem(subMenuItem, i)
	}

	return `<div class="list-group-item">
  <a href="#submenu-` + menuItem.ID + `" data-bs-toggle="collapse" class="d-flex align-items-center">
    ` + icon + text + `
    <i class="ti ti-chevron-down ms-auto"></i>
  </a>
  <div class="collapse" id="submenu-` + menuItem.ID + `">
    <div class="list-group list-group-flush mt-2">
      ` + submenuItems + `
    </div>
  </div>
</div>`
}

// topNavigation returns the top navigation HTML
func (d *Dashboard) topNavigation() string {
	navbarClass := "navbar navbar-expand-md"
	if d.navbarHasBackgroundThemeClass() {
		navbarClass += " " + d.navbarBackgroundThemeClass()
	} else if !lo.IsEmpty(d.navbarBackgroundColor) {
		navbarClass += " bg-" + d.navbarBackgroundColor
	}

	if !lo.IsEmpty(d.navbarTextColor) {
		navbarClass += " navbar-" + d.navbarTextColor
	}

	// Logo
	logo := ""
	if !lo.IsEmpty(d.logoRawHtml) {
		logo = d.logoRawHtml
	} else if !lo.IsEmpty(d.logoImageURL) {
		logoLink := d.logoRedirectURL
		if lo.IsEmpty(logoLink) {
			logoLink = "#"
		}

		logo = `<a href="` + logoLink + `" class="navbar-brand navbar-brand-autodark">
          <img src="` + d.logoImageURL + `" height="36" alt="Logo">
        </a>`
	}

	// Menu button
	menuButton := ""
	if len(d.menuItems) > 0 {
		buttonClass := "btn"
		if d.navbarHasBackgroundThemeClass() {
			buttonClass += " " + d.navbarButtonThemeClass()
		}

		dataTarget := "#offcanvasMenu"
		dataBsToggle := "offcanvas"
		if d.menuType == MENU_TYPE_MODAL {
			dataTarget = "#menuModal"
			dataBsToggle = "modal"
		}

		menuButton = `<button class="` + buttonClass + `" type="button" data-bs-toggle="` + dataBsToggle + `" data-bs-target="` + dataTarget + `">
          <i class="ti ti-menu-2 d-md-none"></i>
          <span class="d-none d-md-inline">Menu</span>
        </button>`
	}

	// User dropdown
	userDropdown := d.navbarDropdownUser()

	// Theme switcher
	themeSwitcher := d.navbarDropdownThemeSwitch()

	// Quick access menu
	quickAccess := ""
	if len(d.quickAccessMenu) > 0 {
		quickAccess = d.navbarDropdownQuickAccess()
	}

	return `<header class="navbar navbar-expand-md d-print-none sticky-top">
  <div class="container-fluid">
    <div class="navbar-nav flex-row order-md-last">
      ` + userDropdown + themeSwitcher + quickAccess + `
    </div>
    <div class="d-flex">
      ` + menuButton + `
      ` + logo + `
    </div>
  </div>
</header>`
}

// navbarDropdownUser returns the user dropdown HTML
func (d *Dashboard) navbarDropdownUser() string {
	if d.user.ID == "" {
		// If no user, show login/register links if available
		if !lo.IsEmpty(d.loginURL) || !lo.IsEmpty(d.registerURL) {
			loginLink := ""
			if !lo.IsEmpty(d.loginURL) {
				loginLink = `<a href="` + d.loginURL + `" class="dropdown-item">Login</a>`
			}

			registerLink := ""
			if !lo.IsEmpty(d.registerURL) {
				registerLink = `<a href="` + d.registerURL + `" class="dropdown-item">Register</a>`
			}

			return `<div class="nav-item dropdown">
        <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown" aria-label="Open user menu">
          <span class="avatar avatar-sm"><i class="ti ti-user"></i></span>
        </a>
        <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
          ` + loginLink + registerLink + `
        </div>
      </div>`
		}

		return ""
	}

	// User is logged in
	avatar := `<span class="avatar avatar-sm">` + d.user.Name[:1] + `</span>`
	if !lo.IsEmpty(d.user.AvatarURL) {
		avatar = `<span class="avatar avatar-sm" style="background-image: url(` + d.user.AvatarURL + `)"></span>`
	}

	userMenuItems := ""
	for _, menuItem := range d.userMenu {
		icon := ""
		if !lo.IsEmpty(menuItem.Icon) {
			icon = `<i class="` + menuItem.Icon + ` me-2"></i>`
		}

		target := ""
		if menuItem.NewWindow {
			target = ` target="_blank"`
		}

		onClick := ""
		if !lo.IsEmpty(menuItem.OnClick) {
			onClick = ` onclick="` + menuItem.OnClick + `"`
		}

		userMenuItems += `<a href="` + menuItem.URL + `" class="dropdown-item"` + target + onClick + `>` + icon + menuItem.Text + `</a>`
	}

	return `<div class="nav-item dropdown">
    <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown" aria-label="Open user menu">
      ` + avatar + `
      <div class="d-none d-xl-block ps-2">
        <div>` + d.user.Name + `</div>
        <div class="mt-1 small text-muted">` + d.user.Email + `</div>
      </div>
    </a>
    <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
      ` + userMenuItems + `
    </div>
  </div>`
}

// navbarDropdownThemeSwitch returns the theme switcher dropdown HTML
func (d *Dashboard) navbarDropdownThemeSwitch() string {
	return `<div class="nav-item dropdown">
    <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown" aria-label="Theme options">
      <span class="avatar avatar-sm"><i class="ti ti-palette"></i></span>
    </a>
    <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
      <a href="#" class="dropdown-item" onclick="setTheme('light')">
        <i class="ti ti-sun me-2"></i>Light
      </a>
      <a href="#" class="dropdown-item" onclick="setTheme('dark')">
        <i class="ti ti-moon me-2"></i>Dark
      </a>
    </div>
  </div>`
}

// navbarDropdownQuickAccess returns the quick access dropdown HTML
func (d *Dashboard) navbarDropdownQuickAccess() string {
	if len(d.quickAccessMenu) == 0 {
		return ""
	}

	quickAccessItems := ""
	for _, menuItem := range d.quickAccessMenu {
		icon := ""
		if !lo.IsEmpty(menuItem.Icon) {
			icon = `<i class="` + menuItem.Icon + ` me-2"></i>`
		}

		target := ""
		if menuItem.NewWindow {
			target = ` target="_blank"`
		}

		onClick := ""
		if !lo.IsEmpty(menuItem.OnClick) {
			onClick = ` onclick="` + menuItem.OnClick + `"`
		}

		quickAccessItems += `<a href="` + menuItem.URL + `" class="dropdown-item"` + target + onClick + `>` + icon + menuItem.Text + `</a>`
	}

	return `<div class="nav-item dropdown">
    <a href="#" class="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown" aria-label="Quick access">
      <span class="avatar avatar-sm"><i class="ti ti-apps"></i></span>
    </a>
    <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
      ` + quickAccessItems + `
    </div>
  </div>`
}

// navbarHasBackgroundThemeClass returns whether the navbar has a background theme class
func (d *Dashboard) navbarHasBackgroundThemeClass() bool {
	return !lo.IsEmpty(d.navbarBackgroundColorMode) && (d.navbarBackgroundColorMode == "light" || d.navbarBackgroundColorMode == "dark")
}

// navbarBackgroundThemeClass returns the navbar background theme class
func (d *Dashboard) navbarBackgroundThemeClass() string {
	if d.navbarBackgroundColorMode == "light" {
		return "bg-light"
	}
	return "bg-dark"
}

// navbarButtonThemeClass returns the navbar button theme class
func (d *Dashboard) navbarButtonThemeClass() string {
	if d.navbarBackgroundColorMode == "light" {
		return "btn-dark"
	}
	return "btn-light"
}

// isThemeDark returns whether the theme is dark
func (d *Dashboard) isThemeDark() bool {
	return d.themeName == "dark"
}

// dashboardStyle returns the dashboard custom CSS
func (d *Dashboard) dashboardStyle() string {
	return `
.avatar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 2rem;
  height: 2rem;
  border-radius: 50%;
  background: #f0f0f0;
  color: #666;
}
`
}

// dashboardScript returns the dashboard custom JavaScript
func (d *Dashboard) dashboardScript() string {
	return `
function setTheme(theme) {
  document.body.setAttribute('data-bs-theme', theme);
  localStorage.setItem('theme', theme);
  
  // Add theme cookie
  document.cookie = "theme=" + theme + "; path=/; max-age=31536000";
}

// Initialize theme from cookie or localStorage
document.addEventListener('DOMContentLoaded', function() {
  const savedTheme = localStorage.getItem('theme') || 'light';
  setTheme(savedTheme);
});
`
}
