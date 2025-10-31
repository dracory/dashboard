package adminlte

import (
	"github.com/dracory/dashboard/types"
	"github.com/dracory/hb"
)

// menuModal generates the modal menu HTML
func menuModal(dashboard types.DashboardInterface) *hb.Tag {
	// Create modal container
	modal := hb.Div().Class("modal fade").ID("mainMenuModal")
	modal.Attr("tabindex", "-1")
	modal.Attr("role", "dialog")

	// Modal dialog
	modalDialog := hb.Div().Class("modal-dialog")
	modal.Child(modalDialog)

	// Modal content
	modalContent := hb.Div().Class("modal-content")
	modalDialog.Child(modalContent)

	// Modal header
	header := hb.Div().Class("modal-header")
	header.Child(hb.H5().Class("modal-title").Text(dashboard.GetTitle()))
	header.Child(hb.Button().Class("close").Attr("data-dismiss", "modal").HTML("&times;"))
	modalContent.Child(header)

	// Modal body with menu
	body := hb.Div().Class("modal-body")
	menu := BuildSidebarMenu(dashboard)
	if menu != nil {
		body.Child(menu)
	}
	modalContent.Child(body)

	// Add script to handle modal
	script := hb.Script(`
		document.addEventListener('DOMContentLoaded', function() {
			var menuModal = document.getElementById('mainMenuModal');
			if (menuModal) {
				// Initialize modal
				var modal = new bootstrap.Modal(menuModal);
				
				// Show the modal immediately
				modal.show();
				
				// Handle menu item clicks
				menuModal.addEventListener('click', function(event) {
					var target = event.target.closest('.nav-link');
					if (target) {
						event.preventDefault();
						modal.hide();
					}
				});

				// Handle close button
				var closeButton = menuModal.querySelector('[data-dismiss="modal"]');
				if (closeButton) {
					closeButton.addEventListener('click', function() {
						modal.hide();
					});
				}
			}
		});
	`)

	// Add script to the modal
	modal.Child(script)

	return modal
}
