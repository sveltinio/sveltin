package sveltinerr

import (
	"errors"
	"testing"

	"github.com/matryer/is"
)

func TestErrors(t *testing.T) {

	is := is.New(t)

	err := errors.New("")
	errVar := NewDefaultError(err)
	re := errVar.(*SveltinError)
	is.Equal(1, int(re.Code))
	is.Equal("DefaultError", re.Name)

	errVar = NewFileNotFoundError("test.example")
	re = errVar.(*SveltinError)
	is.Equal("FileNotFoundError", re.Name)

	errVar = NewDirInsteadOfFileError()
	re = errVar.(*SveltinError)
	is.Equal("DirInsteadOfFileError", re.Name)
	is.Equal("please, check the file path. It seems to be a directory, not a file", re.Message)

	errVar = NewDirNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal("DirNotFoundError", re.Name)

	errVar = NewNotImplementYetError()
	re = errVar.(*SveltinError)
	is.Equal("NotImplementYetError", re.Name)

	errVar = NewOptionNotValidError("test", []string{"default", "none"})
	re = errVar.(*SveltinError)
	is.Equal("OptionNotValidError", re.Name)

	errVar = NewNumOfArgsNotValidError()
	re = errVar.(*SveltinError)
	is.Equal("NumOfArgsNotValidError", re.Name)

	err = errors.New("")
	errVar = NewNumOfArgsNotValidErrorWithMessage(err)
	re = errVar.(*SveltinError)
	is.Equal("NumOfArgsNotValidErrorWithMessage", re.Name)

	errVar = NewResourceNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal("ResourceNotFoundError", re.Name)

	errVar = NewContentTemplateTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal("ContentTemplateTypeNotValidError", re.Name)

	errVar = NewPageTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal("PageTypeNotValidError", re.Name)

	errVar = NewMetadataTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal("MetadataTypeNotValidError", re.Name)

	err = errors.New("")
	errVar = NewNPMClientCommandError(err)
	re = errVar.(*SveltinError)
	is.Equal("NPMClientCommandError", re.Name)

	errVar = NewNPMClientCommandNotValidError()
	re = errVar.(*SveltinError)
	is.Equal("NPMClientCommandNotValidError", re.Name)

	errVar = NewExecSystemCommandError("git", "init")
	re = errVar.(*SveltinError)
	is.Equal("ExecSystemCommandError", re.Name)
}
