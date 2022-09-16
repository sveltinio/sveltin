// Package helpers ...
package helpers

import (
	"path/filepath"
	"reflect"
	"testing"
	template "text/template"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/resources"
	"gopkg.in/yaml.v3"
)

func TestTemplates(t *testing.T) {
	is := is.New(t)

	var conf config.SveltinConfig
	osFs := afero.NewOsFs()

	yamlFile, err := afero.ReadFile(osFs, filepath.Join("..", "resources", "sveltin.yaml"))
	is.NoErr(err)
	err = yaml.Unmarshal(yamlFile, &conf)
	is.NoErr(err)

	pathToTplFile := resources.SveltinProjectFS["theme_config"]
	data := &config.TemplateData{
		Name: "white",
	}

	tplConfig := BuildTemplate(pathToTplFile, nil, data)
	is.Equal("internal/templates/themes/theme.config.js.gotxt", tplConfig.PathToTplFile)

	var dummyFuncMap template.FuncMap
	is.Equal(reflect.TypeOf(dummyFuncMap), reflect.TypeOf(tplConfig.Funcs))
}

func TestExecSveltinTpl(t *testing.T) {
	is := is.New(t)

	pathToTplFile := resources.SveltinProjectFS["theme_config"]
	data := config.TemplateData{
		Theme: &tpltypes.ThemeData{
			Name: "white",
		},
	}

	tplConfig := BuildTemplate(pathToTplFile, nil, &data)
	retrievedContent := tplConfig.Run(&resources.SveltinFS)

	validContent := `// theme.config.js file for your sveltin theme
const config = {
   name: 'white',
   version: '0.1',
   license: 'MIT',
   licenselink: 'https://github.com/yourname/yourtheme/blob/master/LICENSE',
   description: '',
   homepage: 'http://example.com/',
   tags: [],
   features: [],
   author: {
      name: 'YOUR_NAME_HERE',
      homepage: '',
   },
};

export default config
`
	is.Equal(validContent, string(retrievedContent))
}
