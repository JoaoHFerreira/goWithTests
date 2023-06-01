package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const ONE_SECOND time.Duration = 1 * time.Second

type terminal struct {
	pwd   string
	dirs  map[string][]string
	files map[string][]string
	// list_content map[string]string
}

func (t *terminal) run() {
	for {
		raw_cmd := t.getCMD()
		cmd := strings.Split(raw_cmd, " ")[0]

		switch cmd {
		case "pwd\n":
			t.currentPath()
		case "ls\n":
			t.listCurrentDir()
		case "touch":
			t.touch(raw_cmd)
		case "echo":
			t.echo(raw_cmd)
		case "cd":
			t.cd(raw_cmd)
		case "mkdir":
			t.mkdir(raw_cmd)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

func (t *terminal) getCMD() string {
	fmt.Printf("/")
	keyWord, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return keyWord
}

func (t *terminal) currentPath() {
	fmt.Println(t.pwd)
}

func (t *terminal) listCurrentDir() {
	totalFiles := append(t.dirs[t.pwd], t.files[t.pwd]...)
	fmt.Println(totalFiles)
}

func (t *terminal) touch(raw_cmd string) {
	filesOnly := strings.Split(raw_cmd, " ")[1:]
	t.files[t.pwd] = append(t.files[t.pwd], filesOnly...)
}

func (t *terminal) echo(raw_cmd string) {
	fmt.Println("here, only the echo key word will be excluded")
}

func (t *terminal) cd(raw_cmd string) {
	splitedCommands := strings.Split(raw_cmd, " ")
	dirToGo := splitedCommands[len(splitedCommands)-1]
	fmt.Println(t.dirs[dirToGo])
}

func (t *terminal) mkdir(raw_cmd string) {
	dirsOnly := strings.Split(raw_cmd, " ")[1:]
	t.dirs[t.pwd] = append(t.dirs[t.pwd], dirsOnly...)
}

func main() {
	goTerminal := terminal{
		pwd: "/",
		dirs: map[string][]string{
			"/": {"/"},
		},
		files: make(map[string][]string),
	}
	goTerminal.run()
}
