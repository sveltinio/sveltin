package composer

/*
import (
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/sveltinio/sveltin/config"
	"github.com/sveltinio/sveltin/resources"
	"github.com/sveltinio/sveltin/sveltinlib/maker"
)

func TestFolder(t *testing.T) {
	is := is.New(t)
	memFS := afero.NewMemMapFs()
	tests := []struct {
		artifact *maker.Artifact
		data     *config.SveltinArtifact
		file     *File
		folder   *Folder
	}{

		{
			artifact: maker.NewThemeArtifact(&resources.SveltinFS, memFS),
			data: &config.SveltinArtifact{
				Name: "white",
			},
			file: &File{
				Name: "website.js",
			},
			folder: &Folder{
				Name: "config",
				Path: "project",
			},
		},
	}

	for _, tc := range tests {
		tc.folder.Add(tc.file)
		tc.folder.Create(tc.artifact, tc.data)

		folderDxists, err := afero.DirExists(memFS, tc.folder.Path)
		is.NoErr(err)
		is.True(folderDxists)

		fileExists, err := afero.Exists(memFS, tc.folder.Path)
		is.NoErr(err)
		is.True(fileExists)
	}

}
*/
