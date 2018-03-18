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

func (ps *State) Close() {
    for rpname, repo := range ps.Repos {
        log.Printf("closing db '%s'", rpname)
        repo.db.Close()
    }
}
