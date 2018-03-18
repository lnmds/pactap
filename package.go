package main

import (
    "log"
    "fmt"
    "errors"

    "gopkg.in/yaml.v2"
)

type Package struct {
    name string

    version string
    build int

    depedencies []Package
}

func GetPackage(ps *State, pkg_name string) (Package, error) {
    if len(ps.Repos) == 0 {
        return Package{}, errors.New("No repositories are initialized")
    }

    for reponame, repo := range ps.Repos {
        log.Printf("finding package %s", pkg_name)
        pkg, err := FindPackage(repo, pkg_name)

        if err == nil {
            return pkg, nil
        } else {
            log.Printf("error fetching %s from repo %s: %s", pkg_name, reponame, err)
        }
    }

    return Package{}, errors.New(fmt.Sprintf("package '%s' not found", pkg_name))
}

func FindPackages(ps *State, packages []string) ([]Package, error) {
    pkgs := make([]Package, 100)

    for idx := range packages {
        pkgstruct, err := GetPackage(ps, packages[idx])

        if err != nil {
            return pkgs, err
        }

        pkgs = append(pkgs, pkgstruct)
    }

    return pkgs, nil
}

func resolvePackages (packages []Package) []Package {
    // Given a list of packages, output another list of packages,
    // Give out a list of packages to install, in order.

    log.Printf("Resolving lmao")

    return []Package{}
}

func OpenPackageData(data string) (Package, error) {
    pkg := Package{}
    err := yaml.Unmarshal([]byte(data), pkg)

    if err != nil {
        return Package{}, err
    }

    return pkg, nil
}
