package tabler

import "github.com/gouniverse/hb"

// TablerCDN provides CDN URLs for Tabler assets
const (
	// TablerCSSCDN is the CDN URL for Tabler CSS
	TablerCSSCDN = "https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler.min.css"

	// TablerJSCDN is the CDN URL for Tabler JS
	TablerJSCDN = "https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/js/tabler.min.js"

	// TablerIconsCDN is the CDN URL for Tabler Icons CSS
	TablerIconsCDN = "https://cdn.jsdelivr.net/npm/@tabler/icons@latest/iconfont/tabler-icons.min.css"

	// TablerDarkModeCSSCDN is the CDN URL for Tabler dark mode CSS
	TablerDarkModeCSSCDN = "https://cdn.jsdelivr.net/npm/@tabler/core@latest/dist/css/tabler-dark.min.css"
)

// GetTablerCDNLinks returns the HTML links for Tabler CDN assets
func GetTablerCDNLinks(includeDarkMode bool) []*hb.Tag {
	links := []*hb.Tag{
		hb.NewLink().Href(TablerCSSCDN).Rel("stylesheet"),
		hb.NewLink().Href(TablerIconsCDN).Rel("stylesheet"),
	}

	if includeDarkMode {
		links = append(links,
			hb.NewLink().Href(TablerDarkModeCSSCDN).Rel("stylesheet"),
		)
	}

	return links
}

// GetTablerCDNScripts returns the HTML scripts for Tabler CDN assets
func GetTablerCDNScripts() []*hb.Tag {
	return []*hb.Tag{
		hb.ScriptURL(TablerJSCDN),
	}
}
