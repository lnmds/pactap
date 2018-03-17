package main

import (
    "io/ioutil"
    "fmt"
    "path/filepath"
    "log"
    "strings"
)

type Repos map[string]Repo

type RemoteType int

const (
    HTTP RemoteType = iota
    HTTPS
    FTP
    FILE

    INVALID
)

func getRepos(c *Main) Repos {
    return c.Repos
}

func getRepoDBPath(c *Main, reponame string) string {
    // e.g ~/.pactap/db/repo_reponame.db
    return filepath.Join(c.MainPath, fmt.Sprintf("db/repo_%s.db", reponame))
}

func getRemoteType(remote string) RemoteType {
    if strings.HasPrefix(remote, "http://") {
        return HTTP
    } else if strings.HasPrefix(remote, "https://") {
        return HTTPS
    } else if strings.HasPrefix(remote, "ftp://") {
        return FTP
    } else if strings.HasPrefix(remote, "file://") {
        return FILE
    } else {
        return INVALID
    }
}

func downloadRemote(remoteType RemoteType, remote string) string {
    log.Printf("[download:remote] '%s' (fake download)", remote)
    return ""
}

func updateSingleRepo(c *Main, reponame string, repo Repo) {
    // First, we check what protocol is the repo using
    // as a remote.

    remote := repo.Remote
    if remote == "" {
        // try remote list
        _ = repo.RemoteList
        // ???
    }

    remoteType := getRemoteType(remote)

    if remoteType == INVALID {
        log.Printf("[repo:%s] This remote is invalid: '%s'", reponame, remote)
        return
    }

    if c.SlowMode {
        // TODO: download patches and all
        // Slow Mode does not work with FILE
    } else {
        dbdata := downloadRemote(remoteType, remote)
        // Write new db data

        repopath := getRepoDBPath(c, reponame)
        ioutil.WriteFile(repopath, []byte(dbdata), 0700)
    }
}

func updateRepos(c *Main) {
    // Query the interwebs.
    repos := getRepos(c)

    for reponame, repo := range repos {
        log.Printf("Updating repo '%s'", reponame)

        updateSingleRepo(c, reponame, repo)
    }
}
