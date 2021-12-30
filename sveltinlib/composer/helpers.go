/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package composer

import (
	"github.com/sveltinio/sveltin/sveltinlib/pathmaker"
)

func GetRootFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetRootFolder(),
	}
}

func GetConfigFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetConfigFolder(),
	}
}

func GetContentFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetContentFolder(),
	}
}

func GetRoutesFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetRoutesFolder(),
	}
}

func GetAPIFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetAPIFolder(),
	}
}

func GetLibFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetLibFolder(),
	}
}

func GetStaticFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetStaticFolder(),
	}
}

func GetThemesFolder(maker *pathmaker.SveltinPathMaker) *Folder {
	return &Folder{
		Name: maker.GetThemesFolder(),
	}
}

// --------------------------------------------------------------------
