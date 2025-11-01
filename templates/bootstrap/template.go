package bootstrap

import (
	"github.com/dracory/cdn"
	"github.com/dracory/dashboard/templates/shared"
	"github.com/dracory/dashboard/types"
	"github.com/dracory/hb"
	"github.com/samber/lo"
)

// Template implements the types.TemplateInterface for Bootstrap-based templates
type Template struct{}

// Ensure Template implements the TemplateInterface
var _ types.TemplateInterface = (*Template)(nil)

// layout generates the main layout structure for the dashboard
func (t *Template) layout(dashboard types.DashboardInterface) string {
	content := dashboard.GetContent()
	layout := hb.NewBorderLayout()
	layout.AddTop(hb.Raw(topNavigation(dashboard)), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_MIDDLE)
	layout.AddCenter(hb.Raw(content), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_TOP)
	return layout.ToHTML()
}

func (t *Template) getStylesAndScripts(dashboard types.DashboardInterface) (
	styleURLs []string,
	scriptURLs []string,
	styles []string,
	scripts []string,
) {
	styleURLs = make([]string, 0)
	scriptURLs = make([]string, 0)
	styles = make([]string, 0)
	scripts = make([]string, 0)

	// Add CSS
	if style := templateStyle(); style != "" {
		styles = append(styles, style)
	}

	// Add JavaScript
	if script := templateScript(); script != "" {
		scripts = append(scripts, script)
	}

	// Add theme CSS or default Bootstrap CSS
	themeName := dashboard.GetTheme()
	themeName = themeNameVerifyAndFix(themeName)

	if themeName != "" && themeName != THEME_DEFAULT {
		// Check if it's a known theme
		_, isLightTheme := themesLight[themeName]
		_, isDarkTheme := themesDark[themeName]

		if isLightTheme || isDarkTheme {
			// Use Bootswatch theme
			styleURLs = append(styleURLs, "https://cdn.jsdelivr.net/npm/bootswatch@5.3.2/dist/"+themeName+"/bootstrap.min.css")
		} else {
			// Fallback to default Bootstrap
			styleURLs = append(styleURLs, cdn.BootstrapCss_5_3_3())
		}
	} else {
		// Use default Bootstrap
		styleURLs = append(styleURLs, cdn.BootstrapCss_5_3_3())
	}

	// Add Bootstrap Icons
	styleURLs = append(styleURLs, cdn.BootstrapIconsCss_1_11_3())

	// Add Bootstrap JS Bundle with Popper
	scriptURLs = append(scriptURLs, cdn.BootstrapJs_5_3_3())

	styleURLs = append(styleURLs, dashboard.GetStyleURLs()...)
	scriptURLs = append(scriptURLs, dashboard.GetScriptURLs()...)

	styles = append(styles, dashboard.GetStyles()...)
	scripts = append(scripts, dashboard.GetScripts()...)

	return lo.Uniq(styleURLs), lo.Uniq(scriptURLs), lo.Uniq(styles), lo.Uniq(scripts)
}

// ToHTML generates the complete HTML for the dashboard page
func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	styleURLs, scriptURLs, styles, scripts := t.getStylesAndScripts(dashboard)

	// Create a new webpage
	webpage := hb.Webpage()

	// Set the page title
	webpage.SetTitle(dashboard.GetTitle())

	// Add favicon
	if favicon := shared.Favicon(); favicon != "" {
		webpage.SetFavicon(favicon)
	}

	// Add CSS URLs
	for _, styleURL := range styleURLs {
		webpage.AddStyleURL(styleURL)
	}

	// Add JavaScript URLs
	for _, scriptURL := range scriptURLs {
		webpage.AddScriptURL(scriptURL)
	}

	// Add CSS
	for _, style := range styles {
		webpage.AddStyle(style)
	}

	// Add JavaScript
	for _, script := range scripts {
		webpage.AddScript(script)
	}

	// Generate the layout
	layoutHTML := t.layout(dashboard)

	// Add the layout to the webpage
	webpage.Body().Child(hb.Raw(layoutHTML))

	// Handle redirect if needed
	if redirectURL := dashboard.GetRedirectUrl(); redirectURL != "" {
		redirectTime := "0"
		if dashboard.GetRedirectTime() != "" {
			redirectTime = dashboard.GetRedirectTime()
		}
		webpage.Meta(hb.Meta().
			Attr("http-equiv", "refresh").
			Attr("content", redirectTime+"; url = "+redirectURL))
	}

	// Generate the final HTML
	return webpage.ToHTML()
}
