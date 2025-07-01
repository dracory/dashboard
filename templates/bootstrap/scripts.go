package bootstrap

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

			// Enable offcanvas - using a more defensive approach
			var offcanvasElementList = document.querySelectorAll('.offcanvas');
			if (offcanvasElementList.length > 0) {
				Array.prototype.slice.call(offcanvasElementList).forEach(function (offcanvasEl) {
					try {
						new bootstrap.Offcanvas(offcanvasEl);
					} catch (e) {
						console.warn('Error initializing offcanvas:', e);
					}
				});
			}

			// Enable modals - using a more defensive approach
			var modalElementList = document.querySelectorAll('.modal');
			if (modalElementList.length > 0) {
				Array.prototype.slice.call(modalElementList).forEach(function (modalEl) {
					try {
						new bootstrap.Modal(modalEl);
					} catch (e) {
						console.warn('Error initializing modal:', e);
					}
				});
			}

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
