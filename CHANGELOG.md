# CHANGELOG

## [v0.10.0](https://github.com/sveltinio/sveltin/releases/tag/v0.10.0) (2022-09-16)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.9.1...v0.10.0)

### Fixed Bugs

- api endpoints were still on old sveltekit routing mechanism. Updated and fixed an issue when building the project due to `prerender=true` on those files. `fallback: '200.html'` on [static adapter configuration](https://github.com/sveltejs/kit/tree/master/packages/adapter-static) made the magic.

### 🚀 New Features

- `new resource` cmd allows to specify the group layout name according to the sveltekit [Advanced layouts](https://kit.svelte.dev/docs/advanced-routing#advanced-layouts) by passing the `--group` flag

	```bash
	sveltin new resource testimonials --group marketing
	```
- `new resource` cmd allow to specify if a different layout for the `slug` pages must be created in addition to the one for the `index` page.

	```bash
	sveltin new resource posts --slug
	```

### 🔧 Code Refactoring

- `config.TemplateData` struct makes use of individual struct for each artifact template data
- file templates updated accordingly
- generate commands (`menu`, `rss`, `sitemap`) simplified and updated to work for grouped layout too
- `GetAllRoutes` refactored to use `afero.Walk`

### Chores

- sveltekit updated to next.483
- go deps updated
- uniform function names

### Pull Requests

- Merge pull request [#119](https://github.com/sveltinio/sveltin/issues/119) from kit-advanced-layout

## [v0.9.1](https://github.com/sveltinio/sveltin/releases/tag/v0.9.1) (2022-09-06)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.9.0...v0.9.1)

### Fixed Bugs

- make generated page variable reactive

### Chores

- sveltekit updated to **next.472**
- upgrade minor npm deps
- indirect go deps added to generate command docs

### 📖  Documentation

- **cmds:** uniforming cobra help strings for commands

### Pull Requests

- Merge pull request [#110](https://github.com/sveltinio/sveltin/issues/110) from cobra-doc-deps
- Merge pull request [#111](https://github.com/sveltinio/sveltin/issues/111) from page-flags
- Merge pull request [#112](https://github.com/sveltinio/sveltin/issues/112) from content-flags
- Merge pull request [#113](https://github.com/sveltinio/sveltin/issues/113) from update-minor-npm-deps
- Merge pull request [#114](https://github.com/sveltinio/sveltin/issues/114) from uniforming-help-messages
- Merge pull request [#115](https://github.com/sveltinio/sveltin/issues/115) from fix-page-variable
- Merge pull request [#116](https://github.com/sveltinio/sveltin/issues/116) from sk-next.472

## [v0.9.0](https://github.com/sveltinio/sveltin/releases/tag/v0.9.0) (2022-09-05)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.12...v0.9.0)

### Fixed Bugs

- variable name when creating new resource with name like customers/projects
- **connection.go:** golangci SA1019
- **layout:** filename fixed to the new `+layout.svelte` when unstyled project
- **lib.gotxt:** unused param removed from list function

### 🚀 New Features

- **sveltekit:** updated to next.470 with adapter-static.42

### Breakings

- `init` command added (alias `create`) to scaffold a new project instead of `new`
- `new` command is now used **only** to create pages and resources (`routes`)
- `add` command added to create new content and metadata

### 🔧 Code Refactoring

- All prompts are now based on [@charmbracelet](https://github.com/charmbracelet) [bubbletea](https://github.com/charmbracelet/bubbletea) and [lipgloss](https://github.com/charmbracelet/lipgloss) through `sveltinio/prompti` package
- **logger:** externalised through `sveltinio/yinlog` based on [@charmbracelet](https://github.com/charmbracelet) [lipgloss](https://github.com/charmbracelet/lipgloss)
- drop dependency from `manifoldco/promptui`
- **errors:** styled
- **themes:** make embed themes folder structure more logical
- user messages styles
- drop dependency from `jww`
- moving pkgs to the right place -> `internal`
- lib files renamed as `load<Resource>` instead of `api<Resource>`

### ⚙️ CI

- github actions updated
- **test.yml:** updated to run tests against multiple OS and go versions
- **Earthfile:** updated include pkg folder

### Chores

- upgrade `github.com/vbauerster/mpb` to v8
- upgrade `github.com/jlaffaye/ftp` to the latest version
- error page (`+error.svelte`) added over project creation
- remove warnings about file is not gofmtd
- cleanup console log
- move CliVersion string to new version cmd
- renaming `internal/styles` package as `internal/markup`
- move builder, composer, css and pathmaker from pkg to internal
- go deps updated
- **go.mod:** tidy
- **npmc:** handle Desc as addition struct field
- **package.json:** svelte-kit sync added to avoid warnings on dev and check
- **wrapper.go:** golint ok

### 📖 Documentation

- **README:** updated to reflect cmd changes
- **cmds:** commands help messages updated, typos fixed for add and new cmds help messages
- **newResource:** typo fixed in code comment

### Pull Requests

- Merge pull request [#88](https://github.com/sveltinio/sveltin/issues/88) from sveltinio/typos-readme
- Merge pull request [#89](https://github.com/sveltinio/sveltin/issues/89) from sveltinio/rename-lib-files
- Merge pull request [#90](https://github.com/sveltinio/sveltin/issues/90) from sveltinio/fix-earthfile
- Merge pull request [#91](https://github.com/sveltinio/sveltin/issues/91) from sveltinio/refactor-user-prompts
- Merge pull request [#92](https://github.com/sveltinio/sveltin/issues/92) from sveltinio/refactor-styling-errors
- Merge pull request [#93](https://github.com/sveltinio/sveltin/issues/93) from sveltinio/refactor-unified-logging
- Merge pull request [#94](https://github.com/sveltinio/sveltin/issues/94) from sveltinio/rename-styles-markup
- Merge pull request [#95](https://github.com/sveltinio/sveltin/issues/95) from sveltinio/externalise-tui-prompts
- Merge pull request [#96](https://github.com/sveltinio/sveltin/issues/96) from sveltinio/externalise-logger
- Merge pull request [#97](https://github.com/sveltinio/sveltin/issues/97) from sveltinio/refactor-pkg-internal
- Merge pull request [#98](https://github.com/sveltinio/sveltin/issues/98) from sveltinio/no-jww
- Merge pull request [#99](https://github.com/sveltinio/sveltin/issues/99) from sveltinio/typos-comments-npmc
- Merge pull request [#100](https://github.com/sveltinio/sveltin/issues/100) from sveltinio/new-kit-routing
- Merge pull request [#101](https://github.com/sveltinio/sveltin/issues/101) from sveltinio/fix-layout-filename
- Merge pull request [#102](https://github.com/sveltinio/sveltin/issues/102) from sveltinio/add-error-page
- Merge pull request [#103](https://github.com/sveltinio/sveltin/issues/103) from sveltinio/reshape-embed-themes
- Merge pull request [#104](https://github.com/sveltinio/sveltin/issues/104) from sveltinio/ci-workflows
- Merge pull request [#105](https://github.com/sveltinio/sveltin/issues/105) from sveltinio/nested-resources
- Merge pull request [#106](https://github.com/sveltinio/sveltin/issues/106) from sveltinio/sveltekit-latest
- Merge pull request [#107](https://github.com/sveltinio/sveltin/issues/107) from sveltinio/go-deps
- Merge pull request [#109](https://github.com/sveltinio/sveltin/issues/109) from bump-vite-sveltekit

## [v0.8.12](https://github.com/sveltinio/sveltin/releases/tag/v0.8.12) (2022-08-04)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.11...v0.8.12)

### Fixed Bugs

- [#84](https://github.com/sveltinio/sveltin/issues/84)

### Chores

- upgrade to sveltekit next.403

### Pull Requests

- Merge pull request [#85](https://github.com/sveltinio/sveltin/issues/85) from sveltinio/fix-mdsvex
- Merge pull request [#86](https://github.com/sveltinio/sveltin/issues/86) from sveltinio/sveltekit-next.403
- Merge pull request [#87](https://github.com/sveltinio/sveltin/issues/87) from sveltinio/typos-readme

## [v0.8.11](https://github.com/sveltinio/sveltin/releases/tag/v0.8.11) (2022-08-02)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.10...v0.8.11)

### Fixed Bugs

- import bootstrap variables for v5.2.0
- **ci:** test workflow syntax
- **manifest.webmanifest:** reference path to manifest.webmanifest file

### Chores

- css libs updated
- readline as direct dependency
- go deps updated
- bump afero to 1.9.2
- bump sveltekit to next.386
- **app.css:** custom prism styles as sample
- **app.html:** Remove initial-scale=1 from <meta name="viewport">
- **ci:** splitting the lint and test github action workflows
- **package.json:** bump sveltekit to next.401 -> removing the prepare command/script
- **vite.config.js:** remove the alias to $lib
- **vite.config.js:** import defineConfig

### Pull Requests

- Merge pull request [#76](https://github.com/sveltinio/sveltin/issues/76) from sveltinio/vite-config
- Merge pull request [#77](https://github.com/sveltinio/sveltin/issues/77) from sveltinio/no-initial-scale
- Merge pull request [#78](https://github.com/sveltinio/sveltin/issues/78) from sveltinio/sk-401-no-prepare
- Merge pull request [#79](https://github.com/sveltinio/sveltin/issues/79) from sveltinio/split-lint-test-workflows
- Merge pull request [#80](https://github.com/sveltinio/sveltin/issues/80) from sveltinio/update-godeps
- Merge pull request [#81](https://github.com/sveltinio/sveltin/issues/81) from sveltinio/update-deps
- Merge pull request [#82](https://github.com/sveltinio/sveltin/issues/82) from sveltinio/fix-bootstrap-vars

## [v0.8.10](https://github.com/sveltinio/sveltin/releases/tag/v0.8.10) (2022-07-16)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.9...v0.8.10)

### Fixed Bugs

- avoid typescript linting errors
- [#68](https://github.com/sveltinio/sveltin/issues/68) param matchers name when '-' in resource and metadata name
- update to sveltekit next.377 with uppercase endpoint methods
- **apiIndex:** wrong import string

### Chores

- update to afero 1.9.0
- unused files for xml generation as endpoints removed
- uppercase endpoint methods as per sveltekit next.377

### Pull Requests

- Merge pull request [#65](https://github.com/sveltinio/sveltin/issues/65) from sveltinio/sveltekit-next-377
- Merge pull request [#66](https://github.com/sveltinio/sveltin/issues/66) from sveltinio/remove-unused-files
- Merge pull request [#67](https://github.com/sveltinio/sveltin/issues/69) from sveltinio/afero-update
- Merge pull request [#69](https://github.com/sveltinio/sveltin/issues/69) from sveltinio/fix-matcher-names

## [v0.8.9](https://github.com/sveltinio/sveltin/releases/tag/v0.8.9) (2022-07-15)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.8...v0.8.9)

### 🔧  Code Refactoring

- actual work function structure for commands
- resources and metadata API endpoints now are fully REST. This is really useful during the development. Instead to open a browser, just use `curl` on the terminal. E.g. `curl http://localshot:5173/api/v1/posts/category/webdev`

### Chores

- shortening and clearing help messages on commands
- sveltekit updated to next.375 with **Vite3** support
- git-ghlog config and template updated
- **adapter-static:** updated to next.36
- **app.html:** make uses of _%sveltekit.assets%_ to reference static files
- **vite.config.js:** clearScreen:false to prevent Vite from clearing the terminal

### Pull Requests

- Merge pull request [#57](https://github.com/sveltinio/sveltin/issues/57) from sveltinio/rest-endpoints
- Merge pull request [#58](https://github.com/sveltinio/sveltin/issues/58) from sveltinio/cmds-refactoring
- Merge pull request [#59](https://github.com/sveltinio/sveltin/issues/59) from sveltinio/deps-update
- Merge pull request [#60](https://github.com/sveltinio/sveltin/issues/60) from sveltinio/sveltekit-assets
- Merge pull request [#61](https://github.com/sveltinio/sveltin/issues/61) from sveltinio/vite3
- Merge pull request [#62](https://github.com/sveltinio/sveltin/issues/62) from sveltinio/git-chglog-revert
- Merge pull request [#63](https://github.com/sveltinio/sveltin/issues/63) from sveltinio/help-messages

## [v0.8.8](https://github.com/sveltinio/sveltin/releases/tag/v0.8.8) (2022-07-13)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.7...v0.8.8)

### Fixed Bugs

- svelte-check results with new tsconfig

### Chores

- update svelte-kit to next.371
- bump cli version to 0.8.8
- format and lint scripts updated to use their own ignore file
- **defaults.js.ts:** semicolon missed

### Pull Requests

- Merge pull request [#55](https://github.com/sveltinio/sveltin/issues/55) from sveltinio/release-0.8.8
- Merge pull request [#54](https://github.com/sveltinio/sveltin/issues/54) from sveltinio/sk-next-371
- Merge pull request [#53](https://github.com/sveltinio/sveltin/issues/53) from sveltinio/typescript

## [v0.8.7](https://github.com/sveltinio/sveltin/releases/tag/v0.8.7) (2022-07-13)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.6...v0.8.7)

### Fixed Bugs

- remove optimizeDeps config from vite.config.js

### Chores

- bump cli version to 0.8.7

### Pull Requests

- Merge pull request [#52](https://github.com/sveltinio/sveltin/issues/52) from sveltinio/release-0.8.7
- Merge pull request [#51](https://github.com/sveltinio/sveltin/issues/51) from sveltinio/fix-vite-optimizeDeps

## [v0.8.6](https://github.com/sveltinio/sveltin/releases/tag/v0.8.6) (2022-07-12)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.5...v0.8.6)

### Chores

- bump cli version to 0.8.6
- sveltekit updated to next.370 and other deps updated
- golang dependency updated

### 📖  Documentation

- **README:** updated

### Pull Requests

- Merge pull request [#50](https://github.com/sveltinio/sveltin/issues/50) from sveltinio/release-0.8.6
- Merge pull request [#49](https://github.com/sveltinio/sveltin/issues/49) from sveltinio/readme-typos

## [v0.8.5](https://github.com/sveltinio/sveltin/releases/tag/v0.8.5) (2022-07-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.4...v0.8.5)

### Fixed Bugs

- **svelte.config.js:** trailingSlash

### Chores

- bump cli version to 0.8.5

### Pull Requests

- Merge pull request [#48](https://github.com/sveltinio/sveltin/issues/48) from sveltinio/release-0.8.5

## [v0.8.4](https://github.com/sveltinio/sveltin/releases/tag/v0.8.4) (2022-07-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.3...v0.8.4)

### Fixed Bugs

- **vite.config.js:** aliases

### Chores

- release v0.8.4
- upgrade to sveltekit next.366 and static-adapter next.35

### Pull Requests

- Merge pull request [#47](https://github.com/sveltinio/sveltin/issues/47) from sveltinio/release-0.8.4
- Merge pull request [#46](https://github.com/sveltinio/sveltin/issues/46) from sveltinio/vite-aliases
- Merge pull request [#45](https://github.com/sveltinio/sveltin/issues/45) from sveltinio/sk-next.366

## [v0.8.3](https://github.com/sveltinio/sveltin/releases/tag/v0.8.3) (2022-07-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.2...v0.8.3)

### Fixed Bugs

- resource and metadata names to kebab case string
- **new:** temporarily hide the possibility to reuse an existing theme

### Chores

- release v0.8.3
- setup vite.config.js file for all CSS libs
- upgrade to cobra v.1.5.0
- **chglog:** config file updated to include merges, refs and issues
- **server:** run and dev aliases added to server cmd
- **tailwindcss:** postcss-load-config updated to ^4.0.1

### Pull Requests

- Merge pull request [#44](https://github.com/sveltinio/sveltin/issues/44) from sveltinio/release-0.8.3
- Merge pull request [#43](https://github.com/sveltinio/sveltin/issues/43) from sveltinio/chglog-include-merges
- Merge pull request [#42](https://github.com/sveltinio/sveltin/issues/42) from sveltinio/postcss-load-config-update
- Merge pull request [#41](https://github.com/sveltinio/sveltin/issues/41) from sveltinio/resource-kebab-case
- Merge pull request [#40](https://github.com/sveltinio/sveltin/issues/40) from sveltinio/hide-reuse
- Merge pull request [#39](https://github.com/sveltinio/sveltin/issues/39) from sveltinio/server-alias
- Merge pull request [#38](https://github.com/sveltinio/sveltin/issues/38) from sveltinio/sk-next.361

## [v0.8.2](https://github.com/sveltinio/sveltin/releases/tag/v0.8.2) (2022-06-02)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.1...v0.8.2)

### Fixed Bugs

- remove dayjs as dep

### Chores

- codeql action updated to v2

### Pull Requests

- Merge pull request [#37](https://github.com/sveltinio/sveltin/issues/37) from sveltinio/0.8.2

## [v0.8.1](https://github.com/sveltinio/sveltin/releases/tag/v0.8.1) (2022-06-02)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.0...v0.8.1)

### 🔧  Code Refactoring

- sveltinlib folder renamed as pkg
- root vars to struct

### Chores

- sveltekit updated to next.347 with latest adapter static
- upgrade to go yaml.v3
- viper updated
- golangci-lint updated
- deps updated to the latest versions

### Pull Requests

- Merge pull request [#36](https://github.com/sveltinio/sveltin/issues/36) from sveltinio/release-0.8.1
- Merge pull request [#35](https://github.com/sveltinio/sveltin/issues/35) from sveltinio/sveltekit-247
- Merge pull request [#34](https://github.com/sveltinio/sveltin/issues/34) from sveltinio/sveltinlib-to-pkg
- Merge pull request [#33](https://github.com/sveltinio/sveltin/issues/33) from sveltinio/application-struct
- Merge pull request [#32](https://github.com/sveltinio/sveltin/issues/32) from sveltinio/deps-update

## [v0.8.0](https://github.com/sveltinio/sveltin/releases/tag/v0.8.0) (2022-04-30)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.7.3...v0.8.0)

### Fixed Bugs

- metadata index page layout styles

### 🚀  New Features

- new flag `--git` to init a git repo on project creation
- make new theme and allow using existing themes
- resource and metadata endpoints created
- sveltin TS namespace created

### 🔧  Code Refactoring

- css lib setup

### ⚙️  CI

- bug report template updated to include the sveltin version number
- lint-test workflows on PR not only against main branch

### Chores

- git-ghlog config and template updated
- deps updated
- repo name for theme starter updated
- sveltekit updated to next.321
- golang.org/x/text as direct dep
- **ghurl_parser:** utility added to parse GitHub repository url
- **website.js.ts:** current year updated

### 📖  Documentation

- license owner updated to sveltin contributors
- **README:** updated

### Pull Requests

- Merge pull request [#31](https://github.com/sveltinio/sveltin/issues/31) from sveltinio/csslib-builder
- Merge pull request [#30](https://github.com/sveltinio/sveltin/issues/30) from sveltinio/readme-cmds
- Merge pull request [#29](https://github.com/sveltinio/sveltin/issues/29) from sveltinio/license-owner
- Merge pull request [#28](https://github.com/sveltinio/sveltin/issues/28) from sveltinio/sk-next-321
- Merge pull request [#27](https://github.com/sveltinio/sveltin/issues/27) from sveltinio/26-init-git-repo
- Merge pull request [#25](https://github.com/sveltinio/sveltin/issues/25) from sveltinio/theme-maker
- Merge pull request [#24](https://github.com/sveltinio/sveltin/issues/24) from sveltinio/api-endpoints
- Merge pull request [#23](https://github.com/sveltinio/sveltin/issues/23) from sveltinio/22-metadata-index-wrong-styles-tailwindcss
- Merge pull request [#21](https://github.com/sveltinio/sveltin/issues/21) from sveltinio/ci-bug-report-template
- Merge pull request [#20](https://github.com/sveltinio/sveltin/issues/20) from sveltinio/sveltin-namespace
- Merge pull request [#19](https://github.com/sveltinio/sveltin/issues/19) from sveltinio/deps-update

## [v0.7.3](https://github.com/sveltinio/sveltin/releases/tag/v0.7.3) (2022-04-04)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.7.2...v0.7.3)

### Fixed Bugs

- interfaces names to match the new ones from the packages

### ⚙️  CI

- Trigger Build

### Chores

- bump CLI version to 0.7.3
- sveltinio/services updated to v0.2.0
- sveltekit updated to 1.0.0-next.302
- sveltekit updated to next-301. fallthrough removed

### Pull Requests

- Merge pull request [#18](https://github.com/sveltinio/sveltin/issues/18) from sveltinio/codeql
- Merge pull request [#17](https://github.com/sveltinio/sveltin/issues/17) from sveltinio/sveltekit-next-301

## [v0.7.2](https://github.com/sveltinio/sveltin/releases/tag/v0.7.2) (2022-03-21)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.7.1...v0.7.2)

### Fixed Bugs

- speed up bulma & bootstrap loadings
- colours and icons on windows

### ⚙️  CI

- lint-test workflow added

### Chores

- bump CLI version to 0.7.2
- linting
- use golangci-lint
- reference to variable.scss file replaced as _variable.scss for svelte.config.js
- logger and prompt select icons updated
- **.chglog:** git-chglog CHANGELOG generator config added
- **commit-msg:** colour and icon added to the error messages

### Pull Requests

- Merge pull request [#16](https://github.com/sveltinio/sveltin/issues/16) from sveltinio/14-speed-up-bulma-bootstrap-loadings
- Merge pull request [#13](https://github.com/sveltinio/sveltin/issues/13) from sveltinio/windows-colors

## [v0.7.1](https://github.com/sveltinio/sveltin/releases/tag/v0.7.1) (2022-03-17)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.7.0...v0.7.1)

### Fixed Bugs

- postcss for tailwind based projects

### ⚙️  CI

- githooks added and simplified release workflow

### Chores

- set scroll behaviour to smooth
- bump CLI version to 0.7.1
- update pre-push hook
- **index.svelte:** use flexbox instead of grid

### 📖  Documentation

- made with svelte shield added
- **README:** project status section updated

### Pull Requests

- Merge pull request [#11](https://github.com/sveltinio/sveltin/issues/11) from sveltinio/10-postcss-and-tailwind-css

## [v0.7.0](https://github.com/sveltinio/sveltin/releases/tag/v0.7.0) (2022-03-14)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.6.1...v0.7.0)

### Fixed Bugs

- cards grid styles
- mispelled fixed

### 🚀  New Features

- support for not styled project creation added

### 🔧  Code Refactoring

- make PromptGetSelect working with slice of strings and PromptObject

### Chores

- bump CLI version to 0.7.0
- sync added as alias to the prepare command
- human readable messages for prompts
- .gitignore updated
- error messages updated. NewOptionNotValidError now takes the used value and the corrects ones as args
- overall liting

### Pull Requests

- Merge pull request [#7](https://github.com/sveltinio/sveltin/issues/7) from sveltinio/skeleton-project

## [v0.6.1](https://github.com/sveltinio/sveltin/releases/tag/v0.6.1) (2022-03-12)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.6.0...v0.6.1)

### Fixed Bugs

- set html scroll-behavior to smooth
- exit if running sveltin commands from a not valid directory

### Chores

- bump CLI version to 0.6.1

## [v0.6.0](https://github.com/sveltinio/sveltin/releases/tag/v0.6.0) (2022-03-12)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.5...v0.6.0)

### Fixed Bugs

- remove type annotation, trivially inferred from a string literal

### 🚀  New Features

- page endpoints support added
- **logger:** log.Plain added

### Chores

- bump CLI version to 0.6.0
- replacing listLogger with direct log calls
- utility function to get underline text added
- use the last sveltekit and static-adapter tested versions
- **root.go:** unused function removed

### Pull Requests

- Merge pull request [#5](https://github.com/sveltinio/sveltin/issues/5) from sveltinio/deps-update
- Merge pull request [#4](https://github.com/sveltinio/sveltin/issues/4) from sveltinio/shadows

## [v0.5.5](https://github.com/sveltinio/sveltin/releases/tag/v0.5.5) (2022-03-10)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.4...v0.5.5)

### Fixed Bugs

- **libList.gotxt:** resource name variable

### Chores

- bump CLI version to 0.5.5

## [v0.5.4](https://github.com/sveltinio/sveltin/releases/tag/v0.5.4) (2022-03-09)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.3...v0.5.4)

### Fixed Bugs

- artifacts names in human readable way support added

### Chores

- bump CLI version to 0.5.4

### Pull Requests

- Merge pull request [#3](https://github.com/sveltinio/sveltin/issues/3) from sveltinio/fix/human-readable-names

## [v0.5.3](https://github.com/sveltinio/sveltin/releases/tag/v0.5.3) (2022-03-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.2...v0.5.3)

### Fixed Bugs

- favicons

### Chores

- bump CLI version to 0.5.3

## [v0.5.2](https://github.com/sveltinio/sveltin/releases/tag/v0.5.2) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.1...v0.5.2)

### Fixed Bugs

- **deploy.go:** --excludeFile flag renamed as --withExcludeFile

### 🚀  New Features

- **deploy.go:** --excludeFile flag added

### Chores

- bump CLI version to 0.5.2
- **collections.go:** unique and union methods added
- **fs.go:** method ReadFileLineByLine added

## [v0.5.1](https://github.com/sveltinio/sveltin/releases/tag/v0.5.1) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.0...v0.5.1)

### Fixed Bugs

- **svelte.config.js:** set `config.kit.prerender.default` to `true`

### Chores

- bump CLI version to 0.5.1

## [v0.5.0](https://github.com/sveltinio/sveltin/releases/tag/v0.5.0) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.4.0...v0.5.0)

### 🚀  New Features

- sveltin `prepare` command

### 🔧  Code Refactoring

- nest_steps.go renamed as help_messages.go
- new underlying logging lib developed
- CheckIfError renamed as ExitIfError to reflect what it does

### Chores

- bump CLI version to 0.5.0
- deploy command places the backup file within backups folder
- retrieve project name from package.json file
- golang deps updated
- dependencies updated
- fatalf string updated
- using the new logging library
- add PromptConfirm util function for asking a yes or no question
- new IsError method added.
- commented code block deleted
- **prompt.go:** Select instead of SelectAdd

### 📖  Documentation

- sveltin root command documentation updated

## [v0.4.0](https://github.com/sveltinio/sveltin/releases/tag/v0.4.0) (2022-02-26)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.3.1...v0.4.0)

### Fixed Bugs

- errors handling
- **newContent.go:** template strings
- **next_steps.go:** typo

### 🚀  New Features

- deploy over FTP command added

### 🔧  Code Refactoring

- **next_steps.go:** interpolate multiline strings

### Chores

- cli version bumped to 0.4.0
- code cleansing
- delete newPage_test file
- struct SiteConfig renamed as ProjectConfig and moved to a specific file
- **pages.go:** lint
- **text.go:** method ToBasePath added

### 📖  Documentation

- code comments
- overall code comments
- typos fixed
- cmd descriptions updated
- **README:** updated
- **generateSitemap:** typo

## [v0.3.1](https://github.com/sveltinio/sveltin/releases/tag/v0.3.1) (2022-02-17)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.3.0...v0.3.1)

### Fixed Bugs

- env file templates with server port number
- trailingSlash to always

### Chores

- bump cli version to 0.3.1

### 📖  Documentation

- typo
- scss added to the list of css libs

## [v0.3.0](https://github.com/sveltinio/sveltin/releases/tag/v0.3.0) (2022-02-15)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.14...v0.3.0)

### Fixed Bugs

- typos
- pages styles fixed to have center text

### 🚀  New Features

- SCSS support added

### 🔧  Code Refactoring

- template execution
- logger and printer
- package.json file and npmclient handling
- package manager handling

### Chores

- cli version bumped to 0.3.0

### 📖  Documentation

- updated

## [v0.2.14](https://github.com/sveltinio/sveltin/releases/tag/v0.2.14) (2022-02-04)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.13...v0.2.14)

### 🚀  New Features

- bulma support implemented

### 🔧  Code Refactoring

- tailwind css and vanilla css themes

### Chores

- cli version bumped to 0.2.14
- editorconfig updated
- dependencies update

## [v0.2.13](https://github.com/sveltinio/sveltin/releases/tag/v0.2.13) (2022-02-02)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.12...v0.2.13)

### Chores

- cli version bumped to 0.2.13
- Remove target option

## [v0.2.12](https://github.com/sveltinio/sveltin/releases/tag/v0.2.12) (2022-02-01)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.11...v0.2.12)

## [v0.2.11](https://github.com/sveltinio/sveltin/releases/tag/v0.2.11) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.10...v0.2.11)

### Fixed Bugs

- **app.html:** path to favicon, prism theme and script file

### Chores

- cli version bumped to 0.2.11

## [v0.2.10](https://github.com/sveltinio/sveltin/releases/tag/v0.2.10) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.9...v0.2.10)

### Fixed Bugs

- generate menu command used js instead of ts as file extension causing errors on loading

## [v0.2.9](https://github.com/sveltinio/sveltin/releases/tag/v0.2.9) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.8...v0.2.9)

### 🔧  Code Refactoring

- clone()` on fetch response workaround to avoid '_body used already_' error building the project removed

### 📖  Documentation

- readme updated

## [v0.2.8](https://github.com/sveltinio/sveltin/releases/tag/v0.2.8) (2022-01-26)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.7...v0.2.8)

### Fixed Bugs

- pages templates and variables names
- image path on seo components

### Chores

- changelog file added
- cli version bumped to 0.2.7

### 📖  Documentation

- readme updated

## [v0.2.7](https://github.com/sveltinio/sveltin/releases/tag/v0.2.7) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.6...v0.2.7)

## [v0.2.6](https://github.com/sveltinio/sveltin/releases/tag/v0.2.6) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.5...v0.2.6)

### Fixed Bugs

- seo for pages
- interfaces names to match the new ones from the packages

### Chores

- cli version bumped to 0.2.6
- dependencies update

### 📖  Documentation

- **README:** aliases added
