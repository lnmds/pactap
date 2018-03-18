package main

import (
    "log"
    "errors"
)

type Package struct {
    name string

    version string
    build int

    depedencies []Package
}

func GetPackage(ps *State, pkg_name string) (Package, error) {
    for _, repo := range ps.Repos {
        pkg, err := FindPackage(repo, pkg_name)

        if err != nil {
            return pkg, nil
        }
    }

    return Package{}, errors.New("package not found")
}

func resolvePackages (packages []Package) []Package {
    // Given a list of packages, output another list of packages,
    // Give out a list of packages to install, in order.

    log.Printf("Resolving lmao")

    return []Package{}
}
