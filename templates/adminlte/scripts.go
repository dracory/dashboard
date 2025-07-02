package adminlte

// templateScript returns the custom JavaScript for the template
func templateScript() string {
	return `
	// Initialize AdminLTE
	document.addEventListener('DOMContentLoaded', function() {
		// Enable tooltips
		$('[data-toggle="tooltip"]').tooltip();

		// Enable popovers
		$('[data-toggle="popover"]').popover();

		// Initialize sidebar menu
		$('.nav-sidebar a').each(function() {
			const location = window.location.href;
			const link = this.href;
			if (location === link) {
				$(this).addClass('active');
				$(this).closest('.has-treeview').addClass('menu-open');
			}
		});

		// Auto-expand active menu items
		$('.nav-sidebar .nav-treeview').each(function() {
			if ($(this).find('.nav-link.active').length > 0) {
				$(this).parent().addClass('menu-open');
			}
		});

		// Toggle sidebar mini
		$('[data-widget="pushmenu"]').on('click', function() {
			setTimeout(function() {
				window.dispatchEvent(new Event('resize'));
			}, 300);
		});

		// Handle theme switching
		$('.theme-switch').on('click', function(e) {
			e.preventDefault();
			const theme = $(this).data('theme');
			
			// Remove all theme classes
			$('body').removeClass((index, className) => {
				return (className.match(/(^|\s)skin-\S+/g) || []).join(' ');
			});

			// Add the selected theme class
			if (theme !== 'default') {
				$('body').addClass('skin-' + theme);
			}

			// Save the theme preference
			localStorage.setItem('adminlte-theme', theme);

			// Refresh the page to apply the theme properly
			window.location.reload();
		});

		// Load saved theme
		const savedTheme = localStorage.getItem('adminlte-theme');
		if (savedTheme && savedTheme !== 'default') {
			$('body').addClass('skin-' + savedTheme);
		}

		// Handle dropdown menus
		$('.dropdown-toggle').on('click', function(e) {
			e.stopPropagation();
			$(this).next('.dropdown-menu').toggleClass('show');
		});

		// Close dropdowns when clicking outside
		$(document).on('click', function(e) {
			if (!$(e.target).closest('.dropdown').length) {
				$('.dropdown-menu').removeClass('show');
			}
		});

		// Initialize any plugins that need it
		if (typeof $.fn.overlayScrollbars === 'function') {
			$('.sidebar').overlayScrollbars({
				scrollbars: {
					autoHide: 'leave',
					clickScrolling: true
				}
			});
		}

		// Handle form submissions with loading states
		$('form[data-ajax="true"]').on('submit', function() {
			const $form = $(this);
			const $submitBtn = $form.find('[type="submit"]');
			
			// Disable the submit button and show loading state
			$submitBtn.prop('disabled', true);
			$submitBtn.html('<i class="fas fa-spinner fa-spin"></i> ' + $submitBtn.text());
		});

		// Handle external links
		$('a[target="_blank"]').each(function() {
			$(this).append(' <i class="fas fa-external-link-alt fa-xs"></i>');
		});

		// Helper function to show loading overlay
		window.showLoadingOverlay = function() {
			$('body').append(
				'<div class="overlay-wrapper">' +
				'<div class="overlay">' +
				'<i class="fas fa-3x fa-sync-alt fa-spin"></i>' +
				'<div class="text-bold pt-2">Loading...</div>' +
				'</div>' +
				'</div>'
			);
		};

		// Helper function to hide loading overlay
		window.hideLoadingOverlay = function() {
			$('.overlay-wrapper').remove();
		};
	});
	`
}
