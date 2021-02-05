# Fetch.ai fetchd Fork

This repository is the fork from the original [wasmd](https://github.com/CosmWasm/wasmd). It contains Fetch.ai specific updates required for the test networks and future mainnet. For this reason it is versioned independantly. Please refer to the [releases](https://github.com/fetchai/fetchd/releases) section for the compatiblity with upstream versions.

**Note**: Requires [Go 1.14+](https://golang.org/dl/)

## Supported Systems

The supported systems are limited by the dlls created in [`go-cosmwasm`](https://github.com/CosmWasm/go-cosmwasm). In particular, **we only support MacOS and Linux**.
For linux, the default is to build for glibc, and we cross-compile with CentOS 7 to provide
backwards compatibility for `glibc 2.12+`. This includes all known supported distributions
using glibc (CentOS 7 uses 2.12, obsolete Debian Jessy uses 2.19).

As of `0.5.x` we support `muslc` Linux systems, in particular **Alpine linux**,
which is popular in docker distributions. Note that we do **not** store the
static `muslc` build in the repo, so you must compile this yourself, and pass `-tags muslc`.
Please look at the [`Dockerfile`](./Dockerfile) for an example of how we build a static Go
binary for `muslc`. (Or just use this Dockerfile for your production setup).

## Quick Start

```
# Ubuntu/Debian based linux distro dependencies
apt install swig libboost-all-dev libgmp-dev

cd ~/Downloads
wget https://github.com/herumi/mcl/archive/v1.05.tar.gz
tar xvf v1.05.tar.gz
cd mcl-1.05
make install
ldconfig

# OS X dependencies
brew install swig boost gmp

cd ~/Downloads
wget https://github.com/herumi/mcl/archive/v1.05.tar.gz
tar xvf v1.05.tar.gz
cd mcl-1.05
make install

```
After installing the reuired dependencies, install golang and execute the following commands:
```
make install
make test
```
if you are using a linux without X or headless linux, look at [this article](https://ahelpme.com/linux/dbusexception-could-not-get-owner-of-name-org-freedesktop-secrets-no-such-name) or [#31](https://github.com/fetchai/fetchd/issues/31#issuecomment-577058321).

## Run a simple local test network

The easiest way to get started with a simple network is to run the [docker-compose](https://docs.docker.com/compose/). The details of this can be found [here](https://github.com/fetchai/fetchd/blob/master/docker-compose.yml). By default it will launch a small 5 validator node network.

## Resources

1. [Website](https://fetch.ai/)
2. [Discord Server](https://discord.gg/UDzpBFa)
3. [Blog](https://fetch.ai/blog)
4. [Community Website](https://community.fetch.ai/)
5. [Community Telegram Group](https://t.me/fetch_ai)
