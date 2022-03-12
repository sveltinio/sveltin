/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package sveltinerr ...
package sveltinerr

import (
	"errors"
	"fmt"
	"path/filepath"
)

// SveltinError is the struct representing the way Sveltin handles errors.
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

// NewNotValidProjectError ...
func NewNotValidProjectError(pathToFile string) error {
	placeholderText := `

This is related to sveltin not being able to find teh package.json file
within the current directory (%s).

Are you sure you are running the sveltin command from within a valid project directory?
`

	errN := fmt.Errorf(`no package.json file!%s `, fmt.Sprintf(placeholderText, filepath.Dir(pathToFile)))
	return newSveltinError(1, errN, "SVELTIN NotValidProjectError")
}

// NewDefaultError ...
func NewDefaultError(err error) error {
	return newSveltinError(10, err, "SVELTIN Error")
}

// NewFileNotFoundError ...
func NewFileNotFoundError() error {
	err := errors.New("please, check the file path")
	return newSveltinError(11, err, "SVELTIN FileNotFoundError")
}

// NewDirInsteadOfFileError ...
func NewDirInsteadOfFileError() error {
	err := errors.New("please, check the file path. It seems to be a directory, not a file")
	return newSveltinError(12, err, "SVELTIN DirInsteadOfFileError")
}

// NewDirNotFoundError ...
func NewDirNotFoundError() error {
	err := errors.New("please, check the directory path")
	return newSveltinError(13, err, "SVELTIN DirNotFoundError")
}

// NewNotImplementYetError ...
func NewNotImplementYetError() error {
	err := errors.New("not implemented yet. Pure CSS and Tailwindcss are the only available options so far")
	return newSveltinError(20, err, "SVELTIN NotImplementYetError")
}

// NewOptionNotValidError ...
func NewOptionNotValidError() error {
	err := errors.New("it seems you used an invalid option")
	return newSveltinError(30, err, "SVELTIN OptionNotValidError")
}

// NewNumOfArgsNotValidError ...
func NewNumOfArgsNotValidError() error {
	err := errors.New("it seems you used a wrong number of arguments")
	return newSveltinError(31, err, "SVELTIN NumOfArgsNotValidError")
}

// NewNumOfArgsNotValidErrorWithMessage ...
func NewNumOfArgsNotValidErrorWithMessage(err error) error {
	return newSveltinError(32, err, "SVELTIN NumOfArgsNotValidErrorWithMessage")
}

// NewResourceNotFoundError ...
func NewResourceNotFoundError() error {
	err := errors.New("it seems you used an invalid resource")
	return newSveltinError(40, err, "SVELTIN ResourceNotFoundError")
}

// NewContentTemplateTypeNotValidError ...
func NewContentTemplateTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for content template")
	return newSveltinError(50, err, "SVELTIN ContentTemplateTypeNotValidError")
}

// NewPageTypeNotValidError ...
func NewPageTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for page")
	return newSveltinError(60, err, "SVELTIN PageTypeNotValidError")
}

// NewMetadataTypeNotValidError ...
func NewMetadataTypeNotValidError() error {
	err := errors.New("it seems you used an invalid type for metadata")
	return newSveltinError(70, err, "SVELTIN MetadataTypeNotValidError")
}

// NewPackageManagerCommandError ...
func NewPackageManagerCommandError(err error) error {
	return newSveltinError(80, err, "SVELTIN PackageManagerCommandError")
}

// NewPackageManagerCommandNotValidError ...
func NewPackageManagerCommandNotValidError() error {
	err := errors.New("it seems the operation on the package manager is not a valid one")
	return newSveltinError(81, err, "SVELTIN PackageManagerCommandNotValidError")
}

// NewExecSystemCommandError ...
func NewExecSystemCommandError() error {
	err := errors.New("cannot exec the system command. please, check it and its arguments")
	return newSveltinError(82, err, "SVELTIN ExecSystemCommandError")
}

// NewExecSystemCommandErrorWithMsg ...
func NewExecSystemCommandErrorWithMsg(err error) error {
	errN := errors.New("cannot exec the system command. please, check it and its arguments: " + err.Error())
	return newSveltinError(82, errN, "SVELTIN ExecSystemCommandError")
}

// NewPackageManagerKeyNotFoundOnPackageJSONFile ...
func NewPackageManagerKeyNotFoundOnPackageJSONFile() error {
	errN := errors.New(`

did not find the "packageManager" key in your package.json file

[HINT]: add "packageManager": "<your_npm_client>@<version>" to it and run the command again`)
	return newSveltinError(83, errN, "SVELTIN PackageManagerCommandNotValidError")
}

// NewProjectNameNotFoundError ...
func NewProjectNameNotFoundError() error {
	errN := errors.New(`cannot find property "name" in your package.json file`)
	return newSveltinError(84, errN, "SVELTIN ProjectNameNotFoundError")
}
