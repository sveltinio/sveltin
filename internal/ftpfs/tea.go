package ftpfs

import (
	"archive/tar"
	"bytes"
	"io"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/afero"
	"github.com/sveltinio/prompti/progressbar"
	"github.com/sveltinio/sveltin/utils"
)

func mkDirTeaCmd(s *FTPServerConnection, path string, dryRun bool) tea.Cmd {
	if !dryRun {
		if err := s.client.MakeDir(path); err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
	}

	return func() tea.Msg {
		return progressbar.IncrementMsg(path)
	}
}

func uploadFileTeaCmd(s *FTPServerConnection, appFs afero.Fs, file, path string, replaceBasePath, dryRun bool) tea.Cmd {
	if !dryRun {
		fileBytes, err := afero.ReadFile(appFs, file)
		if err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
		var remoteFile string
		if replaceBasePath {
			remoteFile = utils.ToBasePath(file, path)
		} else {
			remoteFile = file
		}

		if err = s.uploadSingle(remoteFile, bytes.NewBuffer(fileBytes), dryRun); err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}

	}
	return func() tea.Msg {
		return progressbar.IncrementMsg(path)
	}
}

func createTarballTeaCmd(s *FTPServerConnection, memFs afero.Fs, tarWriter *tar.Writer, file string, dryRun bool) tea.Cmd {
	fPath := filepath.Dir(file)
	fName := filepath.Base(file)

	if err := s.client.ChangeDir(filepath.Join(s.serverFolder, fPath)); err != nil {
		return func() tea.Msg {
			return progressbar.IncrementErrMsg{Err: err}
		}
	}

	if !dryRun {
		// fetch the file from the remote FTP server
		r, err := s.client.Retr(fName)
		if err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
		defer r.Close()
		// retrieve the file content
		buf, err := io.ReadAll(r)
		if err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
		r.Close()
		// save file in the memory backed filesystem
		if err := afero.WriteFile(memFs, file, buf, 0777); err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
		// add the file to the tar archive
		if err := addToTarWriter(memFs, file, tarWriter); err != nil {
			return func() tea.Msg {
				return progressbar.IncrementErrMsg{Err: err}
			}
		}
	}

	return func() tea.Msg {
		return progressbar.IncrementMsg(fName)
	}
}
