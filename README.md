
[![version](https://img.shields.io/github/release/cybercongress/cyber-wiki-index.svg?style=flat-square)](https://github.com/cybercongress/cyber-wiki-index/releases/latest)
[![CircleCI](https://img.shields.io/circleci/project/github/cybercongress/cyber-wiki-index.svg?style=flat-square)](https://circleci.com/gh/cybercongress/cyber-wiki-index/tree/master)
[![license](https://img.shields.io/badge/License-Cyber-brightgreen.svg?style=flat-square)](https://github.com/cybercongress/cyber-wiki-index/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/cybercongress/cyber-wiki-index)](https://github.com/cybercongress/cyber-wiki-index)
[![contributors](https://img.shields.io/github/contributors/cybercongress/cyber-wiki-index.svg?style=flat-square)](https://github.com/cybercongress/cyber-wiki-index/graphs/contributors)
[![discuss](https://img.shields.io/badge/Join%20Us%20On-Telegram-2599D2.svg?style=flat-square)](https://t.me/fuckgoogle)
[![contribute](https://img.shields.io/badge/contributions-welcome-orange.svg?style=flat-square)](https://github.com/cybercongress/cyber-wiki-index/blob/master/CONTRIBUTING.md)

[://cyber](https://github.com/cybercongress/cyberd) wiki index
==================

  - [Installation](#installation)
  - [Usage](#usage)
  - [Issues](#issues)
  - [Contributing](#contributing)
  - [Changelog](#changelog)

## Installation

Note: Requires Go 1.12+

```
git clone https://github.com/cybercongress/crawler
cd crawler
go build -o crawler
```

## Preparation

1. IPFS daemon should be launched
2. Download enwiki-latest-all-titles to crawler root dir: 

``` 
ipfs get QmddV5QP87BZGiSUCf9x9hsqM73b83rsPC6AYMNqkjKMGx -o enwiki-latest-all-titles
```

3. Add account to cyberdcli: 

```
docker exec -ti cyberd cyberdcli keys add <name> --recover
```

## Usage

### Submit links
Basically, there are two main functions provided by `crawler` tool. 
The first one is to parse wiki titles and submit links between keywords and wiki pages. 
```
./crawler submit-links-to-cyber ./enwiki-latest-all-titles --home=<path-to-cyberdcli> --address=<account> --passphrase=<passphrase> --chunk=100
```

> Note: Uses only local cyberd node.

> Note: Submit links do not add duras to IPFS.

> Note: Chunk - how many links messages added to one tx

> Note: There is --help command, for example 

```
./crawler submit-links-to-cyber --help
```

Here, **enwiki-latest-all-titles** is titles file obtained from 
 [official Wiki dumps](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-all-titles-in-ns0.gz).  

### Uploading duras to IPFS 
Also, `crawler` has separate command `upload-duras-to-ipfs` to upload files to local IPFS node. 
All DURAs are collected under single root unixfs directory.
```
./crawler upload-duras-to-ipfs enwiki-latest-all-titles
```

## Issues

If you have any problems with or questions about search, please contact us through a
 [GitHub issue](https://github.com/cybercongress/crawler/issues).

## Contributing

You are invited to contribute new features, fixes, or updates, large or small; We are always thrilled to receive pull
 requests, and do our best to process them as fast as We can. You can find detailed information in our
 [contribution guide](./docs/contributing/contributing.md).


## Changelog

Stay tuned with our [Changelog](./CHANGELOG.md).

<div align="center">
  <sub>Built by
  <a href="https://twitter.com/cyber_devs">cyber•Congress</a> and
  <a href="https://github.com/cybercongress/crawler/graphs/contributors">contributors</a>
</div>
