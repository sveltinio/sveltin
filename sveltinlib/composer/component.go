/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package composer

import (
	"github.com/sveltinio/sveltin/helpers/factory"
)

type component interface {
	GetName() string
	SetName(string)
	SetPath(string)
	Create(*factory.Artifact) error
}
