/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package ftpfs

import "github.com/spf13/afero"

// Command interface declares just the single method for executing the command.
type Command interface {
	execute() error
}

// DialCommand implements the dial request.
type DialCommand struct {
	Server RemoteServer
}

func (c *DialCommand) execute() error {
	return c.Server.Dial()
}

// LoginCommand implements the login request.
type LoginCommand struct {
	Server RemoteServer
}

func (c *LoginCommand) execute() error {
	return c.Server.Login()
}

// LogoutCommand implements the logout request.
type LogoutCommand struct {
	Server RemoteServer
}

func (c *LogoutCommand) execute() error {
	return c.Server.Logout()
}

// IdleCommand implements the no-operation (idle) request.
type IdleCommand struct {
	Server RemoteServer
}

func (c *IdleCommand) execute() error {
	return c.Server.Idle()
}

// MakeDirsCommand implements the make dirs request.
type MakeDirsCommand struct {
	Server  RemoteServer
	Dirs    []string
	DryRun  bool
	Verbose bool
}

func (c *MakeDirsCommand) execute() error {
	return c.Server.MakeDirs(c.Dirs, c.DryRun)
}

// UploadCommand implements the upload request.
type UploadCommand struct {
	Server       RemoteServer
	AppFs        afero.Fs
	LocalDirname string
	Files        []string
	DryRun       bool
	Verbose      bool
}

func (c *UploadCommand) execute() error {
	return c.Server.Upload(c.AppFs, c.LocalDirname, c.Files, c.DryRun)
}

// DeleteAllCommand implements the delete all request.
type DeleteAllCommand struct {
	Server      RemoteServer
	ExcludeList []string
	DryRun      bool
}

func (c *DeleteAllCommand) execute() error {
	return c.Server.DeleteAll(c.ExcludeList, c.DryRun)
}

// BackupCommand implements the backup request.
type BackupCommand struct {
	Server RemoteServer
	AppFs  afero.Fs
	Name   string
	DryRun bool
}

func (c *BackupCommand) execute() error {
	return c.Server.DoBackup(c.AppFs, c.Name, c.DryRun)
}
