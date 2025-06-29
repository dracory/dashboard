package main

import (
	"github.com/dracory/dashboard/components"
	"github.com/gouniverse/hb"
)

// dashboardContent creates the content for the dashboard page
func dashboardContent() string {
	// Create welcome section
	welcomeSection := createWelcomeSection()
	
	// Create statistics section
	statsSection := createStatsSection()
	
	// Create charts section
	chartsSection := createChartsSection()
	
	// Create recent activity section
	activitySection := createActivitySection()
	
	// Create welcome card with user notifications
	welcomeCardHTML := components.NewCard(components.CardConfig{
		Content: `
			<div class="row">
				<div class="col-md-6">
					<h3 class="mb-1">Welcome back, Admin</h3>
					<p class="text-muted">You have 5 new messages and 2 new notifications.</p>
				</div>
				<div class="col-md-6 text-end">
					<img src="https://preview.tabler.io/static/illustrations/undraw_welcome_3gvl.svg" alt="Welcome illustration" class="img-fluid" style="max-height: 150px;">
				</div>
			</div>
		`,
		Margin: 15,
	}).ToHTML()

	// Combine all sections
	return welcomeSection + welcomeCardHTML + statsSection + chartsSection + activitySection
}

// createWelcomeSection creates the welcome section with user notifications
func createWelcomeSection() string {
	// Create page header with overview title
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
									hb.Div().
										Class("page-pretitle").
										Text("OVERVIEW"),
								).
								AddChild(
									hb.H2().
										Class("page-title").
										Text("Dashboard"),
								),
						).
						AddChild(
							hb.Div().
								Class("col-auto ms-auto d-print-none").
								AddChild(
									hb.Div().
										Class("btn-list").
										AddChild(
											hb.Span().
												Class("d-none d-sm-inline").
												AddChild(
													hb.A().
														Class("btn btn-white").
														Href("#").
														AddChild(
															hb.I().
																Class("ti ti-file-report me-2"),
														).
														Text("Generate report"),
												),
										).
										AddChild(
											hb.A().
												Class("btn btn-primary d-none d-sm-inline-block").
												Href("#").
												AddChild(
													hb.I().
														Class("ti ti-plus"),
												).
												Text(" Create new report"),
										).
										AddChild(
											hb.A().
												Class("btn btn-primary d-sm-none btn-icon").
												Href("#").
												Aria("label", "Create new report").
												AddChild(
													hb.I().
														Class("ti ti-plus"),
												),
										),
								),
						),
				),
		)
	
	return pageHeader.ToHTML()
}

// createStatsSection creates the statistics section with modern cards layout
func createStatsSection() string {
	// Create a container for the stats cards with row and columns
	statsContainer := hb.Div().
		Class("container-xl").
		AddChild(
			hb.Div().
				Class("row row-deck row-cards").
				AddChild(
					// First column - Today's sales with growth rate
					hb.Div().
						Class("col-sm-6 col-lg-3").
						AddChild(
							hb.Div().
								Class("card").
								AddChild(
									hb.Div().
										Class("card-body").
										AddChild(
											hb.Div().
												Class("d-flex align-items-center").
												AddChild(
													hb.Div().
														Class("subheader").
														Text("TODAY'S SALES"),
												),
										).
										AddChild(
											hb.Div().
												Class("d-flex align-items-baseline").
												AddChild(
													hb.Div().
														Class("h1 mb-0 me-2").
														Text("6,782"),
												).
												AddChild(
													hb.Div().
														Class("text-success d-inline-flex align-items-center lh-1").
														HTML("7% <i class='ti ti-trending-up ms-1'></i>"),
												),
										),
								),
						),
				).
				AddChild(
					// Second column - Growth rate
					hb.Div().
						Class("col-sm-6 col-lg-3").
						AddChild(
							hb.Div().
								Class("card").
								AddChild(
									hb.Div().
										Class("card-body").
										AddChild(
											hb.Div().
												Class("d-flex align-items-center").
												AddChild(
													hb.Div().
														Class("subheader").
														Text("GROWTH RATE"),
												),
										).
										AddChild(
											hb.Div().
												Class("d-flex align-items-baseline").
												AddChild(
													hb.Div().
														Class("h1 mb-0 me-2").
														Text("78.4%"),
												).
												AddChild(
													hb.Div().
														Class("text-danger d-inline-flex align-items-center lh-1").
														HTML("-1% <i class='ti ti-trending-down ms-1'></i>"),
												),
										),
								),
						),
				).
				AddChild(
					// Total users column
					hb.Div().
						Class("col-sm-6 col-lg-3").
						AddChild(
							hb.Div().
								Class("card").
								AddChild(
									hb.Div().
										Class("card-body").
										AddChild(
											hb.Div().
												Class("d-flex align-items-center").
												AddChild(
													hb.Div().
														Class("subheader").
														Text("TOTAL USERS"),
												),
										).
										AddChild(
											hb.Div().
												Class("d-flex align-items-baseline").
												AddChild(
													hb.Div().
														Class("h1 mb-0 me-2").
														Text("75,782"),
												).
												AddChild(
													hb.Div().
														Class("text-success d-inline-flex align-items-center lh-1").
														HTML("2% <i class='ti ti-trending-up ms-1'></i>"),
												),
										).
										AddChild(
											hb.Div().
												Class("text-muted mt-1").
												Text("24,635 users increased from last month"),
										),
								),
						),
				).
				AddChild(
					// Active users column
					hb.Div().
						Class("col-sm-6 col-lg-3").
						AddChild(
							hb.Div().
								Class("card").
								AddChild(
									hb.Div().
										Class("card-body").
										AddChild(
											hb.Div().
												Class("d-flex align-items-center").
												AddChild(
													hb.Div().
														Class("subheader").
														Text("ACTIVE USERS"),
												),
										).
										AddChild(
											hb.Div().
												Class("d-flex align-items-baseline").
												AddChild(
													hb.Div().
														Class("h1 mb-0 me-2").
														Text("25,782"),
												).
												AddChild(
													hb.Div().
														Class("text-danger d-inline-flex align-items-center lh-1").
														HTML("-1% <i class='ti ti-trending-down ms-1'></i>"),
												),
										),
								),
						),
				),
		).
		AddChild(
			hb.Div().
				Class("row row-deck row-cards mt-3").
				AddChild(
					hb.Div().
						Class("col-lg-6").
						AddChild(
							hb.Div().
								Class("card").
								AddChild(
									hb.Div().
										Class("card-body").
										AddChild(
											hb.Div().
												Class("d-flex align-items-center").
												AddChild(
													hb.Div().
														Class("subheader").
														Text("ACTIVE SUBSCRIPTIONS"),
												),
										).
										AddChild(
											hb.Div().
												Class("d-flex").
												AddChild(
													hb.Div().
														Class("col-6").
														AddChild(
															hb.Div().
																Class("d-flex align-items-baseline").
																AddChild(
																	hb.Div().
																		Class("h1 mb-0 me-2").
																		Text("2,986"),
																).
																AddChild(
																	hb.Div().
																		Class("text-success d-inline-flex align-items-center lh-1").
																		HTML("4% <i class='ti ti-trending-up ms-1'></i>"),
																),
														),
												).
												AddChild(
													hb.Div().
														Class("text-muted mt-1").
														Text("Last 7 days"),
												),
										).
										AddChild(
											hb.Div().
												Class("col-6").
												AddChild(
													hb.Div().
														Class("chart-container").
														HTML(`
															<div class="chart-circle" style="height: 120px; width: 120px;">
																<div class="chart-circle-value">78%</div>
															</div>
														`),
												),
										),
								),
						),
				),
		)

	return statsContainer.ToHTML()
}

// createChartsSection creates the charts section
func createChartsSection() string {
	// Create chart card
	chartCard := components.NewCard(components.CardConfig{
		Title: "Traffic summary",
		Content: `
			<div id="chart-traffic-summary" style="height: 300px;"></div>
			<div class="mt-3 d-flex">
				<div class="me-3">
					<span class="status-indicator bg-primary"></span>
					<span class="text-muted">Organic Search</span>
				</div>
				<div class="me-3">
					<span class="status-indicator bg-azure"></span>
					<span class="text-muted">Direct Traffic</span>
				</div>
				<div class="me-3">
					<span class="status-indicator bg-green"></span>
					<span class="text-muted">Referral Traffic</span>
				</div>
				<div class="me-3">
					<span class="status-indicator bg-yellow"></span>
					<span class="text-muted">Social Traffic</span>
				</div>
			</div>
			<script>
				// This is a placeholder for the chart
				// In a real implementation, you would use a charting library like ApexCharts
				document.addEventListener("DOMContentLoaded", function() {
					const element = document.getElementById('chart-traffic-summary');
					if (element) {
						element.innerHTML = '<div class="d-flex align-items-center justify-content-center" style="height: 100%;">Chart placeholder - would use ApexCharts in production</div>';
					}
				});
			</script>
		`,
		Margin: 15,
	})
	
	// Create chart grid
	chartGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: chartCard.ToHTML(), Width: 12},
		},
	})
	
	return chartGrid.ToHTML()
}

// createActivitySection creates the recent activity section
func createActivitySection() string {
	// Create activity card
	activityCard := components.NewCard(components.CardConfig{
		Title: "Recent Activity",
		Content: `
			<div class="list-group list-group-flush list-group-hoverable">
				<div class="list-group-item">
					<div class="row align-items-center">
						<div class="col-auto">
							<span class="avatar bg-blue text-white">JL</span>
						</div>
						<div class="col text-truncate">
							<a href="#" class="text-body d-block">John Layman</a>
							<div class="d-block text-muted text-truncate mt-n1">
								Updated the dashboard layout and fixed responsive issues
							</div>
							<div class="text-muted mt-1">2 hours ago</div>
						</div>
						<div class="col-auto">
							<a href="#" class="list-group-item-actions">
								<i class="ti ti-dots"></i>
							</a>
						</div>
					</div>
				</div>
				<div class="list-group-item">
					<div class="row align-items-center">
						<div class="col-auto">
							<span class="avatar" style="background-image: url(https://www.gravatar.com/avatar/2c7d99fe281ecd3bcd65ab915bac6dd5?s=250)"></span>
						</div>
						<div class="col text-truncate">
							<a href="#" class="text-body d-block">Sarah Miller</a>
							<div class="d-block text-muted text-truncate mt-n1">
								Added new analytics charts and improved data visualization
							</div>
							<div class="text-muted mt-1">5 hours ago</div>
						</div>
						<div class="col-auto">
							<a href="#" class="list-group-item-actions">
								<i class="ti ti-dots"></i>
							</a>
						</div>
					</div>
				</div>
				<div class="list-group-item">
					<div class="row align-items-center">
						<div class="col-auto">
							<span class="avatar bg-green text-white">RW</span>
						</div>
						<div class="col text-truncate">
							<a href="#" class="text-body d-block">Robert Wilson</a>
							<div class="d-block text-muted text-truncate mt-n1">
								Deployed new version of the application to production
							</div>
							<div class="text-muted mt-1">Yesterday</div>
						</div>
						<div class="col-auto">
							<a href="#" class="list-group-item-actions">
								<i class="ti ti-dots"></i>
							</a>
						</div>
					</div>
				</div>
				<div class="list-group-item">
					<div class="row align-items-center">
						<div class="col-auto">
							<span class="avatar bg-purple text-white">EP</span>
						</div>
						<div class="col text-truncate">
							<a href="#" class="text-body d-block">Emma Parker</a>
							<div class="d-block text-muted text-truncate mt-n1">
								Fixed bug in user authentication module
							</div>
							<div class="text-muted mt-1">2 days ago</div>
						</div>
						<div class="col-auto">
							<a href="#" class="list-group-item-actions">
								<i class="ti ti-dots"></i>
							</a>
						</div>
					</div>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	// Create activity grid
	activityGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: activityCard.ToHTML(), Width: 12},
		},
	})
	
	return activityGrid.ToHTML()
}
