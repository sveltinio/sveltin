package resources

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetAsciiArt(t *testing.T) {
	is := is.New(t)
	sveltinASCIIArt := `
                _ _   _
               | | | (_)
  _____   _____| | |_ _ _ __
 / __\ \ / / _ \ | __| | '_ \
 \__ \\ V /  __/ | |_| | | | |
 |___/ \_/ \___|_|\__|_|_| |_|

`
	is.Equal(sveltinASCIIArt, GetASCIIArt())
}

func TestSveltinSiteFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/site/defaults.js.ts.gotxt", ProjectFilesMap["defaults"])
	is.Equal("internal/templates/site/externals.js.ts.gotxt", ProjectFilesMap["externals"])
	is.Equal("internal/templates/site/website.js.ts.gotxt", ProjectFilesMap["website"])
	is.Equal("internal/templates/site/init_menu.js.ts.gotxt", ProjectFilesMap["init_menu"])
	is.Equal("internal/templates/site/menu.js.ts.gotxt", ProjectFilesMap["menu"])
}

func TestSveltinResourceFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/lib.gotxt", ResourceFilesMap["lib"])
	is.Equal("internal/templates/resource/themes/blank/page.svelte.gotxt", ResourceFilesMap["index_blank"])
	is.Equal("internal/templates/resource/page.server.ts.gotxt", ResourceFilesMap["indexendpoint"])
	is.Equal("internal/templates/resource/themes/sveltin/slug.svelte.gotxt", ResourceFilesMap["slug_sveltin"])
	is.Equal("internal/templates/resource/slug.ts.gotxt", ResourceFilesMap["slugendpoint"])
}

func TestSveltinAPIFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/api/apiIndex.gotxt", APIFilesMap["api_index"])
	is.Equal("internal/templates/resource/api/apiSlug.gotxt", APIFilesMap["api_slug"])
	is.Equal("internal/templates/resource/api/apiMetadataIndex.gotxt", APIFilesMap["api_metadata_index"])
	is.Equal("internal/templates/resource/api/apiMetadataSingle.gotxt", APIFilesMap["api_metadata_single"])
	is.Equal("internal/templates/resource/api/apiMetadataList.gotxt", APIFilesMap["api_metadata_list"])
}

func TestSveltinMetadataFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/metadata/libList.gotxt", MetadataFilesMap["lib_list"])
	is.Equal("internal/templates/resource/metadata/themes/blank/page.svelte.gotxt", MetadataFilesMap["index_blank"])
	is.Equal("internal/templates/resource/metadata/page.server.ts.gotxt", MetadataFilesMap["indexendpoint"])
	is.Equal("internal/templates/resource/metadata/themes/sveltin/slug.svelte.gotxt", MetadataFilesMap["slug_sveltin"])
	is.Equal("internal/templates/resource/metadata/slug.ts.gotxt", MetadataFilesMap["slugendpoint"])
}

func TestSveltinPageFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/page/themes/blank/page.svelte.gotxt", PageFilesMap["svelte_blank"])
	is.Equal("internal/templates/page/themes/sveltin/page.svx.gotxt", PageFilesMap["markdown_sveltin"])
}

func TestSveltinContentFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/content/blank.svx.gotxt", ContentFilesMap["blank"])
	is.Equal("internal/templates/content/sample.svx.gotxt", ContentFilesMap["sample"])
}

func TestSveltinXMLFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/xml/sitemap.xml.gotxt", XMLFilesMap["sitemap_static"])
	is.Equal("internal/templates/xml/ssr_sitemap.xml.ts.gotxt", XMLFilesMap["sitemap_ssr"])
	is.Equal("internal/templates/xml/rss.xml.gotxt", XMLFilesMap["rss_static"])
	is.Equal("internal/templates/xml/ssr_rss.xml.ts.gotxt", XMLFilesMap["rss_ssr"])
}

func TestBootstrapThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/bootstrap/package.json.gotxt", BootstrapSveltinThemeFilesMap["package_json"])
	is.Equal("internal/templates/themes/sveltin/bootstrap/app.scss", BootstrapSveltinThemeFilesMap["app_css"])
	is.Equal("internal/templates/themes/blank/bootstrap/variables.scss", BootstrapBlankThemeFilesMap["variables_scss"])
}

func TestBulmaThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/bulma/layout.svelte.gotxt", BulmaSveltinThemeFilesMap["layout"])
	is.Equal("internal/templates/themes/sveltin/bulma/app.scss", BulmaSveltinThemeFilesMap["app_css"])
	is.Equal("internal/templates/themes/blank/bulma/variables.scss", BulmaBlankThemeFilesMap["variables_scss"])
}

func TestSCSSThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/scss/package.json.gotxt", SassSveltinThemeFilesMap["package_json"])
	is.Equal("internal/templates/themes/sveltin/scss/app.scss", SassSveltinThemeFilesMap["app_css"])
	is.Equal("internal/templates/themes/blank/scss/variables.scss", SassBlankThemeFilesMap["variables_scss"])
}

func TestTailwindSveltinThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/tailwindcss/postcss.config.cjs", TailwindSveltinThemeFilesMap["postcss"])
	is.Equal("internal/templates/themes/sveltin/tailwindcss/app.css", TailwindSveltinThemeFilesMap["app_css"])
	is.Equal("internal/templates/themes/blank/tailwindcss/tailwind.config.cjs", TailwindBlankThemeFilesMap["tailwind_css_config"])
}

func TestVanillaThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/vanillacss/package.json.gotxt", VanillaSveltinThemeFilesMap["package_json"])
	is.Equal("internal/templates/themes/sveltin/vanillacss/app.css", VanillaSveltinThemeFilesMap["app_css"])
	is.Equal("internal/templates/themes/blank/vanillacss/vite.config.ts.gotxt", VanillaBlankThemeFilesMap["vite_config"])
}
