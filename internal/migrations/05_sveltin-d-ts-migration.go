/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"

	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/resources"
)

// SveltinDTSMigration is the struct representing the migration add the sveltin.json file.
type SveltinDTSMigration struct {
	Mediator IMigrationMediator
	Services *MigrationServices
	Data     *MigrationData
}

// MakeMigration implements IMigrationFactory interface.
func (m *SveltinDTSMigration) MakeMigration(migrationManager *MigrationManager, services *MigrationServices, data *MigrationData) IMigration {
	return &SveltinDTSMigration{
		Mediator: migrationManager,
		Services: services,
		Data:     data,
	}
}

// MakeMigration implements IMigration interface.
func (m *SveltinDTSMigration) getServices() *MigrationServices { return m.Services }
func (m *SveltinDTSMigration) getData() *MigrationData         { return m.Data }

// Execute return error if migration execution over up and down methods fails.
func (m SveltinDTSMigration) Execute() error {
	if err := m.up(); err != nil {
		return err
	}
	if err := m.down(); err != nil {
		return err
	}
	return nil
}

func (m *SveltinDTSMigration) up() error {
	if !m.Mediator.canRun(m) {
		return nil
	}

	exists, err := common.FileExists(m.getServices().fs, m.Data.TargetPath)
	if !exists {
		return err
	}

	if exists {
		fileContent, err := retrieveFileContent(m.getServices().fs, m.getData().TargetPath)
		if err != nil {
			return err
		}

		if !bytes.Contains(fileContent, []byte(patterns[sveltindts])) {
			m.getServices().logger.Info(fmt.Sprintf("Migrating %s", filepath.Base(m.Data.TargetPath)))
			saveTo := path.Join(m.Services.pathMaker.GetSrcFolder())
			return m.Services.fsManager.CopyFileFromEmbed(&resources.SveltinStaticFS, m.Services.fs, resources.SveltinFilesFS, "sveltin_d_ts", saveTo)
		}
	}

	return nil
}

func (m *SveltinDTSMigration) down() error {
	if err := m.Mediator.notifyAboutCompletion(); err != nil {
		return err
	}
	return nil
}

func (m *SveltinDTSMigration) allowUp() error {
	if err := m.up(); err != nil {
		return err
	}
	return nil
}

func (m *SveltinDTSMigration) migrate(content []byte, file string) ([]byte, error) {
	return nil, nil
}

//=============================================================================
