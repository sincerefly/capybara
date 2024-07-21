# Capybara 

<p>
    <picture>
      <img src="docs/logo/capybara.png" width="180"  alt="capybara"/>
    </picture>
</p>

![](https://github.com/sincerefly/capybara/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/sincerefly/capybara)](https://goreportcard.com/report/github.com/sincerefly/capybara)
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/sincerefly/capybara/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
[![codecov](https://codecov.io/gh/sincerefly/capybara/graph/badge.svg?token=D8RGT9H0TU)](https://codecov.io/gh/sincerefly/capybara)
<img src="https://img.shields.io/github/repo-size/sincerefly/capybara?style=flat-square&color=328657" alt="Repo Size">

A convenient tool for adding borders, watermarks to images, with possible future additions such as format conversion, image compression, and more.

**Node: The project is in its early development stage, so parameters and commands may be subject to change.**

### Usage

build binary from source

```bash
$ git clone https://github.com/sincerefly/capybara
$ cd capybara
$ go build
```

### Quick Start

add border with default parameters

```bash
$ capybara border
```

using 'simple' style, with 100 border size, process 'input' folder images, save to 'output' folder

```bash
$ capybara border -s simple -w 100 -i input -o output
```

border color 

```bash
# with html color names, case-insensitive
$ capybara border --color AliceBlue

# with hex
$ capybara border --color "#228B22"

# with rgb
$ capybara border --color "rgb(238, 130, 238)"

# with rgba
$ capybara border --color "rgba(238, 130, 238, 255)"
```

with 'text_bottom' style

```bash
$ capybara border text_bottom

$ capybara border text_bottom --color AliceBlue

$ capybara border text_bottom --container-height 360 --without-subtitle
```

Different styles support different parameters. Please refer to the documentation for details

### Style Example

```bash
$ capybara border simple
$ capybara border text_bottom
$ capybara border melon
```

<table>
  <tr>
    <td>style: simple</td>
    <td>style: text_bottom</td>
    <td>style: melon</td>
  </tr>
  <tr>
    <td><img src="docs/image/style-simple.webp" width=270></td>
    <td><img src="docs/image/style-text_bottom.webp" width=270></td>
    <td><img src="docs/image/style-melon.webp" width=270></td>
  </tr>
</table>

### Help 

You can use `subcommand --help` to view the default parameters.

```bash
$ ./capybara border text_bottom --help
Style: Footer text, with photo exif

Usage:
  border border text_bottom [flags]

Flags:
  -c, --color string           specify border color (default "white")
      --container-height int   bottom text container height (default 300)
  -h, --help                   help for text_bottom
  -i, --input string           specify input folder (default "input")
  -o, --output string          specify output folder (default "output")
  -w, --width int              specify border width (default 100)
      --without-subtitle       without subtitle

Global Flags:
      --debug            enables detailed logging for debugging.
      --no-parallelism   disables parallel processing, without goroutine.
```