package bootstrap

// templateStyle returns the CSS styles for the template
func templateStyle() string {
	return `
		:root {
			--bs-primary: #0d6efd;
			--bs-secondary: #6c757d;
			--bs-success: #198754;
			--bs-info: #0dcaf0;
			--bs-warning: #ffc107;
			--bs-danger: #dc3545;
			--bs-light: #f8f9fa;
			--bs-dark: #212529;
		}

		body {
			overflow-x: hidden;
		}

		.navbar {
			padding: 0.5rem 1rem;
		}

		.navbar-brand {
			padding: 0.3125rem 0;
			margin-right: 1rem;
			font-size: 1.25rem;
			text-decoration: none;
			white-space: nowrap;
		}

		.navbar-toggler {
			padding: 0.25rem 0.75rem;
			font-size: 1.25rem;
			line-height: 1;
			background-color: transparent;
			border: 1px solid transparent;
			border-radius: 0.25rem;
		}

		.dropdown-menu {
			position: absolute;
			top: 100%;
			left: 0;
			z-index: 1000;
			display: none;
			min-width: 10rem;
			padding: 0.5rem 0;
			margin: 0.125rem 0 0;
			font-size: 1rem;
			color: #212529;
			text-align: left;
			list-style: none;
			background-color: #fff;
			background-clip: padding-box;
			border: 1px solid rgba(0, 0, 0, 0.15);
			border-radius: 0.25rem;
		}

		.dropdown-item {
			display: block;
			width: 100%;
			padding: 0.25rem 1rem;
			clear: both;
			font-weight: 400;
			color: #212529;
			text-align: inherit;
			text-decoration: none;
			white-space: nowrap;
			background-color: transparent;
			border: 0;
		}

		.dropdown-item:hover,
		.dropdown-item:focus {
			color: #1e2125;
			background-color: #e9ecef;
		}

		.dropdown-divider {
			height: 0;
			margin: 0.5rem 0;
		overflow: hidden;
		border-top: 1px solid rgba(0, 0, 0, 0.15);
		}

		.offcanvas {
			position: fixed;
			bottom: 0;
			z-index: 1045;
			display: flex;
			flex-direction: column;
			max-width: 100%;
			visibility: hidden;
			background-color: #fff;
			background-clip: padding-box;
			outline: 0;
			transition: transform 0.3s ease-in-out;
		}

		.offcanvas-start {
			top: 0;
			left: 0;
			width: 400px;
			border-right: 1px solid rgba(0, 0, 0, 0.2);
			transform: translateX(-100%);
		}

		.offcanvas-header {
			display: flex;
			align-items: center;
			justify-content: space-between;
			padding: 1rem 1rem;
		}

		.offcanvas-body {
			flex-grow: 1;
			padding: 1rem 1rem;
		overflow-y: auto;
		}

		.modal {
			position: fixed;
			top: 0;
			left: 0;
			z-index: 1050;
			display: none;
			width: 100%;
		height: 100%;
		overflow: hidden;
		outline: 0;
		}

		.modal-dialog {
			position: relative;
			width: auto;
			margin: 0.5rem;
			pointer-events: none;
		}

		.modal-content {
			position: relative;
			display: flex;
			flex-direction: column;
			width: 100%;
			pointer-events: auto;
			background-color: #fff;
			background-clip: padding-box;
			border: 1px solid rgba(0, 0, 0, 0.2);
			border-radius: 0.3rem;
			outline: 0;
		}

		.modal-header {
			display: flex;
			align-items: flex-start;
			justify-content: space-between;
			padding: 1rem 1rem;
			border-bottom: 1px solid #dee2e6;
			border-top-left-radius: 0.3rem;
			border-top-right-radius: 0.3rem;
		}

		.modal-body {
			position: relative;
			flex: 1 1 auto;
			padding: 1rem;
		}

		.btn-close {
			box-sizing: content-box;
			width: 1em;
			height: 1em;
			padding: 0.25em 0.25em;
			color: #000;
			background: transparent url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' fill='%23000'%3e%3cpath d='M.293.293a1 1 0 011.414 0L8 6.586 14.293.293a1 1 0 111.414 1.414L9.414 8l6.293 6.293a1 1 0 01-1.414 1.414L8 9.414l-6.293 6.293a1 1 0 01-1.414-1.414L6.586 8 .293 1.707a1 1 0 010-1.414z'/%3e%3c/svg%3e") center/1em auto no-repeat;
			border: 0;
			border-radius: 0.25rem;
			opacity: 0.5;
		}

		.btn-close:hover {
			color: #000;
			text-decoration: none;
			opacity: 0.75;
		}

		.nav-pills .nav-link {
			border-radius: 0.25rem;
		}

		.nav-pills .nav-link.active,
		.nav-pills .show > .nav-link {
			color: #fff;
			background-color: #0d6efd;
		}

		.nav-link {
			display: block;
			padding: 0.5rem 1rem;
			color: #0d6efd;
			text-decoration: none;
			transition: color 0.15s ease-in-out, background-color 0.15s ease-in-out, border-color 0.15s ease-in-out;
		}

		.nav-link:hover,
		.nav-link:focus {
			color: #0a58ca;
		}

		.nav-link.disabled {
			color: #6c757d;
			pointer-events: none;
			cursor: default;
		}

		.nav-tabs .nav-link {
			margin-bottom: -1px;
			background: 0 0;
			border: 1px solid transparent;
			border-top-left-radius: 0.25rem;
			border-top-right-radius: 0.25rem;
		}

		.nav-tabs .nav-link:hover,
		.nav-tabs .nav-link:focus {
			border-color: #e9ecef #e9ecef #dee2e6;
			isolation: isolate;
		}

		.nav-tabs .nav-link.disabled {
			color: #6c757d;
			background-color: transparent;
			border-color: transparent;
		}

		.nav-tabs .nav-item.show .nav-link,
		.nav-tabs .nav-link.active {
			color: #495057;
			background-color: #fff;
			border-color: #dee2e6 #dee2e6 #fff;
		}

		.nav-tabs .dropdown-menu {
			margin-top: -1px;
			border-top-left-radius: 0;
			border-top-right-radius: 0;
		}
	`
}

// templateScript returns the JavaScript for the template
func templateScript() string {
	return `
		document.addEventListener('DOMContentLoaded', function() {
			// Initialize tooltips
			var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
			var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
				return new bootstrap.Tooltip(tooltipTriggerEl);
			});

			// Initialize popovers
			var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'));
			var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
				return new bootstrap.Popover(popoverTriggerEl);
			});

			// Enable dropdowns
			var dropdownElementList = [].slice.call(document.querySelectorAll('.dropdown-toggle'));
			var dropdownList = dropdownElementList.map(function (dropdownToggleEl) {
				return new bootstrap.Dropdown(dropdownToggleEl);
			});

			// Enable offcanvas
			var offcanvasElementList = [].slice.call(document.querySelectorAll('.offcanvas'));
			var offcanvasList = offcanvasElementList.map(function (offcanvasEl) {
				return new bootstrap.Offcanvas(offcanvasEl);
			});

			// Enable modals
			var modalElementList = [].slice.call(document.querySelectorAll('.modal'));
			var modalList = modalElementList.map(function (modalEl) {
				return new bootstrap.Modal(modalEl);
			});

			// Handle form validation
			var forms = document.querySelectorAll('.needs-validation');
			Array.prototype.slice.call(forms).forEach(function (form) {
				form.addEventListener('submit', function (event) {
					if (!form.checkValidity()) {
						event.preventDefault();
						event.stopPropagation();
					}

					form.classList.add('was-validated');
				}, false);
			});
		});
	`
}
