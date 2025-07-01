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
		// Use cookie value if available, otherwise default to bootstrap
		theme = lo.Ternary(err == nil && cookie != nil && cookie.Value != "", cookie.Value, "bootstrap")
	}

	// Create a new dashboard instance
	d := dashboard.New()

	// Set the template to use Bootstrap
	d.SetTemplate(dashboard.TEMPLATE_BOOTSTRAP)
	
	// Set menu type to modal
	d.SetMenuType(types.MENU_TYPE_MODAL)

	// Set dashboard title
	d.SetTitle("Admin Dashboard")

	// Set logo and navbar settings
	d.SetLogoImageURL("https://tabler.bootmb.com/static/logo.svg")
	d.SetLogoRedirectURL("/")
	d.SetNavbarBackgroundColor("#ffffff") // White background
	d.SetNavbarTextColor("#000000")       // Black text
	d.SetNavbarBackgroundColorMode("light")
	
	// Set theme and theme handler URL
	d.SetTheme(theme)
	d.SetThemeHandlerUrl("/")

	// Add required CSS and JS
	d.SetStyles([]string{``})

	d.SetScripts([]string{
		`// Initialize tooltips
		var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
		var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
		  return new bootstrap.Tooltip(tooltipTriggerEl)
		})`,
	})

	// Set main menu items
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
		{
			Title: "Products",
			URL:   "/products",
			Icon:  `<i class="bi bi-box"></i>`,
		},
		{
			Title: "Orders",
			URL:   "/orders",
			Icon:  `<i class="bi bi-cart"></i>`,
		},
		{
			Title: "Reports",
			URL:   "/reports",
			Icon:  `<i class="bi bi-graph-up"></i>`,
		},
	}
	d.SetMenuMainItems(mainMenuItems)

	// Set quick access menu items
	quickAccessItems := []types.MenuItem{
		{
			Title: "Add New User",
			URL:   "/users/new",
			Icon:  `<i class="bi bi-person-plus"></i>`,
		},
		{
			Title: "Create Product",
			URL:   "/products/new",
			Icon:  `<i class="bi bi-plus-circle"></i>`,
		},
	}
	d.SetMenuQuickAccessItems(quickAccessItems)

	// Set user menu items
	userMenuItems := []types.MenuItem{
		{
			Title: "Profile",
			URL:   "/profile",
			Icon:  `<i class="bi bi-person"></i>`,
		},
		{
			Title: "Settings",
			URL:   "/settings",
			Icon:  `<i class="bi bi-gear"></i>`,
		},
		{},
		{
			Title: "Logout",
			URL:   "/logout",
			Icon:  `<i class="bi bi-box-arrow-right"></i>`,
		},
	}
	d.SetMenuUserItems(userMenuItems)

	// Set user information
	user := types.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	d.SetUser(user)

	// Set the dashboard content
	// 	header := `
	// 	<div class="container-fluid">
	// 		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
	// 			<h1 class="h2">Dashboard</h1>
	// 			<div class="btn-toolbar mb-2 mb-md-0">
	// 				<div class="btn-group me-2">
	// 					<button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
	// 					<button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
	// 				</div>
	// 				<button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
	// 					<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-calendar3" viewBox="0 0 16 16">
	// 						<path d="M14 0H2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zM1 3.857C1 3.384 1.448 3 2 3h12c.552 0 1 .384 1 .857v10.286c0 .473-.448.857-1 .857H2c-.552 0-1-.384-1-.857V3.857z"/>
	// 						<path d="M6.5 7a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm-9 3a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm-9 3a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2 1 1 0 0 0 0 2z"/>
	// 					</svg>
	// 					This week
	// 				</button>
	// 			</div>
	// 		</div>
	// 	`

	// 	content := `
	// 	<div class="row">
	// 		<div class="col-md-3 mb-4">
	// 			<div class="card bg-primary text-white">
	// 				<div class="card-body">
	// 					<h5 class="card-title">Users</h5>
	// 					<h2 class="mb-0">1,234</h2>
	// 				</div>
	// 			</div>
	// 		</div>
	// 		<div class="col-md-3 mb-4">
	// 			<div class="card bg-success text-white">
	// 				<div class="card-body">
	// 					<h5 class="card-title">Products</h5>
	// 					<h2 class="mb-0">567</h2>
	// 				</div>
	// 			</div>
	// 		</div>
	// 		<div class="col-md-3 mb-4">
	// 			<div class="card bg-warning text-dark">
	// 				<div class="card-body">
	// 					<h5 class="card-title">Orders</h5>
	// 					<h2 class="mb-0">8,901</h2>
	// 				</div>
	// 			</div>
	// 		</div>
	// 		<div class="col-md-3 mb-4">
	// 			<div class="card bg-danger text-white">
	// 				<div class="card-body">
	// 					<h5 class="card-title">Revenue</h5>
	// 					<h2 class="mb-0">$45,678</h2>
	// 				</div>
	// 			</div>
	// 		</div>
	// 	</div>

	// 	<div class="row">
	// 		<div class="col-md-6">
	// 			<div class="card mb-4">
	// 				<div class="card-header">
	// 					<h5 class="card-title mb-0">Recent Activity</h5>
	// 				</div>
	// 				<div class="card-body">
	// 					<div class="list-group list-group-flush">
	// 						<div class="list-group-item">
	// 							<div class="d-flex w-100 justify-content-between">
	// 								<h6 class="mb-1">New user registered</h6>
	// 								<small class="text-muted">2 minutes ago</small>
	// 							</div>
	// 							<p class="mb-1">John Doe created a new account</p>
	// 						</div>
	// 						<div class="list-group-item">
	// 							<div class="d-flex w-100 justify-content-between">
	// 								<h6 class="mb-1">Order #12345 completed</h6>
	// 								<small class="text-muted">1 hour ago</small>
	// 							</div>
	// 							<p class="mb-1">Order has been shipped to the customer</p>
	// 						</div>
	// 					</div>
	// 				</div>
	// 			</div>
	// 		</div>
	// 		<div class="col-md-6">
	// 			<div class="card">
	// 				<div class="card-header">
	// 					<h5 class="card-title mb-0">Quick Actions</h5>
	// 				</div>
	// 				<div class="card-body">
	// 					<div class="row g-2">
	// 						<div class="col-6">
	// 							<button class="btn btn-primary w-100 mb-2">Add New User</button>
	// 						</div>
	// 						<div class="col-6">
	// 							<button class="btn btn-success w-100 mb-2">Create Product</button>
	// 						</div>
	// 						<div class="col-6">
	// 							<button class="btn btn-info w-100 mb-2">View Reports</button>
	// 						</div>
	// 						<div class="col-6">
	// 							<button class="btn btn-warning w-100 mb-2">Settings</button>
	// 						</div>
	// 					</div>
	// 				</div>
	// 			</div>
	// 		</div>
	// 	</div>
	// </div>
	// 	`

	content := `Hello World`

	d.SetContent(content)

	// Get the HTML output and write it to the response
	htmlOutput := d.ToHTML()

	// Set the content type and write the response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlOutput))
}
