package adminlte

import "github.com/gouniverse/hb"

// AdminLTE provides CDN URLs for AdminLTE assets
const (
	// AdminLTECSS is the CDN URL for AdminLTE CSS
	AdminLTECSS = "https://cdn.jsdelivr.net/npm/admin-lte@3.2/dist/css/adminlte.min.css"

	// AdminLTEJS is the CDN URL for AdminLTE JS
	AdminLTEJS = "https://cdn.jsdelivr.net/npm/admin-lte@3.2/dist/js/adminlte.min.js"

	// FontAwesome is the CDN URL for Font Awesome (required by AdminLTE)
	FontAwesome = "https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css"

	// GoogleFonts is the CDN URL for Google Fonts (used by AdminLTE)
	GoogleFonts = "https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700"

	// BootstrapCSS is the CDN URL for Bootstrap 4 CSS (required by AdminLTE)
	BootstrapCSS = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css"

	// JQuery is the CDN URL for jQuery (required by AdminLTE)
	JQuery = "https://code.jquery.com/jquery-3.6.0.min.js"

	// BootstrapJS is the CDN URL for Bootstrap 4 JS (required by AdminLTE)
	BootstrapJS = "https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/js/bootstrap.bundle.min.js"
)

// GetAdminLTEAssets returns the HTML links for AdminLTE CDN assets
func GetAdminLTEAssets() []*hb.Tag {
	return []*hb.Tag{
		// Google Fonts
		hb.NewLink().Href(GoogleFonts).Rel("stylesheet"),
		// Font Awesome Icons
		hb.NewLink().Href(FontAwesome).Rel("stylesheet"),
		// Bootstrap 4.6.2
		hb.NewLink().Href(BootstrapCSS).Rel("stylesheet"),
		// AdminLTE theme
		hb.NewLink().Href(AdminLTECSS).Rel("stylesheet"),
	}
}

// GetAdminLTEScripts returns the HTML scripts for AdminLTE CDN assets
func GetAdminLTEScripts() []*hb.Tag {
	return []*hb.Tag{
		// jQuery
		hb.ScriptURL(JQuery),
		// Bootstrap 4.6.2 JS Bundle with Popper
		hb.ScriptURL(BootstrapJS),
		// AdminLTE App
		hb.ScriptURL(AdminLTEJS),
	}
}
