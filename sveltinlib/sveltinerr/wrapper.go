/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package sveltinerr

import (
	"errors"
	"fmt"
)

type SveltinError struct {
	Code    int
	Err     error
	Message string
}

func (e *SveltinError) Error() string {
	return fmt.Sprintf("[%s] %v", e.Message, e.Err)
}

func newSveltinError(code int, err error, message string) error {
	return &SveltinError{
		Code:    code,
		Err:     err,
		Message: message,
	}
}

func NewDefaultError(err error) error {
	return newSveltinError(10, err, "SVELTIN Error")
}

func NewFileNotFoundError() error {
	err := errors.New("please, check the file path")
	return newSveltinError(11, err, "SVELTIN FileNotFoundError")
}

func NewDirInsteadOfFileError() error {
	err := errors.New("please, check the file path. It seems to be a directory, not a file")
	return newSveltinError(12, err, "SVELTIN DirInsteadOfFileError")
}

func NewDirNotFoundError() error {
	err := errors.New("please, check the directory path")
	return newSveltinError(13, err, "SVELTIN DirNotFoundError")
}

func NewNotImplementYetError() error {
	err := errors.New("not implemented yet. Pure CSS and Tailwindcss are the only available options so far")
	return newSveltinError(20, err, "SVELTIN NotImplementYetError")
}

func NewOptionNotValidError() error {
	err := errors.New("it seems you used an invalid option")
	return newSveltinError(30, err, "SVELTIN OptionNotValidError")
}

func NewNumOfArgsNotValidError() error {
	err := errors.New("it seems you used a wrong number of arguments")
	return newSveltinError(31, err, "SVELTIN NumOfArgsNotValidError")
}

func NewNumOfArgsNotValidErrorWithMessage(err error) error {
	return newSveltinError(32, err, "SVELTIN NumOfArgsNotValidErrorWithMessage")
}

func NewResourceNotFoundError() error {
	err := errors.New("it seems you used an invalid resource")
	return newSveltinError(40, err, "SVELTIN ResourceNotFoundError")
}

func NewContentTemplateTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for content template")
	return newSveltinError(50, err, "SVELTIN ContentTemplateTypeNotValidError")
}

func NewPageTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for page")
	return newSveltinError(60, err, "SVELTIN PageTypeNotValidError")
}

func NewMetadataTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for metadata")
	return newSveltinError(70, err, "SVELTIN MetadataTypeNotValidError")
}

func NewPackageManagerCommandError(err error) error {
	return newSveltinError(80, err, "SVELTIN PackageManagerCommandError")
}

func NewPackageManagerCommandNotValidError() error {
	err := errors.New("it seems the operation on the package manager is not a valid one")
	return newSveltinError(81, err, "SVELTIN PackageManagerCommandNotValidError")
}

func NewExecSystemCommandError() error {
	err := errors.New("cannot exec the system command. please, check it and its arguments")
	return newSveltinError(82, err, "SVELTIN ExecSystemCommandError")
}

func NewExecSystemCommandErrorWithMsg(err error) error {
	errN := errors.New("cannot exec the system command. please, check it and its arguments: " + err.Error())
	return newSveltinError(82, errN, "SVELTIN ExecSystemCommandError")
}

func NewPackageManagerKeyNotFoundOnPackageJSONFile() error {
	errN := errors.New(`

did not find the "packageManager" key in your package.json file

[HINT]: add "packageManager": "<your_npm_client>@<version>" to it and run the command again
`)
	return newSveltinError(83, errN, "SVELTIN PackageManagerCommandNotValidError")
}
