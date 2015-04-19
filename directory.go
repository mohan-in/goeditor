package main

import (
	"fmt"
	"io/ioutil"
)

type Dir struct {
	Name  string
	Files []string
	Dirs  []Dir
}

func ReadDir(name string) Dir {
	dir := Dir{Name: name}

	c := make(chan Dir)
	go populate(c, dir)
	dir = <-c
	close(c)

	return dir
}

func populate(c chan Dir, d Dir) {

	files, err := ioutil.ReadDir(d.Name)
	if err != nil {
		fmt.Println(err)
	}

	cc := make(chan Dir)
	j := 0

	for _, file := range files {
		if file.IsDir() {
			dir := Dir{Name: d.Name + "/" + file.Name()}
			go populate(cc, dir)
			j++
		} else {
			d.Files = append(d.Files, file.Name())
		}
	}

	for ; j > 0; j-- {
		d.Dirs = append(d.Dirs, <-cc)
	}
	close(cc)

	c <- d
}