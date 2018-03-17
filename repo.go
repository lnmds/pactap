package main

import (
    "io/ioutil"
    "fmt"
    "path/filepath"
    "log"
    "errors"
    "net/http"
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

func downloadRemote(remoteType RemoteType, remote string) (string, error) {
    log.Printf("[download:remote] '%s'", remote)

    if remoteType == FILE {
        data, err := ioutil.ReadFile(remote)
        if err != nil {
            return "", errors.New("Failure reading from remote path")
        }
        return string(data), err

    } else if remoteType == HTTP || remoteType == HTTPS{
        resp, err := http.Get(remote)
        if err != nil {
            log.Printf("[download:remote] error! %s", err)
            return "", errors.New("Error downloading remote")
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Printf("[download:remote] error on download! %s", err)
            return "", errors.New("Error while downloading remote")
        }

        return string(body), nil
    }

    return "", errors.New(fmt.Sprintf("[download:remote] Invalid remote type. %d remote=%s", remoteType, remote))
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
        dbdata, err := downloadRemote(remoteType, remote)
        if err != nil {
            log.Fatalf("Failure downloading remote. %s", err)
        }

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
