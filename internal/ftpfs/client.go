/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package ftpfs

import (
	"github.com/spf13/afero"
)

// Client is the sturct representing the invoker, who is responsible for initiating requests.
type Client struct {
	Command Command
}

// Run execute the method on a receiving object.
func (c *Client) Run() error {
	return c.Command.execute()
}

// DialAction creates and configures the concrete dial command.
func DialAction(conn *FTPServerConnection) *Client {
	return &Client{
		Command: &DialCommand{
			Server: conn,
		},
	}
}

// LoginAction creates and configures the concrete login command.
func LoginAction(conn *FTPServerConnection) *Client {
	return &Client{
		Command: &LoginCommand{
			Server: conn,
		},
	}
}

// LogoutAction creates and configures the concrete logout command.
func LogoutAction(conn *FTPServerConnection) *Client {
	return &Client{
		Command: &LogoutCommand{
			Server: conn,
		},
	}
}

// IdleAction creates and configures the concrete no-operation(idle) command.
func IdleAction(conn *FTPServerConnection) *Client {
	return &Client{
		Command: &IdleCommand{
			Server: conn,
		},
	}
}

// MakeDirsAction creates and configures the concrete dial command.
func MakeDirsAction(conn *FTPServerConnection, dirs []string, dryRun bool) *Client {
	return &Client{
		Command: &MakeDirsCommand{
			Server: conn,
			Dirs:   dirs,
			DryRun: dryRun,
		},
	}
}

// UploadAction creates and configures the concrete upload command.
func UploadAction(conn *FTPServerConnection, appFs afero.Fs, localDirname string, files []string, replaceBasePath, dryRun bool) *Client {
	return &Client{
		Command: &UploadCommand{
			Server:          conn,
			AppFs:           appFs,
			LocalDirname:    localDirname,
			Files:           files,
			ReplaceBasePath: replaceBasePath,
			DryRun:          dryRun,
		},
	}
}

// DeleteAllAction creates and configures the concrete delete all command.
func DeleteAllAction(conn *FTPServerConnection, exludeList []string, dryRun bool) *Client {
	return &Client{
		Command: &DeleteAllCommand{
			Server:      conn,
			ExcludeList: exludeList,
			DryRun:      dryRun,
		},
	}
}

// BackupAction creates and configures the concrete backup command.
func BackupAction(conn *FTPServerConnection, appFs afero.Fs, name string, dryRun bool) *Client {
	return &Client{
		Command: &BackupCommand{
			Server: conn,
			AppFs:  appFs,
			Name:   name,
			DryRun: dryRun,
		},
	}
}
