package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Inode struct {
	fullPath string /* Absolute path of the inode */
	name     string
	size     int      /* Size of the inode, only set if inode is a file */
	isDir    bool     /* True if the inode is a dir, false if it is a file */
	parent   *Inode   /* Containing directory */
	contents []*Inode /* Child inodes, only set if inode is a directory */
}

// Size calculates the total size of the inode, recurring into subdirectories.
func (i Inode) Size() int {
	if i.isDir {
		size := 0
		for _, child := range i.contents {
			size += child.Size()
		}
		return size
	}
	return i.size
}

// String formats the inode for pretty printing, recurs into subdirectories.
func (i Inode) String() string {
	return i.stringInode(0)
}

func (i Inode) stringInode(indent int) string {
	var output string

	for i := 0; i < indent; i++ {
		output += "  "
	}

	typeText := "file"
	if i.isDir {
		typeText = "dir"
	}
	output += fmt.Sprintf("- %s (%s, size=%d)\n", i.name, typeText, i.Size())

	if i.isDir {
		indent++
		for _, child := range i.contents {
			output += child.stringInode(indent)
		}
	}

	return output
}

// TraverseDirectories recursively traverses through all directories rooted below root.
// On every directory predicate is applied.
// The result is a slice of all directories for which predicate returned true.
// Essentially acts as a filter.
func TraverseDirectories(root *Inode, predicate func(*Inode) bool) []*Inode {
	var result []*Inode
	for _, child := range root.contents {
		if !child.isDir {
			continue
		}
		if predicate(child) {
			result = append(result, child)
		}
		result = append(result, TraverseDirectories(child, predicate)...)
	}
	return result
}

const maxSpaceAvailable = 70_000_000
const spaceNeeded = 30_000_000

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
					absolutePath := pwd.fullPath + param + "/"
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
				absolutePath := pwd.fullPath + name + "/"
				if _, ok := fs[absolutePath]; ok {
					// Inode already exists.
					continue
				}
				newDir := &Inode{
					fullPath: absolutePath,
					name:     name,
					isDir:    true,
					parent:   pwd,
				}
				fs[absolutePath] = newDir
				pwd.contents = append(pwd.contents, newDir)
			} else {
				absolutePath := pwd.fullPath + name
				if _, ok := fs[absolutePath]; ok {
					// Inode already exists.
					continue
				}
				size, _ := strconv.Atoi(description)
				newFile := &Inode{
					fullPath: absolutePath,
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

	totalSpaceUsed := root.Size()
	spaceToFree := spaceNeeded - maxSpaceAvailable + totalSpaceUsed

	result := totalSpaceUsed
	deletionCandidates := TraverseDirectories(root, func(inode *Inode) bool {
		size := inode.Size()
		if size >= spaceToFree && size < result {
			result = size
			return true
		}
		return false
	})

	deletionCandidate := deletionCandidates[len(deletionCandidates)-1]

	fmt.Printf("Result: %v\n", deletionCandidate.Size())
}
