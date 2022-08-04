package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path"
)

func DirInfo(name string) []fs.FileInfo {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

// Prints directory tree and return number of files and directories
func tree(name, prefix string) (int, int) {
	dir_num, file_num := 0, 0
	ident := "    "
	branch := "│   "
	tee := "├── "
	last := "└── "

	files := DirInfo(name)
	for i, f := range files {
		pointer := last
		extension := ident
		if i < len(files)-1 {
			extension = branch
			pointer = tee
		}
		fmt.Println(prefix + pointer + f.Name())

		if f.IsDir() {
			dir_num++
			d_, f_ := tree(path.Join(name, f.Name()), prefix+extension)
			dir_num += d_
			file_num += f_
		} else {
			file_num++
		}
	}
	return dir_num, file_num
}

func Tree(name string) {
	_, dir := path.Split(name)
	fmt.Println(dir)

	// Counting files  and directories
	d, f := tree(name, "")
	fmt.Printf("%d directories, %d files.", d, f)
}

func main() {
	Tree(".")
}
