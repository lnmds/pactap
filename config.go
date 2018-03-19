package main

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
)

const defaultConfig string = `# pactap's main configuration file

# Where should things like /bin and /lib go to?
MainPath = "~/.pactap"

# Enable debugging
Debug = false

# Packages to ignore updates from
Ignore = []

# Only download repo patches
# setting this to false will download the entire
# repo file on each repo update
SlowMode = true

CheckIntegrity = true
CheckSignature = false

# Example repository information
[repo]
    [repo.local]
    Remote = "file:///home/luna/pactap/repo.db"

    [repo.core]
    Remote = "https://pactap.lnmds.me"

    # Example of a repository that uses a mirrorlist file
    [repo.community]
    RemoteList = "file:///etc/pactap/mirrorlist"

    # In the case you want some binary-only
    #  package repository but don't want to host
    #  it with a full vps or something
    [repo.bin]
    Remote = "https://localhost:6969"
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

	log.Printf("[config:load] path '%s'", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("[config] Using fallback")

		data = []byte(defaultConfig)
		err = ioutil.WriteFile(path, data, 0755)

		if err != nil {
			log.Fatalf("Failed to write to conf file. %s", err)
		}
	}

	toml.Unmarshal(data, &c)

	log.Printf("[config] success")
	return &c
}
