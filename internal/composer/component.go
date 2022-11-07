/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package composer

import (
	"github.com/sveltinio/sveltin/helpers/factory"
)

// Component describes operations common to elements of the tree.
type Component interface {
	GetName() string
	SetName(string)
	SetPath(string)
	Create(*factory.Artifact) error
}
