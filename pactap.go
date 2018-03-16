package main

import (
    "fmt"
    "os"
)

func usage(){
    fmt.Printf("usage: pactap [main operators] [...]\n")
}

func main(){
    if len(os.Args) < 2 {
        usage()
        return
    }

    op := os.Args[1]
    if op[0] != '-' {
        usage()
        return
    }

    mainop := string(op[1])
    // Depending of mainop being S, Q, whatever
    // we do our operation.

    fmt.Printf(mainop)
}
