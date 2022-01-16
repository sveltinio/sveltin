package resources

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetAsciiArt(t *testing.T) {
	is := is.New(t)
	sveltinAsciiArt := `
                _ _   _
               | | | (_)
  _____   _____| | |_ _ _ __
 / __\ \ / / _ \ | __| | '_ \
 \__ \\ V /  __/ | |_| | | | |
 |___/ \_/ \___|_|\__|_|_| |_|

`
	is.Equal(sveltinAsciiArt, GetAsciiArt())
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
	is.Equal("internal/templates/resource/api.gotxt", SveltinResourceFS["api"])
	is.Equal("internal/templates/resource/lib.gotxt", SveltinResourceFS["lib"])
	is.Equal("internal/templates/resource/index.gotxt", SveltinResourceFS["index"])
	is.Equal("internal/templates/resource/slug.gotxt", SveltinResourceFS["slug"])
}

func TestSveltinMetadataFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/resource/metadata/apiSingle.gotxt", SveltinMetadataFS["api_single"])
	is.Equal("internal/templates/resource/metadata/libSingle.gotxt", SveltinMetadataFS["lib_single"])
	is.Equal("internal/templates/resource/metadata/apiList.gotxt", SveltinMetadataFS["api_list"])
	is.Equal("internal/templates/resource/metadata/libList.gotxt", SveltinMetadataFS["lib_list"])
	is.Equal("internal/templates/resource/metadata/index.gotxt", SveltinMetadataFS["index"])
	is.Equal("internal/templates/resource/metadata/slug.gotxt", SveltinMetadataFS["slug"])
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

func TestSveltinThemeFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/templates/themes/tailwindcss/tailwind.config.cjs", SveltinTailwindCSSThemeFS["tailwind_css_config"])
	is.Equal("internal/templates/themes/tailwindcss/app.css", SveltinTailwindCSSThemeFS["app_css"])
	is.Equal("internal/templates/themes/tailwindcss/postcss.config.cjs", SveltinTailwindCSSThemeFS["postcss"])
}
