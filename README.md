# Package liborg 

## Introduction

This package demonstrates the usual directory and file organization of a Go language
library in a Git repo which can be imported by a Go main package or another package
in a different Git repo. 

Optionally this Git repo can contain a main package as well as the library. This
might be used to contain a command line integration test, or a functional demonstration
of the library.

## Library format  

``` sh
tree ~/go/src/github.com/x0ray/liborg

```

## Git new liborg repo creation

``` sh
cd ~/go/src/github.com/x0ray
mkdir loborg
cd liborg
echo "# liborg" >> README.md
nano README.md
git init
Initialized empty Git repository in /home/david/go/src/github.com/x0ray/liborg/.git/
git add README.md
git commit -m "initial commit"
git remote add origin git@github.com/x0ray/liborg.git
git push -u origin master
```

## Building the liborg library

``` sh
cd ~/go/src/github.com/x0ray/liborg
go build ./...
```

## Building and installing the liborgtest demo command

``` go
export GOBIN ~/bin
cd ~/go/src/github.com/x0ray/liborg
go install cmd/testliborg/main.go
```
