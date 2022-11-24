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
	"sort"
	"strings"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/prompti/confirm"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/internal/ftpfs"
	"github.com/sveltinio/sveltin/internal/markup"
	"github.com/sveltinio/sveltin/internal/tpltypes"
	"github.com/sveltinio/sveltin/tui/activehelps"
	"github.com/sveltinio/sveltin/tui/feedbacks"
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
	Short:   "Deploy your website over FTP",
	Long: `Command used to deploy the project on your hosting platform over FTP.
`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(0),
	Run:                   DeployCmdRun,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		comps = cobra.AppendActiveHelp(comps, activehelps.Hint("[WARN] This command does not take any argument but accepts flags."))
		return comps, cobra.ShellCompDirectiveDefault
	},
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

	err := ftpfs.DialAction(ftpConn).Run()
	utils.ExitIfError(err)

	err = ftpfs.LoginAction(ftpConn).Run()
	utils.ExitIfError(err)

	// prevent the remote FTP server to close the idle connection
	noOpAction := ftpfs.IdleAction(ftpConn)
	err = noOpAction.Run()
	utils.ExitIfError(err)

	feedbacks.ShowDeployCommandWarningMessages(isBackup)

	if isDryRun {
		feedbacks.ShowDryRunMessage()
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
			err = ftpfs.BackupAction(ftpConn, cfg.fs, filepath.Join(backupsFolderPath, projectName), isDryRun).Run()
			utils.ExitIfError(err)
		}

		// delete content from the remote folder with exclude list
		cfg.log.Important(fmt.Sprintf("If present, the following files will not be deleted from the remote folder: %s", strings.Join(withExclude, ", ")))
		err = ftpfs.DeleteAllAction(ftpConn, withExclude, isDryRun).Run()
		utils.ExitIfError(err)

		// create and update content from "kit.adapter.pages" folder
		kitPagesFolder := cfg.projectSettings.SvelteKit.Adapter.Pages
		pagesFoldersList, err := walkLocal(cfg.fs, EntryTypeFolder, kitPagesFolder, true)
		utils.ExitIfError(err)

		cfg.log.Infof("Creating remote folders structure for '%s'", kitPagesFolder)
		err = ftpfs.MakeDirsAction(ftpConn, pagesFoldersList, isDryRun).Run()
		utils.ExitIfError(err)

		// prevent the remote FTP server to close the idle connection
		err = noOpAction.Run()
		utils.ExitIfError(err)

		cfg.log.Infof("Uploading files to the remote folder '%s'", kitPagesFolder)
		pagesFilesList, err := walkLocal(cfg.fs, EntryTypeFile, kitPagesFolder, true)
		utils.ExitIfError(err)

		err = ftpfs.UploadAction(ftpConn, cfg.fs, kitPagesFolder, pagesFilesList, true, isDryRun).Run()
		utils.ExitIfError(err)

		// prevent the remote FTP server to close the idle connection
		err = noOpAction.Run()
		utils.ExitIfError(err)

		/**
		* Check if pages and assets props for adapter-static are differents.
		* If true, upload the entire kit.adapter.assets folder.
		**/
		kitAssetsFolder := cfg.projectSettings.SvelteKit.Adapter.Assets
		if kitPagesFolder != kitAssetsFolder {
			assetsFoldersList, err := walkLocal(cfg.fs, EntryTypeFolder, kitAssetsFolder, false)
			utils.ExitIfError(err)

			cfg.log.Infof("Creating remote folders structure for '%s'", kitAssetsFolder)
			err = ftpfs.MakeDirsAction(ftpConn, assetsFoldersList, isDryRun).Run()
			utils.ExitIfError(err)

			// prevent the remote FTP server to close the idle connection
			err = noOpAction.Run()
			utils.ExitIfError(err)

			cfg.log.Infof("Uploading files to the remote folder '%s'", kitAssetsFolder)
			assetsFilesList, err := walkLocal(cfg.fs, EntryTypeFile, kitAssetsFolder, false)
			utils.ExitIfError(err)

			err = ftpfs.UploadAction(ftpConn, cfg.fs, kitPagesFolder, assetsFilesList, false, isDryRun).Run()
			utils.ExitIfError(err)

			// prevent the remote FTP server to close the idle connection
			err = noOpAction.Run()
			utils.ExitIfError(err)
		}

		// close the connection
		err = ftpfs.LogoutAction(ftpConn).Run()
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

func walkLocal(fs afero.Fs, fType EntryType, dirname string, replaceBasePath bool) ([]string, error) {
	fList := []string{}
	err := afero.Walk(cfg.fs, dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			switch fType {
			case EntryTypeFolder:
				if info.IsDir() {
					// kit.prerender.pages content must be copied without the parent folder name
					if replaceBasePath {
						_path := utils.ToBasePath(path, dirname)
						fList = append(fList, _path)
						// kit.prerender.assets must be copied as whole folder
					} else {
						fList = append(fList, path)
					}
				}
			case EntryTypeFile:
				if !info.IsDir() && info.Name() != dirname {
					fList = append(fList, path)
				}
			}
			return nil
		})
	sort.Strings(fList)
	return fList, err
}
