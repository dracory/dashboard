package tabler

// favicon returns the favicon HTML tag
func favicon() string {
	// Return Tabler favicon from CDN
	return `
	<link rel="apple-touch-icon" sizes="180x180" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/apple-touch-icon.png">
	<link rel="icon" type="image/png" sizes="32x32" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="16x16" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/favicon-16x16.png">
	<link rel="manifest" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/site.webmanifest">
	<link rel="mask-icon" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/safari-pinned-tab.svg" color="#206bc4">
	<link rel="shortcut icon" href="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/favicon.ico">
	<meta name="msapplication-TileColor" content="#206bc4">
	<meta name="msapplication-config" content="https://cdn.jsdelivr.net/npm/@tabler/core@1.0.0/static/favicon/browserconfig.xml">
	<meta name="theme-color" content="#206bc4">
	`
}
