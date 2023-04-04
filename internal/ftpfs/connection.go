/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package ftpfs handle connections and operations to deal with an FTP server.
package ftpfs

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"path/filepath"
	"sort"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/jlaffaye/ftp"
	"github.com/samber/lo"
	"github.com/spf13/afero"
	"github.com/sveltinio/prompti/progressbar"
	"github.com/sveltinio/sveltin/utils"
)

// FTPServerConnection is the struct with all is needed to establish and act on the FTP remote server.
type FTPServerConnection struct {
	Config       FTPConnectionConfig
	serverFolder string
	client       *ftp.ServerConn
	logger       *log.Logger
}

// NewFTPServerConnection returns a new FTPServerConnection struct.
func NewFTPServerConnection(config *FTPConnectionConfig) *FTPServerConnection {
	return &FTPServerConnection{
		Config: FTPConnectionConfig{
			Host:     config.Host,
			Port:     config.Port,
			User:     config.User,
			Password: config.Password,
			Timeout:  config.Timeout,
			IsEPSV:   config.IsEPSV,
		},
	}
}

// SetRootFolder sets the root folder on the FTP remote server.
func (s *FTPServerConnection) SetRootFolder(name string) {
	s.serverFolder = name
}

// SetLogger sets the root folder on the FTP remote server.
func (s *FTPServerConnection) SetLogger(logger *log.Logger) {
	s.logger = logger
}

// Dial contains the logic for the FTP receiver to handle the dial command.
func (s *FTPServerConnection) Dial() error {
	connStr := s.Config.makeConnectionString()
	s.logger.Infof("Connecting to the FTP Server (%s) ", connStr)
	c, err := ftp.Dial(connStr, ftp.DialWithTimeout(time.Duration(s.Config.Timeout)*time.Second), ftp.DialWithDisabledEPSV(s.Config.IsEPSV))
	if err != nil {
		return err
	}
	s.client = c
	return nil
}

// Login contains the logic for the FTP receiver to handle the login command.
func (s *FTPServerConnection) Login() error {
	s.logger.Infof("Login (as %s)\n\n", s.Config.User)
	if err := s.client.Login(s.Config.User, s.Config.Password); err != nil {
		return err
	}
	return nil
}

// Logout contains the logic for the FTP receiver to handle the logout command.
func (s *FTPServerConnection) Logout() error {
	s.logger.Info("Closing the connection to the FTP server")
	if err := s.client.Quit(); err != nil {
		return err
	}
	return nil
}

// Idle contains the logic for the FTP receiver to handle the no-operation (idle) command.
func (s *FTPServerConnection) Idle() error {
	return s.client.NoOp()
}

// MakeDirs contains the logic for the FTP receiver to handle the make dirs command.
func (s *FTPServerConnection) MakeDirs(folders []string, dryRun bool) error {
	sort.Strings(folders)
	if err := s.client.ChangeDir(s.serverFolder); err != nil {
		return err
	}

	pbConfig := &progressbar.Config{
		Items:          folders,
		OnCompletesMsg: fmt.Sprintf("Done! %d folders created", len(folders)),
		OnProgressCmd: func(path string) tea.Cmd {
			return mkDirTeaCmd(s, path, dryRun)
		},
	}

	if _, err := progressbar.Run(pbConfig); err != nil {
		return err
	}
	return nil
}

// UploadFiles contains the logic for the FTP receiver to handle the upload files command.
func (s *FTPServerConnection) UploadFiles(appFs afero.Fs, localDir string, files []string, replaceBasePath, dryRun bool) error {
	sort.Strings(files)

	pbConfig := &progressbar.Config{
		Items:          files,
		OnCompletesMsg: fmt.Sprintf("Done! %d files uploaded", len(files)),
		OnProgressCmd: func(path string) tea.Cmd {
			return uploadFileTeaCmd(s, appFs, path, localDir, replaceBasePath, dryRun)
		},
	}

	if _, err := progressbar.Run(pbConfig); err != nil {
		return err
	}
	return nil
}

// DeleteAll contains the logic for the FTP receiver to handle the delete all command.
func (s *FTPServerConnection) DeleteAll(exclude []string, dryrun bool) error {
	entries, err := s.client.List(s.serverFolder)
	if err != nil {
		return err
	}

	if len(entries) > 0 {
		s.logger.Info("Deleting previous content from the FTP remote folder")
		for _, entry := range entries {
			switch entry.Type {
			case ftp.EntryTypeFolder:
				if !dryrun {
					folder := filepath.Join(s.serverFolder, entry.Name)
					if err := s.client.RemoveDirRecur(folder); err != nil {
						return err
					}
				}
			case ftp.EntryTypeFile:
				if !dryrun {
					file := filepath.Join(s.serverFolder, entry.Name)
					if !lo.Contains(exclude, filepath.Base(file)) {
						if err := s.client.Delete(file); err != nil {
							return nil
						}
					}
				}
			}
		}
	}

	return nil
}

// DoBackup contains the logic for the FTP receiver to handle the backup command.
func (s *FTPServerConnection) DoBackup(appFs afero.Fs, tarballFilePath string, dryRun bool) error {
	archiveFilename := tarballFilePath + "_" + time.Now().Format("20060102_3:4:5PM") + ".tar.gz"
	s.logger.Infof("Reading the remote folder: %s", s.serverFolder)
	remoteFiles := s.walkRemote()

	if !dryRun {
		if len(remoteFiles) > 0 {
			if err := s.createTarball(appFs, archiveFilename, remoteFiles, dryRun); err != nil {
				return err
			}
		} else {
			s.logger.Info("Nothing to backup on the server!")
		}
	}
	return nil
}

//=============================================================================

func (s *FTPServerConnection) walkRemote() []string {
	w := s.client.Walk(s.serverFolder)
	var remoteFiles []string
	for w.Next() {
		if w.Stat().Type == ftp.EntryTypeFile {
			remoteFiles = append(remoteFiles, utils.ToBasePath(w.Path(), s.serverFolder))
		}
	}
	return remoteFiles
}

func (s *FTPServerConnection) uploadSingle(filename string, data *bytes.Buffer, dryRun bool) error {
	saveTo := filepath.Join(s.serverFolder, filepath.Dir(filename))
	saveAs := filepath.Base(filename)

	if !dryRun {
		cwd, _ := s.client.CurrentDir()
		if cwd != saveTo {
			if err := s.client.ChangeDir(saveTo); err != nil {
				return err
			}
		}
		if err := s.client.Stor(saveAs, data); err != nil {
			return err
		}
	}

	return nil
}

//=============================================================================

func (s *FTPServerConnection) createTarball(appFs afero.Fs, tarballFilePath string, filePaths []string, dryRun bool) error {
	s.logger.Info("Creating the backup archive...")
	// In-memory file system
	memFs := afero.NewMemMapFs()
	// Create a new archive file
	file, err := appFs.Create(tarballFilePath)
	if err != nil {
		return fmt.Errorf("could not create tarball file '%s', got error '%s'", tarballFilePath, err.Error())
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	pbConfig := &progressbar.Config{
		Items:          filePaths,
		OnCompletesMsg: fmt.Sprintf("Backup done! Saved as: %s", tarballFilePath),
		OnProgressCmd: func(path string) tea.Cmd {
			return createTarballTeaCmd(s, memFs, tarWriter, path, dryRun)
		},
	}

	if _, err := progressbar.Run(pbConfig); err != nil {
		return err
	}

	return nil
}

func addToTarWriter(memFs afero.Fs, filePath string, tarWriter *tar.Writer) error {
	file, err := memFs.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file '%s', got error '%s'", filePath, err.Error())
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not get stat for file '%s', got error '%s'", filePath, err.Error())
	}

	header := &tar.Header{
		Name:    filePath,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return fmt.Errorf("could not write header for file '%s', got error '%s'", filePath, err.Error())
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return fmt.Errorf("could not copy the file '%s' data to the tarball, got error '%s'", filePath, err.Error())
	}

	return nil
}
