package adminlte

// templateStyle returns the custom CSS for the template
func templateStyle() string {
	return `
	/* Custom AdminLTE styles */
	.main-header {
		position: fixed;
		width: 100%;
		top: 0;
		z-index: 1030;
	}

	.content-wrapper {
		margin-top: calc(3.5rem + 1px);
	}

	.main-sidebar {
		height: calc(100vh - 3.5rem);
		overflow-y: auto;
	}

	.main-sidebar .nav-link {
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	.main-sidebar .nav-link p {
		display: inline-block;
		margin: 0;
	}

	.main-sidebar .nav-link i {
		margin-right: 0.5rem;
	}

	.main-sidebar .nav-treeview .nav-link {
		padding-left: 2.5rem;
	}

	.brand-link {
		padding: 0.8rem 1rem;
	}

	.user-panel {
		padding: 1rem;
	}

	.navbar-nav .user-menu {
		width: 280px;
	}

	/* Custom scrollbar */
	::-webkit-scrollbar {
		width: 8px;
		height: 8px;
	}

	::-webkit-scrollbar-track {
		background: #f1f1f1;
	}

	::-webkit-scrollbar-thumb {
		background: #888;
		border-radius: 4px;
	}

	::-webkit-scrollbar-thumb:hover {
		background: #555;
	}
	`
}
