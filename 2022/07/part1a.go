package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inode struct {
	fullpath string
	name     string
	size     int
	isDir    bool
	parent   *Inode
	contents []*Inode
}

func (i Inode) Size() int {
	if i.isDir {
		size := 0
		for _, child := range i.contents {
			size += child.Size()
		}
		return size
	} else {
		return i.size
	}
}

func (i Inode) String() string {
	return i.StringInode(0)
}

func (inode Inode) StringInode(indent int) string {
	var output string

	for i := 0; i < indent; i++ {
		output += "  "
	}

	typeText := "file"
	if inode.isDir {
		typeText = "dir"
	}
	output += fmt.Sprintf("- %s (%s, size=%d)\n", inode.name, typeText, inode.Size())

	if inode.isDir {
		indent++
		for _, child := range inode.contents {
			output += child.StringInode(indent)
		}
	}

	return output
}

const maxSize = 100_000

func TraverseDirectories(root *Inode, visit func(*Inode) bool) []*Inode {
	var result []*Inode
	for _, child := range root.contents {
		if !child.isDir {
			continue
		}
		if visit(child) {
			result = append(result, child)
		}
		result = append(result, TraverseDirectories(child, visit)...)
	}
	return result
}

func main() {
	root := &Inode{name: "/", isDir: true}
	pwd := root
	fs := map[string]*Inode{
		"/": root,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			// command
			line := strings.TrimLeft(line, "$ ")
			parts := strings.Split(line, " ")
			command := parts[0]
			if command == "ls" {
				// this one is simple: do nothing
				continue
			}
			if command == "cd" {
				param := parts[1]
				switch param {
				case "/":
					pwd = root
				case "..":
					if pwd.parent != nil {
						pwd = pwd.parent
					}
				default:
					absolutePath := pwd.fullpath + param + "/"
					if dir, ok := fs[absolutePath]; ok {
						pwd = dir
					} else {
						panic("I don't know this directory")
					}
				}
			}
		} else {
			// output
			parts := strings.Split(line, " ")
			description := parts[0]
			name := parts[1]

			if description == "dir" {
				absolutePath := pwd.fullpath + name + "/"
				if _, ok := fs[absolutePath]; ok {
					// Inode already exists.
					continue
				}
				newDir := &Inode{
					fullpath: absolutePath,
					name:     name,
					isDir:    true,
					parent:   pwd,
				}
				fs[absolutePath] = newDir
				pwd.contents = append(pwd.contents, newDir)
			} else {
				absolutePath := pwd.fullpath + name
				if _, ok := fs[absolutePath]; ok {
					// Inode already exists.
					continue
				}
				size, _ := strconv.Atoi(description)
				newFile := &Inode{
					fullpath: absolutePath,
					name:     name,
					size:     size,
					isDir:    false,
					parent:   pwd,
				}
				fs[absolutePath] = newFile
				pwd.contents = append(pwd.contents, newFile)
			}
		}
	}

	fmt.Printf("%s", root)

	deletionCandidates := TraverseDirectories(root, func(inode *Inode) bool {
		size := inode.Size()
		return size <= maxSize
	})

	result := 0
	for _, dc := range deletionCandidates {
		result += dc.Size()
	}

	fmt.Printf("Result: %v\n", result)
}
