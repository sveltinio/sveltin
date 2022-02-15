# sveltin

## 0.3.0

### Minor Changes

- fix: public page styles to have centered text for page title
- fix: theme name on the `index.svelte` not replaced correctly. it now uses the new _ThemeName_ variable on the template data structure.
- chore: **next steps** texts added to some commands
- refactor: go templates execution functions
- refactor: logger and pretty printer
- breaking: `--update` flag on `prepare` cmd removed. `prepare` now install the dependencies from `package.json` file only
- breaking: npmClient handling. It is now project specific and make uses of the packageManager key from `package.json`.
- feat: `update` cmd added to update all the dependecies from `package.json` file
- feat: `--port` flag added to `new` cmd to set the port to start the server on
- feat: Bootstrap 5 support implemented
- feat: SCSS support added

## 0.2.14

### Patch Changes

- chore: dependency updated
- feat: bulma support implemented
- refactor: tailwind css and vanilla css themes updated
- refactor: each CSS supported lib has its own `__layout.svelte` file

## 0.2.13

### Patch Changes

- chore: This version adheres to the changes introduced by `@sveltejs/kit 1.0.0-next.257` where the target option has been removed [#3674](https://github.com/sveltejs/kit/pull/3674)

## 0.2.12

### Patch Changes

- feat: json-ld breadcrumbs added to the public and resources pages

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

- SvelteKit 1.0.0-next.244 fixed [#3473](https://github.com/sveltejs/kit/issues/3473) and [#3521](https://github.com/sveltejs/kit/pull/3521). `clone()` on fetch response as workaround to avoid '_body used already_' error when building the project removed

## 0.2.8

### Patch Changes

- string utility functions added to get valid page names and contents names
- variable names fixed on page templates

  **Full Changelog**: https://github.com/sveltinio/sveltin/compare/v0.2.7...v0.2.8

## 0.2.7

### Patch Changes

- fix: image path on seo components
- fix: seo components added to pages
