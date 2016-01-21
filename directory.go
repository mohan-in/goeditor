package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Dir struct {
	Name  string
	Path  string
	Files []File
	Dirs  []Dir
}

type File struct {
	Name string
	Path string
}

var ignoreDirs = []string{"bower_components", ".git", ".gitignore", "LICENSE", "README.md"}

func ReadDir(name string) Dir {
	p := strings.Split(name, "/")
	dir := Dir{Path: name, Name: p[len(p)-1]}

	c := make(chan Dir)
	go populate(c, dir)
	dir = <-c
	close(c)

	return dir
}

func populate(c chan Dir, d Dir) {

	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		fmt.Println(err)
	}

	cc := make(chan Dir)
	j := 0

	for _, file := range files {
		if isIgnoreFile(file.Name()) {
			continue
		}
		if file.IsDir() {
			dir := Dir{Name: file.Name(), Path: d.Path + "/" + file.Name()}
			go populate(cc, dir)
			j++
		} else {
			f := File{Name: file.Name(), Path: strings.TrimPrefix(d.Path, goPath) + "/" + file.Name()}
			d.Files = append(d.Files, f)
		}
	}

	for ; j > 0; j-- {
		d.Dirs = append(d.Dirs, <-cc)
	}
	close(cc)

	c <- d
}

func isIgnoreFile(name string) bool {
	for _, n := range ignoreDirs {
		if n == name {
			return true
		}
	}
	return false
}
