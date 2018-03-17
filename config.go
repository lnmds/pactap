package main

import (
    "github.com/BurntSushi/toml"
    "io/ioutil"
)

const defaultConfig string = `
[main]
MainPath = "~/.pactap"
Debug = false
Ignore = []

# Only download repo patches
# setting this to false will download the entire
# repo file on each repo update
SlowMode = true

CheckIntegrity = true
CheckSignature = false

[repo]
    [repo.core]
    Remote = "https://pactap.lnmds.me"

    [repo.community]
    RemoteList = "/etc/pactap/mirrorlist"

    # In the case you want some binary-only
    # package repository
    [repo.bin]
    Remote = "https://localhost:6969"

    # [repo.myAss] # this can be any path
    # Remote = file:///home/lunae/pactap/myass
`

type Repo struct {
    Remote string

    RemoteList string
}

type Main struct {
    // Main path for EVERYTHING. default "~/.pactap"
    MainPath string

    // enable debug log?
    Debug bool

    // Packages to ignore updates from
    Ignore []string

    // enable pactap's slow mode
    SlowMode bool

    // check hashes of shit, default false
    // USED ONLY FOR PACKAGE BUILDING
    CheckIntegrity bool

    // TODO: do we really do this with gpg and the shit
    CheckSignature bool

    Repos map[string]Repo `toml:"repo"`
}

func ReadConfig(path string) *Main {
    var c Main

    data, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    if _, err := toml.Decode(string(data), &c); err != nil {
        panic(err)
    }

    return &c
}
