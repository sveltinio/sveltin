/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
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
	"strings"
)

// SveltinError is the struct representing the way Sveltin handles errors.
type SveltinError struct {
	Code int
	Name string
	Err  error
}

func (e *SveltinError) Error() string {
	return fmt.Sprintf("[SveltinError: %s (Code=%d)] %v", e.Name, e.Code, e.Err)
}

func newSveltinError(code int, err error, message string) error {
	return &SveltinError{
		Code: code,
		Name: message,
		Err:  err,
	}
}

// NewNotValidProjectError ...
func NewNotValidProjectError(pathToFile string) error {
	placeholderText := `

This is related to sveltin not being able to find the package.json file
within the current directory (%s).

Are you sure you are running the sveltin command from within a valid project directory?
`

	errN := fmt.Errorf(`no package.json file!%s `, fmt.Sprintf(placeholderText, filepath.Dir(pathToFile)))
	return newSveltinError(1, errN, "NotValidProjectError")
}

// NewNotEmptyProjectError ...
func NewNotEmptyProjectError(pathToFile string) error {
	placeholderText := `

This is related to sveltin and an existing package.json file
within the current directory (%s).

Are you sure you are running the new theme command from within a not existing project directory?
`

	errN := fmt.Errorf(`no package.json file!%s `, fmt.Sprintf(placeholderText, filepath.Dir(pathToFile)))
	return newSveltinError(2, errN, "NotValidProjectError")
}

// NewNotValidURL ...
func NewNotValidURL(input string) error {
	errN := fmt.Errorf("'%s' seems to be a not valid url", input)
	return newSveltinError(3, errN, "NotValidURL")
}

// NewNotValidGitHubURL ...
func NewNotValidGitHubURL(input string) error {
	errN := fmt.Errorf("'%s' seems to be a not valid github url", input)
	return newSveltinError(4, errN, "NotValidGitHubURL")
}

// NewNotValidGitHubRepoURL ...
func NewNotValidGitHubRepoURL(input string) error {
	errN := fmt.Errorf("<user>/<repo> not in url path, received: '%s'", input)
	return newSveltinError(5, errN, "NotValidGitHubRepo")
}

// NewDefaultError ...
func NewDefaultError(err error) error {
	return newSveltinError(10, err, "Error")
}

// NewFileNotFoundError ...
func NewFileNotFoundError() error {
	err := errors.New("please, check the file path")
	return newSveltinError(11, err, "FileNotFoundError")
}

// NewDirInsteadOfFileError ...
func NewDirInsteadOfFileError() error {
	err := errors.New("please, check the file path. It seems to be a directory, not a file")
	return newSveltinError(12, err, "DirInsteadOfFileError")
}

// NewDirNotFoundError ...
func NewDirNotFoundError() error {
	err := errors.New("please, check the directory path")
	return newSveltinError(13, err, "DirNotFoundError")
}

// NewMoveFileError ...
func NewMoveFileError(sourceFile, saveTo string) error {
	placeholderText := `

Something went wrong trying to save %s as %s
`
	err := fmt.Errorf("please, check the file path: %s", fmt.Sprintf(placeholderText, sourceFile, saveTo))
	return newSveltinError(14, err, "FileNotFoundError")
}

// NewNotImplementYetError ...
func NewNotImplementYetError() error {
	err := errors.New("not implemented yet. Pure CSS and Tailwindcss are the only available options so far")
	return newSveltinError(20, err, "NotImplementYetError")
}

// NewOptionNotValidError ...
func NewOptionNotValidError(value string, options []string) error {
	err := fmt.Errorf("it seems a not valid option has been used! Your choice was '%s'. Valid ones are: %s", value, strings.Join(options, ", "))
	return newSveltinError(30, err, "OptionNotValidError")
}

// NewNumOfArgsNotValidError ...
func NewNumOfArgsNotValidError() error {
	err := errors.New("it seems a wrong number of arguments have been used")
	return newSveltinError(31, err, "NumOfArgsNotValidError")
}

// NewNumOfArgsNotValidErrorWithMessage ...
func NewNumOfArgsNotValidErrorWithMessage(err error) error {
	return newSveltinError(32, err, "NumOfArgsNotValidErrorWithMessage")
}

// NewNotValidArgumentsError ...
func NewNotValidArgumentsError() error {
	err := errors.New("some of the provided arguments seem to be not a valid one")
	return newSveltinError(33, err, "NumOfArgsNotValidErrorWithMessage")
}

// NewResourceNotFoundError ...
func NewResourceNotFoundError() error {
	err := errors.New("it seems a not exisiting resource has been used")
	return newSveltinError(40, err, "ResourceNotFoundError")
}

// NewContentTemplateTypeNotValidError ...
func NewContentTemplateTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as content template")
	return newSveltinError(50, err, "ContentTemplateTypeNotValidError")
}

// NewPageTypeNotValidError ...
func NewPageTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as page")
	return newSveltinError(60, err, "PageTypeNotValidError")
}

// NewMetadataTypeNotValidError ...
func NewMetadataTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as metadata")
	return newSveltinError(70, err, "MetadataTypeNotValidError")
}

// NewPackageManagerCommandError ...
func NewPackageManagerCommandError(err error) error {
	return newSveltinError(80, err, "PackageManagerCommandError")
}

// NewPackageManagerCommandNotValidError ...
func NewPackageManagerCommandNotValidError() error {
	err := errors.New("it seems the operation is not a valid one for the package manager")
	return newSveltinError(81, err, "PackageManagerCommandNotValidError")
}

// NewPackageManagerKeyNotFoundOnPackageJSONFile ...
func NewPackageManagerKeyNotFoundOnPackageJSONFile() error {
	errN := errors.New(`

did not find the "packageManager" key in your package.json file

[HINT]: add "packageManager": "<your_npm_client>@<version>" to it and run the command again`)
	return newSveltinError(82, errN, "PackageManagerCommandNotValidError")
}

// NewProjectNameNotFoundError ...
func NewProjectNameNotFoundError() error {
	errN := errors.New(`cannot find property "name" in your package.json file`)
	return newSveltinError(83, errN, "ProjectNameNotFoundError")
}

// NewExecSystemCommandError ...
func NewExecSystemCommandError(cmdName, opts string) error {
	placeholderText := `

Here is the string representing the command line to be executed:

%s %s
`
	errN := fmt.Errorf("cannot exec the system command. please, check it and its arguments: %s", fmt.Sprintf(placeholderText, cmdName, opts))
	return newSveltinError(90, errN, "ExecSystemCommandError")
}

// NewExecSystemCommandErrorWithMsg ...
func NewExecSystemCommandErrorWithMsg(err error) error {
	errN := errors.New("cannot exec the system command. please, check it and its arguments: " + err.Error())
	return newSveltinError(91, errN, "ExecSystemCommandError")
}
