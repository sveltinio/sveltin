<h1 align="center">
    <img src="resources/sveltin-logo.png" width="224px" alt="sveltin logo"/>

</h1>
<p align="center">
    Sveltin is a CLI (Command Line Interface) to get easily started with <strong>SvelteKit powered static websites.</strong>
    <br />
    If you want to build a super fast static website based on the new SvelteKit framework, this is the project you need.
</p>
<p align="center">
    <a href="https://github.com/sveltinio/sveltin/actions/workflows/release.yml" target="_blank">
        <img src="https://github.com/sveltinio/sveltin/actions/workflows/release.yml/badge.svg" alt="CI" />
    </a>
    &nbsp;
    <a href="https://github.com/sveltinio/sveltin/releases" target="_blank">
        <img src="https://img.shields.io/badge/version-v0.1.0-success?style=flat-square&logo=none" alt="sveltin cli version" />
    </a>
    &nbsp;
    <a href="https://github.com/sveltinio/sveltin/blob/main/LICENSE" target="_blank">
        <img src="https://img.shields.io/badge/license-apache_2.0-blue?style=flat-square&logo=none" alt="license" />
    </a>
</p>

> Sveltin is in early development. Please, give it a try and let it evolve. If you are interesting on it please, see the **Contributing** section.

# Sveltin

## Documentation

You can read more on sveltin [here](https://docs.sveltin.io).

## Installation

Sveltin works on OSX, Windows, Linux, and others. Different installation options can be used

### OSX and Linux via Homebrew

```bash
# Tap a new formula:
brew tap sveltinio/sveltin

# Install:
brew install sveltin
```

### Windows via Scoop

```bash
# Tap a new bucket:
scoop bucket add sveltinio https://github.com/sveltinio/sveltin.git

# Install:
scoop install sveltinio/sveltin
```

### From source

#### Prerequisites

- Git
- Node (v16.13.1 or higher is required)

Installation is done by using the `go install` command and rename installed binary in $GOPATH/bin. In this case, ensure to have **Go (v1.17 or higher) installed** on your machine:

```bash
go install github.com/sveltinio/sveltin@latest
```

## Directory Structure

A Sveltin app will have a number of generated files and folders that make up the structure of the project.

```bash
.
├── config
├── content
├── src
│   ├── components
│   ├── lib
│   ├── routes
│   │   ├── api
│   │   │   └── v1
├── themes
└── static
```

Please, refer to [Sveltin application structure](https://docs.sveltin.io/application-structure) for a complete description.

## Tutorials

- [Create a new Sveltin based project](https://docs.sveltin.io//tutorials/your-first-project)
- [Personal Blog](https://docs.sveltin.io//tutorials/personal-blog)

## Commands & Options

> By default, sveltin uses `pnpm` as package manager. You can choose `npm` or `yarn` simply passing it via the `–-package-manager` flag (in short `-p`) to the sveltin commands supporting it.

```bash
$ sveltin -h

sveltin is the main command to work with SvelteKit powered static website.

Usage:
  sveltin [command]

Available Commands:
  build       Builds a production version of your static website
  completion  generate the autocompletion script for the specified shell
  generate    Command to generate static files like sitemap, rss etc
  help        Help about any command
  new         Command to create projects, resources, contents, pages and metadata
  prepare     Get all the dependencies
  preview     Preview the production version locally
  serve       Run the server

Flags:
  -h, --help      help for sveltin
  -v, --version   version for sveltin

Use "sveltin [command] --help" for more information about a command.
```

sveltin comes with a set of commands and subcommands to help dealing with your SvelteKit project.

Each command can be executed with inline arguments or interactivly.

```bash
# Create myWebsite project
sveltin new myWebsite

# or let sveltin prompt your inputs
sveltin new
```

### sveltin new

`sveltin new` is the main command to generate artifacts for your website.

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand | Alias | Description                                                   |
| :--------- | :---: | :------------------------------------------------------------ |
| [project]  | site  | Create a new sveltin based project.                           |
| [resource] |       | Create new resources.                                         |
| [content]  |       | Create a new content for existing resource.                   |
| [metadata] |       | Add a new metadata from your content as a Sveltekit resource. |
| [page]     |       | Create a new public page.                                     |
| [theme]    |       | Create a new theme.                                           |

</details>

Read more [here][new].

### sveltin generate

`sveltin generate` is used to generate static files like sitemap, menu structure or rss feed file.

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand | Alias | Description                                             |
| :--------- | :---: | :------------------------------------------------------ |
| [menu]     |       | Generate the menu config file for your Sveltin project. |
| [sitemap]  |       | Generate a sitemap.xml file for your Sveltin project.   |
| [rss]      |       | Generate a rss.xml file for your Sveltin project.       |

</details>

Read more [here][generate].

### sveltin prepare

`sveltin prepare` is used to initialize the Sveltin project getting all depencencies from the `package.json` file. It acts as wrapper for `pnpm install`.

Read more [here][prepare].

### sveltin serve

`sveltin serve` is used to run the VITE server. It wraps svelte-kit defined commands to run the server.

Read more [here][serve].

### sveltin build

`sveltin build` is used to build a production version of your static website. It wraps `sveltekit-build` command.

Read more [here][build].

### sveltin preview

`sveltin preview` is used to run a preview for the production version locally.

Read more [here][preview].

## Contributing

I would love if you like the idea behind sveltin and you decide to contribute to it. A **community driven project**, that is the end-goal.

Contribution of any kind including documentation, themes, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests are more than welcome.

Read more [here][contributing].

## License

Sveltin is free and open-source software licensed under the Apache 2.0 License.

[new]: https://docs.sveltin.io/cli/new
[resource]: https://docs.sveltin.io/cli/new-resource
[content]: https://docs.sveltin.io/cli/new-content
[metadata]: https://docs.sveltin.io/cli/new-metadata
[page]: https://docs.sveltin.io/cli/new-page
[theme]: https://docs.sveltin.io/cli/new-theme
[generate]: https://docs.sveltin.io/cli/generate
[menu]: https://docs.sveltin.io/cli/generate-menu
[sitemap]: https://docs.sveltin.io/cli/generate-sitemap
[rss]: https://docs.sveltin.io/cli/generate-rss
[serve]: https://docs.sveltin.io/cli/server
[prepare]: https://docs.sveltin.io/cli/prepare
[build]: https://docs.sveltin.io/cli/build
[preview]: https://docs.sveltin.io/cli/preview
[contributing]: CONTRIBUTING.md
