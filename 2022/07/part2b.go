package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Inode represents a filesystem entry.
// An inode can be either a file or a directory (isDir).
type Inode struct {
	// Absolute path of the inode
	fullPath string
	// Inode name
	name string
	// Size of the inode, only set if inode is a file
	size int
	// True if the inode is a dir, false if it is a file
	isDir bool
	// Containing directory
	parent *Inode
	// Child inodes, only set if inode is a directory
	contents []*Inode
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

// FileSystem simulates a filesystem
type FileSystem struct {
	// Root of the filesystem
	root *Inode
	// Pwd references the current working directory
	pwd *Inode
	// fs stores all inodes of the filesystem.
	// The map allows for easy access based on an absolute path.
	fs map[string]*Inode
}

// CreateDirectory creates a new directory rooted in the current working directory.
func (fs *FileSystem) CreateDirectory(name string) error {
	absolutePath := fs.pwd.fullPath + name + "/"

	if _, ok := fs.fs[absolutePath]; ok {
		return fmt.Errorf("directory already exists: %s", absolutePath)
	}

	newDir := &Inode{
		fullPath: absolutePath,
		name:     name,
		isDir:    true,
		parent:   fs.pwd,
	}
	fs.fs[absolutePath] = newDir
	fs.pwd.contents = append(fs.pwd.contents, newDir)
	return nil
}

// CreateFile creates a new file rooted in the current working directory.
func (fs *FileSystem) CreateFile(name string, size int) error {
	absolutePath := fs.pwd.fullPath + name
	if _, ok := fs.fs[absolutePath]; ok {
		return fmt.Errorf("file already exists: %s", absolutePath)
	}
	newFile := &Inode{
		fullPath: absolutePath,
		name:     name,
		size:     size,
		isDir:    false,
		parent:   fs.pwd,
	}
	fs.fs[absolutePath] = newFile
	fs.pwd.contents = append(fs.pwd.contents, newFile)
	return nil
}

// ChangeDirectory changes the current working directory.
// Path must be either '..', '/', or a subdirectory (below pwd).
func (fs *FileSystem) ChangeDirectory(path string) error {
	switch path {
	case "/":
		fs.pwd = fs.root
	case "..":
		if fs.pwd.parent != nil {
			fs.pwd = fs.pwd.parent
		}
	default: // relative path
		absolutePath := fs.pwd.fullPath + path + "/"
		if dir, ok := fs.fs[absolutePath]; ok {
			fs.pwd = dir
		} else {
			return fmt.Errorf("directory does not exist: %s", absolutePath)
		}
	}
	return nil
}

func findDeletionCandidate(root *Inode, smallest *Inode, spaceToFree int) *Inode {
	for _, child := range root.contents {
		if !child.isDir {
			continue
		}
		size := child.Size()
		if size >= spaceToFree && size < smallest.Size() {
			smallest = child
		}
		smaller := findDeletionCandidate(child, smallest, spaceToFree)
		if smaller != nil {
			smallest = smaller
		}
	}
	return smallest
}

const maxSpaceAvailable = 70_000_000
const spaceNeeded = 30_000_000

func main() {
	root := &Inode{name: "/", isDir: true}
	fs := &FileSystem{
		root: root,
		pwd:  root,
		fs: map[string]*Inode{
			"/": root,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			// line is a command
			line := strings.TrimLeft(line, "$ ")
			parts := strings.Split(line, " ")
			command := parts[0]
			if command == "cd" {
				_ = fs.ChangeDirectory(parts[1])
			}
		} else {
			// line is output
			parts := strings.Split(line, " ")
			description := parts[0]
			name := parts[1]

			if description == "dir" {
				_ = fs.CreateDirectory(name)
			} else {
				size, _ := strconv.Atoi(description)
				_ = fs.CreateFile(name, size)
			}
		}
	}

	fmt.Printf("%s", root)

	spaceToFree := spaceNeeded - maxSpaceAvailable + root.Size()
	deletionCandidate := findDeletionCandidate(root, root, spaceToFree)
	fmt.Printf("Result: %v\n", deletionCandidate.Size())
}
