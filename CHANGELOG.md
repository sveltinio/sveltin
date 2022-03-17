# Changelog

## [v0.7.1](https://github.com/sveltinio/sveltin/releases/tag/v0.7.1) (2022-03-17)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.7.0...v0.7.1)

**Closed issues:**

- PostCSS and Tailwind CSS [#10](https://github.com/sveltinio/sveltin/issues/10)

**Miscellaneous chores:**

- Index page for resources now use a flexbox instead of a grid, both on styled and unstyled projects. This gives a better and more controlled result.

## [v0.7.0](https://github.com/sveltinio/sveltin/releases/tag/v0.7.0) (2022-03-14)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.6.1...v0.7.0)

**Closed issues:**

- not-styled project [#6](https://github.com/sveltinio/sveltin/issues/6)

**Merged pull requests:**

- Allow unstyled projects creation by @indaco in https://github.com/sveltinio/sveltin/pull/7

**Miscellaneous chores:**

- cards grid styles
- mispelled word
- `sync` added as alias to the `prepare` command
- overall linting
- simplifing the `sveltinlib/css` implementation and allow styled and not-styled project creation
- more checks added (staticcheck, lint, tests)

## [v0.6.1](https://github.com/sveltinio/sveltin/releases/tag/v0.6.1) (2022-03-12)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.6.0...v0.6.1)

**Fixed bugs:**

- running sveltin commands from a not valid project directory
- set html scroll-behavior to smooth

## [v0.6.0](https://github.com/sveltinio/sveltin/releases/tag/v0.6.0) (2022-03-12)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.5...v0.6.0)

**Breaking changes:**

- `/api/v1/sveltin/version.json` endpoint renamed as `/api/v1/sveltin/info.json`, sveltekit version and build time exposed to it

**Fixed bugs:**

- remove type annotation, trivially inferred from a string literal on `default.js.ts`

**Miscellaneous chores:**

- moving from standalone endpoints to the [page endpoints](https://kit.svelte.dev/docs/routing#endpoints-page-endpoints) for resources and metadata
- new sveltin project will always use the latest **tested** version for sveltekit instead of the latest **released** one
- golang libraries (`cobra` and `afero`) updated the the lastest versions
- `go-pretty` library removed as dependency

## [v0.5.5](https://github.com/sveltinio/sveltin/releases/tag/v0.5.5) (2022-03-10)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.5...v0.5.5)

**Fixed bugs:**

- the lib template file for metadata (list) had a mistyped variable name causing runtime error

## [v0.5.4](https://github.com/sveltinio/sveltin/releases/tag/v0.5.4) (2022-03-09)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.3...v0.5.4)

**Closed issues:**

- [#1](https://github.com/sveltinio/sveltin/issues/1) - it is now possible to use readable name when running command in interactive way
- pages templates still used a wrong import for IWebPageMetadata from @sveltinio/seo
- [#2](https://github.com/sveltinio/sveltin/issues/2) - EndpointOutput import removed from api template files

**Miscellaneous chores:**

- string utilities methods refactored

## [v0.5.3](https://github.com/sveltinio/sveltin/releases/tag/v0.5.3) (2022-03-08)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.2...v0.5.3)

**Fixed bugs:**

- favicons with multiple formats and webmanifest file

## [v0.5.2](https://github.com/sveltinio/sveltin/releases/tag/v0.5.2) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.1...v0.5.2)

**Implemented enhancements:**

- `--withExcludeFile` flag added, to specify a file containing the list of files to not be deleted from the FTP server

## [v0.5.1](https://github.com/sveltinio/sveltin/releases/tag/v0.5.1) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.5.0...v0.5.1)

- fix: set `config.kit.prerender.default` to `true`

## [v0.5.0](https://github.com/sveltinio/sveltin/releases/tag/v0.5.0) (2022-03-07)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.4.0...v0.5.0)

**Breaking changes:**

- breaking: sveltin `prepare` command renamed as `install`. With [#4182](https://github.com/sveltejs/kit/pull/4182), SvelteKit added support for a new CLI command called `sync` wrapped as `prepare` script in `package.json` file.

**Fixed bugs:**

- import type IWebPageMetadata for index.svelte

**Implemented enhancements:**

- sveltin `prepare` command wraps the new SvelteKit `sync` command, [#4182](https://github.com/sveltejs/kit/pull/4182)
- sveltin `deploy` command places the TAR archive within the `backups` folder at project root level.

**Miscellaneous chores:**

- refactor: new underlying logging lib developed

## [v0.4.0](https://github.com/sveltinio/sveltin/releases/tag/v0.4.0) (2022-02-26)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.3.1...v0.4.0)

**Implemented enhancements:**

- `deploy` over FTP command added. [docs](https://docs.sveltin.io/cli/deploy)

**Miscellaneous chores:**

- overall code comments
- code cleansing
- better error handling

## [v0.3.1](https://github.com/sveltinio/sveltin/releases/tag/v0.3.1) (2022-02-17)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.3.0...v0.3.1)

**Fixed bugs:**

- Routing in app built by adapter-static and next.265 [#3801](https://github.com/sveltejs/kit/pull/3801)
- `.env` file templates with server port number

## [v0.3.0](https://github.com/sveltinio/sveltin/releases/tag/v0.3.0) (2022-02-15)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.14...v0.3.0)

**Breaking changes:**

- `--update` flag on `prepare` cmd removed. `prepare` now install the dependencies from `package.json` file only
- npmClient handling. It is now project specific and make uses of the packageManager key from `package.json`.

**Implemented enhancements:**

- `update` cmd added to update all the dependecies from `package.json` file
- `--port` flag added to `new` cmd to set the port to start the server on
- Bootstrap 5 support implemented
- SCSS support added

**Fixed bugs:**

- public page styles to have centered text for page title
- theme name on the `index.svelte` not replaced correctly. it now uses the new _ThemeName_ variable on the template data structure.

**Miscellaneous chores:**

- chore: **next steps** texts added to some commands
- refactor: go templates execution functions
- refactor: logger and pretty printer

## [v0.2.14](https://github.com/sveltinio/sveltin/releases/tag/v0.2.14) (2022-02-04)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.13...v0.2.14)

**Implemented enhancements:**

- Bulma support implemented

**Miscellaneous chores:**

- dependencies updated
- tailwind css and vanilla css themes updated
- each CSS supported lib has its own `__layout.svelte` file

## [v0.2.13](https://github.com/sveltinio/sveltin/releases/tag/v0.2.13) (2022-02-02)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.12...v0.2.13)

**Miscellaneous chores:**

- This version adheres to the changes introduced by `@sveltejs/kit 1.0.0-next.257` where the target option has been removed [#3674](https://github.com/sveltejs/kit/pull/3674)

## [v0.2.12](https://github.com/sveltinio/sveltin/releases/tag/v0.2.12) (2022-02-01)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.11...v0.2.12)

**Implemented enhancements:**

- json-ld breadcrumbs added to the public and resources pages

## [v0.2.11](https://github.com/sveltinio/sveltin/releases/tag/v0.2.11) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.10...v0.2.11)

**Fixed bugs:**

- paths to favicon, prism theme and script file

## [v0.2.10](https://github.com/sveltinio/sveltin/releases/tag/v0.2.10) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.9...v0.2.10)

**Fixed bugs:**

- `generate menu` command used `js` instead of `ts` as file extension causing errors on loading
- with tailwindcss used as css lib, the typography plugin's prose class was not rendered correctly
- postcss and its config file provided for tailwindcss only
- with vanillacss a github markdown theme added as default to render markdown content

## [v0.2.9](https://github.com/sveltinio/sveltin/releases/tag/v0.2.9) (2022-01-27)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.8...v0.2.9)

**Fixed bugs:**

- SvelteKit 1.0.0-next.244 fixed [#3473](https://github.com/sveltejs/kit/issues/3473) and [#3521](https://github.com/sveltejs/kit/pull/3521). `clone()` on fetch response as workaround to avoid '_body used already_' error when building the project removed

## [v0.2.8](https://github.com/sveltinio/sveltin/releases/tag/v0.2.8) (2022-01-26)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.7...v0.2.8)

**Fixed bugs:**

- valid page names and contents names
- variable names on page templates

## [v0.2.7](https://github.com/sveltinio/sveltin/releases/tag/v0.2.7) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.6...v0.2.7)

**Fixed bugs:**

- image path on seo components
- seo components added to pages

## [v0.2.6](https://github.com/sveltinio/sveltin/releases/tag/v0.2.6) (2022-01-25)

[Full Changelog](https://github.com/sveltinio/sveltin/compare/v0.2.6...v0.2.7)

**Fixed bugs:**

- seo for pages
- typescript interface names to match the new ones exposed from the components packages

**Miscellaneous chores:**

- dependencies updated
- command aliases added
