package main

import (
    "log"
)

// Main program state
type State struct {
    Config *Main

    // Raw repository configuration
    RepoConfig Repos

    // Fully complete repositories
    // with their db connections
    Repos map[string]Repository
}

func (ps *State) Start() {
    ps.Repos = make(map[string]Repository)

    for rpname, _ := range ps.RepoConfig {
        repo, err := RepoOpen(ps.Config, rpname)

        if err != nil {
            log.Printf("oops! %s", err)
            continue
        }

        ps.Repos[rpname] = repo
    }
}

func (ps *State) Close() {
    for rpname, repo := range ps.Repos {
        log.Printf("closing db '%s'", rpname)
        repo.db.Close()
    }
}
