/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/confirm"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/internal/ftpfs"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
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
	Short:   "Deploy your website over FTP.",
	Long: `Command used to deploy the project on your hosting platform over FTP.
`,
	Run: DeployCmdRun,
}

// DeployCmdRun is the actual work function.
func DeployCmdRun(cmd *cobra.Command, args []string) {
	// Exit if running sveltin commands either from a not valid directory or not latest sveltin version.
	isValidProject(true)

	cfg.log.Plain(markup.H1("Deploy your website to the FTP server"))

	// if --excludeFile is set, combines its lines with values from the --exclude flag.
	if len(withExcludeFile) != 0 {
		lines, err := common.ReadFileLineByLine(cfg.fs, withExcludeFile)
		utils.ExitIfError(err)
		withExclude = common.Union(withExclude, lines)
	}

	ftpConnectionConfig := newFTPConnectionConfig(cfg.prodData)

	ftpConn := ftpfs.NewFTPServerConnection(ftpConnectionConfig)
	ftpConn.SetRootFolder(cfg.prodData.FTPServerFolder)
	ftpConn.SetLogger(cfg.log)

	err := ftpfs.DialAction(&ftpConn).Run()
	utils.ExitIfError(err)

	err = ftpfs.LoginAction(&ftpConn).Run()
	utils.ExitIfError(err)

	// prevent the remote FTP server to close the idle connection
	noOpAction := ftpfs.IdleAction(&ftpConn)
	err = noOpAction.Run()
	utils.ExitIfError(err)

	common.ShowDeployCommandWarningMessages(isBackup)

	if isDryRun {
		common.PrintHelperTextDryRunFlag()
	}

	isConfirm, err := confirm.Run(&confirm.Config{Question: "Continue?"})
	utils.ExitIfError(err)

	if isConfirm {
		// create a local tar archive as backup for the remote folder content
		if isBackup {
			backupsFolderPath := filepath.Join(cfg.pathMaker.GetRootFolder(), BackupsFolder)
			utils.ExitIfError(common.MkDir(cfg.fs, backupsFolderPath))
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

		/**
		* The folder where adaper-static stores the output for the build process.
		* Check if pages and assets props for adapter-static are differents.
		**/
		kitBuildFolders := []string{cfg.projectSettings.SvelteKit.Adapter.Pages}
		if cfg.projectSettings.SvelteKit.Adapter.Pages != cfg.projectSettings.SvelteKit.Adapter.Assets {
			kitBuildFolders = append(kitBuildFolders, cfg.projectSettings.SvelteKit.Adapter.Assets)
		}

		for _, folder := range kitBuildFolders {
			foldersList := walkLocal(cfg.fs, EntryTypeFolder, folder)

			cfg.log.Info("Creating remote folders structure")
			err = ftpfs.MakeDirsAction(&ftpConn, foldersList, isDryRun).Run()
			utils.ExitIfError(err)

			// prevent the remote FTP server to close the idle connection
			err = noOpAction.Run()
			utils.ExitIfError(err)

			cfg.log.Info("Uploading files to the remote folders")
			filesList := walkLocal(cfg.fs, EntryTypeFile, folder)

			err = ftpfs.UploadAction(&ftpConn, cfg.fs, folder, filesList, isDryRun).Run()
			utils.ExitIfError(err)

			// prevent the remote FTP server to close the idle connection
			err = noOpAction.Run()
			utils.ExitIfError(err)
		}

		// close the connection
		err = ftpfs.LogoutAction(&ftpConn).Run()
		utils.ExitIfError(err)

		cfg.log.Success("Done\n")
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

func newFTPConnectionConfig(data tpltypes.EnvProductionData) *ftpfs.FTPConnectionConfig {
	return &ftpfs.FTPConnectionConfig{
		Host:     data.FTPHost,
		Port:     data.FTPPort,
		User:     data.FTPUser,
		Password: data.FTPPassword,
		Timeout:  data.FTPDialTimeout,
		IsEPSV:   data.FTPEPSVMode,
	}
}

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
