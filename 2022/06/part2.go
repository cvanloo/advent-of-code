package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const size = 14

type WrapBuffer struct {
	wrapped  bool
	contents [size]rune
	next     int
}

func (w *WrapBuffer) Add(r rune) {
	if w.wrapped {
		w.contents[w.next%size] = r
		w.next++
	} else {
		w.contents[w.next] = r
		w.next++

		if w.next >= size {
			w.wrapped = true
		}
	}
}

func (w WrapBuffer) HasWrapped() bool {
	return w.wrapped
}

func (w WrapBuffer) Get(idx int) rune {
	return w.contents[idx]
}

func (w WrapBuffer) GetAll() [size]rune {
	return w.contents
}

func (w WrapBuffer) Index() int {
	return w.next
}

func (w WrapBuffer) String() string {
	var rep string
	for _, r := range w.contents {
		rep += string(r)
	}
	return fmt.Sprintf("%s", rep)
}

func copyToClipboard(text string) {
	panicIf := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	if runtime.GOOS != "linux" {
		panic("have you heard of our lord and saviour, GNU/Linux?")
	}

	if _, err := exec.LookPath("xsel"); err != nil {
		panic("paru -Syu xsel")
	}

	command := exec.Command("xsel", "-bi")
	stdin, err := command.StdinPipe()
	panicIf(err)

	err = command.Start()
	panicIf(err)

	io.WriteString(stdin, text)
	stdin.Close() // Command won't exit until stdin is closed.
	err = command.Wait()
	panicIf(err)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		buffer := &WrapBuffer{}

		for _, r := range line {
			buffer.Add(r)
			if buffer.HasWrapped() {
				repeated := false
				m := map[rune]int{}
				for _, r := range buffer.GetAll() {
					m[r]++
					if m[r] > 1 {
						//log.Printf("repeated at %d (%s)\n", buffer.Index(), string(r))
						repeated = true
						break
					}
				}

				if !repeated {
					copyToClipboard(fmt.Sprintf("%d", buffer.Index()))
					log.Printf("marker at %d (copied to clipboard)", buffer.Index())
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
}
