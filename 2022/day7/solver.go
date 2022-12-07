package day7

import (
	"mrose.de/aoc/utility"
	"strconv"
	"strings"
)

type directory struct {
	subdirectories []*subdirectory
	size           int
	name           string
	files          []file
}

type file struct {
	size int
	name string
}

type subdirectory struct {
	directory
	parentDirectory *subdirectory
}

var rootDirectory *directory
var currentDirectory *subdirectory
var allDirectories []*subdirectory

func Solve() (result int) {
	lines := utility.StrArr(utility.Input(2022, 7))
	rootDirectory = &directory{
		subdirectories: make([]*subdirectory, 0),
		name:           "/",
	}
	currentDirectory = &subdirectory{
		directory:       *rootDirectory,
		parentDirectory: nil,
	}
	addDirectoryToOverview(currentDirectory)
	analyze(lines)

	free := 0
	needed := 0
	rootSize := 0
	for _, d := range allDirectories {
		if d.name == "/" {
			//needed = 70000000 - free
			rootSize = d.size
			free = 70000000 - rootSize
			needed = 30000000 - free
		}
	}
	tmp := 0
	for _, d := range allDirectories {
		if d.size > needed {
			if tmp == 0 {
				tmp = d.size
			}
			if d.size < tmp {
				tmp = d.size
			}
		}
	}
	return tmp
}

func addDirectoryToOverview(d *subdirectory) {
	allDirectories = append(allDirectories, d)
}

func analyze(lines []string) {
	for _, line := range lines {
		components := strings.Split(line, " ")
		switch components[0] {
		case "$":
			execute(line)
		case "dir":
			exists := false
			for _, sd := range currentDirectory.subdirectories {
				if sd.name == components[1] {
					exists = true
					break
				}
			}
			if !exists {
				newDirectory := &subdirectory{
					directory: directory{
						subdirectories: make([]*subdirectory, 0),
						size:           0,
						name:           components[1],
						files:          make([]file, 0),
					},
					parentDirectory: currentDirectory}
				currentDirectory.subdirectories = append(currentDirectory.subdirectories, newDirectory)
				addDirectoryToOverview(newDirectory)
			}
		default:
			exists := false
			for _, f := range currentDirectory.files {
				if f.name == components[1] {
					exists = true
					break
				}
			}
			if !exists {
				sizeToAdd, _ := strconv.Atoi(components[0])
				currentDirectory.files = append(currentDirectory.files, file{
					size: sizeToAdd,
					name: components[1],
				})
				currentDirectory.size += sizeToAdd
				tmp := currentDirectory
				for {
					if tmp.name == "/" {
						break
					}
					tmp = tmp.parentDirectory
					tmp.size += sizeToAdd
				}
			}
		}
	}
}

func execute(line string) {
	if strings.Contains(line, "$ cd") {
		components := strings.Split(line, " ")
		if components[2] == ".." {
			currentDirectory = currentDirectory.parentDirectory
		} else {
			for _, s := range currentDirectory.subdirectories {
				if s.name == components[2] {
					currentDirectory = s
					return
				}
			}
			newDirectory := &subdirectory{
				directory: directory{
					subdirectories: make([]*subdirectory, 0),
					size:           0,
					name:           components[2],
				},
				parentDirectory: currentDirectory,
			}
			addDirectoryToOverview(newDirectory)
			currentDirectory.subdirectories = append(currentDirectory.subdirectories, newDirectory)
			currentDirectory = newDirectory
		}
	}
}
