package main

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
    "strings"

    "github.com/jessevdk/go-flags"
    "github.com/mitchellh/go-homedir"
)

const VERSION string = "0.0.1"

var opts struct {
    Arch bool `long:"arch" description:"Specify an alternate architecture."`
    Fancy bool `long:"fancy" description:"Force fancy mode on non-tty systems."` // TODO this is for CLI stuff, will be disabled on detection of non-tty
    Config string `long:"config" description:"Specify an alternate config file."`
    Debug bool `long:"debug" description:"Display debug messages. Use when reporting bugs."`
}

func version(){
    fmt.Printf("Pactap v%s\n" +
               "Copyright (C) 2018 Luna Mendes\n", VERSION)
}

func main(){
    if len(os.Args) < 2 {
        os.Args = append(os.Args, "-h")
    }

    args, err := flags.Parse(&opts)
    if err != nil {
        return
    }

    operator := args[0]

    if operator == "version" {
        version()
        return
    }

    homedir, err := homedir.Dir()
    if err != nil {
        panic(err)
    }

    // TODO: get ConfigPath from opts
    configPath := ".config/pactap/config.toml"

    if runtime.GOOS == "darwin" {
        configPath = "Library/Application Support/pactap/config.toml"
    }

    conf := ReadConfig(filepath.Join(homedir, configPath))

    // Start main program state
    state := &State{
        Config: conf,
        RepoConfig: conf.Repos,
    }

    defer state.Close()

    // TODO: initialize repos

    if operator == "update" {
        UpdateRepos(conf)
    } else if operator == "install" {
        packages := args[1:]
        fmt.Printf("packages to install: %s\n", strings.Join(packages, " "))
        // TODO: install lmao
    } else {
        fmt.Printf("Invalid operator: %s\n", operator)
    }
}


func Filter(vs []bool,) []bool {
    vsf := make([]bool, 0)
    for _, v := range vs {
        if v {
            vsf = append(vsf, v)
        }
    }
    return vsf
}
