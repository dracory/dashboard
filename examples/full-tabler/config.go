package main

import (
	"fmt"
	"net/http"

	"github.com/dracory/dashboard"
)

// createDashboardConfig creates a dashboard configuration
func createDashboardConfig(r *http.Request, content string) dashboard.Config {
	// Create and return the config
	config := dashboard.Config{
		Content:                   content,
		FaviconURL:                "https://tabler.io/favicon.ico",
		HTTPRequest:               r,
		MenuShowText:              true,
		// MenuType removed as we're now using the two-row navigation layout
		LogoImageURL:              "https://tabler.io/img/logo.svg",
		LogoRedirectURL:           "/dashboard",
		NavbarBackgroundColorMode: "light", // This is just for the navbar appearance
		NavbarTextColor:           "dark",  // This is just for the navbar text color
		LoginURL:                  "/login",
		RegisterURL:              "/register",
		MenuItems:                 createMenuItems(),
		QuickAccessMenu:           createQuickAccessMenu(),
		User:                      createUser(),
		UserMenu:                  createUserMenu(),
		TemplateName:              "tabler", // Explicitly set Tabler template
	}

	// Debug log the config
	if r != nil {
		fmt.Printf("[DEBUG] Creating dashboard config with template: tabler\n")
	}

	return config
}

// createMenuItems creates menu items for the dashboard
func createMenuItems() []dashboard.MenuItem {
	return []dashboard.MenuItem{
		{
			ID:     "dashboard",
			Icon:   "ti ti-home",
			Text:   "Dashboard",
			URL:    "/dashboard",
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
					URL:  "/profile",
				},
				{
					ID:   "security",
					Icon: "ti ti-lock",
					Text: "Security",
					URL:  "/settings/security",
				},
				{
					ID:   "notifications",
					Icon: "ti ti-bell",
					Text: "Notifications",
					URL:  "/settings/notifications",
				},
			},
		},
		{
			ID:   "analytics",
			Icon: "ti ti-chart-bar",
			Text: "Analytics",
			URL:  "#",
		},
		{
			ID:   "reports",
			Icon: "ti ti-file-report",
			Text: "Reports",
			URL:  "#",
		},
		{
			ID:   "help",
			Icon: "ti ti-help",
			Text: "Help & Support",
			URL:  "#",
		},
	}
}

// createQuickAccessMenu creates quick access menu items
func createQuickAccessMenu() []dashboard.MenuItem {
	return []dashboard.MenuItem{
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
			URL:  "#",
		},
		{
			Icon: "ti ti-mail",
			Text: "Messages",
			URL:  "#",
		},
	}
}

// createUser creates a user for the dashboard
func createUser() dashboard.User {
	return dashboard.User{
		ID:        "1",
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		AvatarURL: "https://www.gravatar.com/avatar/00000000000000000000000000000000?d=mp&f=y",
	}
}

// createUserMenu creates user menu items
func createUserMenu() []dashboard.MenuItem {
	return []dashboard.MenuItem{
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
			Icon:      "ti ti-mail",
			Text:      "Messages",
			URL:       "#",
			BadgeText: "5",
		},
		{
			Text: "",
			URL:  "#",
			Icon: "dropdown-divider", // Using Icon field to indicate this is a divider
		},
		{
			Icon: "ti ti-logout",
			Text: "Logout",
			URL:  "#",
		},
	}
}
