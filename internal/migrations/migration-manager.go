/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

// MigrationManager is the struct for the concrete mediator.
type MigrationManager struct {
	isFree         bool
	migrationQueue []Migration
}

// NewMigrationManager is the concrete Mediator.
func NewMigrationManager() *MigrationManager {
	return &MigrationManager{
		isFree: true,
	}
}

func (mm *MigrationManager) canRun(m Migration) bool {
	if mm.isFree {
		mm.isFree = false
		return true
	}
	mm.migrationQueue = append(mm.migrationQueue, m)
	return false
}

func (mm *MigrationManager) notifyAboutCompletion() error {
	if !mm.isFree {
		mm.isFree = true
	}

	if len(mm.migrationQueue) > 0 {
		firstMigrationInQueue := mm.migrationQueue[0]
		mm.migrationQueue = mm.migrationQueue[1:]
		if err := firstMigrationInQueue.allowUp(); err != nil {
			return err
		}
	}

	return nil
}
