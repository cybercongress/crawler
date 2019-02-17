
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

## Usage

### Submit links
Basically, there are two main functions provided by `cyber-wiki` tool. 
The first one is to parse wiki titles and submit links between keywords and wiki pages. 
```
submit-links-to-cyber /home/user/enwiki-latest-all-titles --address=cli_acc_y0y
```

Here, **enwiki-latest-all-titles** is titles file obtained from 
 [official Wiki dumps](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-all-titles-in-ns0.gz).  

> Note: Uses only local cyber node.

> Note: Submit links do not add duras to ipfs.

### Uploading duras to IPFS 
Also, `cyber-wiki` has separate command `upload-duras-to-ipfs` to upload files to local IPFS. 
All duras are collected under single root unixfs directory.
```
upload-duras-to-ipfs /home/user/enwiki-latest-all-titles
```
> Note: We already upload duras. 
They can be downloaded and pined on yor local node by **Qmdwsryu8HskLzBspzPwJbL8UZ1ZPZF8VemtW1ja1GMXGp** hash.

## Issues

If you have any problems with or questions about search, please contact us through a
 [GitHub issue](https://github.com/cybercongress/cyber-wiki-index/issues).

## Contributing

You are invited to contribute new features, fixes, or updates, large or small; We are always thrilled to receive pull
 requests, and do our best to process them as fast as We can. You can find detailed information in our
 [contribution guide](./docs/contributing/contributing.md).


## Changelog

Stay tuned with our [Changelog](./CHANGELOG.md).

<div align="center">
  <sub>Built by
  <a href="https://twitter.com/cyber_devs">cyber•Congress</a> and
  <a href="https://github.com/cybercongress/cyber-wiki-index/graphs/contributors">contributors</a>
</div>
