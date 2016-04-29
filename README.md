[![Circle CI](https://circleci.com/gh/eris-ltd/eris-cm/tree/master.svg?style=svg)](https://circleci.com/gh/eris-ltd/eris-cm/tree/master)

[![GoDoc](https://godoc.org/github.com/eris-cm?status.png)](https://godoc.org/github.com/eris-ltd/eris-cm)

# Eris Chain Manager

```
The Eris Chain Manager is a utility for performing complex operations on eris chains
```

`eris:chain_manager` is a set of high-level tooling for working with `eris chains`. It is similar in nature, design, and level as the `eris:package_manager` which is built to handle smart contract packages and other packages necessary for building blockchain backed applications.

## Install

Usually this repository should be run in a docker container via the [eris-cli](https://docs.erisindustries.com/tutorials/).

Should you want/desire/need to install this repository natively on your host make sure you have go installed and then:

1. [Install go](https://golang.org/doc/install)
2. Ensure you have gmp installed (sudo apt-get install libgmp3-dev || brew install gmp)
3. `go get github.com/eris-ltd/eris-cm/cmd/eris-cm`

## Overview

Currently `eris:chain_manager` provides the following functionality:

* `maker` -- wizard, config file, or csv based utility for creating keys and genesis.json's necessary for a range of different chain types.

(More functionality coming soon.)

## Usage

```
The Eris Chain Manager is a utility for performing complex operations on eris chains.

Made with <3 by Eris Industries.

Complete documentation is available at https://docs.erisindustries.com

Version:
  0.11.0

Usage:
  eris-cm [flags]
  eris-cm [command]

Available Commands:
  make        The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains

Flags:
  -d, --debug[=false]: debug level output; the most output available for eris-cm; if it is too chatty use verbose flag; default respects $ERIS_CHAINMANAGER_DEBUG
  -h, --help[=false]: help for eris-cm
  -o, --output[=true]: should eris-cm provide an output of its job; default respects $ERIS_CHAINMANAGER_OUTPUT
  -v, --verbose[=false]: verbose output; more output than no output flags; less output than debug level; default respects $ERIS_CHAINMANAGER_VERBOSE
```

or

```
The Eris Chain Maker is a utility for easily creating the files necessary to build eris chains.

Usage:
  eris-cm make [flags]

Examples:
$ eris-cm make myChain -- will use the chain-making wizard and make your chain named myChain using eris-keys defaults (available via localhost) (interactive)
$ eris-cm make myChain --chain-type=simplechain --  will use the chain type definition files to make your chain named myChain using eris-keys defaults (non-interactive)
$ eris-cm make myChain --account-types=Root:1,Developer:0,Validator:0,Participant:1 -- will use the flag to make your chain named myChain using eris-keys defaults (non-interactive)
$ eris-cm make myChain --account-types=Root:1,Developer:0,Validator:0,Participant:1 --chain-type=simplechain -- account types trump chain types, this command will use the flags to make the chain (non-interactive)
$ eris-cm make myChain --csv /path/to/csv -- will use the csv file to make your chain named myChain using eris-keys defaults (non-interactive)

Flags:
  -t, --account-types=[]: what number of account types should we use? find these in ~/.eris/chains/account-types; incompatible with and overrides chain-type; default respects $ERIS_CHAINMANAGER_ACCOUNTTYPES
  -c, --chain-type="": which chain type definition should we use? find these in ~/.eris/chains/chain-types; default respects $ERIS_CHAINMANAGER_CHAINTYPE
  -s, --csv-file="": csv file in the form `account-type,number,tokens,toBond,perms; default respects $ERIS_CHAINMANAGER_CSVFILE
  -h, --help[=false]: help for make
  -k, --keys-server="http://localhost:4767": keys server which should be used to generate keys; default respects $ERIS_KEYS_PATH
  -r, --tar[=false]: instead of making directories in ~/.chains, make tarballs; incompatible with and overrides zip; default respects $ERIS_CHAINMANAGER_TARBALLS
  -z, --zip[=false]: instead of making directories in ~/.chains, make zip files; default respects $ERIS_CHAINMANAGER_ZIPFILES

Global Flags:
  -d, --debug[=false]: debug level output; the most output available for eris-cm; if it is too chatty use verbose flag; default respects $ERIS_CHAINMANAGER_DEBUG
  -o, --output[=true]: should eris-cm provide an output of its job; default respects $ERIS_CHAINMANAGER_OUTPUT
  -v, --verbose[=false]: verbose output; more output than no output flags; less output than debug level; default respects $ERIS_CHAINMANAGER_VERBOSE
```

# Contributions

Are Welcome! Before submitting a pull request please:

* go fmt your changes
* have tests
* pull request
* be awesome

That's pretty much it (for now).

Please note that this repository is GPLv3.0 per the LICENSE file. Any code which is contributed via pull request shall be deemed to have consented to GPLv3.0 via submission of the code (were such code accepted into the repository).

# Bug Reporting

Found a bug in our stack? Make an issue!

Issues should contain four things:

* The operating system. Please be specific.
* The reproduction steps. Starting from a fresh environment, what are all the steps that lead to the bug? Also include the branch you're working from.
* What doyou expected to happen. Provide a sample output.
* What actually happened. Error messages, logs, etc. Use `-d` to provide the most information. For lengthy outputs, link to a gist or pastebin please.

Finally, add a label to your bug (critical or minor). Critical bugs will likely be addressed quickly while minor ones may take awhile. Pull requests welcome for either, just let us know you're working on one in the issue (we use the in-progress label accordingly).

# License

[Proudly GPL-3](http://www.gnu.org/philosophy/enforcing-gpl.en.html). See [license file](https://github.com/eris-ltd/eris-pm/blob/master/LICENSE.md).
