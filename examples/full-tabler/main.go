package main

import (
	"fmt"
	"net/http"
	
	_ "github.com/dracory/dashboard" // Import dashboard package to initialize themes
	dashboard "github.com/dracory/dashboard"
)

func main() {
	// Start the web server
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/dashboard", handleDashboard)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/settings", handleSettings)
	http.HandleFunc("/profile", handleProfile)
	
	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("Navigate to http://localhost:8080/dashboard for the full dashboard example")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	// Create dashboard content
	content := dashboardContent()
	
	// Create dashboard configuration
	config := createDashboardConfig(r, content)
	
	// Create dashboard from config
	dash := dashboard.NewFromConfig(config)
	
	// Render dashboard
	html := dash.ToHTML()
	
	// Write response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Create users content
	content := usersContent()
	
	// Create dashboard configuration
	config := createDashboardConfig(r, content)
	
	// Set active menu item
	for i := range config.MenuItems {
		config.MenuItems[i].Active = config.MenuItems[i].ID == "users"
	}
	
	// Create dashboard from config
	dash := dashboard.NewFromConfig(config)
	
	// Render dashboard
	html := dash.ToHTML()
	
	// Write response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleSettings(w http.ResponseWriter, r *http.Request) {
	// Create settings content
	content := settingsContent()
	
	// Create dashboard configuration
	config := createDashboardConfig(r, content)
	
	// Set active menu item
	for i := range config.MenuItems {
		config.MenuItems[i].Active = config.MenuItems[i].ID == "settings"
	}
	
	// Create dashboard from config
	dash := dashboard.NewFromConfig(config)
	
	// Render dashboard
	html := dash.ToHTML()
	
	// Write response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	// Create profile content
	content := profileContent()
	
	// Create dashboard configuration
	config := createDashboardConfig(r, content)
	
	// Create dashboard from config
	dash := dashboard.NewFromConfig(config)
	
	// Render dashboard
	html := dash.ToHTML()
	
	// Write response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
