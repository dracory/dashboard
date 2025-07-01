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
