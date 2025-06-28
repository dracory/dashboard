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
	
	// Combine all sections
	return welcomeSection + statsSection + chartsSection + activitySection
}

// createWelcomeSection creates the welcome section
func createWelcomeSection() string {
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
										Text("Dashboard"),
								).
								AddChild(
									hb.Div().
										Class("text-muted mt-1").
										Text("Welcome to your admin dashboard"),
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

// createStatsSection creates the statistics section
func createStatsSection() string {
	// Create stats cards
	statsCard1 := components.NewCard(components.CardConfig{
		Content: `
			<div class="d-flex align-items-center">
				<div class="subheader">Sales</div>
				<div class="ms-auto lh-1">
					<div class="dropdown">
						<a class="dropdown-toggle text-muted" href="#" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Last 7 days</a>
						<div class="dropdown-menu dropdown-menu-end">
							<a class="dropdown-item active" href="#">Last 7 days</a>
							<a class="dropdown-item" href="#">Last 30 days</a>
							<a class="dropdown-item" href="#">Last 3 months</a>
						</div>
					</div>
				</div>
			</div>
			<div class="h1 mb-3">$4,300</div>
			<div class="d-flex mb-2">
				<div>Conversion rate</div>
				<div class="ms-auto">
					<span class="text-green d-inline-flex align-items-center lh-1">
						7% <i class="ti ti-trending-up"></i>
					</span>
				</div>
			</div>
			<div class="progress progress-sm">
				<div class="progress-bar bg-primary" style="width: 75%" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" aria-label="75% Complete">
					<span class="visually-hidden">75% Complete</span>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	statsCard2 := components.NewCard(components.CardConfig{
		Content: `
			<div class="d-flex align-items-center">
				<div class="subheader">Revenue</div>
				<div class="ms-auto lh-1">
					<div class="dropdown">
						<a class="dropdown-toggle text-muted" href="#" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Last 7 days</a>
						<div class="dropdown-menu dropdown-menu-end">
							<a class="dropdown-item active" href="#">Last 7 days</a>
							<a class="dropdown-item" href="#">Last 30 days</a>
							<a class="dropdown-item" href="#">Last 3 months</a>
						</div>
					</div>
				</div>
			</div>
			<div class="d-flex align-items-baseline">
				<div class="h1 mb-0 me-2">$8,942</div>
				<div class="me-auto">
					<span class="text-green d-inline-flex align-items-center lh-1">
						8% <i class="ti ti-trending-up"></i>
					</span>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	statsCard3 := components.NewCard(components.CardConfig{
		Content: `
			<div class="d-flex align-items-center">
				<div class="subheader">New clients</div>
				<div class="ms-auto lh-1">
					<div class="dropdown">
						<a class="dropdown-toggle text-muted" href="#" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Last 7 days</a>
						<div class="dropdown-menu dropdown-menu-end">
							<a class="dropdown-item active" href="#">Last 7 days</a>
							<a class="dropdown-item" href="#">Last 30 days</a>
							<a class="dropdown-item" href="#">Last 3 months</a>
						</div>
					</div>
				</div>
			</div>
			<div class="d-flex align-items-baseline">
				<div class="h1 mb-3">6,782</div>
				<div class="me-auto">
					<span class="text-yellow d-inline-flex align-items-center lh-1">
						0% <i class="ti ti-minus"></i>
					</span>
				</div>
			</div>
			<div class="progress progress-sm">
				<div class="progress-bar bg-primary" style="width: 45%" role="progressbar" aria-valuenow="45" aria-valuemin="0" aria-valuemax="100" aria-label="45% Complete">
					<span class="visually-hidden">45% Complete</span>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	statsCard4 := components.NewCard(components.CardConfig{
		Content: `
			<div class="d-flex align-items-center">
				<div class="subheader">Active users</div>
				<div class="ms-auto lh-1">
					<div class="dropdown">
						<a class="dropdown-toggle text-muted" href="#" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">Last 7 days</a>
						<div class="dropdown-menu dropdown-menu-end">
							<a class="dropdown-item active" href="#">Last 7 days</a>
							<a class="dropdown-item" href="#">Last 30 days</a>
							<a class="dropdown-item" href="#">Last 3 months</a>
						</div>
					</div>
				</div>
			</div>
			<div class="d-flex align-items-baseline">
				<div class="h1 mb-3">2,986</div>
				<div class="me-auto">
					<span class="text-red d-inline-flex align-items-center lh-1">
						-2% <i class="ti ti-trending-down"></i>
					</span>
				</div>
			</div>
			<div class="progress progress-sm">
				<div class="progress-bar bg-danger" style="width: 75%" role="progressbar" aria-valuenow="75" aria-valuemin="0" aria-valuemax="100" aria-label="75% Complete">
					<span class="visually-hidden">75% Complete</span>
				</div>
			</div>
		`,
		Margin: 15,
	})
	
	// Create statistics grid
	statsGrid := components.NewGrid(components.GridConfig{
		Columns: []components.GridColumn{
			{Content: statsCard1.ToHTML(), Width: 3},
			{Content: statsCard2.ToHTML(), Width: 3},
			{Content: statsCard3.ToHTML(), Width: 3},
			{Content: statsCard4.ToHTML(), Width: 3},
		},
	})
	
	return statsGrid.ToHTML()
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
