/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/pkg/ftpfs"
	"github.com/sveltinio/sveltin/utils"
)

// EntryType describes the different types of an Entry
type EntryType int

// The differents types of an Entry
const (
	EntryTypeFolder EntryType = 0
	EntryTypeFile   EntryType = 1
)

var (
	isDryRun        bool
	isBackup        bool
	withExclude     []string
	withExcludeFile string
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"publish"},
	Short:   "Command to deploy your website over FTP",
	Long: `This command deploy your project on your hosting platform over FTP.
`,
	Run: DeployCmdRun,
}

// DeployCmdRun is the actual work function.
func DeployCmdRun(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands from a not valid directory.
	isValidProject()

	cfg.log.Plain(utils.Underline("Deploy your website to the FTP server"))

	// if --excludeFile is set, combines its lines with values from the --exclude flag.
	if len(withExcludeFile) != 0 {
		lines, err := common.ReadFileLineByLine(cfg.fs, withExcludeFile)
		utils.ExitIfError(err)
		withExclude = common.Union(withExclude, lines)

	}

	ftpConnectionConfig := &ftpfs.FTPConnectionConfig{
		Host:     cfg.project.FTPHost,
		Port:     cfg.project.FTPPort,
		User:     cfg.project.FTPUser,
		Password: cfg.project.FTPPassword,
		Timeout:  cfg.project.FTPDialTimeout,
		IsEPSV:   cfg.project.FTPEPSVMode,
	}
	ftpConn := ftpfs.NewFTPServerConnection(ftpConnectionConfig)
	ftpConn.SetRootFolder(cfg.project.FTPServerFolder)

	err := ftpfs.DialAction(&ftpConn).Run()
	utils.ExitIfError(err)

	err = ftpfs.LoginAction(&ftpConn).Run()
	utils.ExitIfError(err)

	// prevent the remote FTP server to close the idle connection
	noOpAction := ftpfs.IdleAction(&ftpConn)
	err = noOpAction.Run()
	utils.ExitIfError(err)

	common.ShowDeployCommandWarningMessages(cfg.log)
	//setDefaultLoggerOptions()

	confirmStr := promptBackupConfirm()

	if isConfirm(confirmStr) {
		// create a local tar archive as backup for the remote folder content
		if isBackup {
			backupsFolderPath := filepath.Join(cfg.pathMaker.GetRootFolder(), BackupsFolder)
			utils.ExitIfError(common.MkDir(cfg.fs))
			pathToPkgFile := filepath.Join(cfg.pathMaker.GetRootFolder(), "package.json")
			projectName, err := utils.RetrieveProjectName(cfg.fs, pathToPkgFile)
			utils.ExitIfError(err)
			err = ftpfs.BackupAction(&ftpConn, cfg.fs, filepath.Join(backupsFolderPath, projectName), isDryRun).Run()
			utils.ExitIfError(err)
		}

		// delete content from the remote folder with exclude list
		cfg.log.Important(fmt.Sprintf("The following files will not be deleted from the remote folder: %s", strings.Join(withExclude, ", ")))
		err = ftpfs.DeleteAllAction(&ftpConn, withExclude, isDryRun).Run()
		utils.ExitIfError(err)

		// Create the folders structure
		cfg.log.Info("Creating remote folders structure")
		foldersList := walkLocal(cfg.fs, EntryTypeFolder, cfg.project.SvelteKitBuildFolder)
		err = ftpfs.MakeDirsAction(&ftpConn, foldersList, isDryRun).Run()
		utils.ExitIfError(err)

		// prevent the remote FTP server to close the idle connection
		err = noOpAction.Run()
		utils.ExitIfError(err)

		// upload files
		cfg.log.Info("Uploading files to the remote folders")
		filesList := walkLocal(cfg.fs, EntryTypeFile, cfg.project.SvelteKitBuildFolder)
		err = ftpfs.UploadAction(&ftpConn, cfg.fs, cfg.project.SvelteKitBuildFolder, filesList, isDryRun).Run()
		utils.ExitIfError(err)

		// prevent the remote FTP server to close the idle connection
		err = noOpAction.Run()
		utils.ExitIfError(err)

		// close the connection
		err = ftpfs.LogoutAction(&ftpConn).Run()
		utils.ExitIfError(err)

		// LOG SUMMARY TO THE STDOUT
		cfg.log.Info(common.HelperTextDeploySummary(len(foldersList), len(filesList)))
		cfg.log.Success("Done")
	}
}

func deployCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&isBackup, "backup", "b", true, "create a tar archive for the existing content on the remote FTP server")
	cmd.Flags().BoolVarP(&isDryRun, "dryRun", "d", false, "dry run")
	cmd.Flags().StringArrayVarP(&withExclude, "exclude", "e", []string{".htaccess"}, "list of files to not be deleted from the FTP server. Default: .htaccess")
	cmd.Flags().StringVar(&withExcludeFile, "withExcludeFile", "", "path to the file containing the list of files to not be deleted from the FTP server.")
}

func init() {
	deployCmdFlags(deployCmd)
	rootCmd.AddCommand(deployCmd)
}

//=============================================================================

func walkLocal(fs afero.Fs, fType EntryType, dirname string) []string {
	fList := []string{}
	err := afero.Walk(cfg.fs, dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			switch fType {
			case EntryTypeFolder:
				if info.IsDir() && info.Name() != dirname {
					_path := utils.ToBasePath(path, dirname)
					fList = append(fList, _path)
				}
			case EntryTypeFile:
				if !info.IsDir() && info.Name() != dirname {
					fList = append(fList, path)
				}
			}
			return nil
		})
	utils.ExitIfError(err)
	return fList
}

func promptBackupConfirm() string {
	return common.PromptConfirm("Do you wish to continue?")
}

func isConfirm(value string) bool {
	return value == "y"
}
