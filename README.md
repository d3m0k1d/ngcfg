# ngcfg

A universal CLI utility for configuring [Nginx](https://nginx.org/)

# Table of contents
1. [Overview](#overview)
2. [Requirements](#requirements)
3. [Installation](#installation)
4. [Usage](#usage)
5. [License](#license)

## Overview
A config generator for [Nginx](https://nginx.org/) writen in golang and using [Go Template](https://pkg.go.dev/text/template), [Cobra](https://github.com/spf13/cobra) and more libs.
### Features
  - Generate simple server block from yaml config files to nginx format

### Roadmap
  - [x] Generate simple server block from yaml config files to nginx format
  - [ ] Add support json config files
  - [ ] Configurating Nginx with CLI flags and args
  - [ ] SSL sertificate generation support with certbot

## Requirements
 - Go 1.24.6

## Installation

```bash
git clone https://github.com/d3m0k1d/ngcfg.git
cd ngcfg
go mod tidy
go build -o ngcfg
```
## License
This project is licensed under the MIT License. 
