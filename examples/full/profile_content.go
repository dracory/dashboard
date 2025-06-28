package main

import (
	"github.com/dracory/dashboard/components"
	"github.com/gouniverse/hb"
)

// profileContent creates the content for the profile page
func profileContent() string {
	// Create page header
	pageHeader := createProfilePageHeader()
	
	// Create profile sections
	profileOverview := createProfileOverview()
	profileDetails := createProfileDetails()
	
	// Combine all sections
	return pageHeader + profileOverview + profileDetails
}

// createProfilePageHeader creates the page header for the profile page
func createProfilePageHeader() string {
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
										Text("Profile"),
								).
								AddChild(
									hb.Div().
										Class("text-muted mt-1").
										Text("Manage your account settings and profile information"),
								),
						),
				),
		)
	
	return pageHeader.ToHTML()
}

// createProfileOverview creates the profile overview section
func createProfileOverview() string {
	// Create profile overview card
	profileOverviewCard := components.NewCard(components.CardConfig{
		Content: `
			<div class="row">
				<div class="col-auto">
					<span class="avatar avatar-xl" style="background-image: url(https://www.gravatar.com/avatar/2c7d99fe281ecd3bcd65ab915bac6dd5?s=250)"></span>
				</div>
				<div class="col">
					<div class="d-flex align-items-center">
						<h2 class="mb-0">Sarah Miller</h2>
						<div class="ms-3">
							<span class="badge bg-purple-lt">Administrator</span>
						</div>
					</div>
					<div class="text-muted">sarah.miller@example.com</div>
					<div class="mt-3">
						<div class="row g-2 align-items-center">
							<div class="col-auto">
								<button class="btn btn-primary">
									<i class="ti ti-edit"></i> Edit profile
								</button>
							</div>
							<div class="col-auto">
								<button class="btn">
									<i class="ti ti-message"></i> Send message
								</button>
							</div>
						</div>
					</div>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	// Create profile overview grid
	profileOverviewGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: profileOverviewCard.ToHTML(), Width: 12},
		},
	})
	
	return profileOverviewGrid.ToHTML()
}

// createProfileDetails creates the profile details section
func createProfileDetails() string {
	// Create profile tabs
	profileTabs := components.NewTab(components.TabConfig{
		Tabs: []components.Tab{
			{
				ID:      "personal-info",
				Title:   "Personal Information",
				Icon:    "ti ti-user",
				Content: createPersonalInfoContent(),
				Active:  true,
			},
			{
				ID:      "account-activity",
				Title:   "Account Activity",
				Icon:    "ti ti-history",
				Content: createAccountActivityContent(),
			},
			{
				ID:      "connected-accounts",
				Title:   "Connected Accounts",
				Icon:    "ti ti-link",
				Content: createConnectedAccountsContent(),
			},
		},
	})
	
	// Create profile details card
	profileDetailsCard := components.NewCard(components.CardConfig{
		Content: profileTabs.ToHTML(),
		Margin:  15,
	})
	
	// Create profile details grid
	profileDetailsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: profileDetailsCard.ToHTML(), Width: 12},
		},
	})
	
	return profileDetailsGrid.ToHTML()
}

// createPersonalInfoContent creates the personal information tab content
func createPersonalInfoContent() string {
	return `
		<div class="row">
			<div class="col-md-6 col-lg-6">
				<div class="mb-3">
					<label class="form-label">First Name</label>
					<input type="text" class="form-control" name="first-name" value="Sarah">
				</div>
			</div>
			<div class="col-md-6 col-lg-6">
				<div class="mb-3">
					<label class="form-label">Last Name</label>
					<input type="text" class="form-control" name="last-name" value="Miller">
				</div>
			</div>
			<div class="col-md-6 col-lg-6">
				<div class="mb-3">
					<label class="form-label">Email</label>
					<input type="email" class="form-control" name="email" value="sarah.miller@example.com">
				</div>
			</div>
			<div class="col-md-6 col-lg-6">
				<div class="mb-3">
					<label class="form-label">Phone</label>
					<input type="tel" class="form-control" name="phone" value="+1 (555) 123-4567">
				</div>
			</div>
			<div class="col-md-12">
				<div class="mb-3">
					<label class="form-label">Bio</label>
					<textarea class="form-control" name="bio" rows="5">Senior administrator with over 10 years of experience in system management and team leadership. Passionate about technology and innovation.</textarea>
				</div>
			</div>
			<div class="col-md-6">
				<div class="mb-3">
					<label class="form-label">Country</label>
					<select class="form-select">
						<option value="us" selected>United States</option>
						<option value="ca">Canada</option>
						<option value="uk">United Kingdom</option>
						<option value="au">Australia</option>
						<option value="de">Germany</option>
					</select>
				</div>
			</div>
			<div class="col-md-6">
				<div class="mb-3">
					<label class="form-label">Language</label>
					<select class="form-select">
						<option value="en" selected>English</option>
						<option value="es">Spanish</option>
						<option value="fr">French</option>
						<option value="de">German</option>
						<option value="zh">Chinese</option>
					</select>
				</div>
			</div>
			<div class="col-md-12">
				<div class="mb-3">
					<label class="form-label">Profile Picture</label>
					<input type="file" class="form-control">
				</div>
			</div>
			<div class="col-md-12">
				<div class="form-group">
					<button type="submit" class="btn btn-primary">Save changes</button>
				</div>
			</div>
		</div>
	`
}

// createAccountActivityContent creates the account activity tab content
func createAccountActivityContent() string {
	return `
		<div class="card">
			<div class="card-body p-0">
				<div class="table-responsive">
					<table class="table table-vcenter card-table table-striped">
						<thead>
							<tr>
								<th>Activity</th>
								<th>Device</th>
								<th>Location</th>
								<th>Date/Time</th>
							</tr>
						</thead>
						<tbody>
							<tr>
								<td>
									<div class="d-flex align-items-center">
										<span class="avatar bg-green-lt me-2"><i class="ti ti-login"></i></span>
										<div>Login successful</div>
									</div>
								</td>
								<td>Chrome on Windows</td>
								<td>San Francisco, CA</td>
								<td>Today, 10:24 AM</td>
							</tr>
							<tr>
								<td>
									<div class="d-flex align-items-center">
										<span class="avatar bg-blue-lt me-2"><i class="ti ti-settings"></i></span>
										<div>Settings updated</div>
									</div>
								</td>
								<td>Chrome on Windows</td>
								<td>San Francisco, CA</td>
								<td>Yesterday, 3:15 PM</td>
							</tr>
							<tr>
								<td>
									<div class="d-flex align-items-center">
										<span class="avatar bg-red-lt me-2"><i class="ti ti-logout"></i></span>
										<div>Logout</div>
									</div>
								</td>
								<td>Chrome on Windows</td>
								<td>San Francisco, CA</td>
								<td>Yesterday, 2:43 PM</td>
							</tr>
							<tr>
								<td>
									<div class="d-flex align-items-center">
										<span class="avatar bg-green-lt me-2"><i class="ti ti-login"></i></span>
										<div>Login successful</div>
									</div>
								</td>
								<td>Chrome on Windows</td>
								<td>San Francisco, CA</td>
								<td>Yesterday, 9:30 AM</td>
							</tr>
							<tr>
								<td>
									<div class="d-flex align-items-center">
										<span class="avatar bg-yellow-lt me-2"><i class="ti ti-key"></i></span>
										<div>Password changed</div>
									</div>
								</td>
								<td>Chrome on Windows</td>
								<td>San Francisco, CA</td>
								<td>Jul 13, 2023, 11:24 AM</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
			<div class="card-footer d-flex align-items-center">
				<p class="m-0 text-muted">Showing <span>1</span> to <span>5</span> of <span>25</span> entries</p>
				<ul class="pagination m-0 ms-auto">
					<li class="page-item disabled">
						<a class="page-link" href="#" tabindex="-1" aria-disabled="true">
							<i class="ti ti-chevron-left"></i>
							prev
						</a>
					</li>
					<li class="page-item active"><a class="page-link" href="#">1</a></li>
					<li class="page-item"><a class="page-link" href="#">2</a></li>
					<li class="page-item"><a class="page-link" href="#">3</a></li>
					<li class="page-item"><a class="page-link" href="#">4</a></li>
					<li class="page-item"><a class="page-link" href="#">5</a></li>
					<li class="page-item">
						<a class="page-link" href="#">
							next <i class="ti ti-chevron-right"></i>
						</a>
					</li>
				</ul>
			</div>
		</div>
	`
}

// createConnectedAccountsContent creates the connected accounts tab content
func createConnectedAccountsContent() string {
	return `
		<div class="card">
			<div class="card-body">
				<div class="mb-4">
					<div class="d-flex align-items-center mb-3">
						<div class="me-3">
							<span class="avatar bg-google text-white"><i class="ti ti-brand-google"></i></span>
						</div>
						<div>
							<div class="font-weight-medium">Google</div>
							<div class="text-muted">sarah.miller@gmail.com</div>
						</div>
						<div class="ms-auto">
							<button class="btn btn-sm btn-danger">Disconnect</button>
						</div>
					</div>
					<div class="text-muted">
						<p>Connected on July 15, 2023</p>
					</div>
				</div>
				<div class="mb-4">
					<div class="d-flex align-items-center mb-3">
						<div class="me-3">
							<span class="avatar bg-azure text-white"><i class="ti ti-brand-twitter"></i></span>
						</div>
						<div>
							<div class="font-weight-medium">Twitter</div>
							<div class="text-muted">@sarahmiller</div>
						</div>
						<div class="ms-auto">
							<button class="btn btn-sm btn-danger">Disconnect</button>
						</div>
					</div>
					<div class="text-muted">
						<p>Connected on August 3, 2023</p>
					</div>
				</div>
				<div class="mb-4">
					<div class="d-flex align-items-center mb-3">
						<div class="me-3">
							<span class="avatar bg-facebook text-white"><i class="ti ti-brand-facebook"></i></span>
						</div>
						<div>
							<div class="font-weight-medium">Facebook</div>
							<div class="text-muted">Not connected</div>
						</div>
						<div class="ms-auto">
							<button class="btn btn-sm btn-primary">Connect</button>
						</div>
					</div>
				</div>
				<div class="mb-4">
					<div class="d-flex align-items-center mb-3">
						<div class="me-3">
							<span class="avatar bg-github text-white"><i class="ti ti-brand-github"></i></span>
						</div>
						<div>
							<div class="font-weight-medium">GitHub</div>
							<div class="text-muted">Not connected</div>
						</div>
						<div class="ms-auto">
							<button class="btn btn-sm btn-primary">Connect</button>
						</div>
					</div>
				</div>
				<div>
					<div class="d-flex align-items-center mb-3">
						<div class="me-3">
							<span class="avatar bg-linkedin text-white"><i class="ti ti-brand-linkedin"></i></span>
						</div>
						<div>
							<div class="font-weight-medium">LinkedIn</div>
							<div class="text-muted">Not connected</div>
						</div>
						<div class="ms-auto">
							<button class="btn btn-sm btn-primary">Connect</button>
						</div>
					</div>
				</div>
			</div>
		</div>
	`
}
