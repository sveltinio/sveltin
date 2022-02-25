/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package ftpfs ...
package ftpfs

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"sort"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/spf13/afero"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/utils"
)

// FTPServerConnection is the struct with all is needed to establish and act on the FTP remote server.
type FTPServerConnection struct {
	Config       FTPConnectionConfig
	serverFolder string
	client       *ftp.ServerConn
}

// NewFTPServerConnection returns a new FTPServerConnection struct.
func NewFTPServerConnection(config *FTPConnectionConfig) FTPServerConnection {
	return FTPServerConnection{
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

// Dial contains the logic for the FTP receiver to handle the dial command.
func (s *FTPServerConnection) Dial() error {
	connStr := s.Config.makeConnectionString()
	jww.FEEDBACK.Printf("* Connecting to the FTP Server (%s) ", connStr)
	c, err := ftp.Dial(connStr, ftp.DialWithTimeout(time.Duration(s.Config.Timeout)*time.Second), ftp.DialWithDisabledEPSV(s.Config.IsEPSV))
	if err != nil {
		return err
	}
	s.client = c
	return nil
}

// Login contains the logic for the FTP receiver to handle the login command.
func (s *FTPServerConnection) Login() error {
	jww.FEEDBACK.Printf("* Login (as %s)", s.Config.User)
	if err := s.client.Login(s.Config.User, s.Config.Password); err != nil {
		return err
	}
	return nil
}

// Logout contains the logic for the FTP receiver to handle the logout command.
func (s *FTPServerConnection) Logout() error {
	jww.FEEDBACK.Println("* Closing the connection to the FTP server")
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
	if dryRun {
		jww.FEEDBACK.Println(common.HelperTextDryRunFlag())
	}

	pb := utils.NewProgressBar(len(folders))

	for _, folder := range folders {
		if dryRun {
			jww.FEEDBACK.Println("  ✔ " + folder + " -> would be created")
		} else {
			if err := s.client.MakeDir(folder); err != nil {
				return err
			}
			pb.Increment()
		}
	}
	pb.Wait()
	return nil
}

// Upload contains the logic for the FTP receiver to handle the upload files command.
func (s *FTPServerConnection) Upload(appFs afero.Fs, localDir string, files []string, dryRun bool) error {
	sort.Strings(files)

	// initialize progress container, with custom width
	pb := utils.NewProgressBar(len(files))

	for _, file := range files {
		fileBytes, err := afero.ReadFile(appFs, file)
		if err != nil {
			return err
		}

		remoteFile := utils.ToBasePath(file, localDir)
		if err = s.uploadSingle(remoteFile, bytes.NewBuffer(fileBytes), dryRun); err != nil {
			return err
		}
		pb.Increment()
	}
	pb.Wait()
	return nil
}

// DeleteAll contains the logic for the FTP receiver to handle the delete all command.
func (s *FTPServerConnection) DeleteAll(exclude []string, dryrun bool) error {
	entries, err := s.client.List(s.serverFolder)
	if err != nil {
		return err
	}

	if len(entries) > 0 {
		jww.FEEDBACK.Println("* Deleting previous content from the FTP remote folder")

		for _, entry := range entries {
			switch entry.Type {
			case ftp.EntryTypeFolder:
				if dryrun {
					jww.FEEDBACK.Println("  ✔ " + entry.Name + " -> folder would be recursively deleted")
				} else {
					folder := filepath.Join(s.serverFolder, entry.Name)
					if err := s.client.RemoveDirRecur(folder); err != nil {
						return err
					}
				}
			case ftp.EntryTypeFile:
				if dryrun {
					jww.FEEDBACK.Println("  ✔ " + entry.Name + " -> would be deleted")
				} else {
					file := filepath.Join(s.serverFolder, entry.Name)
					if !common.Contains(exclude, filepath.Base(file)) {
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
	jww.FEEDBACK.Printf("* Reading the remote folder '%s' ", s.serverFolder)
	remoteFiles := s.walkRemote()
	if err := s.createTarball(appFs, archiveFilename, remoteFiles, dryRun); err != nil {
		return err
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

	if dryRun {
		jww.FEEDBACK.Println("  ✔ " + filename + " -> would be uploaded")
	} else {
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
	jww.FEEDBACK.Printf("* Creating the archive file '%s' as backup", tarballFilePath)
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

	pb := utils.NewProgressBar(len(filePaths))

	for _, f := range filePaths {
		fPath := filepath.Dir(f)
		fName := filepath.Base(f)

		if err := s.client.ChangeDir(filepath.Join(s.serverFolder, fPath)); err != nil {
			return err
		}

		if dryRun {
			jww.FEEDBACK.Println("  ✔ " + f + " -> would be added to the archive file")
		} else {
			// fetch the file from the remote FTP server
			r, err := s.client.Retr(fName)
			if err != nil {
				return err
			}
			defer r.Close()
			// retrieve the file content
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			r.Close()
			// save file in the memory backed filesystem
			if err := afero.WriteFile(memFs, f, buf, 0777); err != nil {
				return err
			}
			// add the file to the tar archive
			if err := addToTarWriter(memFs, f, tarWriter); err != nil {
				return err
			}

			pb.Increment()

		}
	}

	pb.Wait()

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
