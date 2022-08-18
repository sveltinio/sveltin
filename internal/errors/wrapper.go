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

	"github.com/charmbracelet/lipgloss"
)

// ErrorType represents a specific error.
type ErrorType int8

const (
	defaultError ErrorType = iota + 1
	notImplementYetError
	notValidProjectError
	notEmptyProjectError
	notValidURLError
	notValidGitHubURLError
	notValidGitHubRepoError
	fileNotFoundError
	dirInsteadOfFileError
	existingDirectoryError
	dirNotFoundError
	moveFileError
	optionNotValidError
	numOfArgsNotValidError
	numOfArgsNotValidErrorWithMessage
	notValidArgumentsError
	resourceNotFoundError
	contentTemplateTypeNotValidError
	pageTypeNotValidError
	metadataTypeNotValidError
	npmClientCommandError
	npmClientCommandNotValidError
	packageManagerCommandNotValidError
	execSystemCommandError
	execSystemCommandErrorWithMsg
)

var (
	// Colors
	defaultColor   = lipgloss.Color("#f87171")
	highlightColor = lipgloss.Color("#ef4444")

	// Styles
	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(highlightColor)
	headingStyle = lipgloss.NewStyle().BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(highlightColor).
			BorderBottom(true)
	messageStyle = lipgloss.NewStyle().Foreground(defaultColor)
	footerStyle  = lipgloss.NewStyle().MarginTop(1).Foreground(highlightColor)
	boxStyle     = lipgloss.NewStyle().Margin(1, 0, 1, 0).PaddingLeft(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(highlightColor).BorderLeft(true)
)

// SveltinError is the struct representing the way Sveltin handles errors.
type SveltinError struct {
	Code    ErrorType
	Name    string
	Title   string
	Message string
	Err     error
}

func (e *SveltinError) Error() string {
	title := titleStyle.Render(e.Title)
	header := headingStyle.Render(title)
	message := messageStyle.Render(e.Message)
	footer := footerStyle.Render(fmt.Sprintf("[ERROR INFO] Type(%s) · Code(%s)", e.Name, fmt.Sprint(e.Code)))

	box := boxStyle.Render(lipgloss.JoinVertical(lipgloss.Left, fmt.Sprintf("%s\n%s\n%s", header, message, footer)))
	return fmt.Sprint(box)
}

func newSveltinError(code ErrorType, name, title, message string, err error) error {
	return &SveltinError{
		Code:    code,
		Name:    name,
		Title:   title,
		Message: message,
		Err:     err,
	}
}

// NewDefaultError ...
func NewDefaultError(err error) error {
	return newSveltinError(defaultError, "DefaultError", "Default Error", err.Error(), err)
}

// NewNotImplementYetError ...
func NewNotImplementYetError() error {
	err := errors.New("not implemented yet. Pure CSS and Tailwindcss are the only available options so far")
	return newSveltinError(notImplementYetError, "NotImplementYetError", "Sorry, This Is Not implemented Yet", err.Error(), err)
}

// NewNotValidProjectError ...
func NewNotValidProjectError(pathToFile string) error {
	placeholderText := `

This is related to sveltin not being able to find the package.json file
within the current directory (%s).

Are you sure you are running the sveltin command from within a valid project directory?
`

	msg := fmt.Sprintf(placeholderText, filepath.Dir(pathToFile))
	err := fmt.Errorf(`no package.json file!%s `, msg)

	return newSveltinError(notValidProjectError, "NotValidProjectError", "Sveltin Project Not Found", msg, err)
}

// NewNotEmptyProjectError ...
func NewNotEmptyProjectError(pathToFile string) error {
	placeholderText := `

This is related to sveltin and an existing package.json file
within the current directory (%s).

Are you sure you are running the new theme command from within a not existing project directory?
`

	msg := fmt.Sprintf(placeholderText, filepath.Dir(pathToFile))
	err := fmt.Errorf(`an existing package.json file found!%s `, msg)
	return newSveltinError(notEmptyProjectError, "NotEmptyProjectError", "Existing Project Found", msg, err)
}

// NewNotValidURL ...
func NewNotValidURL(input string) error {
	msg := fmt.Sprintf("'%s' seems to be a not valid url", input)
	err := fmt.Errorf("'%s' seems to be a not valid url", input)
	return newSveltinError(notValidURLError, "NotValidURL", "Not A Valid URL", msg, err)
}

// NewNotValidGitHubURL ...
func NewNotValidGitHubURL(input string) error {
	msg := fmt.Sprintf("'%s' seems to be a not valid github url", input)
	err := fmt.Errorf("'%s' seems to be a not valid github url", input)
	return newSveltinError(notValidGitHubURLError, "NotValidGitHubURL", "Not A Valid GitHub URL", msg, err)
}

// NewNotValidGitHubRepoURL ...
func NewNotValidGitHubRepoURL(input string) error {
	msg := fmt.Sprintf("<user>/<repo> not in url path, received: '%s'", input)
	err := fmt.Errorf("<user>/<repo> not in url path, received: '%s'", input)
	return newSveltinError(notValidGitHubRepoError, "NotValidGitHubRepo", "Not A Valid GitHub Repository", msg, err)
}

// NewFileNotFoundError ...
func NewFileNotFoundError() error {
	err := errors.New("please, check the file path")
	return newSveltinError(fileNotFoundError, "FileNotFoundError", "File Not Found", err.Error(), err)
}

// NewDirInsteadOfFileError ...
func NewDirInsteadOfFileError() error {
	err := errors.New("please, check the file path. It seems to be a directory, not a file")
	return newSveltinError(dirInsteadOfFileError, "DirInsteadOfFileError", "It Should Be A Directory Not A File", err.Error(), err)
}

// NewExistingDirectoryError ...
func NewExistingDirectoryError() error {
	err := errors.New("destination path already exists and is not an empty directory")
	return newSveltinError(existingDirectoryError, "ExistingDirectoryError", "Directory Already Exists", err.Error(), err)
}

// NewDirNotFoundError ...
func NewDirNotFoundError() error {
	err := errors.New("please, check the directory path")
	return newSveltinError(dirNotFoundError, "DirNotFoundError", "Directory Not Found", err.Error(), err)
}

// NewMoveFileError ...
func NewMoveFileError(sourceFile, saveTo string) error {
	placeholderText := `

Something went wrong trying to save %s as %s
`
	msg := fmt.Sprintf(placeholderText, sourceFile, saveTo)
	err := fmt.Errorf("please, check the file path: %s", msg)
	return newSveltinError(moveFileError, "MoveFileError", "File To Be Moved Not Found", msg, err)
}

// NewOptionNotValidError ...
func NewOptionNotValidError(value string, options []string) error {
	err := fmt.Errorf("it seems a not valid option has been used! Your choice was '%s'. Valid ones are: %s", value, strings.Join(options, ", "))
	return newSveltinError(optionNotValidError, "OptionNotValidError", "Option Not Valid", err.Error(), err)
}

// NewNumOfArgsNotValidError ...
func NewNumOfArgsNotValidError() error {
	err := errors.New("it seems a wrong number of arguments have been used")
	return newSveltinError(numOfArgsNotValidError, "NumOfArgsNotValidError", "Wrong Number Of Arguments", err.Error(), err)
}

// NewNumOfArgsNotValidErrorWithMessage ...
func NewNumOfArgsNotValidErrorWithMessage(err error) error {
	return newSveltinError(numOfArgsNotValidErrorWithMessage, "NumOfArgsNotValidErrorWithMessage", "Wrong Number Of Arguments", err.Error(), err)
}

// NewNotValidArgumentsError ...
func NewNotValidArgumentsError() error {
	err := errors.New("some of the provided arguments seem to be not a valid one")
	return newSveltinError(notValidArgumentsError, "NotValidArgumentsError", "Not A Valid Argument", err.Error(), err)
}

// NewResourceNotFoundError ...
func NewResourceNotFoundError() error {
	err := errors.New("it seems a not exisiting resource has been used")
	return newSveltinError(resourceNotFoundError, "ResourceNotFoundError", "Resource Not Found", err.Error(), err)
}

// NewContentTemplateTypeNotValidError ...
func NewContentTemplateTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as content template")
	return newSveltinError(contentTemplateTypeNotValidError, "ContentTemplateTypeNotValidError", "Not A Valid Content Template", err.Error(), err)
}

// NewPageTypeNotValidError ...
func NewPageTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as page")
	return newSveltinError(pageTypeNotValidError, "PageTypeNotValidError", "Not A Valid Page Type", err.Error(), err)
}

// NewMetadataTypeNotValidError ...
func NewMetadataTypeNotValidError() error {
	err := errors.New("it seems a not valid type has been used as metadata")
	return newSveltinError(metadataTypeNotValidError, "MetadataTypeNotValidError", "Not A Valid Metadata Type", err.Error(), err)
}

// NewNPMClientCommandError ...
func NewNPMClientCommandError(err error) error {
	return newSveltinError(npmClientCommandError, "NPMClientCommandError", "NPM Client Error", err.Error(), err)
}

// NewNPMClientCommandNotValidError ...
func NewNPMClientCommandNotValidError() error {
	err := errors.New("it seems the operation is not a valid one for the npm client")
	return newSveltinError(npmClientCommandNotValidError, "NPMClientCommandNotValidError", "Not A Valid NPM Client Operation", err.Error(), err)
}

// NewPackageManagerKeyNotFoundOnPackageJSONFile ...
func NewPackageManagerKeyNotFoundOnPackageJSONFile() error {
	err := errors.New(`

did not find the "packageManager" key in your package.json file

[HINT]: add "packageManager": "<your_npm_client>@<version>" to it and run the command again`)
	return newSveltinError(packageManagerCommandNotValidError, "PackageManagerCommandNotValidError", "Not packageManager key Found", err.Error(), err)
}

// NewProjectNameNotFoundError ...
func NewProjectNameNotFoundError() error {
	err := errors.New(`cannot find property "name" in your package.json file`)
	return newSveltinError(83, "ProjectNameNotFoundError", "Project Name Not Found", err.Error(), err)
}

// NewExecSystemCommandError ...
func NewExecSystemCommandError(cmdName, opts string) error {
	placeholderText := `

Here is the string representing the command line executed:

%s %s
`
	msg := fmt.Sprintf(placeholderText, cmdName, opts)
	err := fmt.Errorf("cannot exec the system command. please, check it and its arguments: %s", msg)
	return newSveltinError(execSystemCommandError, "ExecSystemCommandError", "System Command Execution Failure", msg, err)
}

// NewExecSystemCommandErrorWithMsg ...
func NewExecSystemCommandErrorWithMsg(err error) error {
	errN := errors.New("cannot exec the system command. please, check it and its arguments: " + err.Error())
	return newSveltinError(execSystemCommandErrorWithMsg, "ExecSystemCommandErrorWithMsg", "System Command Execution Failure", errN.Error(), errN)
}
