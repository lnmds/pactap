# pactap
the memework package manager - short for "packaging tape" (as you are free to pack anything)

## Motivation

Memework (a group of friends) have a server, however not everyone can have root.

**but** we want to let users able to install their packages.

the first version of the solution to that idea was `rsudo`, a program made
by us so people can request execution of a command as root.

this is the second iteration, called `pactap`. It was made to be the actual
second iteration of the solution but due to internal issues within the group
it became more of a "learning" project.

we will, very likely, use http://linuxbrew.sh/ over this project.

## Installing

```bash
# make sure $GOPATH/bin is in your path
mkdir -p ~/.pactap && go get -u github.com/lnmds/pactap
# ??? profit ???
```

## Updating

```bash
go get -u github.com/lnmds/pactap
```

# Repository management

## how do make repo???

the basic is making a repository db file:
```bash
sqlite3 -init pkgtools/repo_start.sql my_repo_file.db
```
