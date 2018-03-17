package main

import (
    "fmt"
    "os"
)

const VERSION string = "0.0.1"

func usage(){
    fmt.Printf("usage: pactap [main operators] [...]\n")
}

func version(){
    fmt.Printf("Pactap v%s\n" +
               "Copyright (C) 2018 Luna Mendes\n", VERSION)
}

func main(){
    if len(os.Args) < 2 {
        usage()
        return
    }

    // TODO: We should start reading our config files
    conf := ReadConfig("/home/luna/.pactap/config.toml")
    fmt.Println(*conf)
    // TODO: We should start reading our db files
}
