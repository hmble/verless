package config

var (
	// GitTag is injected when building a new release.
	GitTag string = "UNDEFINED"
	// GitCommit stores the latest Git commit.
	GitCommit string = "UNKNOWN"
)

const (
	// Filename is the name of the config file without extension.
	Filename string = "verless"

	// ContentDir is the directory for Markdown content.
	ContentDir string = "content"

	// ThemesDir is the directory for verless themes.
	ThemesDir string = "themes"

	// TemplateDir is the directory for templates inside ThemesDir.
	TemplateDir string = "templates"

	// GeneratedDir is the directory which can be used by hook-commands
	// and which gets ignored by the serve command.
	// The directory can exist in each theme directory and in the StaticDir.
	GeneratedDir string = "generated"

	// CssDir is the directory for CSS files.
	CssDir string = "css"

	// JsDir is the directory for JavaScript files.
	JsDir string = "js"

	// AssetsDir is the directory containing CSS and JavaScript files.
	AssetsDir string = "assets"

	// DefaultTheme is the name of the default theme.
	DefaultTheme string = "default"

	// StaticDir is the directory for static files.
	StaticDir string = "static"

	// OutputDir is the default output directory.
	OutputDir string = "target"

	// IndexFile is the filename used as directory index.
	IndexFile string = "index.html"

	// PageTpl is the template file used for model.Page.
	PageTpl string = "page.html"

	// ListPageTpl is the template file used for model.ListPage.
	ListPageTpl string = "list-page.html"

	// ListPageID is the ID for custom list pages that overwrite
	// a auto-generated list page.
	ListPageID string = "index"
)
