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
        <img src="https://img.shields.io/badge/version-v0.2.3-success?style=flat-square&logo=none" alt="sveltin cli version" />
    </a>
    &nbsp;
    <a href="https://github.com/sveltinio/sveltin/blob/main/LICENSE" target="_blank">
        <img src="https://img.shields.io/badge/license-apache_2.0-blue?style=flat-square&logo=none" alt="license" />
    </a>
</p>

# Sveltin

The Smartest Way to Build SvelteKit static websites.

## :mega: Features

- Simple, quick and powerful CLI
- Scaffold project
- Generate resources, libs and endpoints
- Generate content
- Generate menu/sitemap/rss
- Ready to use Svelte components
- Out-of-the-box support for TailwindCSS, Bulma and Bootstrap
- SEO Ready (Metadata, Json-LD, OpenGraph)

## :rocket: Quick Start

```bash
# Create myBlog project
sveltin new myBlog

# Or sreate myBlog project with TailwWindCSS as CSS Lib
sveltin new myBlog --css tailwindcss

# Move to the project folder
cd myBlog 

# Install all the dependencies
sveltin prepare

# Create new public page as Svelte component
# (http://localhost:3000/contact)
sveltin new page contact --type svelte

# Create new public page as Markdown file
# (http://localhost:3000/about)
sveltin new page about -t markdown

# Create new 'posts' resource
sveltin new resource posts

# Add new content to the posts resource
# (http://localhost:3000/posts/getting-started)
sveltin new content posts/getting-started

# Add new 'category' metadata 
# (http://localhost:3000/posts/category)
sveltin new metadata category --resource posts --type single

# Add new 'tags' metadata
# (http://localhost:3000/posts/tags)
sveltin new metadata tags -r posts --t lists

# Run the server
sveltin server
```

## :book: Documentation

Please see the [documentation](https://docs.sveltin.io) for information about getting started with Sveltin.

## :computer: Installation

### :wrench: Prerequisites

- Git
- Node (v16.13.1 or higher is required)

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

### From the source

Installation is done by using the `go install` command. In this case, ensure to have **Go (v1.17 or higher) installed** on your machine:

```bash
go install github.com/sveltinio/sveltin@latest
```

## :gear: CLI Commands & Options

```bash
$ sveltin -h

sveltin is the main command to work with SvelteKit powered static website.

Usage:
  sveltin [command]

Available Commands:
  build        Builds a production version of your static website
  generate     Command to generate static files like sitemap, rss etc
  help         Help about any command
  new          Command to create projects, resources, contents, pages and metadata
  prepare      Get all the dependencies
  preview      Preview the production version locally
  server       Run the server

Flags:
  -h, --help      help for sveltin
  -v, --version   version for sveltin

Use "sveltin [command] --help" for more information about a command.
```

sveltin comes with a set of commands and subcommands to help dealing with your SvelteKit project.

Each command can be executed with inline arguments or interactivly.

### sveltin new

`sveltin new` is the main command to generate both the project and the artifacts for your website.

<details>
    <summary>(Click to expand the list of avilable subcommands)</summary>

| Subcommand | Alias | Description                                                   |
| :--------- | :---: | :------------------------------------------------------------ |
| [resource] |       | Create new resources.                                         |
| [content]  |       | Create a new content for existing resource.                   |
| [metadata] |       | Add a new metadata from your content as a Sveltekit resource. |
| [page]     |       | Create a new public page.                                     |

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

`sveltin prepare` is used to initialize the Sveltin project getting all depencencies from the `package.json` file.

Read more [here][prepare].

### sveltin server

`sveltin server` is used to run the VITE server. It wraps svelte-kit defined commands to run the server.

Read more [here][server].

### sveltin build

`sveltin build` is used to build a production version of your static website. It wraps `sveltekit-build` command.

Read more [here][build].

### sveltin preview

`sveltin preview` is used to run a preview for the production version locally.

Read more [here][preview].

## :bulb: Contributing

Contribution of any kind including documentation, themes, tutorials, blog posts, bug reports, issues, feature requests, feature implementations, pull requests are more than welcome.

Read more [here][contributing].

## :free: License

Sveltin is free and open-source software licensed under the Apache 2.0 License.

[new]: https://docs.sveltin.io/cli/new
[resource]: https://docs.sveltin.io/cli/new-resource
[content]: https://docs.sveltin.io/cli/new-content
[metadata]: https://docs.sveltin.io/cli/new-metadata
[page]: https://docs.sveltin.io/cli/new-page
[generate]: https://docs.sveltin.io/cli/generate
[menu]: https://docs.sveltin.io/cli/generate-menu
[sitemap]: https://docs.sveltin.io/cli/generate-sitemap
[rss]: https://docs.sveltin.io/cli/generate-rss
[server]: https://docs.sveltin.io/cli/server
[prepare]: https://docs.sveltin.io/cli/prepare
[build]: https://docs.sveltin.io/cli/build
[preview]: https://docs.sveltin.io/cli/preview
[contributing]: CONTRIBUTING.md
