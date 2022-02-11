/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package packagejson

type Script map[string]string
type Engine map[string]string
type PublishConfig map[string]string
type Dependency map[string]string

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
type Bugs struct {
	URL string `json:"url"`
}

type PackageJson struct {
	Name            string        `json:"name"`
	Version         string        `json:"version"`
	Description     string        `json:"description"`
	Author          string        `json:"author"`
	License         string        `json:"license"`
	Keywords        []string      `json:"keywords"`
	Private         bool          `json:"private"`
	PackageManager  string        `json:"packageManager"`
	PublishConfig   PublishConfig `json:"publishConfig"`
	Engines         Engine        `json:"engines"`
	Workspace       []string      `json:"workspace"`
	Scripts         Script        `json:"scripts"`
	DevDependencies Dependency    `json:"devDependencies"`
	Dependencies    Dependency    `json:"dependencies"`
	Homepage        string        `json:"homepage"`
	Repository      Repository    `json:"repository"`
	Bugs            Bugs          `json:"bugs"`
	Type            string        `json:"type"`
}
