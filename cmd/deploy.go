/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package cmd ...
package cmd

import (
	"os"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/sveltinlib/ftpfs"
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
	isDryRun    bool
	isBackup    bool
	withExclude []string
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
	textLogger.Reset()
	textLogger.SetTitle("Deploy your website to the FTP server")
	// LOG TO STDOUT
	utils.PrettyPrinter(textLogger).Print()

	ftpConnectionConfig := &ftpfs.FTPConnectionConfig{
		Host:     projectConfig.FTPHost,
		Port:     projectConfig.FTPPort,
		User:     projectConfig.FTPUser,
		Password: projectConfig.FTPPassword,
		Timeout:  projectConfig.FTPDialTimeout,
		IsEPSV:   projectConfig.FTPEPSVMode,
	}
	ftpConn := ftpfs.NewFTPServerConnection(ftpConnectionConfig)
	ftpConn.SetRootFolder(projectConfig.FTPServerFolder)

	err := ftpfs.DialAction(&ftpConn).Run()
	utils.CheckIfError(err)

	err = ftpfs.LoginAction(&ftpConn).Run()
	utils.CheckIfError(err)

	// prevent the remote FTP server to close the idle connection
	noOpAction := ftpfs.IdleAction(&ftpConn)
	err = noOpAction.Run()
	utils.CheckIfError(err)

	// create a tar archive on the FTP remote folder as backup
	if isBackup {
		err = ftpfs.BackupAction(&ftpConn, AppFs, "backup", isDryRun).Run()
		utils.CheckIfError(err)
	}

	// delete content from the remote folder with exclude list
	err = ftpfs.DeleteAllAction(&ftpConn, withExclude, isDryRun).Run()
	utils.CheckIfError(err)

	// Create the folders structure
	jww.FEEDBACK.Println("* Creating remote folders structure")
	foldersList := walkLocal(AppFs, EntryTypeFolder, projectConfig.SvelteKitBuildFolder)
	err = ftpfs.MakeDirsAction(&ftpConn, foldersList, isDryRun).Run()
	utils.CheckIfError(err)

	// prevent the remote FTP server to close the idle connection
	err = noOpAction.Run()
	utils.CheckIfError(err)

	// upload files
	jww.FEEDBACK.Println("* Uploading files to the remote folders")
	filesList := walkLocal(AppFs, EntryTypeFile, projectConfig.SvelteKitBuildFolder)
	err = ftpfs.UploadAction(&ftpConn, AppFs, projectConfig.SvelteKitBuildFolder, filesList, isDryRun).Run()
	utils.CheckIfError(err)

	// prevent the remote FTP server to close the idle connection
	err = noOpAction.Run()
	utils.CheckIfError(err)

	// close the connection
	err = ftpfs.LogoutAction(&ftpConn).Run()
	utils.CheckIfError(err)

	// LOG SUMMARY TO THE STDOUT
	jww.FEEDBACK.Println(common.HelperTextDeploySummary(len(foldersList), len(filesList)))

}

func deployCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&isBackup, "backup", "b", false, "create a tar archive for the existing content on the remote FTP server")
	cmd.Flags().BoolVarP(&isDryRun, "dryRun", "d", false, "dry run")
	cmd.Flags().StringArrayVarP(&withExclude, "exclude", "e", []string{".htaccess"}, "list of files to not be deleted from the FTP server. Default: .htaccess")
}

func init() {
	deployCmdFlags(deployCmd)
	rootCmd.AddCommand(deployCmd)
}

//=============================================================================

func walkLocal(fs afero.Fs, fType EntryType, dirname string) []string {
	fList := []string{}
	err := afero.Walk(AppFs, dirname,
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
	utils.CheckIfError(err)
	return fList
}
