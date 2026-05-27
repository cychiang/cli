![CI](https://github.com/crossplane/cli/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/crossplane/cli)](https://goreportcard.com/report/github.com/crossplane/cli)

![Crossplane CLI](banner.png)

The Crossplane CLI is a command-line tool for working with [Crossplane], the
cloud-native framework for platform engineering. It provides tools for building
platforms on top of Crossplane and working with Crossplane clusters.

Crossplane is a [Cloud Native Computing Foundation][cncf] project.

## Installation

The Crossplane CLI is a single binary, which we build for macOS, Linux, and
Windows. You can download the latest version using the install script:

```shell
curl -sfL "https://cli.crossplane.io/install.sh" | sh
```

The script detects your operating system and CPU architecture and downloads the
appropriate binary to the current directory. Note that it does not attempt to
place the binary in your shell's `$PATH`, so you may want to move it.

To install a different version of the CLI, set `XP_VERSION` when running the
install script:

```shell
curl -sfL "https://cli.crossplane.io/install.sh" | XP_VERSION=v2.3.1 sh
```

To install the latest build from our main branch, set `XP_CHANNEL=master`:

```shell
curl -sfL "https://cli.crossplane.io/install.sh" | XP_CHANNEL=master sh
```

## Reference Documentation

Command reference documentation for the CLI can be found on
[docs.crossplane.io](https://docs.crossplane.io/latest/cli/command-reference/).

## License

Crossplane is under the Apache 2.0 license.

<!-- Named links -->

[Crossplane]: https://crossplane.io
[cncf]: https://www.cncf.io/
