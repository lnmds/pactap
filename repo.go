package main

import (
    "io/ioutil"
    "fmt"
    "path/filepath"
    "log"
    "errors"
    "net/http"
    "strings"

    "github.com/mitchellh/go-homedir"
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
    exp, err := homedir.Expand(c.MainPath)

    if err != nil {
        log.Panic(err)
    }

    return filepath.Join(exp, fmt.Sprintf("db/%s.db", reponame))
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
        remoteFiltered := strings.TrimPrefix(remote, "file://")

        remoteFiltered, err := homedir.Expand(remoteFiltered)

        if err != nil {
            return "", err
        }

        log.Printf("[download:file] expanded path is '%s'", remoteFiltered)

        data, err := ioutil.ReadFile(remoteFiltered)
        if err != nil {
            return "", err
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

func UpdateSingleRepo(c *Main, reponame string, repo Repo) {
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

    if c.SlowMode && remoteType != FILE {
        // TODO: download patches and all
        // Slow Mode does not work with FILE

        /*
        latestPatch := fetchPatchLatest(reponame)
        db := RepoOpen(c, reponame)
        version := getRepoDBVersion(db)

        for idx := range verison, latestPatch {
            downloadAndApplyPatch(remoteType, remote, idx)
        }
        */
    } else {
        dbdata, err := downloadRemote(remoteType, remote)
        if err != nil {
            log.Fatalf("Failure downloading remote. %s", err)
        }

        // Write new db data
        repopath := getRepoDBPath(c, reponame)

        log.Printf("Writing to '%s'", repopath)
        err = ioutil.WriteFile(repopath, []byte(dbdata), 0700)
        if err != nil {
            log.Fatalf("Error writing to repo db! %s", err)
        }

        log.Printf("Updated repo %s, success!", reponame)
    }
}

func UpdateRepos(c *Main) {
    log.Printf("Starting repository update")
    for reponame, repo := range c.Repos {
        log.Printf("Updating repo '%s'", reponame)
        UpdateSingleRepo(c, reponame, repo)
    }
}
