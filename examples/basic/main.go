package main

import (
	"fmt"
	"net/http"

	"github.com/dracory/dashboard"
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
	// Create a new dashboard instance
	d := dashboard.New()

	// Set the dashboard content
	header := `
	<div class="container-fluid">
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Dashboard</h1>
		</div>
	`

	content := `
	<div class="row">
		<div class="col-md-3 mb-4">
			<div class="card bg-primary text-white">
				<div class="card-body">
					<h5 class="card-title">Users</h5>
					<h2 class="mb-0">1,234</h2>
				</div>
			</div>
		</div>
		<div class="col-md-3 mb-4">
			<div class="card bg-success text-white">
				<div class="card-body">
					<h5 class="card-title">Products</h5>
					<h2 class="mb-0">567</h2>
				</div>
			</div>
		</div>
		<div class="col-md-3 mb-4">
			<div class="card bg-warning text-dark">
				<div class="card-body">
					<h5 class="card-title">Orders</h5>
					<h2 class="mb-0">8,901</h2>
				</div>
			</div>
		</div>
		<div class="col-md-3 mb-4">
			<div class="card bg-danger text-white">
				<div class="card-body">
					<h5 class="card-title">Revenue</h5>
					<h2 class="mb-0">$45,678</h2>
				</div>
			</div>
		</div>
	</div>

	<div class="row">
		<div class="col-md-6">
			<div class="card mb-4">
				<div class="card-header">
					<h5 class="card-title mb-0">Recent Activity</h5>
				</div>
				<div class="card-body">
					<div class="list-group list-group-flush">
						<div class="list-group-item">
							<div class="d-flex w-100 justify-content-between">
								<h6 class="mb-1">New user registered</h6>
								<small class="text-muted">2 minutes ago</small>
							</div>
							<p class="mb-1">John Doe created a new account</p>
						</div>
						<div class="list-group-item">
							<div class="d-flex w-100 justify-content-between">
								<h6 class="mb-1">Order #12345 completed</h6>
								<small class="text-muted">1 hour ago</small>
							</div>
							<p class="mb-1">Order has been shipped to the customer</p>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="col-md-6">
			<div class="card">
				<div class="card-header">
					<h5 class="card-title mb-0">Quick Actions</h5>
				</div>
				<div class="card-body">
					<div class="row g-2">
						<div class="col-6">
							<button class="btn btn-primary w-100 mb-2">Add New User</button>
						</div>
						<div class="col-6">
							<button class="btn btn-success w-100 mb-2">Create Product</button>
						</div>
						<div class="col-6">
							<button class="btn btn-info w-100 mb-2">View Reports</button>
						</div>
						<div class="col-6">
							<button class="btn btn-warning w-100 mb-2">Settings</button>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
	`

	d.SetContent(header + content)

	// Get the HTML output and write it to the response
	htmlOutput := d.ToHTML()

	// Set the content type and write the response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlOutput))
}
