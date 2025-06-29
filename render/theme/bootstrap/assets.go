package bootstrap

import "github.com/gouniverse/hb"

// BootstrapCDN provides CDN URLs for Bootstrap 5 assets
const (
	// BootstrapCSSCDN is the CDN URL for Bootstrap 5 CSS
	BootstrapCSSCDN = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css"

	// BootstrapJSCDN is the CDN URL for Bootstrap 5 JS Bundle (includes Popper)
	BootstrapJSCDN = "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"

	// BootstrapIconsCDN is the CDN URL for Bootstrap Icons
	BootstrapIconsCDN = "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.0/font/bootstrap-icons.css"
)

// GetBootstrapCDNLinks returns the HTML links for Bootstrap 5 CDN assets
func GetBootstrapCDNLinks() []*hb.Tag {
	return []*hb.Tag{
		hb.NewLink().Href(BootstrapCSSCDN).Rel("stylesheet"),
		hb.NewLink().Href(BootstrapIconsCDN).Rel("stylesheet"),
	}
}

// GetBootstrapCDNScripts returns the HTML scripts for Bootstrap 5 CDN assets
func GetBootstrapCDNScripts() []*hb.Tag {
	return []*hb.Tag{
		hb.ScriptURL(BootstrapJSCDN),
	}
}
