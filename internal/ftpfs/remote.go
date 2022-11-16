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

// RemoteServer is the interface defining the list of actions
// can be performed on a RemoteServer implementation.
type RemoteServer interface {
	Dial() error
	Login() error
	Logout() error
	Idle() error
	MakeDirs([]string, bool) error
	UploadFiles(afero.Fs, string, []string, bool) error
	DeleteAll([]string, bool) error
	DoBackup(afero.Fs, string, bool) error
}
