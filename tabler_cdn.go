package dashboard

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
func GetTablerCDNLinks(includeDarkMode bool) string {
	links := `<link href="` + TablerCSSCDN + `" rel="stylesheet">
<link href="` + TablerIconsCDN + `" rel="stylesheet">`
	
	if includeDarkMode {
		links += `
<link href="` + TablerDarkModeCSSCDN + `" rel="stylesheet">`
	}
	
	return links
}

// GetTablerCDNScripts returns the HTML scripts for Tabler CDN assets
func GetTablerCDNScripts() string {
	return `<script src="` + TablerJSCDN + `"></script>`
}
