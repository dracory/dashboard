package tabler

// templateStyle returns the custom CSS styles for the template
func templateStyle() string {
	return `
	/* Custom styles for Tabler template */
	:root {
		--tblr-font-sans-serif: 'Inter', -apple-system, BlinkMacSystemFont, San Francisco, Segoe UI, Roboto, Helvetica Neue, sans-serif;
	}

	/* Avatar styles */
	.avatar {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		width: 2.5rem;
		height: 2.5rem;
		font-size: 1rem;
		font-weight: 600;
		border-radius: 0.375rem;
	}

	.avatar.avatar-sm {
		width: 2rem;
		height: 2rem;
		font-size: 0.75rem;
	}

	.avatar.avatar-md {
		width: 2.5rem;
		height: 2.5rem;
		font-size: 1rem;
	}

	.avatar.avatar-lg {
		width: 3rem;
		height: 3rem;
		font-size: 1.25rem;
	}

	.avatar.rounded-circle {
		border-radius: 50% !important;
	}

	.avatar-initials {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	/* Custom scrollbar */
	::-webkit-scrollbar {
		width: 8px;
		height: 8px;
	}

	::-webkit-scrollbar-track {
		background: #f1f5f9;
	}

	::-webkit-scrollbar-thumb {
		background: #cbd5e1;
		border-radius: 4px;
	}

	::-webkit-scrollbar-thumb:hover {
		background: #94a3b8;
	}

	/* Custom form controls */
	.form-control:focus, .form-select:focus, .form-check-input:focus {
		box-shadow: 0 0 0 0.2rem rgba(32, 107, 196, 0.25);
		border-color: #206bc4;
	}

	/* Custom buttons */
	.btn-primary {
		background-color: #206bc4;
		border-color: #1a5da3;
	}

	.btn-primary:hover, .btn-primary:focus {
		background-color: #1a5da3;
		border-color: #164d87;
	}

	.btn-outline-primary {
		color: #206bc4;
		border-color: #206bc4;
	}

	.btn-outline-primary:hover, .btn-outline-primary:focus {
		background-color: #206bc4;
		border-color: #206bc4;
	}

	/* Custom cards */
	.card {
		border: 1px solid rgba(0, 0, 0, 0.1);
		box-shadow: 0 0.125rem 0.25rem rgba(0, 0, 0, 0.075);
		margin-bottom: 1.5rem;
	}

	.card-header {
		background-color: #f8f9fa;
		border-bottom: 1px solid rgba(0, 0, 0, 0.1);
		padding: 1rem 1.25rem;
	}

	.card-title {
		margin-bottom: 0;
		font-weight: 600;
	}

	/* Custom tables */
	.table {
		margin-bottom: 0;
	}

	.table > :not(:last-child) > :last-child > * {
		border-bottom-color: #e9ecef;
	}

	.table > thead {
		background-color: #f8f9fa;
	}

	/* Custom alerts */
	.alert {
		border: 1px solid transparent;
		border-radius: 0.25rem;
	}

	/* Custom badges */
	.badge {
		font-weight: 500;
		padding: 0.35em 0.65em;
	}

	/* Custom dropdowns */
	.dropdown-menu {
		border: 1px solid rgba(0, 0, 0, 0.1);
		box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.1);
	}

	.dropdown-item {
		padding: 0.5rem 1rem;
	}

	/* Custom tabs */
	.nav-tabs .nav-link {
		border: 1px solid transparent;
		border-radius: 0.25rem 0.25rem 0 0;
		padding: 0.5rem 1rem;
	}

	.nav-tabs .nav-link.active {
		background-color: #fff;
		border-color: #dee2e6 #dee2e6 #fff;
	}

	/* Custom pagination */
	.pagination .page-link {
		color: #206bc4;
	}

	.pagination .page-item.active .page-link {
		background-color: #206bc4;
		border-color: #1a5da3;
	}

	/* Custom tooltips */
	.tooltip {
		pointer-events: none;
	}

	/* Custom popovers */
	.popover {
		border: 1px solid rgba(0, 0, 0, 0.1);
		box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.1);
	}

	/* Custom modals */
	.modal-content {
		border: 1px solid rgba(0, 0, 0, 0.1);
		box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.1);
	}

	/* Custom toasts */
	.toast {
		border: 1px solid rgba(0, 0, 0, 0.1);
		box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.1);
	}

	/* Responsive adjustments */
	@media (max-width: 767.98px) {
		.navbar-brand {
			font-size: 1rem;
		}

		.navbar-nav .nav-link {
			padding: 0.5rem 0.75rem;
		}

		.card {
			margin-bottom: 1rem;
		}
	}

	/* Dark mode support */
	[data-bs-theme=dark] {
		--tblr-bg-surface: #1a1a1a;
		--tblr-body-bg: #121212;
		--tblr-body-color: #e0e0e0;
		--tblr-border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .card {
		background-color: #1e1e1e;
		border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .card-header {
		background-color: #252525;
		border-bottom-color: #2d2d2d;
	}

	[data-bs-theme=dark] .table {
		--tblr-table-color: #e0e0e0;
		--tblr-table-bg: transparent;
		--tblr-table-border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .table > :not(:last-child) > :last-child > * {
		border-bottom-color: #2d2d2d;
	}

	[data-bs-theme=dark] .table > thead {
		--tblr-table-bg: #252525;
	}

	[data-bs-theme=dark] .dropdown-menu {
		background-color: #252525;
		border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .dropdown-item {
		color: #e0e0e0;
	}

	[data-bs-theme=dark] .dropdown-item:hover, 
	[data-bs-theme=dark] .dropdown-item:focus {
		background-color: #2d2d2d;
		color: #ffffff;
	}

	[data-bs-theme=dark] .dropdown-divider {
		border-top-color: #2d2d2d;
	}

	[data-bs-theme=dark] .nav-tabs .nav-link.active {
		background-color: #1e1e1e;
		border-color: #2d2d2d #2d2d2d #1e1e1e;
	}

	[data-bs-theme=dark] .form-control,
	[data-bs-theme=dark] .form-select,
	[data-bs-theme=dark] .form-check-input {
		background-color: #252525;
		border-color: #2d2d2d;
		color: #e0e0e0;
	}

	[data-bs-theme=dark] .form-control:focus,
	[data-bs-theme=dark] .form-select:focus,
	[data-bs-theme=dark] .form-check-input:focus {
		background-color: #252525;
		border-color: #206bc4;
		color: #e0e0e0;
	}

	[data-bs-theme=dark] .form-control:disabled,
	[data-bs-theme=dark] .form-select:disabled {
		background-color: #1a1a1a;
	}

	[data-bs-theme=dark] .form-control::placeholder {
		color: #6c757d;
	}

	[data-bs-theme=dark] .input-group-text {
		background-color: #252525;
		border-color: #2d2d2d;
		color: #e0e0e0;
	}

	[data-bs-theme=dark] .modal-content {
		background-color: #1e1e1e;
		border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .modal-header,
	[data-bs-theme=dark] .modal-footer {
		border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .btn-close {
		filter: invert(1) grayscale(100%) brightness(200%);
	}

	[data-bs-theme=dark] .toast {
		background-color: #1e1e1e;
		border-color: #2d2d2d;
	}

	[data-bs-theme=dark] .toast-header {
		background-color: #252525;
		border-bottom-color: #2d2d2d;
	}

	/* Print styles */
	@media print {
		.navbar,
		.footer,
		.btn-print {
			display: none !important;
		}

		.card {
			border: 1px solid #e9ecef;
			box-shadow: none;
		}

		@page {
			size: A4;
			margin: 1cm;
		}

		body {
			padding: 0 !important;
		}

		.container {
			width: 100%;
			max-width: 100%;
			padding: 0;
		}

		.table {
			page-break-inside: auto;
		}

		tr {
			page-break-inside: avoid;
			page-break-after: auto;
		}

		td, th {
			page-break-inside: avoid;
			page-break-after: auto;
		}
	}
	.navbar {
		transition: all 0.2s ease-in-out;
	}
	
	.navbar-brand {
		font-weight: 600;
	}
	
	.navbar-nav .nav-link {
		display: flex;
		align-items: center;
	}
	
	.avatar {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		width: 2rem;
		height: 2rem;
		border-radius: 50%;
		overflow: hidden;
		background-color: #f5f5f5;
	}
	
	.avatar-img {
		width: 100%;
		height: 100%;
		object-fit: cover;
	}
	
	.avatar-initials {
		font-weight: 600;
		font-size: 0.875rem;
		color: #495057;
	}
	
	.dropdown-menu {
		border: 0;
		box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
	}
	
	.dropdown-item {
		display: flex;
		align-items: center;
		padding: 0.5rem 1rem;
	}
	
	.dropdown-item i {
		margin-right: 0.5rem;
	}
	
	/* Responsive adjustments */
	@media (max-width: 767.98px) {
		.navbar-collapse {
			padding: 1rem;
			background: #fff;
			border-radius: 0.25rem;
			box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
		}
		
		.navbar-nav {
			margin-top: 1rem;
		}
		
		.nav-item {
			margin-bottom: 0.25rem;
		}
	}
	`
}
