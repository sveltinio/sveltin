<h1 align="center">
    <img src="resources/sveltin-logo.png" width="224px" alt="sveltin logo"/>
</h1>
<h2 align="center">
The Smartest Way to Create SvelteKit powered static websites.
</h2>
<p align="center">
<a href="https://sveltin.io/" target="_blank">
        Homepage
    </a> |
    <a href="https://docs.sveltin.io/release-notes" target="_blank">
        Release Notes
    </a> |
    <a href="https://docs.sveltin.io/quick-start" target="_blank">
        Quick Start
    </a> |
    <a href="https://docs.sveltin.io/" target="_blank">
        Documentation
    </a> |
    <a href="https://github.com/sveltinio/sveltin/blob/main/CONTRIBUTING.md" target="_blank">
        Contributing
    </a>
</p>
<p align="center">
    <a href="https://github.com/sveltinio/sveltin/blob/main/LICENSE" target="_blank">
        <img src="https://img.shields.io/badge/license-apache_2.0-blue?style=flat-square&logo=none" alt="license" />
    </a>
    &nbsp;
     <a href="https://goreportcard.com/report/github.com/sveltinio/sveltin/" target="_blank">
        <img src="https://goreportcard.com/badge/github.com/sveltinio/sveltin" alt="go report card" />
    </a>
    &nbsp;
    <a href="https://pkg.go.dev/github.com/sveltinio/sveltin/" target="_blank">
        <img src="https://pkg.go.dev/badge/github.com/sveltinio/sveltin/.svg" alt="go reference" />
    </a>
    &nbsp;
    <a href="https://github.com/sveltinio/sveltin/releases" target="_blank">
        <img src="https://img.shields.io/badge/version-v0.12.0-success?style=flat-square&logo=none" alt="sveltin cli version" />
    </a>
    &nbsp;
    <a href="https://github.com/sveltinio/sveltin/actions/workflows/release.yml" target="_blank">
        <img src="https://github.com/sveltinio/sveltin/actions/workflows/release.yml/badge.svg" alt="CI" />
    </a>
    &nbsp;
    <a href="https://madewithsvelte.com/p/sveltin/shield-link" target="_blank">
        <img src="https://madewithsvelte.com/storage/repo-shields/3613-shield.svg" alt="made with svelte shield" />
    </a>
</p>

# Sveltin

Sveltin is a CLI (Command Line Interface) created to boost the developers productivity working on <strong>SvelteKit v1.x.x powered static websites</strong>.

## :information_source: SvelteKit versions

> Latest tested SvelteKit version is [1.12.0](https://github.com/sveltejs/kit/releases/tag/%40sveltejs%2Fkit%401.12.0). SvelteKit reached v1.x.x then no more breakings are expected until a new major release. Although we decided to stay sticked to the tested version, you should be able to upgrade SvelteKit to the upcoming minor versions without disruptions.

## :mega: Overview

Sveltin is a simple, quick and powerful CLI to:

- Scaffold SvelteKit powered websites
- Generate resources, libs and endpoints
- Add content to the resources
- Generate menu structure, sitemap and rss
- Make your site SEO Ready (Metadata, Json-LD, OpenGraph) in a easy way

Sveltin provides:

- Out-of-the-box support for Bootstrap, Bulma, Sass/SCSS, Tailwind CSS, UnoCSS and vanilla CSS
- Ready to use Svelte [components](https://github.com/sveltinio/sveltin-components-library)

## :rocket: Quick Start

With few commands Sveltin flex the muscles 💪

> **NOTE**: each command can be executed in interactive way so do not need to pass arguments and flags to it.

```bash
# Create a project with TailwindCSS support
sveltin init myBlog --css tailwindcss

# Move to the project folder
cd myBlog

# Install all the dependencies
sveltin install

# Create a public page and compose it with Svelte
# (http://localhost:5173/contact)
sveltin new page contact --svelte

# Create a public page and compose it with Markdown
# (http://localhost:5173/about)
sveltin new page about --markdown

# Create a 'posts' resource
sveltin new resource posts

# Add new content to the posts resource
# (http://localhost:5173/posts/getting-started)
sveltin add content getting-started --to posts

# Add a 'category' metadata
# (http://localhost:5173/posts/category)
sveltin add metadata category --to posts --as single

# Run the server
sveltin server
```

## :book: Documentation

Please see the [documentation](https://docs.sveltin.io) for more information about Sveltin.

## :computer: Installation

### :wrench: Prerequisites

- Git
- Node (v16.19 LTS or higher is required)

### OSX and Linux via Homebrew

Homebrew will also install Git and Node.

```bash
# Tap a new formula:
brew tap sveltinio/sveltin

# Install:
brew install sveltin
```

### Windows via Scoop

```bash
# Tap a new bucket:
scoop bucket add sveltinio https://github.com/sveltinio/scoop-sveltin.git

# Install:
scoop install sveltinio/sveltin
```

### Go Install

Installation is done by using the `go install` command. In this case, ensure to have **Go (v1.18 or higher) installed** on your machine:

```bash
go install github.com/sveltinio/sveltin@latest
```

### Manually

You can download the pre-compiled binary for you specific OS from the [releases page](https://github.com/sveltinio/sveltin/releases). You will need to copy the and extract the binary, then move it to your local bin folder. Please, refer to the example below:

```bash
curl https://github.com/sveltinio/sveltin/releases/download/${VERSION}/${PACKAGE_NAME} -o ${PACKAGE_NAME}
sudo tar -xvf ${PACKAGE_NAME} -C /usr/local/bin/
sudo chmod +x /usr/local/bin/sveltin
```

## :gear: CLI Commands & Options

sveltin comes with a set of commands and subcommands to help dealing with your SvelteKit project.

Each command can be executed with inline arguments or interactivly.

```bash
$ sveltin -h

sveltin is the main command to work with SvelteKit powered static website.

Usage:
  sveltin [command]

Available Commands:
  add         Add content and metadata to a resource
  build       Builds a production version of your static website
  completion  Generate the autocompletion script for the specified shell
  deploy      Deploy the website over FTP
  generate    Generate static files (sitemap, rss, menu)
  help        Help about any command
  init        Initialize a new sveltin project
  install     Install the project dependencies
  migrate     Migrate existing sveltin project files to the latest sveltin version ones
  new         Create nee resources, pages and themes
  preview     Preview the production version locally
  server      Run the development server
  update      Update your project dependencies

Flags:
  -h, --help      help for sveltin
  -v, --version   version for sveltin

Use "sveltin [command] --help" for more information about a command.
```

### sveltin init

`sveltin init` is the main command to scaffold a project.

Alias: `create`

Read more [here][init].

### sveltin new

`sveltin new` is the main command to generate pages, resources (routes) and themes for your project.

Alias: `n`

<details>
    <summary>(Click to expand the list of available subcommands)</summary>

| Subcommand     | Aliases | Description                          |
| :------------- | :-----: | :----------------------------------- |
| [new-page]     |    p    | Command to create a new public page. |
| [new-resource] |    r    | Command to create a new resource.    |

</details>

Read more [here][new].

### sveltin add

`sveltin add` is the main command to add content and metadata to existing resources.

Alias: `a`

<details>
    <summary>(Click to expand the list of available subcommands)</summary>

| Subcommand     | Aliases | Description                                                            |
| :------------- | :-----: | :--------------------------------------------------------------------- |
| [add-content]  |    c    | Command to create a new content for existing resource.                 |
| [add-metadata] |    m    | Command to add a new metadata to your content as a Sveltekit resource. |

</details>

Read more [here][add].

### sveltin generate

`sveltin generate` is used to generate static files like sitemap, menu structure or rss feed file.

Alias: `g`

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand         | Description                    |
| :----------------- | :----------------------------- |
| [generate-menu]    | Generate the menu config file. |
| [generate-sitemap] | Generate a sitemap.xml.        |
| [generate-rss]     | Generate a rss.xml file.       |

</details>

Read more [here][generate].

### sveltin install

`sveltin install` is used to initialize the Sveltin project getting all depencencies from the `package.json` file.

Alias: `i`

Read more [here][install].

### sveltin update

`sveltin update` is used to update all depencencies from the `package.json` file.

Read more [here][update].

### sveltin migrate

`sveltin migrate` is used to migrate existing sveltin project files to the latest Sveltin version ones.

Read more [here][migrate].

### sveltin server

`sveltin server` is used to run the VITE server. It wraps svelte-kit defined commands to run the server.

Alias: `s`, `serve`, `run`, `dev`

Read more [here][server].

### sveltin build

`sveltin build` is used to build a production version of your static website. It wraps `sveltekit-build` command.

Alias: `b`

Read more [here][build].

### sveltin preview

`sveltin preview` is used to run a preview for the production version locally.

Read more [here][preview].

### sveltin deploy

`sveltin deploy` is used to deploy your website over FTP on your hosting platform.

Read more [here][deploy].

### sveltin completion

`sveltin completion` generates the autocompletion script for the specified shell (bash|zsh|fish|powershell).

Read more [here][completion].

## :bulb: Contributing

Contribution of any kind including documentation, themes, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests are more than welcome.

Read more [here][contributing].

## Dependencies

Sveltin leverages many great open source libraries:

| Name                                                    | Version   | License      |
| :------------------------------------------------------ | :-------: | :----------- |
| [bubble](https://github.com/charmbracelet/bubbles)      | `0.15.0`  | MIT          |
| [bubbletea](https://github.com/charmbracelet/bubbletea) | `0.23.2`  | MIT          |
| [lipgloss](https://github.com/charmbracelet/lipgloss)   | `0.6.0`   | MIT          |
| [validator](https://github.com/go-playground/validator) | `10.12.0` | MIT          |
| [slug](https://github.com/gosimple/slug)                | `1.13.1`  | MPL-2.0      |
| [ftp](https://github.com/jlaffaye/ftp)                  | `0.1.0`   | ISC          |
| [is](https://github.com/matryer/is)                     | `1.4.1`   | MIT          |
| [afero](https://github.com/spf13/afero)                 | `1.9.5`   | Apache-2.0   |
| [cobra](https://github.com/spf13/cobra)                 | `1.6.1`   | Apache-2.0   |
| [viper](https://github.com/spf13/viper)                 | `1.15.0`  | MIT          |
| [prompti](https://github.com/sveltinio/prompti)         | `0.2.2`   | MIT          |
| [gjson](https://github.com/tidwall/gjson)               | `1.14.4`  | MIT          |
| [sjson](https://github.com/tidwall/sjson)               | `1.2.5`   | MIT          |
| [text](https://golang.org/x/text)                       | `0.8.0`   | BSD-3-Clause |

## :free: License

Sveltin is free and open-source software licensed under the Apache 2.0 License.

[add]: https://docs.sveltin.io/cli/add/
[add-content]: https://docs.sveltin.io/cli/add-content/
[add-metadata]: https://docs.sveltin.io/cli/add-metadata/
[build]: https://docs.sveltin.io/cli/build/
[completion]: https://docs.sveltin.io/cli/completion/
[contributing]: CONTRIBUTING.md
[deploy]: https://docs.sveltin.io/cli/deploy/
[generate]: https://docs.sveltin.io/cli/generate/
[generate-menu]: https://docs.sveltin.io/cli/generate-menu/
[generate-rss]: https://docs.sveltin.io/cli/generate-rss/
[generate-sitemap]: https://docs.sveltin.io/cli/generate-sitemap/
[init]: https://docs.sveltin.io/cli/init/
[install]: https://docs.sveltin.io/cli/install/
[migrate]: https://docs.sveltin.io/cli/migrate/
[new]: https://docs.sveltin.io/cli/new/
[new-page]: https://docs.sveltin.io/cli/new-page/
[new-resource]: https://docs.sveltin.io/cli/new-resource/
[preview]: https://docs.sveltin.io/cli/preview/
[server]: https://docs.sveltin.io/cli/server/
[update]: https://docs.sveltin.io/cli/update/
