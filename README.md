<h1 align="center">
    <img src="resources/sveltin-logo.png" width="224px" alt="sveltin logo"/>
</h1>
<h2 align="center">
The Smartest Way to Build SvelteKit powered static websites.
</h2>
<p align="center">
<a href="https://sveltin.io/" target="_blank">
        Website
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
        <img src="https://img.shields.io/badge/version-v0.8.11-success?style=flat-square&logo=none" alt="sveltin cli version" />
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

Sveltin is a CLI (Command Line Interface) created to boost the developers productivity working on <strong>SvelteKit powered static websites</strong>.

## :warning: Project Status

> Sveltin is under active development and some changes are expected before we hit version 1.0. At the same time, we will do our best to follow the progress toward SvelteKit v1.0 (Latest SvelteKit tested version is **1.0.0-next-403**). If you are interesting on it please, give it a try and let it evolves, see the **Contributing** section. If you get stuck, reach out for help in the [discussions tab](https://github.com/sveltinio/sveltin/discussions) or open an [issue](https://github.com/sveltinio/sveltin/issues).

## :mega: Overview

Sveltin is a simple, quick and powerful CLI to:

- Scaffold SvelteKit powered websites
- Generate resources, libs and endpoints
- Add content to the resources
- Generate menu structure, sitemap and rss
- Make your site SEO Ready (Metadata, Json-LD, OpenGraph) in a easy way

Sveltin provides:

- Out-of-the-box support for vanilla CSS, Sass/SCSS, Tailwind CSS, Bulma and Bootstrap
- Ready to use Svelte [components](https://github.com/sveltinio/sveltin-components-library)

## :rocket: Quick Start

With few commands Sveltin flex the muscles 💪

```bash
# Create a project with TailwindCSS support
sveltin new myBlog --css tailwindcss

# Move to the project folder
cd myBlog

# Install all the dependencies
sveltin install

# Create a public page as Svelte component
# (http://localhost:5173//contact)
sveltin new page contact --type svelte

# Create a 'posts' resource
sveltin new resource posts

# Add new content to the posts resource
# (http://localhost:5173//posts/getting-started)
sveltin new content posts/getting-started

# Add a 'category' metadata
# (http://localhost:5173//posts/category)
sveltin new metadata category --resource posts --type single

# Run the server
sveltin server
```

## :book: Documentation

Please see the [documentation](https://docs.sveltin.io) for more information about Sveltin.

## :computer: Installation

### :wrench: Prerequisites

- Git
- Node (v16.9.0 or higher is required)

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

Installation is done by using the `go install` command. In this case, ensure to have **Go (v1.17 or higher) installed** on your machine:

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

```bash
$ sveltin -h

sveltin is the main command to work with SvelteKit powered static website.

Usage:
  sveltin [command]

Available Commands:
  build       Builds a production version of your static website
  deploy      Command to deploy your website over FTP
  generate    Command to generate static files (sitemap, rss, menu)
  help        Help about any command
  install     Get all the dependencies from the `package.json` file
  new         Command to create projects, resources, contents, pages, metadata and themes.
  preview     Preview the production version locally
  server      Run the server
  update      Update the dependencies from the `package.json` file

Flags:
  -h, --help      help for sveltin
  -v, --version   version for sveltin

Use "sveltin [command] --help" for more information about a command.
```

sveltin comes with a set of commands and subcommands to help dealing with your SvelteKit project.

Each command can be executed with inline arguments or interactivly.

### sveltin new

`sveltin new` is the main command to generate projects, artifacts and themes for your website.

Alias: `create`

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand |   Aliases    | Description                                                            |
| :--------- | :----------: | :--------------------------------------------------------------------- |
| [content]  |      c       | Command to create a new content for existing resource.                 |
| [metadata] | m, groupedBy | Command to add a new metadata to your content as a Sveltekit resource. |
| [page]     |      p       | Command to create a new public page.                                   |
| [resource] |      r       | Command to create a new resource.                                      |
| [theme]    |      t       | Command to create a new theme.                                         |

</details>

Read more [here][new].

### sveltin generate

`sveltin generate` is used to generate static files like sitemap, menu structure or rss feed file.

Alias: `g`, `gen`

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand | Description                                             |
| :--------- | :------------------------------------------------------ |
| [menu]     | Generate the menu config file for your Sveltin project. |
| [sitemap]  | Generate a sitemap.xml file for your Sveltin project.   |
| [rss]      | Generate a rss.xml file for your Sveltin project.       |

</details>

Read more [here][generate].

### sveltin install

`sveltin install` is used to initialize the Sveltin project getting all depencencies from the `package.json` file.

Alias: `i`, `init`

Read more [here][install].

### sveltin update

`sveltin update` is used to update all depencencies from the `package.json` file.

Alias: `u`

Read more [here][update].

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

## :bulb: Contributing

Contribution of any kind including documentation, themes, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests are more than welcome.

Read more [here][contributing].

## :free: License

Sveltin is free and open-source software licensed under the Apache 2.0 License.

[new]: https://docs.sveltin.io/cli/new/
[resource]: https://docs.sveltin.io/cli/new-resource/
[content]: https://docs.sveltin.io/cli/new-content/
[metadata]: https://docs.sveltin.io/cli/new-metadata/
[page]: https://docs.sveltin.io/cli/new-page/
[theme]: https://docs.sveltin.io/cli/new-theme/
[generate]: https://docs.sveltin.io/cli/generate/
[menu]: https://docs.sveltin.io/cli/generate-menu/
[sitemap]: https://docs.sveltin.io/cli/generate-sitemap/
[rss]: https://docs.sveltin.io/cli/generate-rss/
[server]: https://docs.sveltin.io/cli/server/
[install]: https://docs.sveltin.io/cli/install/
[update]: https://docs.sveltin.io/cli/update/
[build]: https://docs.sveltin.io/cli/build/
[preview]: https://docs.sveltin.io/cli/preview/
[deploy]: https://docs.sveltin.io/cli/deploy/
[contributing]: CONTRIBUTING.md
