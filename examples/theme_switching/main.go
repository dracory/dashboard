package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/components"
	"github.com/dracory/dashboard/render/templates"
	_ "github.com/dracory/dashboard/render/templates/adminlte"
	_ "github.com/dracory/dashboard/render/templates/bootstrap"
	_ "github.com/dracory/dashboard/render/templates/tabler"
	"github.com/gouniverse/hb"
)

// createSidebar creates the sidebar navigation
func createSidebar() string {
	sidebar := hb.NewNav()
	sidebar.Class("nav flex-column")

	// Dashboard link
	dashboardLink := hb.NewLink()
	dashboardLink.Attr("href", "#")
	dashboardLink.Class("nav-link active")
	dashboardLink.AddChild(hb.NewTag("span").Class("nav-link-icon d-md-none d-lg-inline-block me-1").AddChild(hb.NewTag("i").Class("ti ti-home")))
	dashboardLink.AddChild(hb.NewTag("span").Text("Dashboard"))

	// Theme switcher menu
	themeMenu := hb.NewDiv()
	themeMenu.Class("nav-item dropdown")
	themeLink := hb.NewLink()
	themeLink.Attr("href", "#")
	themeLink.Class("nav-link")
	themeLink.Attr("data-bs-toggle", "dropdown")
	themeLink.AddChild(hb.NewTag("span").Class("nav-link-icon d-md-none d-lg-inline-block me-1").AddChild(hb.NewTag("i").Class("ti ti-palette")))
	themeLink.AddChild(hb.NewTag("span").Text("Themes"))
	themeLink.AddChild(hb.NewTag("span").Class("nav-link-toggle"))

	themeDropdown := hb.NewDiv()
	themeDropdown.Class("dropdown-menu")

	tablerLink := hb.NewLink()
	tablerLink.Attr("href", "/?theme=tabler")
	tablerLink.Class("dropdown-item")
	tablerLink.AddChild(hb.NewTag("span").Text("Tabler"))
	themeDropdown.AddChild(tablerLink)

	bootstrapLink := hb.NewLink()
	bootstrapLink.Attr("href", "/?theme=bootstrap")
	bootstrapLink.Class("dropdown-item")
	bootstrapLink.AddChild(hb.NewTag("span").Text("Bootstrap 5"))
	themeDropdown.AddChild(bootstrapLink)

	adminlteLink := hb.NewLink()
	adminlteLink.Attr("href", "/?theme=adminlte")
	adminlteLink.Class("dropdown-item")
	adminlteLink.AddChild(hb.NewTag("span").Text("AdminLTE"))
	themeDropdown.AddChild(adminlteLink)

	themeMenu.AddChild(themeLink)
	themeMenu.AddChild(themeDropdown)

	// Add items to sidebar
	sidebar.AddChild(hb.NewTag("li").Class("nav-item").AddChild(dashboardLink))
	sidebar.AddChild(hb.NewTag("li").Class("nav-item").AddChild(themeMenu))

	return sidebar.ToHTML()
}

// createTopNav creates the top navigation bar
func createTopNav() string {
	nav := hb.NewNav()
	nav.Class("navbar-nav")

	// Theme switcher in top nav
	themeDropdown := hb.NewDiv()
	themeDropdown.Class("nav-item dropdown d-none d-md-flex me-3")

	themeLink := hb.NewLink()
	themeLink.Attr("href", "#")
	themeLink.Class("nav-link px-0")
	themeLink.Attr("data-bs-toggle", "dropdown")
	themeLink.AddChild(hb.NewTag("i").Class("ti ti-palette"))

	dropdownMenu := hb.NewDiv()
	dropdownMenu.Class("dropdown-menu dropdown-menu-end")
	header := hb.NewTag("h6")
	header.Class("dropdown-header")
	header.AddChild(hb.NewTag("span").Text("Select Theme"))
	dropdownMenu.AddChild(header)

	tablerLink := hb.NewTag("a")
	tablerLink.Class("dropdown-item")
	tablerLink.Attr("href", "/?theme=tabler")
	tablerLink.AddChild(hb.NewTag("span").Text("Tabler"))
	dropdownMenu.AddChild(tablerLink)

	bootstrapLink := hb.NewTag("a")
	bootstrapLink.Class("dropdown-item")
	bootstrapLink.Attr("href", "/?theme=bootstrap")
	bootstrapLink.AddChild(hb.NewTag("span").Text("Bootstrap 5"))
	dropdownMenu.AddChild(bootstrapLink)

	adminlteLink := hb.NewTag("a")
	adminlteLink.Class("dropdown-item")
	adminlteLink.Attr("href", "/?theme=adminlte")
	adminlteLink.AddChild(hb.NewTag("span").Text("AdminLTE"))
	dropdownMenu.AddChild(adminlteLink)

	themeDropdown.AddChild(themeLink)
	themeDropdown.AddChild(dropdownMenu)
	nav.AddChild(themeDropdown)

	return nav.ToHTML()
}

// createDashboardContent creates the main dashboard content
func createDashboardContent() string {
	// Welcome card
	welcomeCard := components.NewCard(components.CardConfig{
		Content: `
			<div class="row">
				<div class="col-md-8">
					<h3 class="mb-1">Welcome to Theme Switcher</h3>
					<p class="text-muted">This is a demo of the theme switching functionality. Try changing themes using the menu.</p>
				</div>
				<div class="col-md-4 text-end">
					<img src="https://preview.tabler.io/static/illustrations/undraw_theme_switch_re_8x7x.svg" alt="Theme switching" class="img-fluid" style="max-height: 100px;">
				</div>
			</div>
		`,
	}).ToHTML()

	// Theme cards
	themeCards := `
	<div class="row row-cards mt-3">
		<div class="col-md-4">
			<div class="card">
				<div class="card-body text-center">
					<i class="ti ti-brand-tabler text-primary mb-3" style="font-size: 2.5rem;"></i>
					<h3>Tabler</h3>
					<p class="text-muted">Modern and clean dashboard theme</p>
					<a href="/?theme=tabler" class="btn btn-primary">Select Tabler</a>
				</div>
			</div>
		</div>
		<div class="col-md-4">
			<div class="card">
				<div class="card-body text-center">
					<i class="ti ti-brand-bootstrap text-success mb-3" style="font-size: 2.5rem;"></i>
					<h3>Bootstrap 5</h3>
					<p class="text-muted">Popular CSS framework</p>
					<a href="/?theme=bootstrap" class="btn btn-success">Select Bootstrap</a>
				</div>
			</div>
		</div>
		<div class="col-md-4">
			<div class="card">
				<div class="card-body text-center">
					<i class="ti ti-brand-windows text-warning mb-3" style="font-size: 2.5rem;"></i>
					<h3>AdminLTE</h3>
					<p class="text-muted">Admin dashboard template</p>
					<a href="/?theme=adminlte" class="btn btn-warning">Select AdminLTE</a>
				</div>
			</div>
		</div>
	</div>
	`

	return welcomeCard + themeCards
}

func main() {
	// Initialize all registered themes
	templates.InitializeRegisteredTemplates()

	// Create a new dashboard instance
	d := dashboard.New()
	d.SetThemeName("tabler") // Set default theme

	// Set up the dashboard with content and navigation
	content := `
	<div class="container-fluid">
		<div class="row">
			<!-- Sidebar -->
			<nav class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3">` + createSidebar() + `
				</div>
			</nav>

			<!-- Main content -->
			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<!-- Top navigation -->
				<nav class="navbar navbar-expand-lg navbar-light bg-white border-bottom mb-4">
					<div class="container-fluid">
						<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
							<span class="navbar-toggler-icon"></span>
						</button>
						<div class="collapse navbar-collapse" id="navbarSupportedContent">` +
		createTopNav() + `
						</div>
					</div>
				</nav>

				<!-- Page content -->
				` + createDashboardContent() + `
			</main>
		</div>
	</div>
	`

	d.SetContent(content)

	// Set up a simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the theme from query parameter
		theme := r.URL.Query().Get("theme")
		if theme != "" {
			d.SetThemeName(theme)
		}

		// Render the dashboard
		html := dashboard.RenderDashboardToHTML(d)

		// Write the response
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	})

	// Start the server
	port := "8080"
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
