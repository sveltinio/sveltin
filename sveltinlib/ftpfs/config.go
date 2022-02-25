/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package ftpfs ...
package ftpfs

import (
	"strconv"
	"strings"
)

// FTPConnectionConfig is the struct with all is needed to
// establish an FTP connection to a remote server.
type FTPConnectionConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Timeout  int
	IsEPSV   bool
}

func (d *FTPConnectionConfig) makeConnectionString() string {
	return strings.Join([]string{d.Host, strconv.Itoa(d.Port)}, ":")
}
