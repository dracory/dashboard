package main

import (
	"fmt"
	"net/http"

	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/types"
	"github.com/samber/lo"
)

func main() {
	// Start the web server
	http.HandleFunc("/", handleHome)
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Get theme from query parameter
	theme := r.URL.Query().Get("theme")

	// If theme is provided in query, set cookie
	lo.If(theme != "", func() {
		cookie := &http.Cookie{
			Name:     "theme",
			Value:    theme,
			Path:     "/",
			MaxAge:   86400 * 30, // 30 days
			HttpOnly: false,
		}
		http.SetCookie(w, cookie)
	})

	// If no theme in query, try to get from cookie
	if theme == "" {
		cookie, err := r.Cookie("theme")
		// Use cookie value if available, otherwise default to tabler
		theme = lo.Ternary(err == nil && cookie != nil && cookie.Value != "", cookie.Value, "tabler")
	}

	// Create a new dashboard instance
	d := dashboard.New()

	// Set the template to use Tabler
	d.SetTemplate(dashboard.TEMPLATE_TABLER)

	// Set dashboard title
	d.SetTitle("Tabler Dashboard")

	// Set logo and navbar settings
	d.SetLogoImageURL("https://tabler.bootmb.com/static/logo.svg")
	d.SetLogoRedirectURL("/")
	d.SetNavbarBackgroundColor("#206bc4") // Tabler primary blue
	d.SetNavbarTextColor("#ffffff")       // White text
	d.SetNavbarBackgroundColorMode("light")

	// Set theme and theme handler URL
	d.SetTheme(theme)
	d.SetThemeHandlerUrl("/")

	// Add custom styles for Tabler
	d.SetStyles([]string{
		`/* Custom styles */
		.navbar { box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.05); }
		.content { padding: 1.5rem; }
		.card { margin-bottom: 1.5rem; }
		`,
	})

	// Add custom scripts
	d.SetScripts([]string{
		`// Initialize tooltips
		document.addEventListener('DOMContentLoaded', function() {
			var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
			var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
				return new bootstrap.Tooltip(tooltipTriggerEl);
			});
		});`,
	})

	// Set main menu items to match Tabler preview
	mainMenuItems := []types.MenuItem{
		{
			Title:    "Home",
			URL:      "/",
			Icon:     `<i class="ti ti-home"></i>`,
			IsActive: true,
		},
		{
			Title: "Interface",
			Icon:  `<i class="ti ti-layout-2"></i>`,
			Children: []types.MenuItem{
				{
					Title: "Accordion",
					URL:   "/components/accordion",
					Icon:  `<i class="ti ti-layout-accordion-rows"></i>`,
				},
				{
					Title: "Alerts",
					URL:   "/components/alerts",
					Icon:  `<i class="ti ti-alert-triangle"></i>`,
				},
				{
					Title: "Avatars",
					URL:   "/components/avatars",
					Icon:  `<i class="ti ti-user-circle"></i>`,
				},
				{
					Title: "Badges",
					URL:   "/components/badges",
					Icon:  `<i class="ti ti-bell"></i>`,
				},
				{
					Title: "Breadcrumbs",
					URL:   "/components/breadcrumbs",
					Icon:  `<i class="ti ti-arrow-narrow-right"></i>`,
				},
				{
					Title: "Buttons",
					URL:   "/components/buttons",
					Icon:  `<i class="ti ti-square"></i>`,
				},
				{
					Title: "Cards",
					URL:   "/components/cards",
					Icon:  `<i class="ti ti-cards"></i>`,
				},
				{
					Title: "Carousel",
					URL:   "/components/carousel",
					Icon:  `<i class="ti ti-layout-grid"></i>`,
				},
				{
					Title: "Charts",
					URL:   "/components/charts",
					Icon:  `<i class="ti ti-chart-line"></i>`,
				},
				{
					Title: "Dropdowns",
					URL:   "/components/dropdowns",
					Icon:  `<i class="ti ti-chevron-down"></i>`,
				},
				{
					Title: "Empty states",
					URL:   "/components/empty-states",
					Icon:  `<i class="ti ti-file-off"></i>`,
				},
				{
					Title: "Flags",
					URL:   "/components/flags",
					Icon:  `<i class="ti ti-flag"></i>`,
				},
				{
					Title: "Icons",
					URL:   "/components/icons",
					Icon:  `<i class="ti ti-brand-tabler"></i>`,
				},
				{
					Title: "Lists",
					URL:   "/components/lists",
					Icon:  `<i class="ti ti-list"></i>`,
				},
				{
					Title: "Maps",
					URL:   "/components/maps",
					Icon:  `<i class="ti ti-map-2"></i>`,
				},
				{
					Title: "Modals",
					URL:   "/components/modals",
					Icon:  `<i class="ti ti-square-plus"></i>`,
				},
				{
					Title: "Offcanvas",
					URL:   "/components/offcanvas",
					Icon:  `<i class="ti ti-layout-sidebar-right"></i>`,
				},
				{
					Title: "Page headers",
					URL:   "/components/page-headers",
					Icon:  `<i class="ti ti-layout-header"></i>`,
				},
				{
					Title: "Pagination",
					URL:   "/components/pagination",
					Icon:  `<i class="ti ti-arrow-narrow-right"></i>`,
				},
				{
					Title: "Placeholders",
					URL:   "/components/placeholders",
					Icon:  `<i class="ti ti-box-margin"></i>`,
				},
				{
					Title: "Pricing cards",
					URL:   "/components/pricing-cards",
					Icon:  `<i class="ti ti-currency-dollar"></i>`,
				},
				{
					Title: "Progress",
					URL:   "/components/progress",
					Icon:  `<i class="ti ti-progress"></i>`,
				},
				{
					Title: "Social cards",
					URL:   "/components/social-cards",
					Icon:  `<i class="ti ti-share"></i>`,
				},
				{
					Title: "Spinners",
					URL:   "/components/spinners",
					Icon:  `<i class="ti ti-loader"></i>`,
				},
				{
					Title: "Tables",
					URL:   "/components/tables",
					Icon:  `<i class="ti ti-table"></i>`,
				},
				{
					Title: "Tabs",
					URL:   "/components/tabs",
					Icon:  `<i class="ti ti-layout-navbar"></i>`,
				},
				{
					Title: "Timelines",
					URL:   "/components/timelines",
					Icon:  `<i class="ti ti-timeline"></i>`,
				},
				{
					Title: "Toasts",
					URL:   "/components/toasts",
					Icon:  `<i class="ti ti-bell"></i>`,
				},
				{
					Title: "Tooltips",
					URL:   "/components/tooltips",
					Icon:  `<i class="ti ti-message-chatbot"></i>`,
				},
			},
		},
		{
			Title: "Components",
			Icon:  `<i class="ti ti-components"></i>`,
			Children: []types.MenuItem{
				{
					Title: "Buttons",
					URL:   "/components/buttons",
					Icon:  `<i class="ti ti-square"></i>`,
				},
				{
					Title: "Forms",
					URL:   "/components/forms",
					Icon:  `<i class="ti ti-forms"></i>`,
				},
				{
					Title: "Typography",
					URL:   "/components/typography",
					Icon:  `<i class="ti ti-text-size"></i>`,
				},
			},
		},
		{
			Title: "Pages",
			Icon:  `<i class="ti ti-file"></i>`,
			Children: []types.MenuItem{
				{
					Title: "Blank",
					URL:   "/blank",
					Icon:  `<i class="ti ti-file-text"></i>`,
				},
				{
					Title: "Profile",
					URL:   "/profile",
					Icon:  `<i class="ti ti-user"></i>`,
				},
				{
					Title: "Settings",
					URL:   "/settings",
					Icon:  `<i class="ti ti-settings"></i>`,
				},
			},
		},
	}

	d.SetMenuMainItems(mainMenuItems)

	// Set quick access menu items
	quickAccessItems := []types.MenuItem{
		{
			Title: "Add New User",
			URL:   "/users/new",
			Icon:  `<i class="ti ti-user-plus"></i>`,
		},
		{
			Title: "Create Product",
			URL:   "/products/new",
			Icon:  `<i class="ti ti-plus"></i>`,
		},
	}
	d.SetMenuQuickAccessItems(quickAccessItems)

	// Set user menu items
	userMenuItems := []types.MenuItem{
		{
			Title: "Profile",
			URL:   "/profile",
			Icon:  `<i class="ti ti-user"></i>`,
		},
		{
			Title: "Settings",
			URL:   "/settings",
			Icon:  `<i class="ti ti-settings"></i>`,
		},
		{},
		{
			Title: "Logout",
			URL:   "/logout",
			Icon:  `<i class="ti ti-logout"></i>`,
		},
	}
	d.SetMenuUserItems(userMenuItems)

	// Set user information
	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	d.SetUser(user)

	// Create dashboard content
	content := `
	<div class="container-fluid">
		<div class="row">
			<div class="col-12">
				<div class="page-header">
					<h1 class="page-title">Dashboard</h1>
					<div class="text-muted mt-1">Welcome back, John!</div>
				</div>
			</div>
		</div>

		<div class="row row-cards">
			<div class="col-sm-6 col-lg-3">
				<div class="card">
					<div class="card-body p-3 text-center">
						<div class="h1 m-0">1,234</div>
						<div class="text-muted mb-3">Total Users</div>
					</div>
				</div>
			</div>
			<div class="col-sm-6 col-lg-3">
				<div class="card">
					<div class="card-body p-3 text-center">
						<div class="h1 m-0">$23,987</div>
						<div class="text-muted mb-3">Total Revenue</div>
					</div>
				</div>
			</div>
			<div class="col-sm-6 col-lg-3">
				<div class="card">
					<div class="card-body p-3 text-center">
						<div class="h1 m-0">89</div>
						<div class="text-muted mb-3">New Orders</div>
					</div>
				</div>
			</div>
			<div class="col-sm-6 col-lg-3">
				<div class="card">
					<div class="card-body p-3 text-center">
						<div class="h1 m-0">98%</div>
						<div class="text-muted mb-3">Satisfaction</div>
					</div>
				</div>
			</div>
		</div>

		<div class="row">
			<div class="col-12">
				<div class="card">
					<div class="card-header">
						<h3 class="card-title">Recent Activity</h3>
					</div>
					<div class="card-body">
						<div class="table-responsive">
							<table class="table table-vcenter">
								<thead>
									<tr>
										<th>User</th>
										<th>Action</th>
										<th>Time</th>
									</tr>
								</thead>
								<tbody>
									<tr>
										<td>John Doe</td>
										<td>Created a new product</td>
										<td>2 minutes ago</td>
									</tr>
									<tr>
										<td>Jane Smith</td>
										<td>Updated user profile</td>
										<td>15 minutes ago</td>
									</tr>
									<tr>
										<td>Bob Johnson</td>
										<td>Processed 5 orders</td>
										<td>1 hour ago</td>
									</tr>
								</tbody>
							</table>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div class="row">
			<div class="col-12">
				<div class="card">
					<div class="card-body">
						<h3 class="card-title">Theme Switcher</h3>
						<p>Try different themes:</p>
						<div class="btn-list">
							<a href="/?theme=tabler" class="btn btn-primary">Default</a>
							<a href="/?theme=dark" class="btn btn-dark">Dark</a>
							<a href="/?theme=light" class="btn btn-light">Light</a>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	`

	d.SetContent(content)

	// Render the dashboard
	html := d.ToHTML()

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
