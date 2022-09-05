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
	is.Equal("internal/templates/site/defaults.js.ts.gotxt", SveltinProjectFS["defaults"])
	is.Equal("internal/templates/site/externals.js.ts.gotxt", SveltinProjectFS["externals"])
	is.Equal("internal/templates/site/website.js.ts.gotxt", SveltinProjectFS["website"])
	is.Equal("internal/templates/site/init_menu.js.ts.gotxt", SveltinProjectFS["init_menu"])
	is.Equal("internal/templates/site/menu.js.ts.gotxt", SveltinProjectFS["menu"])
}

func TestSveltinResourceFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/lib.gotxt", SveltinResourceFS["lib"])
	is.Equal("internal/templates/resource/page.svelte.gotxt", SveltinResourceFS["index"])
	is.Equal("internal/templates/resource/page.server.ts.gotxt", SveltinResourceFS["indexendpoint"])
	is.Equal("internal/templates/resource/slug.svelte.gotxt", SveltinResourceFS["slug"])
	is.Equal("internal/templates/resource/slug.ts.gotxt", SveltinResourceFS["slugendpoint"])
}

func TestSveltinAPIFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/api/apiIndex.gotxt", SveltinAPIFS["api_index"])
	is.Equal("internal/templates/resource/api/apiSlug.gotxt", SveltinAPIFS["api_slug"])
	is.Equal("internal/templates/resource/api/apiMetadataIndex.gotxt", SveltinAPIFS["api_metadata_index"])
	is.Equal("internal/templates/resource/api/apiMetadataSingle.gotxt", SveltinAPIFS["api_metadata_single"])
	is.Equal("internal/templates/resource/api/apiMetadataList.gotxt", SveltinAPIFS["api_metadata_list"])
}

func TestSveltinMetadataFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/metadata/libList.gotxt", SveltinMetadataFS["lib_list"])
	is.Equal("internal/templates/resource/metadata/page.svelte.gotxt", SveltinMetadataFS["index"])
	is.Equal("internal/templates/resource/metadata/page.server.ts.gotxt", SveltinMetadataFS["indexendpoint"])
	is.Equal("internal/templates/resource/metadata/slug.svelte.gotxt", SveltinMetadataFS["slug"])
	is.Equal("internal/templates/resource/metadata/slug.ts.gotxt", SveltinMetadataFS["slugendpoint"])
}

func TestSveltinPageFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/page/page.svelte.gotxt", SveltinPageFS["svelte"])
	is.Equal("internal/templates/page/page.svx.gotxt", SveltinPageFS["markdown"])
}

func TestSveltinContentFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/content/blank.svx.gotxt", SveltinContentFS["blank"])
	is.Equal("internal/templates/content/sample.svx.gotxt", SveltinContentFS["sample"])
}

func TestSveltinXMLFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/xml/sitemap.xml.gotxt", SveltinXMLFS["sitemap_static"])
	is.Equal("internal/templates/xml/ssr_sitemap.xml.ts.gotxt", SveltinXMLFS["sitemap_ssr"])
	is.Equal("internal/templates/xml/rss.xml.gotxt", SveltinXMLFS["rss_static"])
	is.Equal("internal/templates/xml/ssr_rss.xml.ts.gotxt", SveltinXMLFS["rss_ssr"])
}

func TestBootstrapThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/bootstrap/package.json.gotxt", BootstrapSveltinThemeFS["package_json"])
	is.Equal("internal/templates/themes/sveltin/bootstrap/app.scss", BootstrapSveltinThemeFS["app_css"])
	is.Equal("internal/templates/themes/blank/bootstrap/variables.scss", BootstrapBlankThemeFS["variables_scss"])
}

func TestBulmaThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/bulma/layout.svelte.gotxt", BulmaSveltinThemeFS["layout"])
	is.Equal("internal/templates/themes/sveltin/bulma/app.scss", BulmaSveltinThemeFS["app_css"])
	is.Equal("internal/templates/themes/blank/bulma/variables.scss", BulmaBlankThemeFS["variables_scss"])
}

func TestSCSSThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/scss/package.json.gotxt", SCSSSveltinThemeFS["package_json"])
	is.Equal("internal/templates/themes/sveltin/scss/app.scss", SCSSSveltinThemeFS["app_css"])
	is.Equal("internal/templates/themes/blank/scss/variables.scss", SCSSBlankThemeFS["variables_scss"])
}

func TestTailwindSveltinThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/tailwindcss/postcss.config.cjs", TailwindSveltinThemeFS["postcss"])
	is.Equal("internal/templates/themes/sveltin/tailwindcss/app.css", TailwindSveltinThemeFS["app_css"])
	is.Equal("internal/templates/themes/blank/tailwindcss/tailwind.config.cjs", TailwindBlankThemeFS["tailwind_css_config"])
}

func TestVanillaThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/sveltin/vanillacss/package.json.gotxt", VanillaSveltinThemeFS["package_json"])
	is.Equal("internal/templates/themes/sveltin/vanillacss/app.css", VanillaSveltinThemeFS["app_css"])
	is.Equal("internal/templates/themes/blank/vanillacss/vite.config.ts.gotxt", VanillaBlankThemeFS["vite_config"])
}
