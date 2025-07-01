package tabler

// templateScript returns any additional JavaScript for the template
func templateScript() string {
	return `
	// Custom JavaScript for Tabler template
	document.addEventListener('DOMContentLoaded', function() {
		// Initialize tooltips
		var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
		var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
			return new bootstrap.Tooltip(tooltipTriggerEl);
		});

		// Initialize popovers
		// Enable popovers
		var popoverTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]'));
		var popoverList = popoverTriggerList.map(function (popoverTriggerEl) {
			return new bootstrap.Popover(popoverTriggerEl);
		});

		// Handle dropdown submenus
		document.querySelectorAll('.dropdown-menu a.dropdown-toggle').forEach(function(element) {
			element.addEventListener('click', function(e) {
				e.preventDefault();
				e.stopPropagation();
				
				var submenu = this.nextElementSibling;
				var parentItem = this.parentElement;
				
				// Close all other open submenus at this level
				parentItem.parentElement.querySelectorAll('.dropdown-menu.show').forEach(function(openMenu) {
					if (openMenu !== submenu) {
						openMenu.classList.remove('show');
					}
				});
				
				// Toggle current submenu
				submenu.classList.toggle('show');
			});
		});

		// Close dropdowns when clicking outside
		document.addEventListener('click', function(e) {
			if (!e.target.matches('.dropdown-toggle')) {
				document.querySelectorAll('.dropdown-menu.show').forEach(function(openMenu) {
					openMenu.classList.remove('show');
				});
			}
		});

		// Handle mobile menu toggle
		document.querySelectorAll('.navbar-toggler').forEach(function(button) {
			button.addEventListener('click', function() {
				document.querySelector('.navbar-collapse').classList.toggle('show');
			});
		});
	});

	// Add active class to current nav item
	document.addEventListener('DOMContentLoaded', function() {
		var currentPath = window.location.pathname;
		document.querySelectorAll('.nav-link').forEach(function(link) {
			if (link.getAttribute('href') === currentPath) {
				link.classList.add('active');
				// Also activate parent dropdown if exists
				let parentDropdown = link.closest('.dropdown-menu');
				if (parentDropdown) {
					parentDropdown.previousElementSibling?.classList.add('active');
				}
			}
		});
	});
	`
}
