package helpers

import (
	"bytes"
	"embed"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/sveltinlib/builder"
)

func IsValidFileForContent(f fs.FileInfo) bool {
	acceptedExt := []string{".svelte", ".svx", ".mdx"}

	filePrefix := strings.HasPrefix(f.Name(), "__")
	fileExt := filepath.Ext(f.Name())
	if !filePrefix && common.Contains(acceptedExt, fileExt) {
		return true
	}
	return false
}

func PrepareContent(name string, resources map[string]string, templateId string, data *config.TemplateData) builder.Content {
	contentBuilder := builder.GetContentBuilder(name)
	contentBuilder.SetEmbeddedResources(resources)
	contentBuilder.SetTemplateId(templateId)
	contentBuilder.SetTemplateData(data)

	director := builder.NewDirector(contentBuilder)
	return director.GetContent()
}

func MakeFileContent(efs *embed.FS, content builder.Content) []byte {
	tplConfig := NewTplConfig(content.PathToTplFile, content.Funcs, content.TemplateData)
	return ExecSveltinTpl(efs, tplConfig)
}

func WriteContentToDisk(fs afero.Fs, saveAs string, fileContent []byte) error {
	err := common.WriteToDisk(fs, saveAs, bytes.NewReader(fileContent))
	if err != nil {
		return err
	}
	return nil
}
