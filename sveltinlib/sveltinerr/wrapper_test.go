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
	is.Equal(10, re.Code)
	is.Equal("SVELTIN Error", re.Name)

	errVar = NewFileNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(11, re.Code)
	is.Equal("SVELTIN FileNotFoundError", re.Name)

	errVar = NewDirInsteadOfFileError()
	re = errVar.(*SveltinError)
	is.Equal(12, re.Code)
	is.Equal("SVELTIN DirInsteadOfFileError", re.Name)
	is.Equal("[Error Code 12: SVELTIN DirInsteadOfFileError] please, check the file path. It seems to be a directory, not a file", re.Error())

	errVar = NewDirNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(13, re.Code)
	is.Equal("SVELTIN DirNotFoundError", re.Name)

	errVar = NewNotImplementYetError()
	re = errVar.(*SveltinError)
	is.Equal(20, re.Code)
	is.Equal("SVELTIN NotImplementYetError", re.Name)

	errVar = NewOptionNotValidError("test", []string{"default", "none"})
	re = errVar.(*SveltinError)
	is.Equal(30, re.Code)
	is.Equal("SVELTIN OptionNotValidError", re.Name)

	errVar = NewNumOfArgsNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(31, re.Code)
	is.Equal("SVELTIN NumOfArgsNotValidError", re.Name)

	err = errors.New("")
	errVar = NewNumOfArgsNotValidErrorWithMessage(err)
	re = errVar.(*SveltinError)
	is.Equal(32, re.Code)
	is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Name)

	errVar = NewResourceNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(40, re.Code)
	is.Equal("SVELTIN ResourceNotFoundError", re.Name)

	errVar = NewContentTemplateTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(50, re.Code)
	is.Equal("SVELTIN ContentTemplateTypeNotValidError", re.Name)

	errVar = NewPageTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(60, re.Code)
	is.Equal("SVELTIN PageTypeNotValidError", re.Name)

	errVar = NewMetadataTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(70, re.Code)
	is.Equal("SVELTIN MetadataTypeNotValidError", re.Name)

	err = errors.New("")
	errVar = NewPackageManagerCommandError(err)
	re = errVar.(*SveltinError)
	is.Equal(80, re.Code)
	is.Equal("SVELTIN PackageManagerCommandError", re.Name)

	errVar = NewPackageManagerCommandNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(81, re.Code)
	is.Equal("SVELTIN PackageManagerCommandNotValidError", re.Name)

	errVar = NewExecSystemCommandError()
	re = errVar.(*SveltinError)
	is.Equal(82, re.Code)
	is.Equal("SVELTIN ExecSystemCommandError", re.Name)
}
