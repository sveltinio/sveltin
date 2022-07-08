# CHANGELOG

## [v0.8.4](https://github.com/sveltinio/sveltin/releases/tag/v0.8.4) (2022-07-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.8.3...v0.8.4)

### Fixed Bugs

- **vite.config.js:** aliases

### 📖  Documentation

- **CHANGELOG:** updated

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

### 📖  Documentation

- **CHANGELOG:** updated

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

### 📖  Documentation

- license owner updated to sveltin contributors
- **CHANGELOG:** updated
- **README:** updated

### Chores

- git-ghlog config and template updated
- deps updated
- repo name for theme starter updated
- sveltekit updated to next.321
- golang.org/x/text as direct dep
- **ghurl_parser:** utility added to parse GitHub repository url
- **website.js.ts:** current year updated

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

### 📖  Documentation

- **CHANGELOG:** updated

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

### 📖  Documentation

- made with svelte shield added
- **CHANGELOG:** updated
- **CHANGELOG:** updated
- **README:** project status section updated

### Chores

- set scroll behaviour to smooth
- bump CLI version to 0.7.1
- update pre-push hook
- **index.svelte:** use flexbox instead of grid

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

### Fixed Bugs

- **CHANGELOG:** typo

### 🚀  New Features

- sveltin `prepare` command

### 🔧  Code Refactoring

- nest_steps.go renamed as help_messages.go
- new underlying logging lib developed
- CheckIfError renamed as ExitIfError to reflect what it does

### 📖  Documentation

- sveltin root command documentation updated

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

### 📖  Documentation

- code comments
- overall code comments
- typos fixed
- cmd descriptions updated
- **README:** updated
- **generateSitemap:** typo

### Chores

- cli version bumped to 0.4.0
- code cleansing
- delete newPage_test file
- struct SiteConfig renamed as ProjectConfig and moved to a specific file
- **pages.go:** lint
- **text.go:** method ToBasePath added

## [v0.3.1](https://github.com/sveltinio/sveltin/releases/tag/v0.3.1) (2022-02-17)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.3.0...v0.3.1)

### Fixed Bugs

- env file templates with server port number
- trailingSlash to always

### 📖  Documentation

- typo
- scss added to the list of css libs

### Chores

- bump cli version to 0.3.1

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

### 📖  Documentation

- updated

### Chores

- cli version bumped to 0.3.0

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

### 📖  Documentation

- readme updated

### Chores

- changelog file added
- cli version bumped to 0.2.7

## [v0.2.7](https://github.com/sveltinio/sveltin/releases/tag/v0.2.7) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.6...v0.2.7)

## [v0.2.6](https://github.com/sveltinio/sveltin/releases/tag/v0.2.6) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.5...v0.2.6)

### Fixed Bugs

- seo for pages
- interfaces names to match the new ones from the packages

### 📖  Documentation

- **README:** aliases added

### Chores

- cli version bumped to 0.2.6
- dependencies update
