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
	fmt.Printf(t.pwd)
	keyWord, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return keyWord
}

func (t *terminal) currentPath() {
	fmt.Println(t.pwd)
}

func (t *terminal) listCurrentDir() {
	var totalFiles []string
	totalFiles = append(totalFiles, t.dirs[t.pwd]...)
	totalFiles = append(totalFiles, t.files[t.pwd]...)
	fmt.Println(totalFiles)
}

func (t *terminal) touch(raw_cmd string) {
	filesOnly := strings.Split(strings.TrimSpace(raw_cmd), " ")[1:]
	t.files[t.pwd] = append(t.files[t.pwd], filesOnly...)
}

func (t *terminal) echo(raw_cmd string) {
	echoPart := strings.Split(raw_cmd, " ")[1:]
	joinedString := strings.Join(echoPart, " ")
	fmt.Println(joinedString)
}

func (t *terminal) cd(raw_cmd string) {
	desiredPath := strings.TrimSpace(strings.Split(raw_cmd, " ")[1])

	if desiredPath == "/" {
		t.pwd = desiredPath
		return
	}

	_, exists := t.dirs[t.pwd+desiredPath+"/"]

	if exists {
		t.pwd = t.pwd + desiredPath + "/"
	} else {
		fmt.Println("cd: no such file or directory:", desiredPath)
	}
}

func (t *terminal) mkdir(raw_cmd string) {
	dirsOnly := strings.Split(strings.TrimSpace(raw_cmd), " ")[1:]
	for _, path := range dirsOnly {
		composeString := t.pwd + path + "/"
		t.dirs[t.pwd] = append(t.dirs[t.pwd], composeString)
		t.dirs[composeString] = []string{} // Initialize the dirs slice for each new directory
	}
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
