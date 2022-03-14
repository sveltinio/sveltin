# sveltin

## 0.7.0

### Minor Changes

- feature:
- fix: cards grid styles
- fix: mispelled word
- chore: `sync` added as alias to the `prepare` command
- chore: overall linting
- refactor: simplifing the `sveltinlib/css` implementation and allow styled and not-styled project creation
- ci: more checks added (staticcheck, lint, tests)

## 0.6.1

### Patch Changes

- fix: no more possible run sveltin commands from a not valid project directory
- fix: set html scroll-behavior to smooth

## 0.6.0

### Minor Changes

- breaking: `/api/v1/sveltin/version.json` endpoint renamed as `/api/v1/sveltin/info.json`, sveltekit version and build time exposed to it
- refactor: moving from standalone endpoints to the [page endpoints](https://kit.svelte.dev/docs/routing#endpoints-page-endpoints) for resources and metadata
- chore: new sveltin project will always use the latest **tested** version for sveltekit instead of the latest **released** one
- chore: golang libraries (`cobra` and `afero`) updated the the lastest versions
- chore: `go-pretty` library removed as dependency
- fix: remove type annotation, trivially inferred from a string literal on `default.js.ts`

## 0.5.5

### Patch Changes

- fix: the lib template file for metadata (list) had a mistyped variable name causing runtime error

## 0.5.4

### Patch Changes

- fix: [#1](https://github.com/sveltinio/sveltin/issues/1) - it is now possible to use readable name when running command in interactive way
- fix: pages templates still used a wrong import for IWebPageMetadata from @sveltinio/seo
- fix: [#2](https://github.com/sveltinio/sveltin/issues/2) - EndpointOutput import removed from api template files
- refactor: string utilities methods

## 0.5.3

### Patch Changes

- fix: favicons with multiple formats and webmanifest file

## 0.5.2

### Patch Changes

- feature: `--withExcludeFile` flag added, to specify a file containing the list of files to not be deleted from the FTP server

## 0.5.1

### Patch Changes

- fix: set `config.kit.prerender.default` to `true`

## 0.5.0

### Minor Changes

- fix: import type IWebPageMetadata for index.svelte
- breaking: sveltin `prepare` command renamed as `install`. With [#4182](https://github.com/sveltejs/kit/pull/4182), SvelteKit added support for a new CLI command called `sync` wrapped as `prepare` script in `package.json` file.
- feature: sveltin `prepare` command wraps the new SvelteKit `sync` command, [#4182](https://github.com/sveltejs/kit/pull/4182)
- feature: sveltin `deploy` command places the TAR archive within the `backups` folder at project root level.
- refactor: new underlying logging lib developed

## 0.4.0

### Minor Changes

- feature: `deploy` over FTP command added. [docs](https://docs.sveltin.io/cli/deploy)
- docs: overall code comments
- chore: code cleansing
- chore: better error handling

## 0.3.1

### Patch Changes

- fix: Routing in app built by adapter-static and next.265 [#3801](https://github.com/sveltejs/kit/pull/3801)
- fix: .env file templates with server port number

## 0.3.0

### Minor Changes

- fix: public page styles to have centered text for page title
- fix: theme name on the `index.svelte` not replaced correctly. it now uses the new _ThemeName_ variable on the template data structure.
- chore: **next steps** texts added to some commands
- refactor: go templates execution functions
- refactor: logger and pretty printer
- breaking: `--update` flag on `prepare` cmd removed. `prepare` now install the dependencies from `package.json` file only
- breaking: npmClient handling. It is now project specific and make uses of the packageManager key from `package.json`.
- feature: `update` cmd added to update all the dependecies from `package.json` file
- feature: `--port` flag added to `new` cmd to set the port to start the server on
- feature: Bootstrap 5 support implemented
- feature: SCSS support added

## 0.2.14

### Patch Changes

- chore: dependency updated
- feature: bulma support implemented
- refactor: tailwind css and vanilla css themes updated
- refactor: each CSS supported lib has its own `__layout.svelte` file

## 0.2.13

### Patch Changes

- chore: This version adheres to the changes introduced by `@sveltejs/kit 1.0.0-next.257` where the target option has been removed [#3674](https://github.com/sveltejs/kit/pull/3674)

## 0.2.12

### Patch Changes

- feature: json-ld breadcrumbs added to the public and resources pages

## 0.2.11

### Patch Changes

- fix: paths to favicon, prism theme and script file

## 0.2.10

### Patch Changes

- fix: `generate menu` command used `js` instead of `ts` as file extension causing errors on loading
- fix: with tailwindcss used as css lib, the typography plugin's prose class was not rendered correctly
- fix: postcss and its config file provided for tailwindcss only
- fix: with vanillacss a github markdown theme added as default to render markdown content

## 0.2.9

### Patch Changes

- fix: SvelteKit 1.0.0-next.244 fixed [#3473](https://github.com/sveltejs/kit/issues/3473) and [#3521](https://github.com/sveltejs/kit/pull/3521). `clone()` on fetch response as workaround to avoid '_body used already_' error when building the project removed

## 0.2.8

### Patch Changes

- fix: string utility functions added to get valid page names and contents names
- fix: variable names fixed on page templates

  **Full Changelog**: https://github.com/sveltinio/sveltin/compare/v0.2.7...v0.2.8

## 0.2.7

### Patch Changes

- fix: image path on seo components
- fix: seo components added to pages
