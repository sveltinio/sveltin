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
	is.Equal("SVELTIN Error", re.Message)

	errVar = NewFileNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(11, re.Code)
	is.Equal("SVELTIN FileNotFoundError", re.Message)

	errVar = NewDirInsteadOfFileError()
	re = errVar.(*SveltinError)
	is.Equal(12, re.Code)
	is.Equal("SVELTIN DirInsteadOfFileError", re.Message)
	is.Equal("SVELTIN DirInsteadOfFileError: please, check the file path. It seems to be a directory, not a file", re.Error())

	errVar = NewDirNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(13, re.Code)
	is.Equal("SVELTIN DirNotFoundError", re.Message)

	errVar = NewNotImplementYetError()
	re = errVar.(*SveltinError)
	is.Equal(20, re.Code)
	is.Equal("SVELTIN NotImplementYetError", re.Message)

	errVar = NewOptionNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(30, re.Code)
	is.Equal("SVELTIN OptionNotValidError", re.Message)

	errVar = NewNumOfArgsNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(31, re.Code)
	is.Equal("SVELTIN NumOfArgsNotValidError", re.Message)

	err = errors.New("")
	errVar = NewNumOfArgsNotValidErrorWithMessage(err)
	re = errVar.(*SveltinError)
	is.Equal(32, re.Code)
	is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)

	errVar = NewResourceNotFoundError()
	re = errVar.(*SveltinError)
	is.Equal(40, re.Code)
	is.Equal("SVELTIN ResourceNotFoundError", re.Message)

	errVar = NewContentTemplateTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(50, re.Code)
	is.Equal("SVELTIN ContentTemplateTypeNotValidError", re.Message)

	errVar = NewPageTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(60, re.Code)
	is.Equal("SVELTIN PageTypeNotValidError", re.Message)

	errVar = NewMetadataTypeNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(70, re.Code)
	is.Equal("SVELTIN MetadataTypeNotValidError", re.Message)

	err = errors.New("")
	errVar = NewPackageManagerCommandError(err)
	re = errVar.(*SveltinError)
	is.Equal(80, re.Code)
	is.Equal("SVELTIN PackageManagerCommandError", re.Message)

	errVar = NewPackageManagerCommandNotValidError()
	re = errVar.(*SveltinError)
	is.Equal(81, re.Code)
	is.Equal("SVELTIN PackageManagerCommandNotValidError", re.Message)

	errVar = NewExecSystemCommandError()
	re = errVar.(*SveltinError)
	is.Equal(82, re.Code)
	is.Equal("SVELTIN ExecSystemCommandError", re.Message)
}
