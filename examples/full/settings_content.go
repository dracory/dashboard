package main

import (
	"github.com/dracory/dashboard/components"
	"github.com/gouniverse/hb"
)

// settingsContent creates the content for the settings page
func settingsContent() string {
	// Create page header
	pageHeader := createSettingsPageHeader()
	
	// Create settings sections
	generalSettings := createGeneralSettings()
	notificationSettings := createNotificationSettings()
	securitySettings := createSecuritySettings()
	
	// Combine all sections
	return pageHeader + generalSettings + notificationSettings + securitySettings
}

// createSettingsPageHeader creates the page header for the settings page
func createSettingsPageHeader() string {
	// Create page header
	pageHeader := hb.Div().
		Class("page-header d-print-none").
		AddChild(
			hb.Div().
				Class("container-xl").
				AddChild(
					hb.Div().
						Class("row g-2 align-items-center").
						AddChild(
							hb.Div().
								Class("col").
								AddChild(
									hb.H2().
										Class("page-title").
										Text("Settings"),
								).
								AddChild(
									hb.Div().
										Class("text-muted mt-1").
										Text("Manage your application settings"),
								),
						),
				),
		)
	
	return pageHeader.ToHTML()
}

// createGeneralSettings creates the general settings section
func createGeneralSettings() string {
	// Create general settings card
	generalSettingsCard := components.NewCard(components.CardConfig{
		Title:   "General Settings",
		Content: `
			<div class="mb-3">
				<label class="form-label">Site Name</label>
				<input type="text" class="form-control" name="site-name" placeholder="Enter site name" value="My Dashboard">
			</div>
			<div class="mb-3">
				<label class="form-label">Site URL</label>
				<input type="url" class="form-control" name="site-url" placeholder="Enter site URL" value="https://example.com">
			</div>
			<div class="mb-3">
				<label class="form-label">Site Description</label>
				<textarea class="form-control" name="site-description" rows="3" placeholder="Enter site description">A powerful dashboard built with Go and Tabler.</textarea>
			</div>
			<div class="mb-3">
				<label class="form-label">Logo</label>
				<div class="row align-items-center">
					<div class="col-auto">
						<img src="https://preview.tabler.io/static/logo.svg" alt="Current logo" class="avatar avatar-lg">
					</div>
					<div class="col">
						<input type="file" class="form-control" name="logo">
					</div>
				</div>
			</div>
			<div class="mb-3">
				<label class="form-label">Favicon</label>
				<div class="row align-items-center">
					<div class="col-auto">
						<img src="https://preview.tabler.io/favicon.ico" alt="Current favicon" width="16" height="16">
					</div>
					<div class="col">
						<input type="file" class="form-control" name="favicon">
					</div>
				</div>
			</div>
			<div>
				<button type="submit" class="btn btn-primary">Save general settings</button>
			</div>
		`,
		Margin: 15,
	})
	
	// Create general settings grid
	generalSettingsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: generalSettingsCard.ToHTML(), Width: 12},
		},
	})
	
	return generalSettingsGrid.ToHTML()
}

// createNotificationSettings creates the notification settings section
func createNotificationSettings() string {
	// Create notification settings card
	notificationSettingsCard := components.NewCard(components.CardConfig{
		Title:   "Notification Settings",
		Content: `
			<div class="form-selectgroup-boxes row mb-3">
				<div class="col-lg-6">
					<div class="form-selectgroup form-selectgroup-boxes d-flex flex-column">
						<label class="form-selectgroup-item flex-fill">
							<input type="checkbox" name="email-notifications" value="1" class="form-selectgroup-input" checked>
							<div class="form-selectgroup-label d-flex align-items-center p-3">
								<div class="me-3">
									<span class="form-selectgroup-check"></span>
								</div>
								<div>
									<span class="form-selectgroup-title strong mb-1">Email Notifications</span>
									<span class="d-block text-muted">Receive daily email notifications</span>
								</div>
							</div>
						</label>
					</div>
				</div>
				<div class="col-lg-6">
					<div class="form-selectgroup form-selectgroup-boxes d-flex flex-column">
						<label class="form-selectgroup-item flex-fill">
							<input type="checkbox" name="push-notifications" value="1" class="form-selectgroup-input">
							<div class="form-selectgroup-label d-flex align-items-center p-3">
								<div class="me-3">
									<span class="form-selectgroup-check"></span>
								</div>
								<div>
									<span class="form-selectgroup-title strong mb-1">Push Notifications</span>
									<span class="d-block text-muted">Receive push notifications</span>
								</div>
							</div>
						</label>
					</div>
				</div>
			</div>
			<div class="row">
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">Email Frequency</label>
						<select class="form-select">
							<option value="daily" selected>Daily</option>
							<option value="weekly">Weekly</option>
							<option value="monthly">Monthly</option>
							<option value="never">Never</option>
						</select>
					</div>
				</div>
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">Notification Types</label>
						<div class="form-selectgroup">
							<label class="form-selectgroup-item">
								<input type="checkbox" name="notification-type" value="updates" class="form-selectgroup-input" checked>
								<span class="form-selectgroup-label">Updates</span>
							</label>
							<label class="form-selectgroup-item">
								<input type="checkbox" name="notification-type" value="alerts" class="form-selectgroup-input" checked>
								<span class="form-selectgroup-label">Alerts</span>
							</label>
							<label class="form-selectgroup-item">
								<input type="checkbox" name="notification-type" value="messages" class="form-selectgroup-input">
								<span class="form-selectgroup-label">Messages</span>
							</label>
						</div>
					</div>
				</div>
			</div>
			<div>
				<button type="submit" class="btn btn-primary">Save notification settings</button>
			</div>
		`,
		Margin: 15,
	})
	
	// Create notification settings grid
	notificationSettingsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: notificationSettingsCard.ToHTML(), Width: 12},
		},
	})
	
	return notificationSettingsGrid.ToHTML()
}

// createSecuritySettings creates the security settings section
func createSecuritySettings() string {
	// Create security settings card
	securitySettingsCard := components.NewCard(components.CardConfig{
		Title:   "Security Settings",
		Content: `
			<div class="row">
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">Current Password</label>
						<input type="password" class="form-control" name="current-password" placeholder="Enter current password">
					</div>
				</div>
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">Two-Factor Authentication</label>
						<div class="form-check form-switch">
							<input class="form-check-input" type="checkbox" id="two-factor-auth">
							<label class="form-check-label" for="two-factor-auth">Enable two-factor authentication</label>
						</div>
					</div>
				</div>
			</div>
			<div class="row">
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">New Password</label>
						<input type="password" class="form-control" name="new-password" placeholder="Enter new password">
					</div>
				</div>
				<div class="col-lg-6">
					<div class="mb-3">
						<label class="form-label">Confirm New Password</label>
						<input type="password" class="form-control" name="confirm-password" placeholder="Confirm new password">
					</div>
				</div>
			</div>
			<div class="mb-3">
				<label class="form-label">Session Management</label>
				<div class="card">
					<div class="card-body p-3">
						<div class="d-flex align-items-center mb-3">
							<div class="me-3">
								<span class="status-indicator status-green status-indicator-animated"></span>
							</div>
							<div>
								<div class="fw-bold">Current Session</div>
								<div class="text-muted">Windows 10 · Chrome · 192.168.1.100</div>
							</div>
							<div class="ms-auto">
								<span class="text-muted">Active now</span>
							</div>
						</div>
						<div class="d-flex align-items-center">
							<div class="me-3">
								<span class="status-indicator status-gray"></span>
							</div>
							<div>
								<div class="fw-bold">Last Session</div>
								<div class="text-muted">Mac OS · Firefox · 192.168.1.200</div>
							</div>
							<div class="ms-auto">
								<span class="text-muted">2 days ago</span>
							</div>
							<div class="ms-3">
								<a href="#" class="btn btn-sm btn-outline-danger">
									Revoke
								</a>
							</div>
						</div>
					</div>
				</div>
			</div>
			<div>
				<button type="submit" class="btn btn-primary">Save security settings</button>
			</div>
		`,
		Margin: 15,
	})
	
	// Create security settings grid
	securitySettingsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: securitySettingsCard.ToHTML(), Width: 12},
		},
	})
	
	return securitySettingsGrid.ToHTML()
}
