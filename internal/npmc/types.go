/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package npmc ...
package npmc

import (
	"strings"
)

// NPMClient is the struct representing an npm client with name and version.
type NPMClient struct {
	Name    string
	Desc    string
	Version string
}

// NPMClientInfoStr is an alias for string
type NPMClientInfoStr string

// ToString returns a valid string representing the package manager name and version
// <pm_name@pm_version>
func (n NPMClient) ToString() string {
	return strings.TrimSuffix(n.Name+"@"+n.Version, "\n")
}

// ToNPMClient takes an NPMClientInfoStr string alias and returns the relative NPMClient struct
func (s NPMClientInfoStr) ToNPMClient() NPMClient {
	splitted := strings.Split(string(s), "@")
	return NPMClient{
		Name:    splitted[0],
		Version: splitted[1],
	}
}

// Script is a map representing the script section of a package.json file
type Script map[string]string

// Engine is a map representing the engine section of a package.json file
type Engine map[string]string

// PublishConfig is a map representing the publishConfig section of a package.json file
type PublishConfig map[string]string

// Dependency is a map representing the dependency section of a package.json file
type Dependency map[string]string

// Repository is a struct representing the repository section of a package.json file
type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Bugs is a struct representing the bugs section of a package.json file
type Bugs struct {
	URL string `json:"url"`
}

// PackageJSON is a struct representing a package.json file
type PackageJSON struct {
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
