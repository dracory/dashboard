package main

import (
	"fmt"
	"net/http"
	
	"github.com/dracory/dashboard"
	"github.com/dracory/dashboard/components"
)

func main() {
	// Start the web server
	http.HandleFunc("/", handleHome)
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Create dashboard content
	content := createDashboardContent()
	
	// Create dashboard configuration
	config := dashboard.Config{
		Content: content,
		FaviconURL: "https://tabler.io/favicon.ico",
		HTTPRequest: r,
		MenuShowText: true,
		MenuType: dashboard.MENU_TYPE_OFFCANVAS,
		LogoImageURL: "https://tabler.io/img/logo.svg",
		LogoRedirectURL: "/",
		NavbarBackgroundColorMode: "dark",
		NavbarTextColor: "light",
		LoginURL: "/login",
		RegisterURL: "/register",
		MenuItems: []dashboard.MenuItem{
			{
				ID:   "dashboard",
				Icon: "ti ti-home",
				Text: "Dashboard",
				URL:  "/",
				Active: true,
			},
			{
				ID:   "users",
				Icon: "ti ti-users",
				Text: "Users",
				URL:  "/users",
			},
			{
				ID:   "settings",
				Icon: "ti ti-settings",
				Text: "Settings",
				URL:  "/settings",
				SubMenu: []dashboard.MenuItem{
					{
						ID:   "profile",
						Icon: "ti ti-user",
						Text: "Profile",
						URL:  "/settings/profile",
					},
					{
						ID:   "security",
						Icon: "ti ti-lock",
						Text: "Security",
						URL:  "/settings/security",
					},
				},
			},
		},
		QuickAccessMenu: []dashboard.MenuItem{
			{
				Icon: "ti ti-user",
				Text: "Profile",
				URL:  "/profile",
			},
			{
				Icon: "ti ti-settings",
				Text: "Settings",
				URL:  "/settings",
			},
			{
				Icon: "ti ti-help",
				Text: "Help",
				URL:  "/help",
			},
		},
		User: dashboard.User{
			ID:    "1",
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		UserMenu: []dashboard.MenuItem{
			{
				Icon: "ti ti-user",
				Text: "Profile",
				URL:  "/profile",
			},
			{
				Icon: "ti ti-settings",
				Text: "Settings",
				URL:  "/settings",
			},
			{
				Icon: "ti ti-logout",
				Text: "Logout",
				URL:  "/logout",
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
}

func createDashboardContent() string {
	// Create welcome card
	welcomeCard := components.NewCard(components.CardConfig{
		Title:   "Welcome to Your Dashboard",
		Content: "<p>This is a basic example of the dashboard using the Tabler template.</p>",
		Margin:  15,
	})
	
	// Create statistics cards
	statsCard1 := components.NewCard(components.CardConfig{
		Title:   "Users",
		Content: "<h3>1,234</h3><p>Total users</p>",
		Margin:  15,
	})
	
	statsCard2 := components.NewCard(components.CardConfig{
		Title:   "Revenue",
		Content: "<h3>$5,678</h3><p>Total revenue</p>",
		Margin:  15,
	})
	
	statsCard3 := components.NewCard(components.CardConfig{
		Title:   "Orders",
		Content: "<h3>9,012</h3><p>Total orders</p>",
		Margin:  15,
	})
	
	// Create statistics grid
	statsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: statsCard1.ToHTML(), Width: 4},
			{Content: statsCard2.ToHTML(), Width: 4},
			{Content: statsCard3.ToHTML(), Width: 4},
		},
	})
	
	// Create tabs
	tabs := components.NewTab(components.TabConfig{
		Tabs: []components.Tab{
			{
				ID:      "overview",
				Title:   "Overview",
				Content: "<p>This is the overview tab content.</p>",
				Active:  true,
				Icon:    "ti ti-chart-bar",
			},
			{
				ID:      "details",
				Title:   "Details",
				Content: "<p>This is the details tab content.</p>",
				Icon:    "ti ti-list",
			},
			{
				ID:      "settings",
				Title:   "Settings",
				Content: "<p>This is the settings tab content.</p>",
				Icon:    "ti ti-settings",
			},
		},
		Class: "mt-3 mb-3", // Using class instead of margin
	})
	
	// Combine all components
	content := welcomeCard.ToHTML() + 
		statsGrid.ToHTML() + 
		components.NewShadowBox(components.ShadowBoxConfig{
			Content:     tabs.ToHTML(),
			ShadowLevel: 2,
			Padding:     15,
			Margin:      15,
		}).ToHTML()
	
	return content
}
