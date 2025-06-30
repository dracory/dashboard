package main

import (
	"github.com/dracory/dashboard/components"
	"github.com/gouniverse/hb"
)

// usersContent creates the content for the users page
func usersContent() string {
	// Create page header
	pageHeader := createUsersPageHeader()
	
	// Create users table
	usersTable := createUsersTable()
	
	// Combine all sections
	return pageHeader + usersTable
}

// createUsersPageHeader creates the page header for the users page
func createUsersPageHeader() string {
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
										Text("Users"),
								).
								AddChild(
									hb.Div().
										Class("text-muted mt-1").
										Text("Manage your users and permissions"),
								),
						).
						AddChild(
							hb.Div().
								Class("col-auto ms-auto d-print-none").
								AddChild(
									hb.Div().
										Class("d-flex").
										AddChild(
											hb.Div().
												Class("me-3").
												AddChild(
													hb.Div().
														Class("input-icon").
														AddChild(
															hb.Span().
																Class("input-icon-addon").
																AddChild(
																	hb.I().
																		Class("ti ti-search"),
																),
														).
														AddChild(
															hb.Input().
																Type("text").
																Class("form-control").
																Placeholder("Search users..."),
														),
												),
										).
										AddChild(
											hb.A().
												Class("btn btn-primary").
												Href("#").
												AddChild(
													hb.I().
														Class("ti ti-plus"),
												).
												Text(" New user"),
										),
								),
						),
				),
		)
	
	return pageHeader.ToHTML()
}

// createUsersTable creates the users table
func createUsersTable() string {
	// Create users card
	usersCard := components.NewCard(components.CardConfig{
		Content: `
			<div class="table-responsive">
				<table class="table table-vcenter card-table table-striped">
					<thead>
						<tr>
							<th>Name</th>
							<th>Email</th>
							<th>Role</th>
							<th>Status</th>
							<th>Created</th>
							<th class="w-1"></th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>
								<div class="d-flex py-1 align-items-center">
									<span class="avatar me-2" style="background-image: url(https://www.gravatar.com/avatar/2c7d99fe281ecd3bcd65ab915bac6dd5?s=250)"></span>
									<div class="flex-fill">
										<div class="font-weight-medium">Sarah Miller</div>
										<div class="text-muted"><small>Admin</small></div>
									</div>
								</div>
							</td>
							<td>
								<div>sarah.miller@example.com</div>
							</td>
							<td>
								<span class="badge bg-purple-lt">Administrator</span>
							</td>
							<td>
								<span class="badge bg-success">Active</span>
							</td>
							<td>
								<div>15 Mar 2023</div>
								<div class="text-muted">10:24 AM</div>
							</td>
							<td>
								<div class="btn-list flex-nowrap">
									<a href="#" class="btn btn-sm btn-outline-primary">
										Edit
									</a>
									<div class="dropdown">
										<button class="btn btn-sm btn-outline-secondary dropdown-toggle" data-bs-toggle="dropdown">
											<i class="ti ti-dots"></i>
										</button>
										<div class="dropdown-menu dropdown-menu-end">
											<a class="dropdown-item" href="#">
												View profile
											</a>
											<a class="dropdown-item" href="#">
												Reset password
											</a>
											<a class="dropdown-item text-danger" href="#">
												Disable account
											</a>
										</div>
									</div>
								</div>
							</td>
						</tr>
						<tr>
							<td>
								<div class="d-flex py-1 align-items-center">
									<span class="avatar me-2 bg-blue text-white">JL</span>
									<div class="flex-fill">
										<div class="font-weight-medium">John Layman</div>
										<div class="text-muted"><small>User</small></div>
									</div>
								</div>
							</td>
							<td>
								<div>john.layman@example.com</div>
							</td>
							<td>
								<span class="badge bg-green-lt">Editor</span>
							</td>
							<td>
								<span class="badge bg-success">Active</span>
							</td>
							<td>
								<div>22 Apr 2023</div>
								<div class="text-muted">3:15 PM</div>
							</td>
							<td>
								<div class="btn-list flex-nowrap">
									<a href="#" class="btn btn-sm btn-outline-primary">
										Edit
									</a>
									<div class="dropdown">
										<button class="btn btn-sm btn-outline-secondary dropdown-toggle" data-bs-toggle="dropdown">
											<i class="ti ti-dots"></i>
										</button>
										<div class="dropdown-menu dropdown-menu-end">
											<a class="dropdown-item" href="#">
												View profile
											</a>
											<a class="dropdown-item" href="#">
												Reset password
											</a>
											<a class="dropdown-item text-danger" href="#">
												Disable account
											</a>
										</div>
									</div>
								</div>
							</td>
						</tr>
						<tr>
							<td>
								<div class="d-flex py-1 align-items-center">
									<span class="avatar me-2 bg-green text-white">RW</span>
									<div class="flex-fill">
										<div class="font-weight-medium">Robert Wilson</div>
										<div class="text-muted"><small>User</small></div>
									</div>
								</div>
							</td>
							<td>
								<div>robert.wilson@example.com</div>
							</td>
							<td>
								<span class="badge bg-blue-lt">Viewer</span>
							</td>
							<td>
								<span class="badge bg-success">Active</span>
							</td>
							<td>
								<div>10 May 2023</div>
								<div class="text-muted">9:45 AM</div>
							</td>
							<td>
								<div class="btn-list flex-nowrap">
									<a href="#" class="btn btn-sm btn-outline-primary">
										Edit
									</a>
									<div class="dropdown">
										<button class="btn btn-sm btn-outline-secondary dropdown-toggle" data-bs-toggle="dropdown">
											<i class="ti ti-dots"></i>
										</button>
										<div class="dropdown-menu dropdown-menu-end">
											<a class="dropdown-item" href="#">
												View profile
											</a>
											<a class="dropdown-item" href="#">
												Reset password
											</a>
											<a class="dropdown-item text-danger" href="#">
												Disable account
											</a>
										</div>
									</div>
								</div>
							</td>
						</tr>
						<tr>
							<td>
								<div class="d-flex py-1 align-items-center">
									<span class="avatar me-2 bg-purple text-white">EP</span>
									<div class="flex-fill">
										<div class="font-weight-medium">Emma Parker</div>
										<div class="text-muted"><small>User</small></div>
									</div>
								</div>
							</td>
							<td>
								<div>emma.parker@example.com</div>
							</td>
							<td>
								<span class="badge bg-orange-lt">Contributor</span>
							</td>
							<td>
								<span class="badge bg-warning">Pending</span>
							</td>
							<td>
								<div>18 Jun 2023</div>
								<div class="text-muted">11:30 AM</div>
							</td>
							<td>
								<div class="btn-list flex-nowrap">
									<a href="#" class="btn btn-sm btn-outline-primary">
										Edit
									</a>
									<div class="dropdown">
										<button class="btn btn-sm btn-outline-secondary dropdown-toggle" data-bs-toggle="dropdown">
											<i class="ti ti-dots"></i>
										</button>
										<div class="dropdown-menu dropdown-menu-end">
											<a class="dropdown-item" href="#">
												View profile
											</a>
											<a class="dropdown-item" href="#">
												Reset password
											</a>
											<a class="dropdown-item text-danger" href="#">
												Disable account
											</a>
										</div>
									</div>
								</div>
							</td>
						</tr>
						<tr>
							<td>
								<div class="d-flex py-1 align-items-center">
									<span class="avatar me-2 bg-red text-white">MB</span>
									<div class="flex-fill">
										<div class="font-weight-medium">Michael Brown</div>
										<div class="text-muted"><small>User</small></div>
									</div>
								</div>
							</td>
							<td>
								<div>michael.brown@example.com</div>
							</td>
							<td>
								<span class="badge bg-blue-lt">Viewer</span>
							</td>
							<td>
								<span class="badge bg-danger">Inactive</span>
							</td>
							<td>
								<div>5 Jul 2023</div>
								<div class="text-muted">2:00 PM</div>
							</td>
							<td>
								<div class="btn-list flex-nowrap">
									<a href="#" class="btn btn-sm btn-outline-primary">
										Edit
									</a>
									<div class="dropdown">
										<button class="btn btn-sm btn-outline-secondary dropdown-toggle" data-bs-toggle="dropdown">
											<i class="ti ti-dots"></i>
										</button>
										<div class="dropdown-menu dropdown-menu-end">
											<a class="dropdown-item" href="#">
												View profile
											</a>
											<a class="dropdown-item" href="#">
												Reset password
											</a>
											<a class="dropdown-item text-success" href="#">
												Enable account
											</a>
										</div>
									</div>
								</div>
							</td>
						</tr>
					</tbody>
				</table>
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
		`,
		Margin: 15,
	})
	
	// Create users grid
	usersGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: usersCard.ToHTML(), Width: 12},
		},
	})
	
	return usersGrid.ToHTML()
}
