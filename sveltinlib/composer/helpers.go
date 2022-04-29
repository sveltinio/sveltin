/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package composer ...
package composer

import (
	"github.com/sveltinio/sveltin/sveltinlib/pathmaker"
)

// GetRootFolder create a Folder struct for the project root folder.
func GetRootFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetRootFolder(),
	}
}

// GetConfigFolder create a Folder struct for the "config" folder of the project.
func GetConfigFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetConfigFolder(),
	}
}

// GetContentFolder create a Folder struct for the "content" folder of the project.
func GetContentFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetContentFolder(),
	}
}

// GetRoutesFolder create a Folder struct for the "src/routes" folder of the project.
func GetRoutesFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetRoutesFolder(),
	}
}

// GetAPIFolder create a Folder struct for the "src/routes/api" folder of the project.
func GetAPIFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetAPIFolder(),
	}
}

// GetLibFolder create a Folder struct for the "src/lib" folder of the project.
func GetLibFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetLibFolder(),
	}
}

// GetStaticFolder create a Folder struct for the "static" folder of the project.
func GetStaticFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetStaticFolder(),
	}
}

// GetThemesFolder create a Folder struct for the "themes" folder of the project.
func GetThemesFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetThemesFolder(),
	}
}
